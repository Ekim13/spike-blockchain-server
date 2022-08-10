package chain

import (
	"github.com/google/uuid"
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
		if !m.counter.Allow() {
			continue
		}
		t := (*m.nlq)[i]
		m.nlq.Remove(i)
		go func(task *NftListReq) {
			m.queryNftListByMoralis(task.uuid, task.walletAddr, task.network)
		}(t)
	}
}

func (m *NftListManager) WaitCall(uuid uuid.UUID) (interface{}, error) {
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
	}
}

func (m *NftListManager) queryNftListByMoralis(uuid uuid.UUID, walletAddr, network string) {
	var res []NftResult
	res, err := QueryWalletNft("", walletAddr, network, res)
	m.workLk.Lock()
	defer m.workLk.Unlock()
	m.callRes[uuid] <- result{
		r:   res,
		err: err,
	}
}

type NativeTxManager struct {
	counter Counter
	ntq     *nativeTxQueue
	listLk  sync.Mutex
	workLk  sync.Mutex
	callRes map[uuid.UUID]chan result
	notify  chan struct{}
}

func NewNativeTxManager() *NativeTxManager {
	var counter Counter
	counter.Set(6, time.Second)
	m := &NativeTxManager{
		counter: counter,
		ntq:     &nativeTxQueue{},
		callRes: map[uuid.UUID]chan result{},
		notify:  make(chan struct{}),
	}
	go m.RunSched()
	return m
}

func (m *NativeTxManager) QueryNativeTxRecord(uuid uuid.UUID, walletAddr string, blockNum uint64) {
	m.listLk.Lock()
	defer m.listLk.Unlock()
	m.ntq.Push(&NativeTxReq{
		uuid:       uuid,
		walletAddr: walletAddr,
		blockNum:   blockNum,
	})
	m.notify <- struct{}{}
}

func (m *NativeTxManager) RunSched() {
	ticker := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case <-ticker.C:

		case <-m.notify:

		}
		m.handle()
	}
}

func (m *NativeTxManager) handle() {
	m.listLk.Lock()
	defer m.listLk.Unlock()
	queueLen := m.ntq.Len()
	for i := 0; i < queueLen; i++ {
		if !m.counter.Allow() {
			continue
		}
		t := (*m.ntq)[i]
		m.ntq.Remove(i)
		go func(task *NativeTxReq) {
			m.queryNativeTxRecordByBscScan(task.uuid, task.walletAddr, task.blockNum)
		}(t)
	}
}

func (m *NativeTxManager) WaitCall(uuid uuid.UUID) (interface{}, error) {
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
	}
}

func (m *NativeTxManager) queryNativeTxRecordByBscScan(uuid uuid.UUID, walletAddr string, blockNum uint64) {
	res, err := queryNativeTxRecord(walletAddr, blockNum)
	m.workLk.Lock()
	defer m.workLk.Unlock()
	m.callRes[uuid] <- result{
		r:   res,
		err: err,
	}
}

type ERC20TxManager struct {
	counter Counter
	etq     *erc20TxQueue
	listLk  sync.Mutex
	workLk  sync.Mutex
	callRes map[uuid.UUID]chan result
	notify  chan struct{}
}

func NewERC20TxManager() *ERC20TxManager {
	var counter Counter
	counter.Set(6, time.Second)
	m := &ERC20TxManager{
		counter: counter,
		etq:     &erc20TxQueue{},
		callRes: map[uuid.UUID]chan result{},
		notify:  make(chan struct{}),
	}
	go m.RunSched()
	return m
}

func (m *ERC20TxManager) QueryERC20TxRecord(uuid uuid.UUID, contractAddr, walletAddr string, blockNum uint64) {
	m.listLk.Lock()
	defer m.listLk.Unlock()
	m.etq.Push(&ERC20TxReq{
		uuid:         uuid,
		walletAddr:   walletAddr,
		contractAddr: contractAddr,
		blockNum:     blockNum,
	})
	m.notify <- struct{}{}
}

func (m *ERC20TxManager) RunSched() {
	ticker := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case <-ticker.C:

		case <-m.notify:

		}
		m.handle()
	}
}

func (m *ERC20TxManager) handle() {
	m.listLk.Lock()
	defer m.listLk.Unlock()
	queueLen := m.etq.Len()
	for i := 0; i < queueLen; i++ {
		if !m.counter.Allow() {
			continue
		}
		t := (*m.etq)[i]
		m.etq.Remove(i)
		go func(task *ERC20TxReq) {
			m.queryErc20TxRecordByBscScan(task.uuid, task.contractAddr, task.walletAddr, task.blockNum)
		}(t)
	}
}

func (m *ERC20TxManager) WaitCall(uuid uuid.UUID) (interface{}, error) {
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
	}
}

func (m *ERC20TxManager) queryErc20TxRecordByBscScan(uuid uuid.UUID, contractAddr, walletAddr string, blockNum uint64) {
	res, err := queryERC20TxRecord(contractAddr, walletAddr, blockNum)
	m.workLk.Lock()
	defer m.workLk.Unlock()
	m.callRes[uuid] <- result{
		r:   res,
		err: err,
	}
}
