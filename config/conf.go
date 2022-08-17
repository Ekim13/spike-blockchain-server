package config

import (
	"github.com/BurntSushi/toml"
	logger "github.com/ipfs/go-log"
	"os"
	"spike-blockchain-server/cache"
)

var log = logger.Logger("config")

var Cfg Config

func Init() {

	path, ok := os.LookupEnv("CONFIG_PATH")
	if !ok {
		panic("config path is not assign")
		return
	}
	file, err := os.Open(path)
	switch {
	case os.IsNotExist(err):
		panic("config is not exist")
		return
	case err != nil:
		panic("config path error")
		return
	}
	_, err = toml.NewDecoder(file).Decode(&Cfg)
	if err != nil {
		panic("init config err")
		return
	}
	log.Infof("cfg : %+v", Cfg)
	cache.Redis(Cfg.Redis.Address, Cfg.Redis.Password)
}

type Config struct {
	Moralis  Moralis  `toml:"moralis"`
	BscScan  BscScan  `toml:"bscscan"`
	Redis    Redis    `toml:"redis"`
	Kafka    Kafka    `toml:"kafka"`
	Contract Contract `toml:"contract"`
	Chain    Chain    `toml:"chain"`
}

type Chain struct {
	NodeAddress string `toml:"node_address""`
}

type Moralis struct {
	XApiKey string `toml:"x_api_key"`
}

type BscScan struct {
	ApiKey    string `toml:"api_key"`
	UrlPrefix string `toml:"url_prefix"`
}

type Redis struct {
	Address   string `toml:"address"`
	Password  string `toml:"password"`
	MachineId string `toml:"machine_id"`
}

type Kafka struct {
	Address string `toml:"address"`
}

type Contract struct {
	GameNftAddress         string `toml:"game_nft_address"`
	GovernanceTokenAddress string `toml:"governance_token_address"`
	GameTokenAddress       string `toml:"game_token_address"`
	GameVaultAddress       string `toml:"game_vault_address"`
	UsdcAddress            string `toml:"usdc_address"`
}
