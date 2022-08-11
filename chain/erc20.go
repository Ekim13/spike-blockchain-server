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

type SKSTarget struct {
	txAddress string
}

func newSKSTarget(address string) *SKSTarget {
	return &SKSTarget{
		txAddress: address,
	}
}

func (t *SKSTarget) Accept(fromAddr, toAddr string) (bool, uint64) {
	if strings.ToLower(t.txAddress) == strings.ToLower(toAddr) {
		return true, SKS_RECHARGE
	}

	if strings.ToLower(t.txAddress) == strings.ToLower(fromAddr) {
		return true, SKS_WITHDRAW
	}

	return false, NOT_EXIST
}

type SKKTarget struct {
	txAddress string
}

func newSKKTarget(address string) *SKKTarget {
	return &SKKTarget{
		txAddress: address,
	}
}

func (t *SKKTarget) Accept(fromAddr, toAddr string) (bool, uint64) {
	if strings.ToLower(t.txAddress) == strings.ToLower(toAddr) {
		return true, SKK_RECHARGE
	}

	if strings.ToLower(t.txAddress) == strings.ToLower(fromAddr) {
		return true, SKK_WITHDRAW
	}

	return false, NOT_EXIST
}

type ERC20Listener struct {
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

func newERC20Listener(filter TxFilter, contractAddr string, tokenType TokenType, ec *ethclient.Client, rc *redis.Client, erc20Notify chan ERC20Tx, newBlockNotify DataChannel, abi abi.ABI, errorHandle chan ErrMsg) *ERC20Listener {
	el := &ERC20Listener{
		filter,
		contractAddr,
		tokenType,
		erc20Notify,
		newBlockNotify,
		ec,
		rc,
		abi,
		errorHandle,
	}
	return el
}

func (el *ERC20Listener) run() {
	go el.NewEventFilter(el.contractAddr)
}

func (el *ERC20Listener) NewEventFilter(contractAddr string) error {
	for {
		select {
		case de := <-el.newBlockNotify:
			height := de.Data.(*big.Int)
			el.handlePastBlock(height, height)
		}
	}
}

func (el *ERC20Listener) handlePastBlock(fromBlockNum, toBlockNum *big.Int) error {
	log.Infof("erc20 past event filter, type : %v, fromBlock : %d, toBlock : %d ", el.tokenType.String(), fromBlockNum, toBlockNum)
	ethClient := el.ec
	contractAddress := common.HexToAddress(el.contractAddr)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		FromBlock: fromBlockNum,
		ToBlock:   toBlockNum,
	}

	sub, err := ethClient.FilterLogs(context.Background(), query)
	if err != nil {
		el.errorHandle <- ErrMsg{
			tp:   el.tokenType,
			from: fromBlockNum,
			to:   toBlockNum,
		}
		log.Errorf("erc20 subscribe err : %+v, from : %d, to : %d, type : %s", err, fromBlockNum.Int64(), toBlockNum.Int64(), el.tokenType.String())
		return err
	}
	for _, logEvent := range sub {
		switch logEvent.Topics[0].String() {
		case EventSignHash(TransferTopic):
			msg := ErrMsg{
				tp:   el.tokenType,
				from: big.NewInt(int64(logEvent.BlockNumber)),
				to:   big.NewInt(int64(logEvent.BlockNumber)),
			}

			input, err := el.abi.Events["Transfer"].Inputs.Unpack(logEvent.Data)
			if err != nil {
				log.Error("erc20 data unpack err : ", err)
				el.errorHandle <- msg
				break
			}
			fromAddr := common.HexToAddress(logEvent.Topics[1].Hex()).String()
			toAddr := common.HexToAddress(logEvent.Topics[2].Hex()).String()
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
				Amount:  input[0].(*big.Int).String(),
			}
		}
	}
	return err
}
