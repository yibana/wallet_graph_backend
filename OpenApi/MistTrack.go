package OpenApi

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
	"time"
	"wallet_graph_backend/db"
)

const Misttrack_url = "https://openapi.misttrack.io/v1"
const MisttrackApiKey = "YourApiKey"

type MisttrackResult struct {
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func OpenApiMistTrack(method, coin, address string) (OpenApiResult, error) {
	if coin == "" {
		coin = "eth"
	}

	var MongoInstance *db.MongoManger
	switch method {
	case "address_labels":
		MongoInstance = db.MongoInstance_address_labels
	case "address_overview":
		MongoInstance = db.MongoInstance_address_overview
	case "risk_score":
		MongoInstance = db.MongoInstance_risk_score
	case "transactions_investigation":
		MongoInstance = db.MongoInstance_transactions_investigation
	}

	// 尝试从数据库读取
	find, err := MongoInstance.MongoFind(bson.M{
		"filter": bson.M{
			"address": address,
			"coin":    strings.ToLower(coin),
		},
	})
	var old_data bson.M
	// if update_time < now - 10min , update else return
	if err == nil && len(find) != 0 {
		old_data = find[0]["data"].(bson.M)
		update_time := find[0]["update_time"]
		if update_time.(int64) > time.Now().UnixMilli()-60*60*1000 {
			return OpenApiResult{
				Error:  "old",
				Result: old_data,
			}, nil
		}
	}

	url := fmt.Sprintf("%s/%s?coin=%s&address=%s&api_key=%s",
		Misttrack_url, method, coin, address, MisttrackApiKey)
	bytes, err := APIClientInstance.Get(url)
	if err != nil {
		return OpenApiResult{}, err
	}
	var result MisttrackResult
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return OpenApiResult{}, err
	}

	if !result.Success {
		if err != nil || len(find) == 0 {
			return OpenApiResult{
				Error:  result.Msg,
				Status: "error",
			}, nil
		}
		return OpenApiResult{
			Error:  "",
			Status: "success",
			Result: old_data,
		}, nil

	}
	// 如果获取成功，写入到MongoInstance_address_labels

	address = strings.ToLower(address)
	err = MongoInstance.MongoUpdateOne(bson.M{
		"filter": bson.M{
			"address": address,
			"coin":    strings.ToLower(coin),
		},
		"update": bson.M{
			"$set": bson.M{
				"address": address,
				"coin":    strings.ToLower(coin),
				"data":    result.Data,
			},
		},
	})
	if err != nil {
		return OpenApiResult{}, err
	}
	return OpenApiResult{
		Error:  "",
		Status: "success",
		Result: result.Data,
	}, nil

}
