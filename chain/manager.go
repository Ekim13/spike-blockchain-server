package chain

import (
	"context"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
	"sync"
	"time"
)

type result struct {
	r   interface{}
	err error
}

type NftListManager struct {
	counter Counter
	nlq     *nftListQueue
	listLk  sync.Mutex
	workLk  sync.Mutex
	callRes map[uuid.UUID]chan result
	notify  chan struct{}
}

func NewNftListManager() *NftListManager {
	var counter Counter
	counter.Set(12, time.Second)
	m := &NftListManager{
		counter: counter,
		nlq:     &nftListQueue{},
		callRes: map[uuid.UUID]chan result{},
		notify:  make(chan struct{}),
	}
	go m.RunSched()
	return m
}

func (m *NftListManager) QueryNftList(uuid uuid.UUID, walletAddr, network string) {
	m.listLk.Lock()
	defer m.listLk.Unlock()
	m.nlq.Push(&NftListReq{
		uuid:       uuid,
		walletAddr: walletAddr,
		network:    network,
	})
	m.notify <- struct{}{}
}

func (m *NftListManager) RunSched() {
	ticker := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case <-ticker.C:

		case <-m.notify:

		}
		m.handle()
	}
}

func (m *NftListManager) handle() {
	m.listLk.Lock()
	defer m.listLk.Unlock()

	queueLen := m.nlq.Len()
	for i := 0; i < queueLen; i++ {
		if !m.counter.Allow(1) {
			continue
		}
		t := (*m.nlq)[i]
		m.nlq.Remove(i)
		go func(task *NftListReq) {
			m.queryNftListByMoralis(task.uuid, task.walletAddr, task.network)
		}(t)
	}
}

func (m *NftListManager) WaitCall(ctx context.Context, uuid uuid.UUID) (interface{}, error) {
	defer func() {
		m.workLk.Lock()
		defer m.workLk.Unlock()
		delete(m.callRes, uuid)
	}()
	m.workLk.Lock()
	ch, ok := m.callRes[uuid]
	if !ok {
		ch = make(chan result, 1)
		m.callRes[uuid] = ch
	}
	m.workLk.Unlock()
	select {
	case res := <-ch:
		return res.r, res.err
	case <-ctx.Done():
		return nil, xerrors.New("query nft list time out")
	}
}

func (m *NftListManager) queryNftListByMoralis(uuid uuid.UUID, walletAddr, network string) {
	var res []NftResult
	res, err := QueryWalletNft("", walletAddr, network, res)
	m.workLk.Lock()
	defer m.workLk.Unlock()
	ch, ok := m.callRes[uuid]
	if ok {
		ch <- result{
			r:   res,
			err: err,
		}
	}
}

func (m *TxRecordManager) queryNativeTxRecordByBscScan(uuid uuid.UUID, walletAddr string, blockNum uint64) {
	res, err := queryNativeTxRecord(walletAddr, blockNum)
	m.workLk.Lock()
	defer m.workLk.Unlock()
	ch, ok := m.callRes[uuid]
	if ok {
		ch <- result{
			r:   res,
			err: err,
		}
	}
}

type TxRecordManager struct {
	counter Counter
	trq     *txRecordQueue
	listLk  sync.Mutex
	workLk  sync.Mutex
	callRes map[uuid.UUID]chan result
	notify  chan struct{}
}

func NewTxRecordManager() *TxRecordManager {
	var counter Counter
	counter.Set(12, time.Second)
	m := &TxRecordManager{
		counter: counter,
		trq:     &txRecordQueue{},
		callRes: map[uuid.UUID]chan result{},
		notify:  make(chan struct{}),
	}
	go m.RunSched()
	return m
}

func (m *TxRecordManager) QueryTxRecord(uuid uuid.UUID, contractAddr, walletAddr string, blockNum uint64) {
	m.listLk.Lock()
	defer m.listLk.Unlock()
	m.trq.Push(&TxRecordReq{
		uuid:         uuid,
		walletAddr:   walletAddr,
		contractAddr: contractAddr,
		blockNum:     blockNum,
	})
	m.notify <- struct{}{}
}

func (m *TxRecordManager) RunSched() {
	ticker := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case <-ticker.C:

		case <-m.notify:

		}
		m.handle()
	}
}

func (m *TxRecordManager) handle() {
	m.listLk.Lock()
	defer m.listLk.Unlock()
	queueLen := m.trq.Len()
	for i := 0; i < queueLen; i++ {
		weight := 1
		t := (*m.trq)[i]
		if t.contractAddr == "" {
			weight = 2
		}
		if !m.counter.Allow(weight) {
			continue
		}
		m.trq.Remove(i)
		go func(task *TxRecordReq) {
			if task.contractAddr == "" {
				m.queryNativeTxRecordByBscScan(task.uuid, task.walletAddr, task.blockNum)
				return
			}
			m.queryErc20TxRecordByBscScan(task.uuid, task.contractAddr, task.walletAddr, task.blockNum)
		}(t)
	}
}

func (m *TxRecordManager) WaitCall(ctx context.Context, uuid uuid.UUID) (interface{}, error) {
	defer func() {
		m.workLk.Lock()
		defer m.workLk.Unlock()
		delete(m.callRes, uuid)
	}()
	m.workLk.Lock()
	ch, ok := m.callRes[uuid]
	if !ok {
		ch = make(chan result, 1)
		m.callRes[uuid] = ch
	}
	m.workLk.Unlock()

	select {
	case res := <-ch:
		return res.r, res.err
	case <-ctx.Done():
		return nil, xerrors.New("query tx record time out")
	}
}

func (m *TxRecordManager) queryErc20TxRecordByBscScan(uuid uuid.UUID, contractAddr, walletAddr string, blockNum uint64) {
	res, err := queryERC20TxRecord(contractAddr, walletAddr, blockNum)
	m.workLk.Lock()
	defer m.workLk.Unlock()
	ch, ok := m.callRes[uuid]
	if ok {
		ch <- result{
			r:   res,
			err: err,
		}
	}
}
