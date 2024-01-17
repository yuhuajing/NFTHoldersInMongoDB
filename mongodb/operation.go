package mongodb

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
	"strconv"
	"time"
)

func InsertNFTDataDB(blockTimeStamp int, collectionname, BlockNumber, TimeStamp, Hash, Nonce, BlockHash, From, ContractAddress, To, TokenID, TokenName, TokenSymbol, TokenDecimal, TransactionIndex, Gas, GasPrice, GasUsed, CumulativeGasUsed, Input, Confirmations string) error {
	startHold, _ := strconv.Atoi(TimeStamp)
	var res = NFTdata{
		BlockNumber:       BlockNumber,
		TimeStamp:         TimeStamp,
		Hash:              Hash,
		Nonce:             Nonce,
		BlockHash:         BlockHash,
		From:              From,
		ContractAddress:   ContractAddress,
		To:                To,
		TokenID:           TokenID,
		TokenName:         TokenName,
		TokenSymbol:       TokenSymbol,
		TokenDecimal:      TokenDecimal,
		TransactionIndex:  TransactionIndex,
		Gas:               Gas,
		GasPrice:          GasPrice,
		GasUsed:           GasUsed,
		CumulativeGasUsed: CumulativeGasUsed,
		Input:             Input,
		Confirmations:     Confirmations,
		VersionModal: VersionModal{
			DefaultModel: DefaultModel{
				ID:        utils.UUIDv4(),
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC(),
			},
			Version_: 0,
		},
		HoldingTime: blockTimeStamp - startHold,
	}

	err := InsertDocument(collectionname, res)
	if err != nil {
		return fmt.Errorf("InsertNFTdataDB:err in inserting NFTData")
	}
	return nil
}

func InsertDocument(collectionname string, data interface{}) error {
	collection := GetCollection(Dbname, collectionname, Mongodburl)
	_, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		return fmt.Errorf("failed to insert document: %v", err)
	}
	return nil
}

func QueryTokenID(collectionname string) (error, []interface{}) {
	//filter := bson.M{"chain": chain, "contract": contract, "collection": collection}
	err, idres := GetDocuments(collectionname, bson.M{}, &NFTdata{})
	if err != nil {
		return fmt.Errorf("err in getting NFTStageData: %v", err), nil
	}
	return nil, idres
}

func GetDocuments(collectionname string, filter bson.M, result interface{}, opts ...*options.FindOptions) (error, []interface{}) {
	collection := GetCollection(Dbname, collectionname, Mongodburl)
	cur, err := collection.Find(context.Background(), filter, opts...)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("failed to get documents: %v", err)), nil
	}
	defer cur.Close(context.Background())

	var results []interface{}
	for cur.Next(context.Background()) {
		// Create a new instance of the result type for each document
		res := reflect.New(reflect.TypeOf(result).Elem()).Interface()
		err := cur.Decode(res)
		if err != nil {
			return fmt.Errorf(fmt.Sprintf("failed to decode document: %v", err)), nil
		}
		results = append(results, res)
	}

	if err := cur.Err(); err != nil {
		return fmt.Errorf(fmt.Sprintf("cursor error: %v", err)), nil
	}

	return nil, results
}

func UpdateLevel(collectionname string, tokenID int, level string) error {
	update := bson.M{"$set": bson.M{"level": level}}
	fmt.Println(level)
	itopkenID := strconv.FormatInt(int64(tokenID), 10)
	fmt.Println(itopkenID)
	filter := bson.M{"tokenID": itopkenID}
	err := UpdateDocument(collectionname, filter, update)
	if err != nil {
		return fmt.Errorf("err in updating order's tx status: %v", err)
	}
	return nil
}

func UpdateDocument(collectionname string, filter bson.M, update bson.M) error {
	fmt.Println("up")
	collection := GetCollection(Dbname, collectionname, Mongodburl)
	_, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return fmt.Errorf("failed to update document: %v", err)
	}
	return nil
}
