package chain

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/go-redis/redis"
	"math/big"
	"strings"
	"sync"
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
	errorHandle chan ErrMsg
}

func newBNBListener(filter TxFilter, ec *ethclient.Client, rc *redis.Client, erc20Notify chan ERC20Tx, errorHandle chan ErrMsg) *BNBListener {
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
		errorHandle,
	}
}

func (bl *BNBListener) run() {
	go bl.NewBlockFilter()
}

func (bl *BNBListener) NewBlockFilter() error {
	newBlockChan := make(chan *types.Header)
	sub, err := bl.ec.SubscribeNewHead(context.Background(), newBlockChan)
	if err != nil {
		log.Error("bnb subscribe new head err : ", err)
		return err
	}
	for {
		select {
		case err = <-sub.Err():
			sub = event.Resubscribe(time.Millisecond, func(ctx context.Context) (event.Subscription, error) {
				return bl.ec.SubscribeNewHead(context.Background(), newBlockChan)
			})
			log.Error("new block subscribe err : ", err)
		case header := <-newBlockChan:
			height := new(big.Int).Sub(header.Number, big.NewInt(blockConfirmHeight))
			cacheHeight, err := bl.rc.Get(BLOCK_NUM).Int64()

			if height.Int64()-1 > cacheHeight {
				for i := cacheHeight; i < height.Int64(); i++ {
					log.Infof("ws node timeout err : height %d", i)
					bl.errorHandle <- ErrMsg{
						tp:   bnb,
						from: big.NewInt(i),
						to:   big.NewInt(i),
					}
					eb.Publish(newBlockTopic, big.NewInt(i))
				}
			}
			eb.Publish(newBlockTopic, height)
			log.Infof("new block num : %d, height : %d", header.Number.Int64(), height.Int64())

			err = bl.SingleBlockFilter(height)
			if err != nil {
				bl.errorHandle <- ErrMsg{
					tp:   bnb,
					from: height,
					to:   height,
				}
				break
			}
			bl.rc.Set(BLOCK_NUM, height.Int64(), 0)
			log.Infof("bnb listen new block %d finished", height)
		}
	}
}

func (bl *BNBListener) handlePastBlock(blockNum, nowBlockNum *big.Int) error {
	throttle := make(chan struct{}, 30)
	var wg sync.WaitGroup
	wg.Add(int(new(big.Int).Sub(nowBlockNum, blockNum).Int64()) + 1)
	for i := blockNum.Int64(); i <= nowBlockNum.Int64(); i++ {
		throttle <- struct{}{}
		go func(height int64) {
			defer func() {
				wg.Done()
				<-throttle
			}()
			h := big.NewInt(height)
			err := bl.SingleBlockFilter(h)
			if err != nil {
				bl.errorHandle <- ErrMsg{
					tp:   bnb,
					from: h,
					to:   h,
				}
			}
		}(i)
	}
	wg.Wait()
	return nil
}

func (bl *BNBListener) SingleBlockFilter(height *big.Int) error {
	block, err := bl.ec.BlockByNumber(context.Background(), height)
	if err != nil {
		log.Errorf("bnb blockByHash heght : %d ,err : %+v", height.Int64(), err)
		return err
	}
	log.Infof("bnb height : %d , tx num :  %d", block.Number(), len(block.Transactions()))
	for _, tx := range block.Transactions() {
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
		recp, err := bl.ec.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Error("bnb TransactionReceipt err : ", err)
			return err
		}
		tx := ERC20Tx{
			From:    fromAddr,
			To:      tx.To().Hex(),
			TxType:  txType,
			TxHash:  tx.Hash().Hex(),
			Status:  recp.Status,
			PayTime: int64(block.Time() * 1000),
			Amount:  tx.Value().String(),
		}
		bl.erc20Notify <- tx
	}
	return nil
}
