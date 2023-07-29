package config

import (
	"wallet_graph_backend/utils"
)

var MongodbUrl = utils.GetEnv("MONGO_URL", "")
var MongodbName = utils.GetEnv("MONGO_NAME", "misttrack")

var RedisUrl = utils.GetEnv("REDIS_URL", "redis://localhost:6379")

var ProxyUrl = utils.GetEnv("HTTPS_PROXY", "")
