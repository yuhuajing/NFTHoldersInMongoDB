package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var (
	Mongodburl       = "mongodb://clay:password@127.0.0.1:27017"
	Dbname           = "holdnft"
	DbcollectionEFES = "efes"
	DbcollectionAG   = "ag"
	Mongoclient      *mongo.Client
)

func init() {
	Mongoclient = GetMongoClient(Mongodburl)
	fmt.Println("ConnedDB")
}

func GetMongoClient(mongourl string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		mongourl,
	))

	if err != nil {
		log.Fatalf("err in conn MonggoDB: %v", err)

	}
	return client
}

func GetCollection(dbname, dbcollectionname, mongourl string) *mongo.Collection {
	err := Mongoclient.Ping(context.TODO(), nil)
	if err != nil {
		Mongoclient = GetMongoClient(mongourl)
		log.Fatalf("error in connecting mongodb: %v", err)
	}
	database := Mongoclient.Database(dbname)
	txdata_collection := database.Collection(dbcollectionname)
	return txdata_collection
}

type DefaultModel struct {
	ID        string    `json:"id" bson:"_id" msg:"id"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt" msg:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt" msg:"updatedAt"`
}

// VersionModal struct contains a model's version field.
type VersionModal struct {
	DefaultModel `bson:",inline"`
	Version_     int64 `json:"_v" bson:"_v" msg:"_v"`
}

type NFTdata struct {
	VersionModal      `bson:",inline"`
	BlockNumber       string `json:"blockNumber" bson:"blockNumber" msg:"blockNumber"`
	TimeStamp         string `json:"timeStamp" bson:"timeStamp" msg:"timeStamp"`
	Hash              string `json:"hash" bson:"hash" msg:"hash"`
	Nonce             string `json:"nonce" bson:"nonce" msg:"nonce"`
	BlockHash         string `json:"blockHash" bson:"blockHash" msg:"blockHash"`
	From              string `json:"from" bson:"from" msg:"from"`
	ContractAddress   string `json:"contractAddress" bson:"contractAddress" msg:"contractAddress"`
	To                string `json:"to" bson:"to" msg:"to"`
	TokenID           string `json:"tokenID" bson:"tokenID" msg:"tokenID"`
	TokenName         string `json:"tokenName" bson:"tokenName" msg:"tokenName"`
	TokenSymbol       string `json:"tokenSymbol" bson:"tokenSymbol" msg:"tokenSymbol"`
	TokenDecimal      string `json:"tokenDecimal" bson:"tokenDecimal" msg:"tokenDecimal"`
	TransactionIndex  string `json:"transactionIndex" bson:"transactionIndex" msg:"transactionIndex"`
	Gas               string `json:"gas" bson:"gas" msg:"gas"`
	GasPrice          string `json:"gasPrice" bson:"gasPrice" msg:"gasPrice"`
	GasUsed           string `json:"gasUsed" bson:"gasUsed" msg:"gasUsed"`
	CumulativeGasUsed string `json:"cumulativeGasUsed" bson:"cumulativeGasUsed" msg:"cumulativeGasUsed"`
	Input             string `json:"input" bson:"input" msg:"input"`
	Confirmations     string `json:"confirmations" bson:"confirmations" msg:"confirmations"`
	HoldingTime       int    `json:"holdingTime" bson:"holdingTime" msg:"holdingTime"`
}
