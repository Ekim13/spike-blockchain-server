package chain

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
	"sort"
	"spike-blockchain-server/config"
	"spike-blockchain-server/serializer"
	"strconv"
	"strings"
	"time"
)

const queryTxRecordTimeout = 10 * time.Second

const BscScanRateLimit = "\"Max rate limit reached\""

type NativeTransactionRecordService struct {
	Address string `form:"address" json:"address" binding:"required"`
}

type ERC20TransactionRecordService struct {
	Address         string `form:"address" json:"address" binding:"required"`
	ContractAddress string `form:"contract_address" json:"contract_address" binding:"required"`
}

type Result struct {
	Hash        string `json:"hash"`
	TimeStamp   string `json:"timeStamp"`
	BlockNumber string `json:"blockNumber"`
	BlockHash   string `json:"blockHash"`
	From        string `json:"from"`
	To          string `json:"to"`
	Value       string `json:"value"`
	Input       string `json:"input"`
	Type        string `json:"type"`
}

type BscRes struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Result  []Result `json:"result"`
}

func (bl *BscListener) NativeTxRecord(c *gin.Context) {
	var service NativeTransactionRecordService
	if err := c.ShouldBind(&service); err == nil {
		res := bl.findFindNativeTxRecord(service.Address)
		c.JSON(200, res)
	} else {
		c.JSON(500, serializer.Response{
			Code: 500,
			Msg:  ErrorParam.Error(),
		})
	}
}

func (bl *BscListener) ERC20TxRecord(c *gin.Context) {
	var service ERC20TransactionRecordService
	if err := c.ShouldBind(&service); err == nil {
		res := bl.findFindERC20TxRecord(service.Address, service.ContractAddress)
		c.JSON(200, res)
	} else {
		c.JSON(500, serializer.Response{
			Code: 500,
			Msg:  ErrorParam.Error(),
		})
	}
}

func (bl *BscListener) findFindERC20TxRecord(address, contractAddr string) serializer.Response {
	if record := bl.GetJson(address + contractAddr + erc20TxRecordSuffix); record != "" {
		var bscRes BscRes
		bnbRecord := make([]Result, 0)
		bscRes.Result = bnbRecord
		err := json.Unmarshal([]byte(record), &bscRes)
		if err != nil {
			return serializer.Response{
				Code: 500,
				Msg:  err.Error(),
			}
		}
		return serializer.Response{
			Code: 200,
			Data: bscRes,
		}
	}
	bscRes, err := bl.FindFindERC20TxRecord(contractAddr, address)
	if err != nil {
		return serializer.Response{
			Code: 500,
			Msg:  err.Error(),
		}
	}

	return serializer.Response{
		Code: 200,
		Data: bscRes,
	}
}

func (bl *BscListener) GetBlockNum() uint64 {
	if bl.rc.Get(BLOCKNUM+config.Cfg.Redis.MachineId).Err() == redis.Nil {
		log.Infof("blockNum is not exist")
		blockNum, err := bl.ec.BlockNumber(context.Background())
		if err != nil {
			return 0
		}
		return blockNum
	} else {
		blockNum, _ := bl.rc.Get(BLOCKNUM + config.Cfg.Redis.MachineId).Uint64()
		return blockNum
	}
}

func queryNativeTxRecord(address string, blockNum uint64) (BscRes, error) {
	bscRes := BscRes{Result: make([]Result, 0)}
	bscInternalRes := BscRes{Result: make([]Result, 0)}
	client := resty.New()
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		Get(getNativeUrl(blockNum, address))
	if err != nil {
		return bscRes, err
	}
	err = json.Unmarshal(resp.Body(), &bscRes)
	if err != nil {
		return bscRes, xerrors.New(BscScanRateLimit)
	}

	resp, err = client.R().
		SetHeader("Accept", "application/json").
		Get(getNativeInternalUrl(blockNum, address))
	if err != nil {
		return bscRes, err
	}
	err = json.Unmarshal(resp.Body(), &bscInternalRes)
	if err != nil {
		return bscRes, xerrors.New(BscScanRateLimit)
	}
	bscRes.Result = append(bscRes.Result, bscInternalRes.Result...)
	return bscRes, nil
}

func (bl *BscListener) findFindNativeTxRecord(address string) serializer.Response {
	if record := bl.GetJson(address + nativeTxRecordSuffix); record != "" {
		var bscRes BscRes
		bnbRecord := make([]Result, 0)
		bscRes.Result = bnbRecord
		err := json.Unmarshal([]byte(record), &bscRes)
		if err != nil {
			return serializer.Response{
				Code: 500,
				Msg:  err.Error(),
			}
		}
		return serializer.Response{
			Code: 200,
			Data: bscRes,
		}
	}
	bscRes, err := bl.FindNativeTransactionRecord(address)
	if err != nil {
		return serializer.Response{
			Code: 500,
			Msg:  err.Error(),
		}
	}

	return serializer.Response{
		Code: 200,
		Data: bscRes,
	}
}

func (bl *BscListener) FindNativeTransactionRecord(address string) (BscRes, error) {
	blockNum := bl.GetBlockNum()
	uuid := uuid.New()
	ctx, cancel := context.WithTimeout(context.TODO(), queryTxRecordTimeout)
	bl.trManager.QueryTxRecord(uuid, "", address, blockNum)
	res, err := bl.trManager.WaitCall(ctx, uuid)
	cancel()
	if err != nil {
		log.Errorf("QueryNativeTxRecord  err :%+v", err)
		return BscRes{}, err
	}
	bscRes := res.(BscRes)
	bnbRecord := make([]Result, 0)

	if len(bscRes.Result) == 0 {
		bscRes.Result = make([]Result, 0)
		cacheData, _ := json.Marshal(bscRes)
		bl.rc.Set(address+nativeTxRecordSuffix, string(cacheData), txRecordDuration)
		return bscRes, nil
	}

	for i := range bscRes.Result {
		if bscRes.Result[i].Input == "0x" {
			bnbRecord = append(bnbRecord, bscRes.Result[i])
			continue
		}
		if strings.ToLower(bscRes.Result[i].From) == strings.ToLower(config.Cfg.Contract.GameVaultAddress) {
			bnbRecord = append(bnbRecord, bscRes.Result[i])
			continue
		}
	}
	sort.Slice(bnbRecord, func(i, j int) bool {
		time1, _ := strconv.Atoi(bnbRecord[i].TimeStamp)
		time2, _ := strconv.Atoi(bnbRecord[j].TimeStamp)
		return time1 > time2
	})
	bscRes.Result = bnbRecord
	cacheData, _ := json.Marshal(bscRes)
	bl.rc.Set(address+nativeTxRecordSuffix, string(cacheData), txRecordDuration)
	return bscRes, nil
}

func (bl *BscListener) FindFindERC20TxRecord(contractAddr, address string) (BscRes, error) {
	blockNum := bl.GetBlockNum()
	uuid := uuid.New()
	ctx, cancel := context.WithTimeout(context.TODO(), queryTxRecordTimeout)
	bl.trManager.QueryTxRecord(uuid, contractAddr, address, blockNum)
	res, err := bl.trManager.WaitCall(ctx, uuid)
	cancel()
	if err != nil {
		log.Errorf("QueryErc20TxRecord  err :%+v", err)
		return BscRes{}, err
	}
	bscRes := res.(BscRes)
	if len(bscRes.Result) == 0 {
		bscRes.Result = make([]Result, 0)
		cacheData, _ := json.Marshal(bscRes)
		bl.rc.Set(address+contractAddr+erc20TxRecordSuffix, string(cacheData), txRecordDuration)
		return bscRes, nil
	}
	cacheData, _ := json.Marshal(bscRes)
	bl.rc.Set(address+contractAddr+erc20TxRecordSuffix, string(cacheData), txRecordDuration)
	return bscRes, nil
}

func getNativeUrl(blockNumber uint64, address string) string {
	return fmt.Sprintf("%s?module=account&action=txlist&address=%s&startblock=%d&endblock=%d&offset=10000&page=1&sort=desc&apikey=%s", config.Cfg.BscScan.UrlPrefix, address, blockNumber-201600, blockNumber, config.Cfg.BscScan.ApiKey)
}

func getNativeInternalUrl(blockNumber uint64, address string) string {
	return fmt.Sprintf("%s?module=account&action=txlistinternal&address=%s&startblock=%d&endblock=%d&offset=10000&page=1&sort=desc&apikey=%s", config.Cfg.BscScan.UrlPrefix, address, blockNumber-201600, blockNumber, config.Cfg.BscScan.ApiKey)
}

func getERC20url(contractAddr, addr string, blockNumber uint64) string {
	return fmt.Sprintf("%s?module=account&action=tokentx&address=%s&startblock=%d&endblock=%d&offset=10000&page=1&sort=desc&apikey=%s&contractaddress=%s", config.Cfg.BscScan.UrlPrefix, addr, blockNumber-201600, blockNumber, config.Cfg.BscScan.ApiKey, contractAddr)
}

func queryERC20TxRecord(contractAddr, address string, blockNum uint64) (BscRes, error) {
	bscRes := BscRes{Result: make([]Result, 0)}
	client := resty.New()
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		Get(getERC20url(contractAddr, address, blockNum))
	if err != nil {
		return bscRes, err
	}
	err = json.Unmarshal(resp.Body(), &bscRes)

	if err != nil {
		return bscRes, xerrors.New(BscScanRateLimit)
	}
	return bscRes, nil
}
