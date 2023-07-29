package OpenApi

import (
	"wallet_graph_backend/APIClient"
	"wallet_graph_backend/config"
)

type OpenApiResult struct {
	Error  string      `json:"error"`
	Status string      `json:"status"`
	Result interface{} `json:"result"`
}

var APIClientInstance = APIClient.NewAPIClient(config.ProxyUrl, map[string]string{
	"User-Agent":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36",
	"Content-Type":     "application/json",
	"X-Requested-With": "XMLHttpRequest",
})

func init() {

}
