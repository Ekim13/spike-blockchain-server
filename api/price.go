package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	logger "github.com/ipfs/go-log"
	"spike-blockchain-server/chain"
	"spike-blockchain-server/config"
	"spike-blockchain-server/constants"
	"spike-blockchain-server/serializer"
	service "spike-blockchain-server/service/price"
)

var log = logger.Logger("api")

func FindERC20TokenPrice(c *gin.Context) {
	var service service.TokenPriceService
	if err := c.ShouldBind(&service); err == nil {
		res := findERC20TokenPrice(service.Token)
		c.JSON(200, res)
	} else {
		c.JSON(500, serializer.Response{
			Code: 500,
			Msg:  chain.ErrorParam.Error(),
		})
	}
}

type priceResp struct {
	UsdPrice float64 `json:"usdPrice"`
}

func findERC20TokenPrice(token string) serializer.Response {
	contractAddr, err := service.GetTokenContractAddrByTokenSymbol(token)
	if err != nil {
		return serializer.Response{
			Code: 500,
			Msg:  err.Error(),
		}
	}
	client := resty.New()
	log.Infof("query erc20 price url : %s", getUrl(contractAddr))

	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("x-api-key", config.Cfg.Moralis.XApiKey).
		Get(getUrl(contractAddr))
	if err != nil {
		return serializer.Response{
			Code: 500,
			Msg:  err.Error(),
		}
	}
	if resp.IsError() {
		return serializer.Response{
			Code: 401,
			Msg:  resp.String(),
		}
	}

	var res priceResp
	err = json.Unmarshal(resp.Body(), &res)
	if err != nil {
		return serializer.Response{
			Code: 500,
			Msg:  err.Error(),
		}
	}
	return serializer.Response{
		Code: 200,
		Data: res.UsdPrice,
	}

}

func getUrl(contractAddr string) string {
	return fmt.Sprintf("%serc20/%s/price?chain=bsc", constants.MORALIS_API, contractAddr)
}
