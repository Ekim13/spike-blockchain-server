package wallet

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-resty/resty/v2"
	"math/big"
	"os"
	"spike-blockchain-server/chain"
	"spike-blockchain-server/config"
	"spike-blockchain-server/serializer"
	"strconv"
	"strings"
)

type NativeTransactionRecordService struct {
	Address string `form:"address" json:"address" binding:"required"`
}

func (s *NativeTransactionRecordService) FindNativeTransactionRecord() serializer.Response {
	client := resty.New()

	apiKey := os.Getenv("BSC_API_KEY")

	bscClient, err := ethclient.Dial(os.Getenv("MORALIS_SPEEDY_NODE"))
	if err != nil {
		return serializer.Response{
			Code:  400,
			Error: err.Error(),
		}
	}
	defer bscClient.Close()
	blockNumber, err := bscClient.BlockNumber(context.Background())
	if err != nil {
		return serializer.Response{
			Code:  402,
			Error: err.Error(),
		}
	}

	resp, err := client.R().
		SetHeader("Accept", "application/json").
		Get(s.url(apiKey, blockNumber))
	if err != nil {
		return serializer.Response{
			Code:  403,
			Error: err.Error(),
		}
	}
	bscRes := &BscRes{Result: make([]Result, 0)}
	bnbRecord := make([]Result, 0)

	err = json.Unmarshal(resp.Body(), &bscRes)
	if err != nil {
		return serializer.Response{
			Code:  405,
			Error: err.Error(),
		}
	}

	if len(bscRes.Result) != 0 {

		for i, result := range bscRes.Result {
			if bscRes.Result[i].Input == "0x" {
				bnbRecord = append(bnbRecord, bscRes.Result[i])
				continue
			}
			methodId := result.Input[0:10]
			switch methodId {
			case hexutil.Encode(chain.GetTxMethodName("swapExactTokensForETHSupportingFeeOnTransferTokens(uint256,uint256,address[],address,uint256)")):
				height, err := strconv.ParseInt(bscRes.Result[i].BlockNumber, 10, 64)
				if err != nil {
					return serializer.Response{
						Code:  406,
						Error: err.Error(),
					}
				}

				query := ethereum.FilterQuery{
					FromBlock: big.NewInt(height),
					ToBlock:   big.NewInt(height),
				}
				sub, err := bscClient.FilterLogs(context.Background(), query)
				if err != nil {
					return serializer.Response{
						Code:  407,
						Error: err.Error(),
					}
				}
				for _, logEvent := range sub {
					if logEvent.Topics[0].String() == chain.EventSignHash(chain.WITHRAWALTOPIC) {
						bscRes.Result[i].Type = "Receive"
						value := new(big.Int)
						value.SetString(strings.Split(hexutil.Encode(logEvent.Data), "0x")[1], 16)

						bscRes.Result[i].Value = value.String()
						bnbRecord = append(bnbRecord, bscRes.Result[i])
						break
					}
				}
			case hexutil.Encode(chain.GetTxMethodName("swapExactETHForTokens(uint256,address[],address,uint256)")):
				bscRes.Result[i].Type = "Send"
				bnbRecord = append(bnbRecord, bscRes.Result[i])
			}

		}
		bscRes.Result = bnbRecord
		return serializer.Response{
			Code: 200,
			Data: bscRes,
		}
	}
	bscRes.Result = make([]Result, 0)
	return serializer.Response{
		Code: 200,
		Data: bscRes,
	}
}

func (s *NativeTransactionRecordService) url(apiKey string, blockNumber uint64) string {
	return fmt.Sprintf("%s?module=account&action=txlist&address=%s&startblock=%d&endblock=%d&offset=10000&page=1&sort=desc&apikey=%s", config.Cfg.BscScan.UrlPrefix, s.Address, blockNumber-201600, blockNumber, apiKey)
}
