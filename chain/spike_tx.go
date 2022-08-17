package chain

import (
	"encoding/json"
	"spike-blockchain-server/game"
)

type ERC20Tx struct {
	From    string `json:"from"`
	To      string `json:"to"`
	TxType  uint64 `json:"txType"`
	TxHash  string `json:"txHash"`
	Status  uint64 `json:"status"`
	PayTime int64  `json:"payTime"`
	Amount  string `json:"amount"`
}

type ERC721Tx struct {
	From    string `json:"from"`
	To      string `json:"to"`
	TxType  uint64 `json:"txType"`
	TxHash  string `json:"txHash"`
	Status  uint64 `json:"status"`
	PayTime int64  `json:"payTime"`
	TokenId uint64 `json:"tokenId"`
}

type SpikeTxMgr struct {
	erc20Notify  chan ERC20Tx
	erc721Notify chan ERC721Tx
	close        chan struct{}
	mqApi        game.MqApi
}

func newSpikeTxMgr(client *game.KafkaClient, erc20Notify chan ERC20Tx, erc721Notify chan ERC721Tx) *SpikeTxMgr {
	s := &SpikeTxMgr{
		erc20Notify:  erc20Notify,
		erc721Notify: erc721Notify,
		mqApi:        client,
	}

	return s
}

func (s *SpikeTxMgr) run() {
	for {
		select {
		case erc20Tx := <-s.erc20Notify:
			txByte, err := json.Marshal(erc20Tx)
			if err != nil {
				log.Errorf("json marshal err : %+v, txByte : %s", err, txByte)
				break
			}
			log.Infof("erc20 value : %s", string(txByte))
			var topic string
			if checkRecharge(int(erc20Tx.TxType)) {
				topic = game.RECHARGETXTOPIC
			} else {
				topic = game.ERC20TXTOPIC
			}
			err = s.mqApi.SendMessage(game.Msg{
				Topic: topic,
				Key:   erc20Tx.TxHash,
				Value: string(txByte),
			})
			if err != nil {
				log.Error("erc20 tx produce err : ", err)
			}
		case erc721Tx := <-s.erc721Notify:
			txByte, err := json.Marshal(erc721Tx)
			log.Infof("value : %s", string(txByte))
			if err != nil {
				log.Error(err)
				break
			}
			var topic string
			if checkImport(int(erc721Tx.TxType)) {
				topic = game.IMPORTNFTTOPIC
			} else {
				topic = game.ERC721TXTOPIC
			}
			err = s.mqApi.SendMessage(game.Msg{
				Topic: topic,
				Key:   erc721Tx.TxHash,
				Value: string(txByte),
			})
			if err != nil {
				log.Error("erc721 tx produce err : ", err)
			}
		case <-s.close:
			//log
			return
		}
	}
}
