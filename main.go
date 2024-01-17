package main

import (
	"awesomeProject/efes"
	"awesomeProject/mongodb"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Trans struct {
	BlockNumber       string `json:"blockNumber"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `json:"hash"`
	Nonce             string `json:"nonce"`
	BlockHash         string `json:"blockHash"`
	From              string `json:"from"`
	ContractAddress   string `json:"contractAddress"`
	To                string `json:"to"`
	TokenID           string `json:"tokenID"`
	TokenName         string `json:"tokenName"`
	TokenSymbol       string `json:"tokenSymbol"`
	TokenDecimal      string `json:"tokenDecimal"`
	TransactionIndex  string `json:"transactionIndex"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	GasUsed           string `json:"gasUsed"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	Input             string `json:"input"`
	Confirmations     string `json:"confirmations"`
}

type NFTData struct {
	Status  string  `json:"status"`
	Message string  `json:"message"`
	Result  []Trans `json:"result"`
}

var (
	efescontract = strings.ToLower("0x1aae1A668c92Eb411eAfD80DD0c60ca67ad17a1c")
	agcontract   = strings.ToLower("0xff2B4721F997c242fF406a626f17df083Bd2C568")
	//address        = "0x00d758Aae6CC209C4977625d76fB986cf6F71f29"
	apikey      = "R5DGF5PJIDEEXQBNW63YFV9FKM664MDIPE"
	startblock  = 17943453
	endblock    = 19024755
	client, err = ethclient.Dial("https://eth-mainnet.g.alchemy.com/v2/cnkFUCDI1vY0c2Xg8aUmSjaK73tIveHy")
	blockNumber = int64(19024755)
	efestotal   = 1155
	agtotal     = 515
)

func main() {
	blocks, err := client.BlockByNumber(context.Background(), big.NewInt(blockNumber))
	if err != nil {
		log.Fatal(err)
	}
	blockTimeStamp := int(blocks.Time())
	CalNFTHoldTime(blockTimeStamp, agcontract, mongodb.DbcollectionAG, 515, agtotal)
	time.Sleep(1 * time.Second)
	CalNFTHoldTime(blockTimeStamp, efescontract, mongodb.DbcollectionEFES, 460, efestotal)
	//getLevel(mongodb.DbcollectionAG)
}

func CalNFTHoldTime(blockTimeStamp int, contract, collection string, start, totalsupply int) {
	for i := start; i < totalsupply; i++ {
		owner, err := OwnerOf(contract, int64(i))
		//fmt.Println(owner)
		if err != nil {
			log.Fatal(err)
		}
		for page := 1; page <= 100; page++ {
			res := getNFTLatestTrans(blockTimeStamp, contract, collection, strings.ToLower(owner), i, startblock, endblock, page, apikey)
			if res {
				break
			}
		}
	}
}

func OwnerOf(contract string, id int64) (string, error) {
	instance, err := efes.NewEfes(common.HexToAddress(contract), client)
	if err != nil {
		return "", nil
	}
	idOwner, err := instance.OwnerOf(&bind.CallOpts{BlockNumber: big.NewInt(blockNumber)}, big.NewInt(id))
	if err != nil {
		return "", nil
	}
	return idOwner.Hex(), nil
}

func getNFTLatestTrans(blockTimeStamp int, contractaddress, collection, owner string, nftid, startblock, endblock, page int, apikey string) bool {
	EtherMainTokenUrl := "https://api.etherscan.io/api?module=account&action=tokennfttx&contractaddress=%s&address=%s&page=%d&offset=100&startblock=%d&endblock=%d&sort=desc&apikey=%s"
	url := fmt.Sprintf(EtherMainTokenUrl, contractaddress, owner, page, startblock, endblock, apikey)
	req, _ := http.NewRequest("GET", url, nil)
	client := &http.Client{Timeout: time.Minute * 1}
	resp, err := client.Do(req)
	for err != nil {
		log.Fatalf("get request failed: %s", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatal(err)
	}

	blocks := NFTData{}
	err = json.NewDecoder(resp.Body).Decode(&blocks)
	if err != nil {
		log.Fatal(err)
	}
	for _, info := range blocks.Result {
		tokenId, _ := strconv.Atoi(info.TokenID)

		if tokenId == nftid && strings.ToLower(info.To) == owner {
			err := mongodb.InsertNFTDataDB(blockTimeStamp, collection, info.BlockNumber, info.TimeStamp, info.Hash, info.Nonce, info.BlockHash, info.From, info.ContractAddress, info.To, info.TokenID, info.TokenName, info.TokenSymbol, info.TokenDecimal, info.TransactionIndex, info.Gas, info.GasPrice, info.GasUsed, info.CumulativeGasUsed, info.Input, info.Confirmations)
			if err != nil {
				log.Fatal(err)
			} //xxx insert int the mongoDB
			return true
		}
	}
	return false
}

func getLevel(collection string) {
	err, res := mongodb.QueryTokenID(collection)
	if err != nil {
		log.Fatalf("err in getting NFTID: %v", err)
	}
	//err = mongodb.InsertNFTDataDB(0, collection, "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "")
	//fmt.Println(len(res))
	for _, value := range res {
		NFTData := value.(*mongodb.NFTdata)
		DoMetaDataReq(collection, NFTData.TokenID)
	}
}

type AlphaGate struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	ExternalURL string      `json:"external_url"`
	TokenID     int         `json:"tokenId"`
	Image       string      `json:"image"`
	Attributes  []Attribute `json:"attributes"`
}

type Attribute struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
}

func DoMetaDataReq(collection string, tokenId string) {
	tokenID, _ := strconv.Atoi(tokenId)
	metadataUrl := fmt.Sprintf("https://metadata.spacenation.online/alpha-gate/%d.json", tokenID)
	req, _ := http.NewRequest("GET", metadataUrl, nil)
	client := &http.Client{Timeout: time.Minute * 1}
	resp, err := client.Do(req)
	for err != nil {
		log.Fatalf("get request failed: %s", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatal(err)
	}

	blocks := AlphaGate{}
	err = json.NewDecoder(resp.Body).Decode(&blocks)
	if err != nil {
		log.Fatal(err)
	}
	err = mongodb.UpdateLevel(collection, tokenID, blocks.Attributes[0].Value)
	if err != nil {
		log.Fatal(err)
	}
}
