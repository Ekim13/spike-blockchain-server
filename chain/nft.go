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

const emptyAddress = "0x0000000000000000000000000000000000000000"

type AUNFTTarget struct {
	txAddress string
}

func newAUNFTTarget(address string) *AUNFTTarget {
	return &AUNFTTarget{
		txAddress: address,
	}
}

func (t *AUNFTTarget) Accept(fromAddr, toAddr string) (bool, uint64) {
	if strings.ToLower(emptyAddress) == strings.ToLower(fromAddr) {
		return true, AUNFT_TRANSFER
	}

	if strings.ToLower(t.txAddress) == strings.ToLower(toAddr) {
		return true, AUNFT_IMPORT
	}
	return true, AUNFT_TRANSFER
}

type AUNFTListener struct {
	TxFilter
	contractAddr   string
	tokenType      TokenType
	erc721Notify   chan ERC721Tx
	newBlockNotify DataChannel
	ec             *ethclient.Client
	rc             *redis.Client
	abi            abi.ABI
	errorHandle    chan ErrMsg
}

func newAUNFTListener(filter TxFilter, contractAddr string, tokenType TokenType, ec *ethclient.Client, rc *redis.Client, erc721Notify chan ERC721Tx, newBlockNotify DataChannel, abi abi.ABI, errorHandle chan ErrMsg) *AUNFTListener {
	return &AUNFTListener{
		filter,
		contractAddr,
		tokenType,
		erc721Notify,
		newBlockNotify,
		ec,
		rc,
		abi,
		errorHandle,
	}
}

func (al *AUNFTListener) run() {
	go al.NewEventFilter()
}

func (al *AUNFTListener) NewEventFilter() error {
	for {
		select {
		case de := <-al.newBlockNotify:
			height := de.Data.(*big.Int)
			al.handlePastBlock(height, height)
		}
	}
}

func (al *AUNFTListener) handlePastBlock(fromBlockNum, toBlockNum *big.Int) error {
	log.Infof("nft past event filter, fromBlock : %d, toBlock : %d ", fromBlockNum, toBlockNum)
	ethClient := al.ec
	contractAddress := common.HexToAddress(al.contractAddr)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		FromBlock: fromBlockNum,
		ToBlock:   toBlockNum,
	}

	sub, err := ethClient.FilterLogs(context.Background(), query)
	if err != nil {
		al.errorHandle <- ErrMsg{
			tp:   al.tokenType,
			from: fromBlockNum,
			to:   toBlockNum,
		}
		log.Errorf("nft subscribe event log, from: %d,to: %d,err : %+v", fromBlockNum.Int64(), toBlockNum.Int64(), err)
		return err
	}
	for _, l := range sub {
		switch l.Topics[0].String() {
		case EventSignHash(TransferTopic):
			msg := ErrMsg{
				tp:   al.tokenType,
				from: big.NewInt(int64(l.BlockNumber)),
				to:   big.NewInt(int64(l.BlockNumber)),
			}
			recp, err := al.ec.TransactionReceipt(context.Background(), l.TxHash)
			if err != nil {
				al.errorHandle <- msg
				log.Error("nft TransactionReceipt err : ", err)
				break
			}
			block, err := al.ec.BlockByNumber(context.Background(), big.NewInt(int64(l.BlockNumber)))
			if err != nil {
				al.errorHandle <- msg
				log.Errorf("query BlockByNumber blockNum : %d, err : %+v", l.BlockNumber, err)
				break
			}

			fromAddr := common.HexToAddress(l.Topics[1].Hex()).String()
			toAddr := common.HexToAddress(l.Topics[2].Hex()).String()
			_, txType := al.Accept(fromAddr, toAddr)
			al.rc.Del(fromAddr + nftTypeSuffix)
			al.rc.Del(toAddr + nftTypeSuffix)
			al.erc721Notify <- ERC721Tx{
				From:    fromAddr,
				To:      toAddr,
				TxType:  txType,
				TxHash:  l.TxHash.Hex(),
				Status:  recp.Status,
				PayTime: int64(block.Time() * 1000),
				TokenId: l.Topics[3].Big().Uint64(),
			}
		}
	}
	return nil
}
