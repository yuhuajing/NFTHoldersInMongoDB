package mongodb

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/utils"
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
