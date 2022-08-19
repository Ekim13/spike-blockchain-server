package chain

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"golang.org/x/xerrors"
	"math"
	"math/big"
	"sort"
	"spike-blockchain-server/chain/contract"
	"spike-blockchain-server/config"
	"spike-blockchain-server/constants"
	"spike-blockchain-server/serializer"
	"strconv"
	"strings"
	"sync"
	"time"
)

const queryNftListTimeout = 20 * time.Second

var ErrorParam = errors.New("params error")
var MoralisRateLimit = "{\"message\":\"Rate limit exceeded.\"}"

const (
	txRecordDuration     = 10 * time.Minute
	nftListDuration      = 10 * time.Minute
	nftTypeSuffix        = "_nftType"
	nativeTxRecordSuffix = "_nativeTxRecord"
	erc20TxRecordSuffix  = "_erc20TxRecord"
)

type balanceShow struct {
	Symbol  string `json:"symbol"`
	Balance string `json:"balance"`
}

type NftType struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type CacheData struct {
	Type        string                 `json:"type"`
	GameId      string                 `json:"gameId"`
	BlockNumber string                 `json:"block_number"`
	TokenId     string                 `json:"token_id"`
	Image       string                 `json:"image"`
	Description string                 `json:"description"`
	SpikeInfo   SpikeInfo              `json:"spike_info"`
	Attributes  map[string]interface{} `json:"attributes"`
}
type CacheDataList struct {
	CD []CacheData `json:"cache_data"`
}

type txQueryService struct {
	TxHash string `json:"txHash" binding:"required"`
}

type balanceService struct {
	Address string `json:"address" binding:"required"`
}

type metadataService struct {
	TokenId string `json:"tokenId"`
	Address string `json:"address"`
}

type nftTypeService struct {
	WalletAddress string `json:"walletAddress" binding:"required"`
}

type nftMetadataService struct {
	WalletAddress string `json:"walletAddress"`
	Type          string `json:"type"`
	Page          int    `json:"page"`
	PageSize      int    `json:"page_size"`
}

type Metadata struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	ExternalUrl string    `json:"external_url"`
	Fid         string    `json:"fid"`
	Mid         string    `json:"mid"`
	SpikeInfo   SpikeInfo `json:"spike_info"`
	Attribute   []Attr    `json:"attributes"`
}

type SpikeInfo struct {
	Version string `json:"version"`
	Tp      string `json:"type"`
	Url     string `json:"url"`
}

type Attr struct {
	TraitType string      `json:"trait_type"`
	Value     interface{} `json:"value"`
}

type NftResults struct {
	Total    int         `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
	Cursor   string      `json:"cursor"`
	Results  []NftResult `json:"result"`
}

type NftResult struct {
	TokenId     string `json:"token_id"`
	BlockNumber string `json:"block_number"`
	TokenUri    string `json:"token_uri"`
	Metadata    string `json:"metadata"`
}

func (bl *BscListener) QueryTxIsPendingByHash(c *gin.Context) {
	var service txQueryService
	if err := c.ShouldBind(&service); err == nil {
		log.Infof("tx: %s", service.TxHash)
		res := bl.queryTxIsPendingByHash(service.TxHash)
		c.JSON(200, res)
	} else {
		c.JSON(500, err.Error())
	}
}

func (bl *BscListener) QueryTxStatusByHash(c *gin.Context) {
	var service txQueryService
	if err := c.ShouldBind(&service); err == nil {
		log.Infof("tx: %s", service.TxHash)
		res := bl.queryTxStatusByHash(service.TxHash)
		c.JSON(200, res)
	} else {
		c.JSON(500, err.Error())
	}
}

func (bl *BscListener) queryTxIsPendingByHash(txHash string) serializer.Response {
	_, isPending, err := bl.ec.TransactionByHash(context.Background(), common.HexToHash(txHash))
	code := 200
	if err != nil {
		code = 500
		return serializer.Response{
			Code:  code,
			Data:  isPending,
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Code: code,
		Data: isPending,
	}
}

func (bl *BscListener) queryTxStatusByHash(txHash string) serializer.Response {
	receipt, err := bl.ec.TransactionReceipt(context.Background(), common.HexToHash(txHash))
	code := 200
	if err != nil {
		code = 500
		return serializer.Response{
			Code:  code,
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Code: code,
		Data: receipt.Status,
	}
}

func (bl *BscListener) QueryNftMetadata(c *gin.Context) {
	var service metadataService

	if err := c.ShouldBind(&service); err == nil {
		tokenId, err := strconv.Atoi(service.TokenId)
		if err != nil {
			c.JSON(500, err.Error())
			return
		}
		res := bl.queryNftMetadata(int64(tokenId), service.Address)
		c.JSON(200, res)
	} else {
		c.JSON(500, err.Error())
	}
}

func (bl *BscListener) queryNftMetadata(tokenId int64, address string) serializer.Response {
	aunft, err := contract.NewGameNft(common.HexToAddress(config.Cfg.Contract.GameNftAddress), bl.ec)
	if err != nil {
		log.Error("new auNft err : ", err)
		return serializer.Response{
			Code:  500,
			Error: err.Error(),
		}
	}
	uri, err := aunft.TokenURI(nil, big.NewInt(tokenId))
	if err != nil {
		log.Error("query tokenUri err : ", err)
		return serializer.Response{
			Code:  500,
			Error: err.Error(),
		}
	}
	owner, err := aunft.OwnerOf(nil, big.NewInt(tokenId))
	if err != nil {
		return serializer.Response{
			Code:  500,
			Error: err.Error(),
		}
	}
	if strings.ToLower(owner.String()) != strings.ToLower(address) {
		return serializer.Response{
			Code:  500,
			Error: xerrors.New("tokenId, nft not match").Error(),
		}
	}

	client := resty.New()

	resp, err := client.R().Get(uri)
	log.Infof("query nft tokenId : %d, uri : %s", tokenId, uri)
	if err != nil {
		return serializer.Response{
			Code:  500,
			Error: err.Error(),
		}
	}

	if resp.IsError() {
		return serializer.Response{
			Code:  500,
			Error: resp.String(),
		}
	}
	var m Metadata
	err = json.Unmarshal(resp.Body(), &m)
	if err != nil {
		return serializer.Response{
			Code:  500,
			Error: err.Error(),
		}
	}
	metadata, err := json.Marshal(m)
	if err != nil {
		return serializer.Response{
			Code:  500,
			Error: err.Error(),
		}
	}
	log.Infof(string(metadata))
	return serializer.Response{
		Code: 200,
		Data: string(metadata),
	}
}

func (bl *BscListener) QueryNftListByType(c *gin.Context) {
	var service nftMetadataService

	if err := c.ShouldBind(&service); err == nil {
		if service.WalletAddress == "" || service.Type == "" {
			c.JSON(500, xerrors.New("param can not be null").Error())
			return
		}
		res := bl.queryNftListByType(service.WalletAddress, bl.network, service.Type, int64(service.Page), int64(service.PageSize))
		c.JSON(200, res)
	} else {
		c.JSON(500, serializer.Response{
			Code: 500,
			Msg:  ErrorParam.Error(),
		})
	}
}

func (bl *BscListener) queryNftListByType(addr, network, tp string, page, pageSize int64) serializer.Response {
	if result := bl.GetJson(addr + tp); result == "" {
		nftType, err := bl.queryNftFromMoralis(addr, network)
		if err != nil {
			return serializer.Response{
				Code: 500,
				Msg:  err.Error(),
			}
		}
		if len(nftType) == 0 {
			return serializer.Response{
				Code: 200,
				Data: nftType,
			}
		}
	}
	nftjson := bl.GetJson(addr + tp)
	var cdList CacheDataList
	err := json.Unmarshal([]byte(nftjson), &cdList)
	if err != nil {
		return serializer.Response{
			Code: 500,
			Msg:  err.Error(),
		}
	}
	dataList := cdList.CD
	sort.Slice(dataList, func(i, j int) bool {
		return dataList[i].BlockNumber < dataList[j].BlockNumber
	})
	start, end := SlicePage(page, pageSize, int64(len(dataList)))
	dataPage := dataList[start:end]
	return serializer.Response{
		Code: 200,
		Data: dataPage,
	}
}

func (bl *BscListener) QueryWalletAddrNft(c *gin.Context) {
	var service nftTypeService

	if err := c.ShouldBind(&service); err == nil {
		if service.WalletAddress == "" {
			c.JSON(500, xerrors.New("param can not be null").Error())
			return
		}
		res := bl.queryWalletAddrNft(service.WalletAddress, bl.network)
		c.JSON(200, res)
	} else {
		c.JSON(500, serializer.Response{
			Code: 500,
			Msg:  ErrorParam.Error(),
		})
	}
}

func (bl *BscListener) queryWalletAddrNft(addr string, network string) serializer.Response {
	if t := bl.GetJson(addr + nftTypeSuffix); t == "" {
		nftType, err := bl.queryNftFromMoralis(addr, network)
		log.Infof("queryWalletAddrNft , wallet : %s , type : %+v", addr, nftType)
		if err != nil {
			return serializer.Response{
				Code: 500,
				Msg:  err.Error(),
			}
		}
		return serializer.Response{
			Code: 200,
			Data: nftType,
		}
	}
	t := bl.GetJson(addr + nftTypeSuffix)
	log.Infof("queryWalletAddrNft , wallet : %s , nft : %s", addr, t)
	nftType := make([]NftType, 0)
	err := json.Unmarshal([]byte(t), &nftType)
	if err != nil {
		return serializer.Response{
			Code: 500,
			Msg:  err.Error(),
		}
	}
	return serializer.Response{
		Code: 200,
		Data: nftType,
	}
}

func (bl *BscListener) queryNftFromMoralis(addr string, network string) ([]NftType, error) {
	uuid := uuid.New()
	bl.nlManager.QueryNftList(uuid, addr, network)
	nftType := make([]NftType, 0)
	ctx, cancel := context.WithTimeout(context.TODO(), queryNftListTimeout)
	result, err := bl.nlManager.WaitCall(ctx, uuid)
	cancel()
	if err != nil {
		return nftType, err
	}
	nr := result.([]NftResult)
	nr = bl.convertNftResult(nr)
	dataList := parseMetadata(nr)
	dataMap := parseCacheData(dataList)

	for k, _ := range dataMap {
		nftType = append(nftType, NftType{
			Name:   k,
			Amount: len(dataMap[k]),
		})
		var cdList CacheDataList
		cdList.CD = dataMap[k]
		cacheByte, err := json.Marshal(cdList)
		if err != nil {
			break
		}
		bl.SetJson(addr+k, string(cacheByte), nftListDuration)
	}

	nftTypeByte, err := json.Marshal(nftType)
	if err != nil {
		return nftType, err
	}

	bl.SetJson(addr+nftTypeSuffix, string(nftTypeByte), nftListDuration)
	return nftType, err
}

func (bl *BscListener) SetJson(key string, value string, duration time.Duration) {
	bl.rc.Set(key, value, duration)
}

func (bl *BscListener) GetJson(key string) string {
	if bl.rc.Get(key).Err() == redis.Nil {
		return ""
	}
	return bl.rc.Get(key).Val()
}

func (bl *BscListener) convertNftResult(res []NftResult) []NftResult {
	aunft, err := contract.NewGameNft(common.HexToAddress(config.Cfg.Contract.GameNftAddress), bl.ec)
	if err != nil {
		log.Error("new auNft err : ", err)
		return res
	}
	throttle := make(chan struct{}, 20)
	var wg sync.WaitGroup
	for i, nftResult := range res {
		wg.Add(1)
		throttle <- struct{}{}
		go func(k int, v NftResult) {
			defer func() {
				wg.Done()
				<-throttle
			}()

			if v.TokenUri == "" || v.Metadata == "" {
				tokenId, err := strconv.Atoi(v.TokenId)
				if err != nil {
					log.Errorf("string %s convert int err : %v", v.TokenId, err)
					return
				}
				uri, err := aunft.TokenURI(nil, big.NewInt(int64(tokenId)))
				if err != nil {
					log.Errorf("query tokenUri tokenId : %d, err : %+v", tokenId, err)
					return
				}
				client := resty.New()
				t3 := time.Now()
				resp, err := client.R().Get(uri)
				if err != nil {
					log.Errorf("req err : uri :%s", uri)
					return
				}
				log.Info("query metadata took :", time.Since(t3))
				log.Infof("query nft tokenId : %d, uri : %s", tokenId, uri)
				var m Metadata
				err = json.Unmarshal(resp.Body(), &m)
				if err != nil {
					log.Errorf("json unmarshal err : %+v, resp : %s", err, string(resp.Body()))
					return
				}
				metadata, err := json.Marshal(m)
				if err != nil {
					log.Errorf("json marshal err : %+v, meatadata : %+v", err, m)
					return
				}
				res[k].TokenUri = uri
				res[k].Metadata = string(metadata)
			}
		}(i, nftResult)
	}
	wg.Wait()
	return res
}

func parseCacheData(cds []CacheData) map[string][]CacheData {
	dataMap := make(map[string][]CacheData)
	for _, v := range cds {
		if _, have := dataMap[v.Type]; have {
			dataMap[v.Type] = append(dataMap[v.Type], v)
		} else {
			var cd []CacheData
			cd = append(cd, v)
			dataMap[v.Type] = cd
		}
	}
	return dataMap
}

func parseMetadata(nr []NftResult) []CacheData {
	var dataList []CacheData
	for _, v := range nr {
		var cd CacheData
		cd.TokenId = v.TokenId
		cd.BlockNumber = v.BlockNumber
		var m Metadata
		err := json.Unmarshal([]byte(v.Metadata), &m)
		if err != nil {
			log.Error("json unmarshal err : ", err, v.TokenId, v.TokenUri)
			continue
		}
		split := strings.Split(m.Name, " #")
		if len(split) != 2 {
			log.Errorf("nft name len != 2")
			continue
		}
		cd.Type = split[0]
		cd.GameId = split[1]
		cd.Image = m.Image
		cd.Description = m.Description
		cd.SpikeInfo = m.SpikeInfo
		attrMap := make(map[string]interface{})
		for _, attr := range m.Attribute {
			attrMap[attr.TraitType] = attr.Value
		}
		cd.Attributes = attrMap
		dataList = append(dataList, cd)
	}
	return dataList
}

func QueryWalletNft(cursor, walletAddr, network string, res []NftResult) ([]NftResult, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("x-api-key", config.Cfg.Moralis.XApiKey).
		Get(getUrl(config.Cfg.Contract.GameNftAddress, walletAddr, network, cursor))
	if err != nil {
		log.Errorf("query wallet nft , wallet : %s, err : %+v", walletAddr, err)
		return res, err
	}
	if string(resp.Body()) == MoralisRateLimit {
		log.Errorf(MoralisRateLimit)
		return res, xerrors.New(MoralisRateLimit)
	}
	var nrs NftResults
	err = json.Unmarshal(resp.Body(), &nrs)
	if err != nil {
		log.Error("json unmarshal err : ", err)
		return res, err
	}
	res = append(res, nrs.Results...)
	if nrs.Page*nrs.PageSize >= nrs.Total {
		return res, nil
	}
	res, err = QueryWalletNft(nrs.Cursor, walletAddr, network, res)
	return res, nil
}

func getUrl(contractAddr, walletAddr, network, cursor string) string {
	return fmt.Sprintf("%s%s/nft/%s?chain=%s&cursor=%s", constants.MORALIS_API, walletAddr, contractAddr, network, cursor)
}

func SlicePage(page, pageSize, nums int64) (sliceStart, sliceEnd int64) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}
	if pageSize > nums {
		return 0, nums
	}
	pageCount := int64(math.Ceil(float64(nums) / float64(pageSize)))
	if page > pageCount {
		return 0, 0
	}
	sliceStart = (page - 1) * pageSize
	sliceEnd = sliceStart + pageSize

	if sliceEnd > nums {
		sliceEnd = nums
	}
	return sliceStart, sliceEnd
}

func (bl *BscListener) QueryWalletBalance(c *gin.Context) {
	var service balanceService

	if err := c.ShouldBind(&service); err == nil {
		res := bl.queryWalletBalance(service.Address)
		c.JSON(200, res)
	} else {
		c.JSON(500, serializer.Response{
			Code: 500,
			Msg:  ErrorParam.Error(),
		})
	}
}

func (bl *BscListener) queryWalletBalance(address string) serializer.Response {
	balanceList := make([]balanceShow, 0)
	addr := common.HexToAddress(address)
	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		defer wg.Done()
		skkContract, err := contract.NewGovernanceToken(common.HexToAddress(config.Cfg.Contract.GovernanceTokenAddress), bl.ec)
		if err != nil {
			return
		}
		skkBalance, err := skkContract.BalanceOf(nil, addr)
		if err != nil {
			return
		}
		balanceList = append(balanceList, balanceShow{
			Symbol:  "SKK",
			Balance: parseBalance(skkBalance),
		})
	}()

	go func() {
		defer wg.Done()
		sksContract, err := contract.NewGameToken(common.HexToAddress(config.Cfg.Contract.GameTokenAddress), bl.ec)
		if err != nil {
			return
		}
		sksBalance, err := sksContract.BalanceOf(nil, addr)
		if err != nil {
			return
		}
		balanceList = append(balanceList, balanceShow{
			Symbol:  "SKS",
			Balance: parseBalance(sksBalance),
		})
	}()

	go func() {
		defer wg.Done()
		usdcContract, err := contract.NewUsdc(common.HexToAddress(config.Cfg.Contract.UsdcAddress), bl.ec)
		if err != nil {
			return
		}
		usdcBalance, err := usdcContract.BalanceOf(nil, addr)
		if err != nil {
			return
		}
		balanceList = append(balanceList, balanceShow{
			Symbol:  "USDC",
			Balance: parseBalance(usdcBalance),
		})
	}()

	go func() {
		defer wg.Done()
		bnbBalance, err := bl.ec.BalanceAt(context.Background(), addr, nil)
		if err != nil {
			return
		}
		balanceList = append(balanceList, balanceShow{
			Symbol:  "BNB",
			Balance: parseBalance(bnbBalance),
		})
	}()
	wg.Wait()
	if len(balanceList) != 4 {
		return serializer.Response{
			Code: 500,
			Msg:  "chain node err ",
		}
	}
	return serializer.Response{
		Code: 200,
		Data: balanceList,
	}
}

// ToDecimal wei to decimals
func ToDecimal(ivalue interface{}, decimals int) decimal.Decimal {
	value := new(big.Int)
	switch v := ivalue.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)

	return result
}

// ToWei decimals to wei
func ToWei(iamount interface{}, decimals int) *big.Int {
	amount := decimal.NewFromFloat(0)
	switch v := iamount.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)

	wei := new(big.Int)
	wei.SetString(result.String(), 10)

	return wei
}

func parseBalance(balance *big.Int) string {
	decimalBalance := ToDecimal(balance, 18)
	return decimalBalance.String()
}
