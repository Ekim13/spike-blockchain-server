package chain

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis"
	"math/big"
	"strings"
)

type GameVaultTarget struct {
	txAddress string
}

func newGameVaultTarget(address string) *GameVaultTarget {
	return &GameVaultTarget{
		txAddress: address,
	}
}

func (t *GameVaultTarget) Accept(fromAddr, toAddr string) (bool, uint64) {
	if strings.ToLower(t.txAddress) == strings.ToLower(fromAddr) {
		return true, BNB_WITHDRAW
	}
	return false, NOT_EXIST
}

type GameVaultListener struct {
	TxFilter
	contractAddr   string
	tokenType      TokenType
	erc20Notify    chan ERC20Tx
	newBlockNotify DataChannel
	ec             *ethclient.Client
	rc             *redis.Client
	abi            abi.ABI
	errorHandle    chan ErrMsg
}

func newGameVaultListener(filter TxFilter, contractAddr string, tokenType TokenType, ec *ethclient.Client, rc *redis.Client, erc20Notify chan ERC20Tx, newBlockNotify DataChannel, abi abi.ABI, errHandle chan ErrMsg) *GameVaultListener {
	return &GameVaultListener{
		filter,
		contractAddr,
		tokenType,
		erc20Notify,
		newBlockNotify,
		ec,
		rc,
		abi,
		errHandle,
	}
}

func (el *GameVaultListener) run() {
	go el.NewEventFilter()
}

func (el *GameVaultListener) NewEventFilter() error {
	for {
		select {
		case de := <-el.newBlockNotify:
			height := de.Data.(*big.Int)
			el.handlePastBlock(height, height)
		}
	}
}

func (el *GameVaultListener) handlePastBlock(fromBlockNum, toBlockNum *big.Int) error {
	log.Infof("erc20 past event filter, type : %s, fromBlock : %d, toBlock : %d ", el.tokenType.String(), fromBlockNum, toBlockNum)
	contractAddress := common.HexToAddress(el.contractAddr)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		FromBlock: fromBlockNum,
		ToBlock:   toBlockNum,
	}

	sub, err := el.ec.FilterLogs(context.Background(), query)
	if err != nil {
		el.errorHandle <- ErrMsg{
			tp:   el.tokenType,
			from: fromBlockNum,
			to:   toBlockNum,
		}
		log.Errorf("game vault subscribe err : %+v, from : %d, to : %d", err, fromBlockNum.Int64(), toBlockNum.Int64())
		return err
	}
	for _, logEvent := range sub {
		switch logEvent.Topics[0].String() {
		case EventSignHash(WITHRAWALTOPIC):
			msg := ErrMsg{
				tp:   el.tokenType,
				from: big.NewInt(int64(logEvent.BlockNumber)),
				to:   big.NewInt(int64(logEvent.BlockNumber)),
			}
			input, err := el.abi.Events["Withdraw"].Inputs.Unpack(logEvent.Data)
			if err != nil {
				log.Error("game vault data unpack err : ", err)
				el.errorHandle <- msg
				break
			}
			if input[0].(common.Address).String() != emptyAddress {
				break
			}
			fromAddr := input[1].(common.Address).String()
			toAddr := input[2].(common.Address).String()
			accept, txType := el.Accept(fromAddr, toAddr)
			if !accept {
				break
			}
			recp, err := el.ec.TransactionReceipt(context.Background(), logEvent.TxHash)
			if err != nil {
				el.errorHandle <- msg
				log.Errorf("query txReceipt txHash : %s, err : %+v", logEvent.TxHash, err)
				break
			}
			block, err := el.ec.BlockByNumber(context.Background(), big.NewInt(int64(logEvent.BlockNumber)))
			if err != nil {
				el.errorHandle <- msg
				log.Errorf("query BlockByNumber blockNum : %d, err : %+v", logEvent.BlockNumber, err)
				break
			}
			el.erc20Notify <- ERC20Tx{
				From:    fromAddr,
				To:      toAddr,
				TxType:  txType,
				TxHash:  logEvent.TxHash.Hex(),
				Status:  recp.Status,
				PayTime: int64(block.Time() * 1000),
				Amount:  input[3].(*big.Int).String(),
			}
		}
	}
	return err
}
