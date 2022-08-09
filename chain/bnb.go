package chain

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/go-redis/redis"
	"math/big"
	"strings"
	"time"
)

type BNBTarget struct {
	txAddress string
}

func newBNBTarget(address string) *BNBTarget {
	return &BNBTarget{
		txAddress: address,
	}
}

func (t *BNBTarget) Accept(fromAddr, toAddr string) (bool, uint64) {
	if strings.ToLower(t.txAddress) == strings.ToLower(toAddr) {
		return true, BNB_RECHARGE
	}

	return false, NOT_EXIST
}

type BNBListener struct {
	TxFilter
	erc20Notify chan ERC20Tx
	ec          *ethclient.Client
	rc          *redis.Client
	chainId     *big.Int
}

func newBNBListener(filter TxFilter, ec *ethclient.Client, rc *redis.Client, erc20Notify chan ERC20Tx) *BNBListener {
	chainId, err := ec.NetworkID(context.Background())
	if err != nil {
		log.Error("query network id err : ", err)
		return nil
	}
	return &BNBListener{
		filter,
		erc20Notify,
		ec,
		rc,
		chainId,
	}
}

func (bl *BNBListener) run() {
	go bl.NewBlockFilter()
}

func (bl *BNBListener) NewBlockFilter() error {
	newBlockChan := make(chan *types.Header)
	ethClient := bl.ec
	sub, err := ethClient.SubscribeNewHead(context.Background(), newBlockChan)
	if err != nil {
		log.Error("bnb subscribe new head err : ", err)
		return err
	}
	for {
		select {
		case err = <-sub.Err():
			sub = event.Resubscribe(time.Millisecond, func(ctx context.Context) (event.Subscription, error) {
				return ethClient.SubscribeNewHead(context.Background(), newBlockChan)
			})
			log.Error("bnb subscribe err : ", err)
		case header := <-newBlockChan:
			height := new(big.Int).Sub(header.Number, big.NewInt(blockConfirmHeight))
			eb.Publish(newBlockTopic, height)
			log.Infof("header num : %d, height : %d", header.Number.Int64(), height.Int64())
			block, err := ethClient.BlockByNumber(context.Background(), height)
			if err != nil {
				log.Errorf("bnb blockByHash err : %+v", err)
				break
			}
			bl.SingleBlockFilter(block)
			log.Infof("bnb listen new block %d finished", block.Number())
			bl.rc.Set(BLOCK_NUM, height.Int64(), 0)
		}
	}
}

func (bl *BNBListener) handlePastBlock(fromBlock, toBlock uint64) {
	go bl.PastBlockFilter(fromBlock, toBlock)
}

func (bl *BNBListener) PastBlockFilter(blockNum, nowBlockNum uint64) error {
	for i := blockNum; i < nowBlockNum; i++ {
		log.Infof("bnb past block num : %d", i)
		//go func(num uint64) {
		//	block, err := bl.ec.BlockByNumber(context.Background(), big.NewInt(int64(num)))
		//	if err != nil {
		//		log.Error("blockByNumber err : ", err)
		//		return
		//	}
		//	bl.SingleBlockFilter(block)
		//}(i)
		block, err := bl.ec.BlockByNumber(context.Background(), big.NewInt(int64(i)))
		if err != nil {
			log.Error("blockByNumber err : ", err)
			break
		}
		bl.SingleBlockFilter(block)
	}
	return nil
}

func (bl *BNBListener) SingleBlockFilter(block *types.Block) error {
	log.Infof("bnb height : %d , tx num :  %d", block.Number(), len(block.Transactions()))
	for _, tx := range block.Transactions() {
		//log.Infof("bnb tx : %s", tx.Hash())
		var fromAddr string
		if msg, err := tx.AsMessage(types.NewEIP155Signer(bl.chainId), nil); err == nil {
			fromAddr = msg.From().Hex()
		}
		if tx.To() == nil {
			continue
		}
		if tx.Value().Int64() == 0 {
			continue
		}
		accept, txType := bl.Accept(fromAddr, tx.To().Hex())
		if !accept {
			continue
		}
		var status uint64
		recp, err := bl.ec.TransactionReceipt(context.Background(), tx.Hash())
		status = recp.Status
		if err != nil {
			log.Error("bnb TransactionReceipt err : ", err)
			status = 0
		}
		tx := ERC20Tx{
			From:    fromAddr,
			To:      tx.To().Hex(),
			TxType:  txType,
			TxHash:  tx.Hash().Hex(),
			Status:  status,
			PayTime: int64(block.Time() * 1000),
			Amount:  tx.Value().String(),
		}
		bl.erc20Notify <- tx
	}
	return nil
}
