package chain

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis"
	logger "github.com/ipfs/go-log"
	"math/big"
	"spike-blockchain-server/cache"
	"spike-blockchain-server/config"
	"spike-blockchain-server/game"
	"sync"
	"time"
)

var log = logger.Logger("chain")

const (
	BNB_BLOCKNUM        = "bnb_blockNum"
	GAME_VAULT_BLOCKNUM = "vault_blockNum"
	SKK_BLOCKNUM        = "skk_blockNum"
	SKS_BLOCKNUM        = "sks_blockNum"
	AUNFT_BLOCKNUM      = "aunft_blockNum"
	BLOCKNUM            = "blockNum"
)

type ErrMsg struct {
	tp   TokenType
	from *big.Int
	to   *big.Int
}

type BscListener struct {
	network     string
	nlManager   *NftListManager
	trManager   *TxRecordManager
	ec          *ethclient.Client
	rc          *redis.Client
	l           map[TokenType]Listener
	errorHandle chan ErrMsg
}

func NewBscListener(speedyNodeAddress string, targetWalletAddr string) (*BscListener, error) {
	log.Infof("bsc listener start")
	bl := &BscListener{}
	bl.nlManager = NewNftListManager()
	bl.trManager = NewTxRecordManager()
	client, err := ethclient.Dial(speedyNodeAddress)
	if err != nil {
		log.Error("eth client dial err : ", err)
		return nil, err
	}
	chainId, err := client.ChainID(context.Background())
	switch chainId.String() {
	case "56":
		bl.network = "bsc"
	case "97":
		bl.network = "bsc testnet"
	default:
		panic("not expected chainId")
	}

	errorHandle := make(chan ErrMsg, 10)
	bl.errorHandle = errorHandle
	bl.rc = cache.RedisClient
	bl.ec = client
	erc20Notify := make(chan ERC20Tx, 10)
	erc721Notify := make(chan ERC721Tx, 10)

	vaultChan := make(DataChannel, 10)
	//skkChan := make(DataChannel, 10)
	sksChan := make(DataChannel, 10)
	usdcChan := make(DataChannel, 10)
	aunftChan := make(DataChannel, 10)
	eb.Subscribe(newBlockTopic, vaultChan)
	//eb.Subscribe(newBlockTopic, skkChan)
	eb.Subscribe(newBlockTopic, sksChan)
	eb.Subscribe(newBlockTopic, sksChan)
	eb.Subscribe(newBlockTopic, aunftChan)

	l := make(map[TokenType]Listener)
	l[bnb] = newBNBListener(newBNBTarget(targetWalletAddr), bl.ec, bl.rc, erc20Notify, errorHandle)
	l[gameVault] = newGameVaultListener(newGameVaultTarget(targetWalletAddr), config.Cfg.Contract.GameVaultAddress, gameVault, bl.ec, bl.rc, erc20Notify, vaultChan, getABI(GameVaultABI), errorHandle)
	//l[governanceToken] = newERC20Listener(newSKKTarget(targetWalletAddr), constants.GOVERNANCE_TOKEN_ADDRESS, governanceToken, bl.ec, bl.rc, erc20Notify, skkChan, getABI(GovernanceTokenABI), errorHandle)
	l[gameToken] = newERC20Listener(newSKSTarget(targetWalletAddr), config.Cfg.Contract.GameTokenAddress, gameToken, bl.ec, bl.rc, erc20Notify, sksChan, getABI(GameTokenABI), errorHandle)
	l[usdc] = newERC20Listener(newUSDCTarget(targetWalletAddr), config.Cfg.Contract.UsdcAddress, usdc, bl.ec, bl.rc, erc20Notify, usdcChan, getABI(USDCContractABI), errorHandle)
	l[gameNft] = newAUNFTListener(newAUNFTTarget(targetWalletAddr), config.Cfg.Contract.GameNftAddress, gameNft, bl.ec, bl.rc, erc721Notify, aunftChan, getABI(GameNftABI), errorHandle)
	bl.l = l
	spikeTxMgr := newSpikeTxMgr(game.NewKafkaClient(config.Cfg.Kafka.Address), erc20Notify, erc721Notify)
	go spikeTxMgr.run()
	return bl, nil
}

func (bl *BscListener) Run() {
	go bl.handleError()
	//sync
	for {
		nowBlockNum, err := bl.ec.BlockNumber(context.Background())
		if err != nil {
			time.Sleep(500 * time.Millisecond)
			log.Error("query now bnb_blockNum err :", err)
			continue
		}
		if bl.rc.Get(BLOCKNUM+config.Cfg.Redis.MachineId).Err() == redis.Nil {
			log.Infof("blockNum is not exist")
			bl.rc.Set(BLOCKNUM+config.Cfg.Redis.MachineId, nowBlockNum-blockConfirmHeight, 0)
			break
		}
		cacheBlockNum, err := bl.rc.Get(BLOCKNUM + config.Cfg.Redis.MachineId).Uint64()
		if err != nil {
			log.Error("query cache bnb_blockNum err : ", err)
			continue
		}
		if cacheBlockNum >= nowBlockNum-blockConfirmHeight {
			log.Infof("sync done")
			break
		}
		var wg sync.WaitGroup
		for _, listener := range bl.l {
			wg.Add(1)
			go func(l Listener) {
				defer wg.Done()
				l.handlePastBlock(big.NewInt(int64(cacheBlockNum+1)), big.NewInt(int64(nowBlockNum-blockConfirmHeight)))
			}(listener)
		}
		wg.Wait()
		bl.rc.Set(BLOCKNUM+config.Cfg.Redis.MachineId, nowBlockNum-blockConfirmHeight, 0)
	}

	for _, listener := range bl.l {
		go func(l Listener) {
			l.run()
		}(listener)
	}
}

func (bl *BscListener) handleError() {
	for {
		select {
		case msg := <-bl.errorHandle:
			log.Infof("handle err ,type : %s, from : %d, to : %d", msg.tp.String(), msg.from.Int64(), msg.to.Int64())
			if _, ok := bl.l[msg.tp]; ok {
				time.Sleep(200 * time.Millisecond)
				bl.l[msg.tp].handlePastBlock(msg.from, msg.to)
			}
		}
	}
}
