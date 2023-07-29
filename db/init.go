package db

var RedisCacheInstance *RedisCacheManger
var MongoInstance_address_labels *MongoManger
var MongoInstance_address_overview *MongoManger
var MongoInstance_risk_score *MongoManger
var MongoInstance_transactions_investigation *MongoManger

func init() {
	var err error
	//RedisCacheInstance, err = NewRedisCacheManger(config.RedisUrl)
	//if err != nil {
	//	panic(err)
	//}

	MongoInstance_address_labels, err = NewMongoManger("address_labels")
	if err != nil {
		panic(err)
	}

	MongoInstance_address_overview, err = NewMongoManger("address_overview")
	if err != nil {
		panic(err)
	}

	MongoInstance_risk_score, err = NewMongoManger("risk_score")
	if err != nil {
		panic(err)
	}

	MongoInstance_transactions_investigation, err = NewMongoManger("transactions_investigation")
	if err != nil {
		panic(err)
	}

}
