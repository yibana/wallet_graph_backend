package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"wallet_graph_backend/routes"
)

var Router *gin.Engine

func main() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}

	r.Use(cors.New(config))
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})
	r.GET("/readme", routes.Readme)

	//https://openapi.misttrack.io/v1/transactions_investigation
	//   ?coin=ETH
	//   &address=0xb3065fe2125c413e973829108f23e872e1db9a6b
	//   &api_key=YourApiKey
	r.GET("/openapi/misttrack/v1/:method", routes.OpenApiMistTrack)
	r.Run(":8081")
}
