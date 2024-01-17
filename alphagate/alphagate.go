package alphagate

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IERC721ATokenOwnership is an auto generated low-level Go binding around an user-defined struct.
type IERC721ATokenOwnership struct {
	Addr           common.Address
	StartTimestamp uint64
	Burned         bool
	ExtraData      *big.Int
}

// IERC721MMintStageInfo is an auto generated low-level Go binding around an user-defined struct.
type IERC721MMintStageInfo struct {
	Price                *big.Int
	WalletLimit          uint32
	MerkleRoot           [32]byte
	MaxStageSupply       *big.Int
	StartTimeUnixSeconds uint64
	EndTimeUnixSeconds   uint64
}

// AlphegateMetaData contains all meta data concerning the Alphegate contract.
var AlphegateMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"collectionName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"collectionSymbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenURISuffix\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"maxMintableSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"globalWalletLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"timestampExpirySeconds\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ApprovalCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceQueryForZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotIncreaseMaxMintableSupply\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotUpdatePermanentBaseURI\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CosignerNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CrossmintAddressNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CrossmintOnly\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalWalletLimitOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientStageTimeGap\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCosignSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidQueryRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStage\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStageArgsLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStartAndEndTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintERC2309QuantityExceedsLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintZeroQuantity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Mintable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoSupplyLeft\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotMintable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnershipNotInitializedForExtraData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StageSupplyExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimestampExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromIncorrectOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToNonERC721ReceiverImplementer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"URIQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WalletGlobalLimitExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WalletStageLimitExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"ConsecutiveTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"}],\"name\":\"PermanentBaseURI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activeStage\",\"type\":\"uint256\"}],\"name\":\"SetActiveStage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"}],\"name\":\"SetBaseURI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"SetCosigner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"crossmintAddress\",\"type\":\"address\"}],\"name\":\"SetCrossmintAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"globalWalletLimit\",\"type\":\"uint256\"}],\"name\":\"SetGlobalWalletLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxMintableSupply\",\"type\":\"uint256\"}],\"name\":\"SetMaxMintableSupply\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"mintable\",\"type\":\"bool\"}],\"name\":\"SetMintable\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"}],\"name\":\"SetTimestampExpirySeconds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stage\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint80\",\"name\":\"price\",\"type\":\"uint80\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"walletLimit\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"maxStageSupply\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"startTimeUnixSeconds\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"endTimeUnixSeconds\",\"type\":\"uint64\"}],\"name\":\"UpdateStage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"qty\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"assertValidCosign\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"qty\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"crossmint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"explicitOwnershipOf\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"burned\",\"type\":\"bool\"},{\"internalType\":\"uint24\",\"name\":\"extraData\",\"type\":\"uint24\"}],\"internalType\":\"structIERC721A.TokenOwnership\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"explicitOwnershipsOf\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"burned\",\"type\":\"bool\"},{\"internalType\":\"uint24\",\"name\":\"extraData\",\"type\":\"uint24\"}],\"internalType\":\"structIERC721A.TokenOwnership[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"getActiveStageFromTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"qty\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"getCosignDigest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"getCosignNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCosigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCrossmintAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGlobalWalletLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxMintableSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMintable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumberStages\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getStageInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint80\",\"name\":\"price\",\"type\":\"uint80\"},{\"internalType\":\"uint32\",\"name\":\"walletLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint24\",\"name\":\"maxStageSupply\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"startTimeUnixSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endTimeUnixSeconds\",\"type\":\"uint64\"}],\"internalType\":\"structIERC721M.MintStageInfo\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTimestampExpirySeconds\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenURISuffix\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"qty\",\"type\":\"uint32\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"qty\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"ownerMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setBaseURIPermanent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"setCosigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"crossmintAddress\",\"type\":\"address\"}],\"name\":\"setCrossmintAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"globalWalletLimit\",\"type\":\"uint256\"}],\"name\":\"setGlobalWalletLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxMintableSupply\",\"type\":\"uint256\"}],\"name\":\"setMaxMintableSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"mintable\",\"type\":\"bool\"}],\"name\":\"setMintable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint80\",\"name\":\"price\",\"type\":\"uint80\"},{\"internalType\":\"uint32\",\"name\":\"walletLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint24\",\"name\":\"maxStageSupply\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"startTimeUnixSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endTimeUnixSeconds\",\"type\":\"uint64\"}],\"internalType\":\"structIERC721M.MintStageInfo[]\",\"name\":\"newStages\",\"type\":\"tuple[]\"}],\"name\":\"setStages\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"}],\"name\":\"setTimestampExpirySeconds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"suffix\",\"type\":\"string\"}],\"name\":\"setTokenURISuffix\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stop\",\"type\":\"uint256\"}],\"name\":\"tokensOfOwnerIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"totalMintedByAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"price\",\"type\":\"uint80\"},{\"internalType\":\"uint32\",\"name\":\"walletLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint24\",\"name\":\"maxStageSupply\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"startTimeUnixSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endTimeUnixSeconds\",\"type\":\"uint64\"}],\"name\":\"updateStage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052600436106102c95760003560e01c80638462151c11610175578063aac5ab1f116100dc578063ce2b0ec011610095578063efdaa2ec1161006f578063efdaa2ec146108eb578063f2fde38b14610900578063f698bceb14610920578063f8d096961461093857600080fd5b8063ce2b0ec01461086f578063e985e9c51461088f578063efb6b11f146108d857600080fd5b8063aac5ab1f146107ba578063b50248e7146107da578063b7a9fa60146107fa578063b88d4fde1461080f578063c23dc68f14610822578063c87b56dd1461084f57600080fd5b8063997556241161012e578063997556241461069e57806399a2557a146106be578063a06c492f1461067e578063a22cb465146106de578063a3759f60146106fe578063a9852bfb1461079a57600080fd5b80638462151c146105e05780638da5cb5b1461060d5780638dcdb09d1461062b5780638f9315111461064b57806395d89b411461066957806397cf84fc1461067e57600080fd5b80633ccfd60b1161023457806362acbd9a116101ed57806370a08231116101c757806370a082311461057657806370da24ee14610596578063715018a6146105ab57806373e1607e146105c057600080fd5b806362acbd9a146105235780636352211e1461053657806367808a341461055657600080fd5b80633ccfd60b1461046257806342842e0e146104775780634ae0402f1461048a5780634b1c53b4146104c157806355f804b3146104d65780635bbb2177146104f657600080fd5b806318160ddd1161028657806318160ddd146103a75780631ce03eed146103ca57806323b872dd146103ea578063285d70d4146103fd57806333bbbf061461041d578063372992e41461044257600080fd5b806301ffc9a7146102ce578063020451381461030357806306fdde0314610325578063081812fc14610347578063095ea7b31461037f5780631053a81514610392575b600080fd5b3480156102da57600080fd5b506102ee6102e936600461345c565b610958565b60405190151581526020015b60405180910390f35b34801561030f57600080fd5b5061032361031e366004613495565b6109aa565b005b34801561033157600080fd5b5061033a610a28565b6040516102fa9190613500565b34801561035357600080fd5b50610367610362366004613513565b610aba565b6040516001600160a01b0390911681526020016102fa565b61032361038d36600461352c565b610afe565b34801561039e57600080fd5b50610323610b9e565b3480156103b357600080fd5b50600154600054035b6040519081526020016102fa565b3480156103d657600080fd5b506103bc6103e5366004613581565b610bf0565b6103236103f83660046135c4565b610d30565b34801561040957600080fd5b5061032361041836600461360e565b610e91565b34801561042957600080fd5b50600a54600160501b90046001600160a01b0316610367565b34801561044e57600080fd5b5061032361045d366004613513565b610eda565b34801561046e57600080fd5b50610323610f3a565b6103236104853660046135c4565b610fe5565b34801561049657600080fd5b50600a546201000090046001600160401b03166040516001600160401b0390911681526020016102fa565b3480156104cd57600080fd5b50600c546103bc565b3480156104e257600080fd5b506103236104f136600461366c565b611136565b34801561050257600080fd5b506105166105113660046136f1565b6111a6565b6040516102fa9190613762565b6103236105313660046137a4565b611271565b34801561054257600080fd5b50610367610551366004613513565b611339565b34801561056257600080fd5b506103bc610571366004613848565b611344565b34801561058257600080fd5b506103bc610591366004613495565b61140c565b3480156105a257600080fd5b506010546103bc565b3480156105b757600080fd5b5061032361145a565b3480156105cc57600080fd5b506103236105db36600461388d565b61146e565b3480156105ec57600080fd5b506106006105fb366004613495565b61170e565b6040516102fa9190613905565b34801561061957600080fd5b506008546001600160a01b0316610367565b34801561063757600080fd5b5061032361064636600461393d565b611816565b34801561065757600080fd5b50600b546001600160a01b0316610367565b34801561067557600080fd5b5061033a611ceb565b34801561068a57600080fd5b506103bc610699366004613495565b611cfa565b3480156106aa57600080fd5b506103236106b9366004613495565b611d24565b3480156106ca57600080fd5b506106006106d93660046139b1565b611d7a565b3480156106ea57600080fd5b506103236106f93660046139e4565b611ef1565b34801561070a57600080fd5b5061071e610719366004613513565b611f5d565b6040805184516001600160501b0316815260208086015163ffffffff90811691830191909152858301519282019290925260608086015162ffffff16908201526080808601516001600160401b039081169183019190915260a095860151169481019490945290911660c083015260e0820152610100016102fa565b3480156107a657600080fd5b506103236107b536600461366c565b6120a1565b3480156107c657600080fd5b506103236107d5366004613a1b565b6120bb565b3480156107e657600080fd5b506103236107f5366004613af0565b612114565b34801561080657600080fd5b5061033a612158565b61032361081d366004613b5e565b612167565b34801561082e57600080fd5b5061084261083d366004613513565b6122bf565b6040516102fa9190613bad565b34801561085b57600080fd5b5061033a61086a366004613513565b612337565b34801561087b57600080fd5b5061032361088a366004613848565b612440565b34801561089b57600080fd5b506102ee6108aa366004613bbb565b6001600160a01b03918216600090815260076020908152604080832093909416825291909152205460ff1690565b6103236108e6366004613bd7565b6124a1565b3480156108f757600080fd5b50600d546103bc565b34801561090c57600080fd5b5061032361091b366004613495565b612514565b34801561092c57600080fd5b50600a5460ff166102ee565b34801561094457600080fd5b50610323610953366004613513565b61258d565b60006301ffc9a760e01b6001600160e01b03198316148061098957506380ac58cd60e01b6001600160e01b03198316145b806109a45750635b5e139f60e01b6001600160e01b03198316145b92915050565b6109b26125ed565b600a80547fffff0000000000000000000000000000000000000000ffffffffffffffffffff16600160501b6001600160a01b038416908102919091179091556040519081527faea1573caf7b4fdd079b947d86c1be6c725642c47582f8f9bd2c7d2a30bf0bd9906020015b60405180910390a150565b606060028054610a3790613c6a565b80601f0160208091040260200160405190810160405280929190818152602001828054610a6390613c6a565b8015610ab05780601f10610a8557610100808354040283529160200191610ab0565b820191906000526020600020905b815481529060010190602001808311610a9357829003601f168201915b5050505050905090565b6000610ac582612647565b610ae2576040516333d1c03960e21b815260040160405180910390fd5b506000908152600660205260409020546001600160a01b031690565b6000610b0982611339565b9050336001600160a01b03821614610b4257610b2581336108aa565b610b42576040516367d9dca160e11b815260040160405180910390fd5b60008281526006602052604080822080546001600160a01b0319166001600160a01b0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b610ba66125ed565b600a805461ff0019166101001790556040517fc6a6c2b165e62c9d37fc51a18ed76e5be22304bc1d337877c98f31c23e40b0f590610be690600e90613ca4565b60405180910390a1565b600a54600090600160501b90046001600160a01b0316610c23576040516353bd4fb360e11b815260040160405180910390fd5b610d26308585600a8054906101000a90046001600160a01b031686610c454690565b610c4e8b611cfa565b604051606097881b6bffffffffffffffffffffffff19908116602083015296881b8716603482015260e09590951b6001600160e01b031916604886015292861b909416604c84015260c01b6001600160c01b031916938201939093526068810191909152608881019190915260a801604051602081830303815290604052805190602001206040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b90505b9392505050565b826daaeb6d7670e522a718067333cd4e3b15610e8057336001600160a01b03821603610d6657610d6184848461266e565b610e8b565b604051633185c44d60e21b81523060048201523360248201526daaeb6d7670e522a718067333cd4e9063c617113490604401602060405180830381865afa158015610db5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dd99190613d2f565b8015610e5c5750604051633185c44d60e21b81523060048201526001600160a01b03821660248201526daaeb6d7670e522a718067333cd4e9063c617113490604401602060405180830381865afa158015610e38573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e5c9190613d2f565b610e8057604051633b79c77360e21b81523360048201526024015b60405180910390fd5b610e8b84848461266e565b50505050565b610e996125ed565b600a805460ff19168215159081179091556040519081527fe717a2bfc51e250b028aaac5eb448e76f4df26b9609956782bff49097bb792cf90602001610a1d565b610ee26125ed565b600c54811115610f0557604051630590c51360e01b815260040160405180910390fd5b600d8190556040518181527f5307de8ad7d34d5ddfd5171435c143bdc645493980f453eb5d7cdb3e494a1b3590602001610a1d565b610f426125ed565b6040514790600090339083908381818185875af1925050503d8060008114610f86576040519150601f19603f3d011682016040523d82523d6000602084013e610f8b565b606091505b5050905080610fad57604051631d42c86760e21b815260040160405180910390fd5b6040518281527f5b6b431d4476a211bb7d41c20d1aab9ae2321deee0d20be3d9fc9b1093fa6e3d906020015b60405180910390a15050565b826daaeb6d7670e522a718067333cd4e3b1561112b57336001600160a01b0382160361101657610d61848484612807565b604051633185c44d60e21b81523060048201523360248201526daaeb6d7670e522a718067333cd4e9063c617113490604401602060405180830381865afa158015611065573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110899190613d2f565b801561110c5750604051633185c44d60e21b81523060048201526001600160a01b03821660248201526daaeb6d7670e522a718067333cd4e9063c617113490604401602060405180830381865afa1580156110e8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061110c9190613d2f565b61112b57604051633b79c77360e21b8152336004820152602401610e77565b610e8b848484612807565b61113e6125ed565b600a54610100900460ff1615611167576040516306ccad4160e41b815260040160405180910390fd5b600e611174828483613d92565b507f23c8c9488efebfd474e85a7956de6f39b17c7ab88502d42a623db2d8e382bbaa8282604051610fd9929190613e51565b6060816000816001600160401b038111156111c3576111c3613a4e565b60405190808252806020026020018201604052801561121557816020015b6040805160808101825260008082526020808301829052928201819052606082015282526000199092019101816111e15790505b50905060005b8281146112685761124386868381811061123757611237613e80565b905060200201356122bf565b82828151811061125557611255613e80565b602090810291909101015260010161121b565b50949350505050565b6002600954036112c35760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610e77565b6002600955600b546001600160a01b03166112f157604051637e9f68eb60e11b815260040160405180910390fd5b600b546001600160a01b0316331461131c5760405163f46fd68360e01b815260040160405180910390fd5b61132b87878787878787612822565b505060016009555050505050565b60006109a482612c7c565b6000805b6010548110156113f2576010818154811061136557611365613e80565b60009182526020909120600260039092020101546001600160401b0363010000009091048116908416108015906113d55750601081815481106113aa576113aa613e80565b60009182526020909120600260039092020101546001600160401b03600160581b9091048116908416105b156113e05792915050565b806113ea81613eac565b915050611348565b5060405163e82a532960e01b815260040160405180910390fd5b60006001600160a01b038216611435576040516323d3ad8160e21b815260040160405180910390fd5b506001600160a01b03166000908152600560205260409020546001600160401b031690565b6114626125ed565b61146c6000612ce3565b565b6114766125ed565b60105487106114985760405163e82a532960e01b815260040160405180910390fd5b6001871061153057600a546201000090046001600160401b031660106114bf60018a613ec5565b815481106114cf576114cf613e80565b9060005260206000209060030201600201600b9054906101000a90046001600160401b03166114fe9190613ed8565b6001600160401b0316826001600160401b0316101561153057604051636bc1af9360e01b815260040160405180910390fd5b61153a8282612d35565b856010888154811061154e5761154e613e80565b906000526020600020906003020160000160006101000a8154816001600160501b0302191690836001600160501b03160217905550846010888154811061159757611597613e80565b9060005260206000209060030201600001600a6101000a81548163ffffffff021916908363ffffffff16021790555083601088815481106115da576115da613e80565b906000526020600020906003020160010181905550826010888154811061160357611603613e80565b906000526020600020906003020160020160006101000a81548162ffffff021916908362ffffff160217905550816010888154811061164457611644613e80565b906000526020600020906003020160020160036101000a8154816001600160401b0302191690836001600160401b03160217905550806010888154811061168d5761168d613e80565b9060005260206000209060030201600201600b6101000a8154816001600160401b0302191690836001600160401b031602179055507fb3268648542a1bb1b2dd12e3b14aeb5a3ab22c592de96bdd3e842154a5b394fa878787878787876040516116fd9796959493929190613eff565b60405180910390a150505050505050565b6060600080600061171e8561140c565b90506000816001600160401b0381111561173a5761173a613a4e565b604051908082528060200260200182016040528015611763578160200160208202803683370190505b50905061179060408051608081018252600080825260208201819052918101829052606081019190915290565b60005b83861461180a576117a381612d6b565b915081604001516118025781516001600160a01b0316156117c357815194505b876001600160a01b0316856001600160a01b03160361180257808387806001019850815181106117f5576117f5613e80565b6020026020010181815250505b600101611793565b50909695505050505050565b61181e6125ed565b60105460005b8181101561189957601080548061183d5761183d613f4f565b60008281526020812060036000199093019283020180546001600160701b03191681556001810191909155600201805472ffffffffffffffffffffffffffffffffffffff1916905590558061189181613eac565b915050611824565b50600a546201000090046001600160401b031660005b83811015611ce45760018110611959578185856118cd600185613ec5565b8181106118dc576118dc613e80565b905060c0020160a00160208101906118f49190613848565b6118fe9190613ed8565b6001600160401b031685858381811061191957611919613e80565b905060c0020160800160208101906119319190613848565b6001600160401b0316101561195957604051636bc1af9360e01b815260040160405180910390fd5b6119b585858381811061196e5761196e613e80565b905060c0020160800160208101906119869190613848565b86868481811061199857611998613e80565b905060c0020160a00160208101906119b09190613848565b612d35565b60106040518060c001604052808787858181106119d4576119d4613e80565b6119ea92602060c0909202019081019150613f65565b6001600160501b03168152602001878785818110611a0a57611a0a613e80565b905060c002016020016020810190611a229190613f80565b63ffffffff168152602001878785818110611a3f57611a3f613e80565b905060c00201604001358152602001878785818110611a6057611a60613e80565b905060c002016060016020810190611a789190613f9b565b62ffffff168152602001878785818110611a9457611a94613e80565b905060c002016080016020810190611aac9190613848565b6001600160401b03168152602001878785818110611acc57611acc613e80565b905060c0020160a0016020810190611ae49190613848565b6001600160401b039081169091528254600181810185556000948552602094859020845160039093020180549585015163ffffffff16600160501b026001600160701b03199096166001600160501b0390931692909217949094178155604083015193810193909355606082015160029093018054608084015160a0909401518316600160581b0267ffffffffffffffff60581b19949093166301000000026affffffffffffffffffffff1990911662ffffff9095169490941793909317919091161790557fb3268648542a1bb1b2dd12e3b14aeb5a3ab22c592de96bdd3e842154a5b394fa81868682818110611bdd57611bdd613e80565b611bf392602060c0909202019081019150613f65565b878785818110611c0557611c05613e80565b905060c002016020016020810190611c1d9190613f80565b888886818110611c2f57611c2f613e80565b905060c0020160400135898987818110611c4b57611c4b613e80565b905060c002016060016020810190611c639190613f9b565b8a8a88818110611c7557611c75613e80565b905060c002016080016020810190611c8d9190613848565b8b8b89818110611c9f57611c9f613e80565b905060c0020160a0016020810190611cb79190613848565b604051611cca9796959493929190613eff565b60405180910390a180611cdc81613eac565b9150506118af565b5050505050565b606060038054610a3790613c6a565b6001600160a01b038116600090815260056020526040808220546001600160401b03911c166109a4565b611d2c6125ed565b600b80546001600160a01b0319166001600160a01b0383169081179091556040519081527ff477d93c015f2a73c2ccc5ed37078d12123b80fc5d12e0014c60b913bc1a1ec490602001610a1d565b6060818310611d9c57604051631960ccad60e11b815260040160405180910390fd5b600080611da860005490565b905080841115611db6578093505b6000611dc18761140c565b905084861015611de05785850381811015611dda578091505b50611de4565b5060005b6000816001600160401b03811115611dfe57611dfe613a4e565b604051908082528060200260200182016040528015611e27578160200160208202803683370190505b50905081600003611e3d579350610d2992505050565b6000611e48886122bf565b905060008160400151611e59575080515b885b888114158015611e6b5750848714155b15611ee057611e7981612d6b565b92508260400151611ed85782516001600160a01b031615611e9957825191505b8a6001600160a01b0316826001600160a01b031603611ed85780848880600101995081518110611ecb57611ecb613e80565b6020026020010181815250505b600101611e5b565b505050928352509095945050505050565b3360008181526007602090815260408083206001600160a01b03871680855290835292819020805460ff191686151590811790915590519081529192917f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a35050565b6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a081019190915260105460009081908410611fd45760405162461bcd60e51b815260206004820152600c60248201526b496e76616c6964537461676560a01b6044820152606401610e77565b60008481526011602090815260408083203384528252808320548784526012909252909120546010805463ffffffff909316928790811061201757612017613e80565b60009182526020918290206040805160c08101825260039390930290910180546001600160501b0381168452600160501b900463ffffffff169383019390935260018301549082015260029091015462ffffff81166060830152630100000081046001600160401b039081166080840152600160581b9091041660a0820152969195509350915050565b6120a96125ed565b600f6120b6828483613d92565b505050565b6120c36125ed565b8163ffffffff16600c54816120db6001546000540390565b6120e59190613fb6565b11156121045760405163800113cb60e01b815260040160405180910390fd5b6120b6828463ffffffff16612da7565b600a5461213c90600160501b90046001600160a01b0316612136868686610bf0565b83612dc1565b610e8b5760405162b7fad960e11b815260040160405180910390fd5b6060600f8054610a3790613c6a565b836daaeb6d7670e522a718067333cd4e3b156122b357336001600160a01b0382160361219e5761219985858585612f03565b611ce4565b604051633185c44d60e21b81523060048201523360248201526daaeb6d7670e522a718067333cd4e9063c617113490604401602060405180830381865afa1580156121ed573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122119190613d2f565b80156122945750604051633185c44d60e21b81523060048201526001600160a01b03821660248201526daaeb6d7670e522a718067333cd4e9063c617113490604401602060405180830381865afa158015612270573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122949190613d2f565b6122b357604051633b79c77360e21b8152336004820152602401610e77565b611ce485858585612f03565b60408051608080820183526000808352602080840182905283850182905260608085018390528551938401865282845290830182905293820181905292810183905290915060005483106123135792915050565b61231c83612d6b565b905080604001511561232e5792915050565b610d2983612f47565b606061234282612647565b61235f57604051630a14c4b560e41b815260040160405180910390fd5b6000600e805461236e90613c6a565b80601f016020809104026020016040519081016040528092919081815260200182805461239a90613c6a565b80156123e75780601f106123bc576101008083540402835291602001916123e7565b820191906000526020600020905b8154815290600101906020018083116123ca57829003601f168201915b50505050509050805160000361240c5760405180602001604052806000815250610d29565b8061241684612f7c565b600f60405160200161242a93929190613fc9565b6040516020818303038152906040529392505050565b6124486125ed565b600a805469ffffffffffffffff00001916620100006001600160401b038416908102919091179091556040519081527f41b9126ccd8cb4505310c40a376055b5ef246bd4c9214de02af31ef4f26b1b5f90602001610a1d565b6002600954036124f35760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610e77565b600260095561250786338787878787612822565b5050600160095550505050565b61251c6125ed565b6001600160a01b0381166125815760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610e77565b61258a81612ce3565b50565b6125956125ed565b600c548111156125b85760405163430b83b160e11b815260040160405180910390fd5b600c8190556040518181527fc7bbc2b288fc13314546ea4aa51f6bcf71b7ba4740beeb3d32e9acef57b6668a90602001610a1d565b6008546001600160a01b0316331461146c5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610e77565b60008054821080156109a4575050600090815260046020526040902054600160e01b161590565b600061267982612c7c565b9050836001600160a01b0316816001600160a01b0316146126ac5760405162a1148160e81b815260040160405180910390fd5b60008281526006602052604090208054338082146001600160a01b038816909114176126f9576126dc86336108aa565b6126f957604051632ce44b5f60e11b815260040160405180910390fd5b6001600160a01b03851661272057604051633a954ecd60e21b815260040160405180910390fd5b801561272b57600082555b6001600160a01b038681166000908152600560205260408082208054600019019055918716808252919020805460010190554260a01b17600160e11b17600085815260046020526040812091909155600160e11b841690036127bd576001840160008181526004602052604081205490036127bb5760005481146127bb5760008181526004602052604090208490555b505b83856001600160a01b0316876001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a45b505050505050565b6120b683838360405180602001604052806000815250612167565b600a5460ff1661284557604051630952c8a960e11b815260040160405180910390fd5b8663ffffffff16600c548161285d6001546000540390565b6128679190613fb6565b11156128865760405163800113cb60e01b815260040160405180910390fd5b426128bf6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a081019190915290565b600a54600160501b90046001600160a01b03161561292657612919338b8888888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061211492505050565b61292286612fc0565b8591505b600061293183611344565b90506010818154811061294657612946613e80565b60009182526020918290206040805160c081018252600390930290910180546001600160501b03811680855263ffffffff600160501b9092048216958501959095526001820154928401929092526002015462ffffff811660608401526001600160401b036301000000820481166080850152600160581b9091041660a08301529093506129d791908d1690614069565b6001600160501b0316341015612a0057604051630717c22560e51b815260040160405180910390fd5b606082015162ffffff1615612a5a57606082015160008281526012602052604090205462ffffff90911690612a3c9063ffffffff8e1690613fb6565b1115612a5a5760405162d0844960e21b815260040160405180910390fd5b600d5415612ac157600d548b63ffffffff16612a988c6001600160a01b03166000908152600560205260409081902054901c6001600160401b031690565b612aa29190613fb6565b1115612ac15760405163751304ed60e11b815260040160405180910390fd5b602082015163ffffffff1615612b335760208083015160008381526011835260408082206001600160a01b038f168352909352919091205463ffffffff91821691612b0e918e9116614098565b63ffffffff161115612b335760405163b4f3729b60e01b815260040160405180910390fd5b604082015115612bda578160400151612bbc8a8a8080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604051612ba192508f915060200160609190911b6bffffffffffffffffffffffff1916815260140190565b60405160208183030381529060405280519060200120613005565b14612bda576040516309bde33960e01b815260040160405180910390fd5b60008181526011602090815260408083206001600160a01b038e168452909152812080548d9290612c1290849063ffffffff16614098565b92506101000a81548163ffffffff021916908363ffffffff1602179055508a63ffffffff16601260008381526020019081526020016000206000828254612c599190613fb6565b90915550612c6f90508a63ffffffff8d16612da7565b5050505050505050505050565b600081600054811015612cca5760008181526004602052604081205490600160e01b82169003612cc8575b80600003610d29575060001901600081815260046020526040902054612ca7565b505b604051636f96cda160e11b815260040160405180910390fd5b600880546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b806001600160401b0316826001600160401b031610612d6757604051631750215560e11b815260040160405180910390fd5b5050565b6040805160808101825260008082526020820181905291810182905260608101919091526000828152600460205260409020546109a490613052565b612d67828260405180602001604052806000815250613099565b6000806000612dd085856130ff565b90925090506000816004811115612de957612de96140b5565b148015612e075750856001600160a01b0316826001600160a01b0316145b15612e1757600192505050610d29565b600080876001600160a01b0316631626ba7e60e01b8888604051602401612e3f9291906140cb565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909416939093179092529051612e7d91906140e4565b600060405180830381855afa9150503d8060008114612eb8576040519150601f19603f3d011682016040523d82523d6000602084013e612ebd565b606091505b5091509150818015612ed0575080516020145b8015612ef757508051630b135d3f60e11b90612ef59083016020908101908401614100565b145b98975050505050505050565b612f0e848484610d30565b6001600160a01b0383163b15610e8b57612f2a84848484613144565b610e8b576040516368d2bf6b60e11b815260040160405180910390fd5b6040805160808101825260008082526020820181905291810182905260608101919091526109a4612f7783612c7c565b613052565b606060a06040510180604052602081039150506000815280825b600183039250600a81066030018353600a900480612f965750819003601f19909101908152919050565b600a54612fdc906201000090046001600160401b031642613ec5565b816001600160401b0316101561258a576040516313634e8d60e11b815260040160405180910390fd5b600081815b845181101561304a576130368286838151811061302957613029613e80565b602002602001015161322f565b91508061304281613eac565b91505061300a565b509392505050565b604080516080810182526001600160a01b038316815260a083901c6001600160401b03166020820152600160e01b831615159181019190915260e89190911c606082015290565b6130a3838361325b565b6001600160a01b0383163b156120b6576000548281035b6130cd6000868380600101945086613144565b6130ea576040516368d2bf6b60e11b815260040160405180910390fd5b8181106130ba578160005414611ce457600080fd5b60008082516041036131355760208301516040840151606085015160001a61312987828585613359565b9450945050505061313d565b506000905060025b9250929050565b604051630a85bd0160e11b81526000906001600160a01b0385169063150b7a0290613179903390899088908890600401614119565b6020604051808303816000875af19250505080156131b4575060408051601f3d908101601f191682019092526131b191810190614156565b60015b613212573d8080156131e2576040519150601f19603f3d011682016040523d82523d6000602084013e6131e7565b606091505b50805160000361320a576040516368d2bf6b60e11b815260040160405180910390fd5b805181602001fd5b6001600160e01b031916630a85bd0160e11b149050949350505050565b600081831061324b576000828152602084905260409020610d29565b5060009182526020526040902090565b60008054908290036132805760405163b562e8dd60e01b815260040160405180910390fd5b6001600160a01b03831660008181526005602090815260408083208054680100000000000000018802019055848352600490915281206001851460e11b4260a01b178317905582840190839083907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8180a4600183015b81811461332f57808360007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef600080a46001016132f7565b508160000361335057604051622e076360e81b815260040160405180910390fd5b60005550505050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115613390575060009050600361343d565b8460ff16601b141580156133a857508460ff16601c14155b156133b9575060009050600461343d565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa15801561340d573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166134365760006001925092505061343d565b9150600090505b94509492505050565b6001600160e01b03198116811461258a57600080fd5b60006020828403121561346e57600080fd5b8135610d2981613446565b80356001600160a01b038116811461349057600080fd5b919050565b6000602082840312156134a757600080fd5b610d2982613479565b60005b838110156134cb5781810151838201526020016134b3565b50506000910152565b600081518084526134ec8160208601602086016134b0565b601f01601f19169290920160200192915050565b602081526000610d2960208301846134d4565b60006020828403121561352557600080fd5b5035919050565b6000806040838503121561353f57600080fd5b61354883613479565b946020939093013593505050565b803563ffffffff8116811461349057600080fd5b80356001600160401b038116811461349057600080fd5b60008060006060848603121561359657600080fd5b61359f84613479565b92506135ad60208501613556565b91506135bb6040850161356a565b90509250925092565b6000806000606084860312156135d957600080fd5b6135e284613479565b92506135f060208501613479565b9150604084013590509250925092565b801515811461258a57600080fd5b60006020828403121561362057600080fd5b8135610d2981613600565b60008083601f84011261363d57600080fd5b5081356001600160401b0381111561365457600080fd5b60208301915083602082850101111561313d57600080fd5b6000806020838503121561367f57600080fd5b82356001600160401b0381111561369557600080fd5b6136a18582860161362b565b90969095509350505050565b60008083601f8401126136bf57600080fd5b5081356001600160401b038111156136d657600080fd5b6020830191508360208260051b850101111561313d57600080fd5b6000806020838503121561370457600080fd5b82356001600160401b0381111561371a57600080fd5b6136a1858286016136ad565b80516001600160a01b031682526020808201516001600160401b03169083015260408082015115159083015260609081015162ffffff16910152565b6020808252825182820181905260009190848201906040850190845b8181101561180a57613791838551613726565b928401926080929092019160010161377e565b600080600080600080600060a0888a0312156137bf57600080fd5b6137c888613556565b96506137d660208901613479565b955060408801356001600160401b03808211156137f257600080fd5b6137fe8b838c016136ad565b909750955085915061381260608b0161356a565b945060808a013591508082111561382857600080fd5b506138358a828b0161362b565b989b979a50959850939692959293505050565b60006020828403121561385a57600080fd5b610d298261356a565b80356001600160501b038116811461349057600080fd5b803562ffffff8116811461349057600080fd5b600080600080600080600060e0888a0312156138a857600080fd5b873596506138b860208901613863565b95506138c660408901613556565b9450606088013593506138db6080890161387a565b92506138e960a0890161356a565b91506138f760c0890161356a565b905092959891949750929550565b6020808252825182820181905260009190848201906040850190845b8181101561180a57835183529284019291840191600101613921565b6000806020838503121561395057600080fd5b82356001600160401b038082111561396757600080fd5b818501915085601f83011261397b57600080fd5b81358181111561398a57600080fd5b86602060c08302850101111561399f57600080fd5b60209290920196919550909350505050565b6000806000606084860312156139c657600080fd5b6139cf84613479565b95602085013595506040909401359392505050565b600080604083850312156139f757600080fd5b613a0083613479565b91506020830135613a1081613600565b809150509250929050565b60008060408385031215613a2e57600080fd5b613a3783613556565b9150613a4560208401613479565b90509250929050565b634e487b7160e01b600052604160045260246000fd5b600082601f830112613a7557600080fd5b81356001600160401b0380821115613a8f57613a8f613a4e565b604051601f8301601f19908116603f01168101908282118183101715613ab757613ab7613a4e565b81604052838152866020858801011115613ad057600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008060008060808587031215613b0657600080fd5b613b0f85613479565b9350613b1d60208601613556565b9250613b2b6040860161356a565b915060608501356001600160401b03811115613b4657600080fd5b613b5287828801613a64565b91505092959194509250565b60008060008060808587031215613b7457600080fd5b613b7d85613479565b9350613b8b60208601613479565b92506040850135915060608501356001600160401b03811115613b4657600080fd5b608081016109a48284613726565b60008060408385031215613bce57600080fd5b613a3783613479565b60008060008060008060808789031215613bf057600080fd5b613bf987613556565b955060208701356001600160401b0380821115613c1557600080fd5b613c218a838b016136ad565b9097509550859150613c3560408a0161356a565b94506060890135915080821115613c4b57600080fd5b50613c5889828a0161362b565b979a9699509497509295939492505050565b600181811c90821680613c7e57607f821691505b602082108103613c9e57634e487b7160e01b600052602260045260246000fd5b50919050565b6000602080835260008454613cb881613c6a565b80848701526040600180841660008114613cd95760018114613cf357613d21565b60ff1985168984015283151560051b890183019550613d21565b896000528660002060005b85811015613d195781548b8201860152908301908801613cfe565b8a0184019650505b509398975050505050505050565b600060208284031215613d4157600080fd5b8151610d2981613600565b601f8211156120b657600081815260208120601f850160051c81016020861015613d735750805b601f850160051c820191505b818110156127ff57828155600101613d7f565b6001600160401b03831115613da957613da9613a4e565b613dbd83613db78354613c6a565b83613d4c565b6000601f841160018114613df15760008515613dd95750838201355b600019600387901b1c1916600186901b178355611ce4565b600083815260209020601f19861690835b82811015613e225786850135825560209485019460019092019101613e02565b5086821015613e3f5760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060018201613ebe57613ebe613e96565b5060010190565b818103818111156109a4576109a4613e96565b6001600160401b03818116838216019080821115613ef857613ef8613e96565b5092915050565b9687526001600160501b0395909516602087015263ffffffff939093166040860152606085019190915262ffffff1660808401526001600160401b0390811660a08401521660c082015260e00190565b634e487b7160e01b600052603160045260246000fd5b600060208284031215613f7757600080fd5b610d2982613863565b600060208284031215613f9257600080fd5b610d2982613556565b600060208284031215613fad57600080fd5b610d298261387a565b808201808211156109a4576109a4613e96565b600084516020613fdc8285838a016134b0565b855191840191613fef8184848a016134b0565b855492019160009061400081613c6a565b60018281168015614018576001811461402d57614059565b60ff1984168752821515830287019450614059565b896000528560002060005b8481101561405157815489820152908301908701614038565b505082870194505b50929a9950505050505050505050565b60006001600160501b038083168185168183048111821515161561408f5761408f613e96565b02949350505050565b63ffffffff818116838216019080821115613ef857613ef8613e96565b634e487b7160e01b600052602160045260246000fd5b828152604060208201526000610d2660408301846134d4565b600082516140f68184602087016134b0565b9190910192915050565b60006020828403121561411257600080fd5b5051919050565b6001600160a01b038581168252841660208201526040810183905260806060820181905260009061414c908301846134d4565b9695505050505050565b60006020828403121561416857600080fd5b8151610d298161344656fea2646970667358221220095950fdd01c5925958d1fb092647deb343994abad487af3a0dfe102b40a135264736f6c63430008100033",
}

// AlphegateABI is the input ABI used to generate the binding from.
// Deprecated: Use AlphegateMetaData.ABI instead.
var AlphegateABI = AlphegateMetaData.ABI

// AlphegateBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AlphegateMetaData.Bin instead.
var AlphegateBin = AlphegateMetaData.Bin

// DeployAlphegate deploys a new Ethereum contract, binding an instance of Alphegate to it.
func DeployAlphegate(auth *bind.TransactOpts, backend bind.ContractBackend, collectionName string, collectionSymbol string, tokenURISuffix string, maxMintableSupply *big.Int, globalWalletLimit *big.Int, cosigner common.Address, timestampExpirySeconds uint64) (common.Address, *types.Transaction, *Alphegate, error) {
	parsed, err := AlphegateMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AlphegateBin), backend, collectionName, collectionSymbol, tokenURISuffix, maxMintableSupply, globalWalletLimit, cosigner, timestampExpirySeconds)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Alphegate{AlphegateCaller: AlphegateCaller{contract: contract}, AlphegateTransactor: AlphegateTransactor{contract: contract}, AlphegateFilterer: AlphegateFilterer{contract: contract}}, nil
}

// Alphegate is an auto generated Go binding around an Ethereum contract.
type Alphegate struct {
	AlphegateCaller     // Read-only binding to the contract
	AlphegateTransactor // Write-only binding to the contract
	AlphegateFilterer   // Log filterer for contract events
}

// AlphegateCaller is an auto generated read-only Go binding around an Ethereum contract.
type AlphegateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AlphegateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AlphegateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AlphegateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AlphegateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AlphegateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AlphegateSession struct {
	Contract     *Alphegate        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AlphegateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AlphegateCallerSession struct {
	Contract *AlphegateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// AlphegateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AlphegateTransactorSession struct {
	Contract     *AlphegateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AlphegateRaw is an auto generated low-level Go binding around an Ethereum contract.
type AlphegateRaw struct {
	Contract *Alphegate // Generic contract binding to access the raw methods on
}

// AlphegateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AlphegateCallerRaw struct {
	Contract *AlphegateCaller // Generic read-only contract binding to access the raw methods on
}

// AlphegateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AlphegateTransactorRaw struct {
	Contract *AlphegateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAlphegate creates a new instance of Alphegate, bound to a specific deployed contract.
func NewAlphegate(address common.Address, backend bind.ContractBackend) (*Alphegate, error) {
	contract, err := bindAlphegate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Alphegate{AlphegateCaller: AlphegateCaller{contract: contract}, AlphegateTransactor: AlphegateTransactor{contract: contract}, AlphegateFilterer: AlphegateFilterer{contract: contract}}, nil
}

// NewAlphegateCaller creates a new read-only instance of Alphegate, bound to a specific deployed contract.
func NewAlphegateCaller(address common.Address, caller bind.ContractCaller) (*AlphegateCaller, error) {
	contract, err := bindAlphegate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AlphegateCaller{contract: contract}, nil
}

// NewAlphegateTransactor creates a new write-only instance of Alphegate, bound to a specific deployed contract.
func NewAlphegateTransactor(address common.Address, transactor bind.ContractTransactor) (*AlphegateTransactor, error) {
	contract, err := bindAlphegate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AlphegateTransactor{contract: contract}, nil
}

// NewAlphegateFilterer creates a new log filterer instance of Alphegate, bound to a specific deployed contract.
func NewAlphegateFilterer(address common.Address, filterer bind.ContractFilterer) (*AlphegateFilterer, error) {
	contract, err := bindAlphegate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AlphegateFilterer{contract: contract}, nil
}

// bindAlphegate binds a generic wrapper to an already deployed contract.
func bindAlphegate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AlphegateMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Alphegate *AlphegateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Alphegate.Contract.AlphegateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Alphegate *AlphegateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Alphegate.Contract.AlphegateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Alphegate *AlphegateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Alphegate.Contract.AlphegateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Alphegate *AlphegateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Alphegate.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Alphegate *AlphegateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Alphegate.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Alphegate *AlphegateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Alphegate.Contract.contract.Transact(opts, method, params...)
}

// AssertValidCosign is a free data retrieval call binding the contract method 0xb50248e7.
//
// Solidity: function assertValidCosign(address minter, uint32 qty, uint64 timestamp, bytes signature) view returns()
func (_Alphegate *AlphegateCaller) AssertValidCosign(opts *bind.CallOpts, minter common.Address, qty uint32, timestamp uint64, signature []byte) error {
	var out []interface{}
	err := _Alphegate.contract.Call(opts, &out, "assertValidCosign", minter, qty, timestamp, signature)

	if err != nil {
		return err
	}

	return err

}

// AssertValidCosign is a free data retrieval call binding the contract method 0xb50248e7.
//
// Solidity: function assertValidCosign(address minter, uint32 qty, uint64 timestamp, bytes signature) view returns()
func (_Alphegate *AlphegateSession) AssertValidCosign(minter common.Address, qty uint32, timestamp uint64, signature []byte) error {
	return _Alphegate.Contract.AssertValidCosign(&_Alphegate.CallOpts, minter, qty, timestamp, signature)
}

// AssertValidCosign is a free data retrieval call binding the contract method 0xb50248e7.
//
// Solidity: function assertValidCosign(address minter, uint32 qty, uint64 timestamp, bytes signature) view returns()
func (_Alphegate *AlphegateCallerSession) AssertValidCosign(minter common.Address, qty uint32, timestamp uint64, signature []byte) error {
	return _Alphegate.Contract.AssertValidCosign(&_Alphegate.CallOpts, minter, qty, timestamp, signature)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Alphegate *AlphegateCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Alphegate.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Alphegate *AlphegateSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Alphegate.Contract.BalanceOf(&_Alphegate.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Alphegate *AlphegateCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Alphegate.Contract.BalanceOf(&_Alphegate.CallOpts, owner)
}

// ExplicitOwnershipOf is a free data retrieval call binding the contract method 0xc23dc68f.
//
// Solidity: function explicitOwnershipOf(uint256 tokenId) view returns((address,uint64,bool,uint24))
func (_Alphegate *AlphegateCaller) ExplicitOwnershipOf(opts *bind.CallOpts, tokenId *big.Int) (IERC721ATokenOwnership, error) {
	var out []interface{}
	err := _Alphegate.contract.Call(opts, &out, "explicitOwnershipOf", tokenId)

	if err != nil {
		return *new(IERC721ATokenOwnership), err
	}

	out0 := *abi.ConvertType(out[0], new(IERC721ATokenOwnership)).(*IERC721ATokenOwnership)

	return out0, err

}

// ExplicitOwnershipOf is a free data retrieval call binding the contract method 0xc23dc68f.
//
// Solidity: function explicitOwnershipOf(uint256 tokenId) view returns((address,uint64,bool,uint24))
func (_Alphegate *AlphegateSession) ExplicitOwnershipOf(tokenId *big.Int) (IERC721ATokenOwnership, error) {
	return _Alphegate.Contract.ExplicitOwnershipOf(&_Alphegate.CallOpts, tokenId)
}

// ExplicitOwnershipOf is a free data retrieval call binding the contract method 0xc23dc68f.
//
// Solidity: function explicitOwnershipOf(uint256 tokenId) view returns((address,uint64,bool,uint24))
func (_Alphegate *AlphegateCallerSession) ExplicitOwnershipOf(tokenId *big.Int) (IERC721ATokenOwnership, error) {
	return _Alphegate.Contract.ExplicitOwnershipOf(&_Alphegate.CallOpts, tokenId)
}

// ExplicitOwnershipsOf is a free data retrieval call binding the contract method 0x5bbb2177.
//
// Solidity: function explicitOwnershipsOf(uint256[] tokenIds) view returns((address,uint64,bool,uint24)[])
func (_Alphegate *AlphegateCaller) ExplicitOwnershipsOf(opts *bind.CallOpts, tokenIds []*big.Int) ([]IERC721ATokenOwnership, error) {
	var out []interface{}
	err := _Alphegate.contract.Call(opts, &out, "explicitOwnershipsOf", tokenIds)

	if err != nil {
		return *new([]IERC721ATokenOwnership), err
	}

	out0 := *abi.ConvertType(out[0], new([]IERC721ATokenOwnership)).(*[]IERC721ATokenOwnership)

	return out0, err

}

// ExplicitOwnershipsOf is a free data retrieval call binding the contract method 0x5bbb2177.
//
// Solidity: function explicitOwnershipsOf(uint256[] tokenIds) view returns((address,uint64,bool,uint24)[])
func (_Alphegate *AlphegateSession) ExplicitOwnershipsOf(tokenIds []*big.Int) ([]IERC721ATokenOwnership, error) {
	return _Alphegate.Contract.ExplicitOwnershipsOf(&_Alphegate.CallOpts, tokenIds)
}

// ExplicitOwnershipsOf is a free data retrieval call binding the contract method 0x5bbb2177.
//
// Solidity: function explicitOwnershipsOf(uint256[] tokenIds) view returns((address,uint64,bool,uint24)[])
func (_Alphegate *AlphegateCallerSession) ExplicitOwnershipsOf(tokenIds []*big.Int) ([]IERC721ATokenOwnership, error) {
	return _Alphegate.Contract.ExplicitOwnershipsOf(&_Alphegate.CallOpts, tokenIds)
}

// GetActiveStageFromTimestamp is a free data retrieval call binding the contract method 0x67808a34.
//
// Solidity: function getActiveStageFromTimestamp(uint64 timestamp) view returns(uint256)
func (_Alphegate *AlphegateCaller) GetActiveStageFromTimestamp(opts *bind.CallOpts, timestamp uint64) (*big.Int, error) {
	var out []interface{}
	err := _Alphegate.contract.Call(opts, &out, "getActiveStageFromTimestamp", timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveStageFromTimestamp is a free data retrieval call binding the contract method 0x67808a34.
//
// Solidity: function getActiveStageFromTimestamp(uint64 timestamp) view returns(uint256)
func (_Alphegate *AlphegateSession) GetActiveStageFromTimestamp(timestamp uint64) (*big.Int, error) {
	return _Alphegate.Contract.GetActiveStageFromTimestamp(&_Alphegate.CallOpts, timestamp)
}

// GetActiveStageFromTimestamp is a free data retrieval call binding the contract method 0x67808a34.
//
// Solidity: function getActiveStageFromTimestamp(uint64 timestamp) view returns(uint256)
func (_Alphegate *AlphegateCallerSession) GetActiveStageFromTimestamp(timestamp uint64) (*big.Int, error) {
	return _Alphegate.Contract.GetActiveStageFromTimestamp(&_Alphegate.CallOpts, timestamp)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Alphegate *AlphegateCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Alphegate.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Alphegate *AlphegateSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Alphegate.Contract.GetApproved(&_Alphegate.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Alphegate *AlphegateCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Alphegate.Contract.GetApproved(&_Alphegate.CallOpts, tokenId)
}

// GetCosignDigest is a free data retrieval call binding the contract method 0x1ce03eed.
//
// Solidity: function getCosignDigest(address minter, uint32 qty, uint64 timestamp) view returns(bytes32)
func (_Alphegate *AlphegateCaller) GetCosignDigest(opts *bind.CallOpts, minter common.Address, qty uint32, timestamp uint64) ([32]byte, error) {
	var out []interface{}
	err := _Alphegate.contract.Call(opts, &out, "getCosignDigest", minter, qty, timestamp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCosignDigest is a free data retrieval call binding the contract method 0x1ce03eed.
//
// Solidity: function getCosignDigest(address minter, uint32 qty, uint64 timestamp) view returns(bytes32)
func (_Alphegate *AlphegateSession) GetCosignDigest(minter common.Address, qty uint32, timestamp uint64) ([32]byte, error) {
	return _Alphegate.Contract.GetCosignDigest(&_Alphegate.CallOpts, minter, qty, timestamp)
}

// GetCosignDigest is a free data retrieval call binding the contract method 0x1ce03eed.
//
// Solidity: function getCosignDigest(address minter, uint32 qty, uint64 timestamp) view returns(bytes32)
func (_Alphegate *AlphegateCallerSession) GetCosignDigest(minter common.Address, qty uint32, timestamp uint64) ([32]byte, error) {
	return _Alphegate.Contract.GetCosignDigest(&_Alphegate.CallOpts, minter, qty, timestamp)
}

// GetCosignNonce is a free data retrieval call binding the contract method 0xa06c492f.
//
// Solidity: function getCosignNonce(address minter) view returns(uint256)
func (_Alphegate *AlphegateCaller) GetCosignNonce(opts *bind.CallOpts, minter common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Alphegate.contract.Call(opts, &out, "getCosignNonce", minter)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCosignNonce is a free data retrieval call binding the contract method 0xa06c492f.
//
// Solidity: function getCosignNonce(address minter) view returns(uint256)
func (_Alphegate *AlphegateSession) GetCosignNonce(minter common.Address) (*big.Int, error) {
	return _Alphegate.Contract.GetCosignNonce(&_Alphegate.CallOpts, minter)
}

// GetCosignNonce is a free data retrieval call binding the contract method 0xa06c492f.
//
// Solidity: function getCosignNonce(address minter) view returns(uint256)
func (_Alphegate *AlphegateCallerSession) GetCosignNonce(minter common.Address) (*big.Int, error) {
	return _Alphegate.Contract.GetCosignNonce(&_Alphegate.CallOpts, minter)
}

// GetCosigner is a free data retrieval call binding the contract method 0x33bbbf06.
//
// Solidity: function getCosigner() view returns(address)
func (_Alphegate *AlphegateCaller) GetCosigner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Alphegate.contract.Call(opts, &out, "getCosigner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCosigner is a free data retrieval call binding the contract method 0x33bbbf06.
//
// Solidity: function getCosigner() view returns(address)
func (_Alphegate *AlphegateSession) GetCosigner() (common.Address, error) {
	return _Alphegate.Contract.GetCosigner(&_Alphegate.CallOpts)
}

// GetCosigner is a free data retrieval call binding the contract method 0x33bbbf06.
//
// Solidity: function getCosigner() view returns(address)
func (_Alphegate *AlphegateCallerSession) GetCosigner() (common.Address, error) {
	return _Alphegate.Contract.GetCosigner(&_Alphegate.CallOpts)
}

// GetCrossmintAddress is a free data retrieval call binding the contract method 0x8f931511.
//
// Solidity: function getCrossmintAddress() view returns(address)
func (_Alphegate *AlphegateCaller) GetCrossmintAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Alphegate.contract.Call(opts, &out, "getCrossmintAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCrossmintAddress is a free data retrieval call binding the contract method 0x8f931511.
//
// Solidity: function getCrossmintAddress() view returns(address)
func (_Alphegate *AlphegateSession) GetCrossmintAddress() (common.Address, error) {
	return _Alphegate.Contract.GetCrossmintAddress(&_Alphegate.CallOpts)
}

// GetCrossmintAddress is a free data retrieval call binding the contract method 0x8f931511.
//
// Solidity: function getCrossmintAddress() view returns(address)
func (_Alphegate *AlphegateCallerSession) GetCrossmintAddress() (common.Address, error) {
	return _Alphegate.Contract.GetCrossmintAddress(&_Alphegate.CallOpts)
}

// GetGlobalWalletLimit is a free data retrieval call binding the contract method 0xefdaa2ec.
//
// Solidity: function getGlobalWalletLimit() view returns(uint256)
func (_Alphegate *AlphegateCaller) GetGlobalWalletLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Alphegate.contract.Call(opts, &out, "getGlobalWalletLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGlobalWalletLimit is a free data retrieval call binding the contract method 0xefdaa2ec.
//
// Solidity: function getGlobalWalletLimit() view returns(uint256)
func (_Alphegate *AlphegateSession) GetGlobalWalletLimit() (*big.Int, error) {
	return _Alphegate.Contract.GetGlobalWalletLimit(&_Alphegate.CallOpts)
}

// GetGlobalWalletLimit is a free data retrieval call binding the contract method 0xefdaa2ec.
//
// Solidity: function getGlobalWalletLimit() view returns(uint256)
func (_Alphegate *AlphegateCallerSession) GetGlobalWalletLimit() (*big.Int, error) {
	return _Alphegate.Contract.GetGlobalWalletLimit(&_Alphegate.CallOpts)
}

// GetMaxMintableSupply is a free data retrieval call binding the contract method 0x4b1c53b4.
//
// Solidity: function getMaxMintableSupply() view returns(uint256)
func (_Alphegate *AlphegateCaller) GetMaxMintableSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Alphegate.contract.Call(opts, &out, "getMaxMintableSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxMintableSupply is a free data retrieval call binding the contract method 0x4b1c53b4.
//
// Solidity: function getMaxMintableSupply() view returns(uint256)
func (_Alphegate *AlphegateSession) GetMaxMintableSupply() (*big.Int, error) {
	return _Alphegate.Contract.GetMaxMintableSupply(&_Alphegate.CallOpts)
}

// GetMaxMintableSupply is a free data retrieval call binding the contract method 0x4b1c53b4.
//
// Solidity: function getMaxMintableSupply() view returns(uint256)
func (_Alphegate *AlphegateCallerSession) GetMaxMintableSupply() (*big.Int, error) {
	return _Alphegate.Contract.GetMaxMintableSupply(&_Alphegate.CallOpts)
}

// GetMintable is a free data retrieval call binding the contract method 0xf698bceb.
//
// Solidity: function getMintable() view returns(bool)
func (_Alphegate *AlphegateCaller) GetMintable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Alphegate.contract.Call(opts, &out, "getMintable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetMintable is a free data retrieval call binding the contract method 0xf698bceb.
//
// Solidity: function getMintable() view returns(bool)
func (_Alphegate *AlphegateSession) GetMintable() (bool, error) {
	return _Alphegate.Contract.GetMintable(&_Alphegate.CallOpts)
}

// GetMintable is a free data retrieval call binding the contract method 0xf698bceb.
//
// Solidity: function getMintable() view returns(bool)
func (_Alphegate *AlphegateCallerSession) GetMintable() (bool, error) {
	return _Alphegate.Contract.GetMintable(&_Alphegate.CallOpts)
}

// GetNumberStages is a free data retrieval call binding the contract method 0x70da24ee.
//
// Solidity: function getNumberStages() view returns(uint256)
func (_Alphegate *AlphegateCaller) GetNumberStages(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Alphegate.contract.Call(opts, &out, "getNumberStages")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumberStages is a free data retrieval call binding the contract method 0x70da24ee.
//
// Solidity: function getNumberStages() view returns(uint256)
func (_Alphegate *AlphegateSession) GetNumberStages() (*big.Int, error) {
	return _Alphegate.Contract.GetNumberStages(&_Alphegate.CallOpts)
}

// GetNumberStages is a free data retrieval call binding the contract method 0x70da24ee.
//
// Solidity: function getNumberStages() view returns(uint256)
func (_Alphegate *AlphegateCallerSession) GetNumberStages() (*big.Int, error) {
	return _Alphegate.Contract.GetNumberStages(&_Alphegate.CallOpts)
}

// GetStageInfo is a free data retrieval call binding the contract method 0xa3759f60.
//
// Solidity: function getStageInfo(uint256 index) view returns((uint80,uint32,bytes32,uint24,uint64,uint64), uint32, uint256)
func (_Alphegate *AlphegateCaller) GetStageInfo(opts *bind.CallOpts, index *big.Int) (IERC721MMintStageInfo, uint32, *big.Int, error) {
	var out []interface{}
	err := _Alphegate.contract.Call(opts, &out, "getStageInfo", index)

	if err != nil {
		return *new(IERC721MMintStageInfo), *new(uint32), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(IERC721MMintStageInfo)).(*IERC721MMintStageInfo)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetStageInfo is a free data retrieval call binding the contract method 0xa3759f60.
//
// Solidity: function getStageInfo(uint256 index) view returns((uint80,uint32,bytes32,uint24,uint64,uint64), uint32, uint256)
func (_Alphegate *AlphegateSession) GetStageInfo(index *big.Int) (IERC721MMintStageInfo, uint32, *big.Int, error) {
	return _Alphegate.Contract.GetStageInfo(&_Alphegate.CallOpts, index)
}

// GetStageInfo is a free data retrieval call binding the contract method 0xa3759f60.
//
// Solidity: function getStageInfo(uint256 index) view returns((uint80,uint32,bytes32,uint24,uint64,uint64), uint32, uint256)
func (_Alphegate *AlphegateCallerSession) GetStageInfo(index *big.Int) (IERC721MMintStageInfo, uint32, *big.Int, error) {
	return _Alphegate.Contract.GetStageInfo(&_Alphegate.CallOpts, index)
}

// GetTimestampExpirySeconds is a free data retrieval call binding the contract method 0x4ae0402f.
//
// Solidity: function getTimestampExpirySeconds() view returns(uint64)
func (_Alphegate *AlphegateCaller) GetTimestampExpirySeconds(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Alphegate.contract.Call(opts, &out, "getTimestampExpirySeconds")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetTimestampExpirySeconds is a free data retrieval call binding the contract method 0x4ae0402f.
//
// Solidity: function getTimestampExpirySeconds() view returns(uint64)
func (_Alphegate *AlphegateSession) GetTimestampExpirySeconds() (uint64, error) {
	return _Alphegate.Contract.GetTimestampExpirySeconds(&_Alphegate.CallOpts)
}

// GetTimestampExpirySeconds is a free data retrieval call binding the contract method 0x4ae0402f.
//
// Solidity: function getTimestampExpirySeconds() view returns(uint64)
func (_Alphegate *AlphegateCallerSession) GetTimestampExpirySeconds() (uint64, error) {
	return _Alphegate.Contract.GetTimestampExpirySeconds(&_Alphegate.CallOpts)
}

// GetTokenURISuffix is a free data retrieval call binding the contract method 0xb7a9fa60.
//
// Solidity: function getTokenURISuffix() view returns(string)
func (_Alphegate *AlphegateCaller) GetTokenURISuffix(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Alphegate.contract.Call(opts, &out, "getTokenURISuffix")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetTokenURISuffix is a free data retrieval call binding the contract method 0xb7a9fa60.
//
// Solidity: function getTokenURISuffix() view returns(string)
func (_Alphegate *AlphegateSession) GetTokenURISuffix() (string, error) {
	return _Alphegate.Contract.GetTokenURISuffix(&_Alphegate.CallOpts)
}

// GetTokenURISuffix is a free data retrieval call binding the contract method 0xb7a9fa60.
//
// Solidity: function getTokenURISuffix() view returns(string)
func (_Alphegate *AlphegateCallerSession) GetTokenURISuffix() (string, error) {
	return _Alphegate.Contract.GetTokenURISuffix(&_Alphegate.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Alphegate *AlphegateCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Alphegate.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Alphegate *AlphegateSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Alphegate.Contract.IsApprovedForAll(&_Alphegate.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Alphegate *AlphegateCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Alphegate.Contract.IsApprovedForAll(&_Alphegate.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Alphegate *AlphegateCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Alphegate.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Alphegate *AlphegateSession) Name() (string, error) {
	return _Alphegate.Contract.Name(&_Alphegate.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Alphegate *AlphegateCallerSession) Name() (string, error) {
	return _Alphegate.Contract.Name(&_Alphegate.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Alphegate *AlphegateCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Alphegate.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Alphegate *AlphegateSession) Owner() (common.Address, error) {
	return _Alphegate.Contract.Owner(&_Alphegate.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Alphegate *AlphegateCallerSession) Owner() (common.Address, error) {
	return _Alphegate.Contract.Owner(&_Alphegate.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Alphegate *AlphegateCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Alphegate.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Alphegate *AlphegateSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Alphegate.Contract.OwnerOf(&_Alphegate.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Alphegate *AlphegateCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Alphegate.Contract.OwnerOf(&_Alphegate.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Alphegate *AlphegateCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Alphegate.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Alphegate *AlphegateSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Alphegate.Contract.SupportsInterface(&_Alphegate.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Alphegate *AlphegateCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Alphegate.Contract.SupportsInterface(&_Alphegate.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Alphegate *AlphegateCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Alphegate.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Alphegate *AlphegateSession) Symbol() (string, error) {
	return _Alphegate.Contract.Symbol(&_Alphegate.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Alphegate *AlphegateCallerSession) Symbol() (string, error) {
	return _Alphegate.Contract.Symbol(&_Alphegate.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Alphegate *AlphegateCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Alphegate.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Alphegate *AlphegateSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Alphegate.Contract.TokenURI(&_Alphegate.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Alphegate *AlphegateCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Alphegate.Contract.TokenURI(&_Alphegate.CallOpts, tokenId)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_Alphegate *AlphegateCaller) TokensOfOwner(opts *bind.CallOpts, owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Alphegate.contract.Call(opts, &out, "tokensOfOwner", owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_Alphegate *AlphegateSession) TokensOfOwner(owner common.Address) ([]*big.Int, error) {
	return _Alphegate.Contract.TokensOfOwner(&_Alphegate.CallOpts, owner)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_Alphegate *AlphegateCallerSession) TokensOfOwner(owner common.Address) ([]*big.Int, error) {
	return _Alphegate.Contract.TokensOfOwner(&_Alphegate.CallOpts, owner)
}

// TokensOfOwnerIn is a free data retrieval call binding the contract method 0x99a2557a.
//
// Solidity: function tokensOfOwnerIn(address owner, uint256 start, uint256 stop) view returns(uint256[])
func (_Alphegate *AlphegateCaller) TokensOfOwnerIn(opts *bind.CallOpts, owner common.Address, start *big.Int, stop *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Alphegate.contract.Call(opts, &out, "tokensOfOwnerIn", owner, start, stop)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokensOfOwnerIn is a free data retrieval call binding the contract method 0x99a2557a.
//
// Solidity: function tokensOfOwnerIn(address owner, uint256 start, uint256 stop) view returns(uint256[])
func (_Alphegate *AlphegateSession) TokensOfOwnerIn(owner common.Address, start *big.Int, stop *big.Int) ([]*big.Int, error) {
	return _Alphegate.Contract.TokensOfOwnerIn(&_Alphegate.CallOpts, owner, start, stop)
}

// TokensOfOwnerIn is a free data retrieval call binding the contract method 0x99a2557a.
//
// Solidity: function tokensOfOwnerIn(address owner, uint256 start, uint256 stop) view returns(uint256[])
func (_Alphegate *AlphegateCallerSession) TokensOfOwnerIn(owner common.Address, start *big.Int, stop *big.Int) ([]*big.Int, error) {
	return _Alphegate.Contract.TokensOfOwnerIn(&_Alphegate.CallOpts, owner, start, stop)
}

// TotalMintedByAddress is a free data retrieval call binding the contract method 0x97cf84fc.
//
// Solidity: function totalMintedByAddress(address a) view returns(uint256)
func (_Alphegate *AlphegateCaller) TotalMintedByAddress(opts *bind.CallOpts, a common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Alphegate.contract.Call(opts, &out, "totalMintedByAddress", a)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalMintedByAddress is a free data retrieval call binding the contract method 0x97cf84fc.
//
// Solidity: function totalMintedByAddress(address a) view returns(uint256)
func (_Alphegate *AlphegateSession) TotalMintedByAddress(a common.Address) (*big.Int, error) {
	return _Alphegate.Contract.TotalMintedByAddress(&_Alphegate.CallOpts, a)
}

// TotalMintedByAddress is a free data retrieval call binding the contract method 0x97cf84fc.
//
// Solidity: function totalMintedByAddress(address a) view returns(uint256)
func (_Alphegate *AlphegateCallerSession) TotalMintedByAddress(a common.Address) (*big.Int, error) {
	return _Alphegate.Contract.TotalMintedByAddress(&_Alphegate.CallOpts, a)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Alphegate *AlphegateCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Alphegate.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Alphegate *AlphegateSession) TotalSupply() (*big.Int, error) {
	return _Alphegate.Contract.TotalSupply(&_Alphegate.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Alphegate *AlphegateCallerSession) TotalSupply() (*big.Int, error) {
	return _Alphegate.Contract.TotalSupply(&_Alphegate.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_Alphegate *AlphegateTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Alphegate.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_Alphegate *AlphegateSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Alphegate.Contract.Approve(&_Alphegate.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_Alphegate *AlphegateTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Alphegate.Contract.Approve(&_Alphegate.TransactOpts, to, tokenId)
}

// Crossmint is a paid mutator transaction binding the contract method 0x62acbd9a.
//
// Solidity: function crossmint(uint32 qty, address to, bytes32[] proof, uint64 timestamp, bytes signature) payable returns()
func (_Alphegate *AlphegateTransactor) Crossmint(opts *bind.TransactOpts, qty uint32, to common.Address, proof [][32]byte, timestamp uint64, signature []byte) (*types.Transaction, error) {
	return _Alphegate.contract.Transact(opts, "crossmint", qty, to, proof, timestamp, signature)
}

// Crossmint is a paid mutator transaction binding the contract method 0x62acbd9a.
//
// Solidity: function crossmint(uint32 qty, address to, bytes32[] proof, uint64 timestamp, bytes signature) payable returns()
func (_Alphegate *AlphegateSession) Crossmint(qty uint32, to common.Address, proof [][32]byte, timestamp uint64, signature []byte) (*types.Transaction, error) {
	return _Alphegate.Contract.Crossmint(&_Alphegate.TransactOpts, qty, to, proof, timestamp, signature)
}

// Crossmint is a paid mutator transaction binding the contract method 0x62acbd9a.
//
// Solidity: function crossmint(uint32 qty, address to, bytes32[] proof, uint64 timestamp, bytes signature) payable returns()
func (_Alphegate *AlphegateTransactorSession) Crossmint(qty uint32, to common.Address, proof [][32]byte, timestamp uint64, signature []byte) (*types.Transaction, error) {
	return _Alphegate.Contract.Crossmint(&_Alphegate.TransactOpts, qty, to, proof, timestamp, signature)
}

// Mint is a paid mutator transaction binding the contract method 0xefb6b11f.
//
// Solidity: function mint(uint32 qty, bytes32[] proof, uint64 timestamp, bytes signature) payable returns()
func (_Alphegate *AlphegateTransactor) Mint(opts *bind.TransactOpts, qty uint32, proof [][32]byte, timestamp uint64, signature []byte) (*types.Transaction, error) {
	return _Alphegate.contract.Transact(opts, "mint", qty, proof, timestamp, signature)
}

// Mint is a paid mutator transaction binding the contract method 0xefb6b11f.
//
// Solidity: function mint(uint32 qty, bytes32[] proof, uint64 timestamp, bytes signature) payable returns()
func (_Alphegate *AlphegateSession) Mint(qty uint32, proof [][32]byte, timestamp uint64, signature []byte) (*types.Transaction, error) {
	return _Alphegate.Contract.Mint(&_Alphegate.TransactOpts, qty, proof, timestamp, signature)
}

// Mint is a paid mutator transaction binding the contract method 0xefb6b11f.
//
// Solidity: function mint(uint32 qty, bytes32[] proof, uint64 timestamp, bytes signature) payable returns()
func (_Alphegate *AlphegateTransactorSession) Mint(qty uint32, proof [][32]byte, timestamp uint64, signature []byte) (*types.Transaction, error) {
	return _Alphegate.Contract.Mint(&_Alphegate.TransactOpts, qty, proof, timestamp, signature)
}

// OwnerMint is a paid mutator transaction binding the contract method 0xaac5ab1f.
//
// Solidity: function ownerMint(uint32 qty, address to) returns()
func (_Alphegate *AlphegateTransactor) OwnerMint(opts *bind.TransactOpts, qty uint32, to common.Address) (*types.Transaction, error) {
	return _Alphegate.contract.Transact(opts, "ownerMint", qty, to)
}

// OwnerMint is a paid mutator transaction binding the contract method 0xaac5ab1f.
//
// Solidity: function ownerMint(uint32 qty, address to) returns()
func (_Alphegate *AlphegateSession) OwnerMint(qty uint32, to common.Address) (*types.Transaction, error) {
	return _Alphegate.Contract.OwnerMint(&_Alphegate.TransactOpts, qty, to)
}

// OwnerMint is a paid mutator transaction binding the contract method 0xaac5ab1f.
//
// Solidity: function ownerMint(uint32 qty, address to) returns()
func (_Alphegate *AlphegateTransactorSession) OwnerMint(qty uint32, to common.Address) (*types.Transaction, error) {
	return _Alphegate.Contract.OwnerMint(&_Alphegate.TransactOpts, qty, to)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Alphegate *AlphegateTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Alphegate.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Alphegate *AlphegateSession) RenounceOwnership() (*types.Transaction, error) {
	return _Alphegate.Contract.RenounceOwnership(&_Alphegate.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Alphegate *AlphegateTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Alphegate.Contract.RenounceOwnership(&_Alphegate.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Alphegate *AlphegateTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Alphegate.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Alphegate *AlphegateSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Alphegate.Contract.SafeTransferFrom(&_Alphegate.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Alphegate *AlphegateTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Alphegate.Contract.SafeTransferFrom(&_Alphegate.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) payable returns()
func (_Alphegate *AlphegateTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Alphegate.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) payable returns()
func (_Alphegate *AlphegateSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Alphegate.Contract.SafeTransferFrom0(&_Alphegate.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) payable returns()
func (_Alphegate *AlphegateTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Alphegate.Contract.SafeTransferFrom0(&_Alphegate.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Alphegate *AlphegateTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Alphegate.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Alphegate *AlphegateSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Alphegate.Contract.SetApprovalForAll(&_Alphegate.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Alphegate *AlphegateTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Alphegate.Contract.SetApprovalForAll(&_Alphegate.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_Alphegate *AlphegateTransactor) SetBaseURI(opts *bind.TransactOpts, baseURI string) (*types.Transaction, error) {
	return _Alphegate.contract.Transact(opts, "setBaseURI", baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_Alphegate *AlphegateSession) SetBaseURI(baseURI string) (*types.Transaction, error) {
	return _Alphegate.Contract.SetBaseURI(&_Alphegate.TransactOpts, baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_Alphegate *AlphegateTransactorSession) SetBaseURI(baseURI string) (*types.Transaction, error) {
	return _Alphegate.Contract.SetBaseURI(&_Alphegate.TransactOpts, baseURI)
}

// SetBaseURIPermanent is a paid mutator transaction binding the contract method 0x1053a815.
//
// Solidity: function setBaseURIPermanent() returns()
func (_Alphegate *AlphegateTransactor) SetBaseURIPermanent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Alphegate.contract.Transact(opts, "setBaseURIPermanent")
}

// SetBaseURIPermanent is a paid mutator transaction binding the contract method 0x1053a815.
//
// Solidity: function setBaseURIPermanent() returns()
func (_Alphegate *AlphegateSession) SetBaseURIPermanent() (*types.Transaction, error) {
	return _Alphegate.Contract.SetBaseURIPermanent(&_Alphegate.TransactOpts)
}

// SetBaseURIPermanent is a paid mutator transaction binding the contract method 0x1053a815.
//
// Solidity: function setBaseURIPermanent() returns()
func (_Alphegate *AlphegateTransactorSession) SetBaseURIPermanent() (*types.Transaction, error) {
	return _Alphegate.Contract.SetBaseURIPermanent(&_Alphegate.TransactOpts)
}

// SetCosigner is a paid mutator transaction binding the contract method 0x02045138.
//
// Solidity: function setCosigner(address cosigner) returns()
func (_Alphegate *AlphegateTransactor) SetCosigner(opts *bind.TransactOpts, cosigner common.Address) (*types.Transaction, error) {
	return _Alphegate.contract.Transact(opts, "setCosigner", cosigner)
}

// SetCosigner is a paid mutator transaction binding the contract method 0x02045138.
//
// Solidity: function setCosigner(address cosigner) returns()
func (_Alphegate *AlphegateSession) SetCosigner(cosigner common.Address) (*types.Transaction, error) {
	return _Alphegate.Contract.SetCosigner(&_Alphegate.TransactOpts, cosigner)
}

// SetCosigner is a paid mutator transaction binding the contract method 0x02045138.
//
// Solidity: function setCosigner(address cosigner) returns()
func (_Alphegate *AlphegateTransactorSession) SetCosigner(cosigner common.Address) (*types.Transaction, error) {
	return _Alphegate.Contract.SetCosigner(&_Alphegate.TransactOpts, cosigner)
}

// SetCrossmintAddress is a paid mutator transaction binding the contract method 0x99755624.
//
// Solidity: function setCrossmintAddress(address crossmintAddress) returns()
func (_Alphegate *AlphegateTransactor) SetCrossmintAddress(opts *bind.TransactOpts, crossmintAddress common.Address) (*types.Transaction, error) {
	return _Alphegate.contract.Transact(opts, "setCrossmintAddress", crossmintAddress)
}

// SetCrossmintAddress is a paid mutator transaction binding the contract method 0x99755624.
//
// Solidity: function setCrossmintAddress(address crossmintAddress) returns()
func (_Alphegate *AlphegateSession) SetCrossmintAddress(crossmintAddress common.Address) (*types.Transaction, error) {
	return _Alphegate.Contract.SetCrossmintAddress(&_Alphegate.TransactOpts, crossmintAddress)
}

// SetCrossmintAddress is a paid mutator transaction binding the contract method 0x99755624.
//
// Solidity: function setCrossmintAddress(address crossmintAddress) returns()
func (_Alphegate *AlphegateTransactorSession) SetCrossmintAddress(crossmintAddress common.Address) (*types.Transaction, error) {
	return _Alphegate.Contract.SetCrossmintAddress(&_Alphegate.TransactOpts, crossmintAddress)
}

// SetGlobalWalletLimit is a paid mutator transaction binding the contract method 0x372992e4.
//
// Solidity: function setGlobalWalletLimit(uint256 globalWalletLimit) returns()
func (_Alphegate *AlphegateTransactor) SetGlobalWalletLimit(opts *bind.TransactOpts, globalWalletLimit *big.Int) (*types.Transaction, error) {
	return _Alphegate.contract.Transact(opts, "setGlobalWalletLimit", globalWalletLimit)
}

// SetGlobalWalletLimit is a paid mutator transaction binding the contract method 0x372992e4.
//
// Solidity: function setGlobalWalletLimit(uint256 globalWalletLimit) returns()
func (_Alphegate *AlphegateSession) SetGlobalWalletLimit(globalWalletLimit *big.Int) (*types.Transaction, error) {
	return _Alphegate.Contract.SetGlobalWalletLimit(&_Alphegate.TransactOpts, globalWalletLimit)
}

// SetGlobalWalletLimit is a paid mutator transaction binding the contract method 0x372992e4.
//
// Solidity: function setGlobalWalletLimit(uint256 globalWalletLimit) returns()
func (_Alphegate *AlphegateTransactorSession) SetGlobalWalletLimit(globalWalletLimit *big.Int) (*types.Transaction, error) {
	return _Alphegate.Contract.SetGlobalWalletLimit(&_Alphegate.TransactOpts, globalWalletLimit)
}

// SetMaxMintableSupply is a paid mutator transaction binding the contract method 0xf8d09696.
//
// Solidity: function setMaxMintableSupply(uint256 maxMintableSupply) returns()
func (_Alphegate *AlphegateTransactor) SetMaxMintableSupply(opts *bind.TransactOpts, maxMintableSupply *big.Int) (*types.Transaction, error) {
	return _Alphegate.contract.Transact(opts, "setMaxMintableSupply", maxMintableSupply)
}

// SetMaxMintableSupply is a paid mutator transaction binding the contract method 0xf8d09696.
//
// Solidity: function setMaxMintableSupply(uint256 maxMintableSupply) returns()
func (_Alphegate *AlphegateSession) SetMaxMintableSupply(maxMintableSupply *big.Int) (*types.Transaction, error) {
	return _Alphegate.Contract.SetMaxMintableSupply(&_Alphegate.TransactOpts, maxMintableSupply)
}

// SetMaxMintableSupply is a paid mutator transaction binding the contract method 0xf8d09696.
//
// Solidity: function setMaxMintableSupply(uint256 maxMintableSupply) returns()
func (_Alphegate *AlphegateTransactorSession) SetMaxMintableSupply(maxMintableSupply *big.Int) (*types.Transaction, error) {
	return _Alphegate.Contract.SetMaxMintableSupply(&_Alphegate.TransactOpts, maxMintableSupply)
}

// SetMintable is a paid mutator transaction binding the contract method 0x285d70d4.
//
// Solidity: function setMintable(bool mintable) returns()
func (_Alphegate *AlphegateTransactor) SetMintable(opts *bind.TransactOpts, mintable bool) (*types.Transaction, error) {
	return _Alphegate.contract.Transact(opts, "setMintable", mintable)
}

// SetMintable is a paid mutator transaction binding the contract method 0x285d70d4.
//
// Solidity: function setMintable(bool mintable) returns()
func (_Alphegate *AlphegateSession) SetMintable(mintable bool) (*types.Transaction, error) {
	return _Alphegate.Contract.SetMintable(&_Alphegate.TransactOpts, mintable)
}

// SetMintable is a paid mutator transaction binding the contract method 0x285d70d4.
//
// Solidity: function setMintable(bool mintable) returns()
func (_Alphegate *AlphegateTransactorSession) SetMintable(mintable bool) (*types.Transaction, error) {
	return _Alphegate.Contract.SetMintable(&_Alphegate.TransactOpts, mintable)
}

// SetStages is a paid mutator transaction binding the contract method 0x8dcdb09d.
//
// Solidity: function setStages((uint80,uint32,bytes32,uint24,uint64,uint64)[] newStages) returns()
func (_Alphegate *AlphegateTransactor) SetStages(opts *bind.TransactOpts, newStages []IERC721MMintStageInfo) (*types.Transaction, error) {
	return _Alphegate.contract.Transact(opts, "setStages", newStages)
}

// SetStages is a paid mutator transaction binding the contract method 0x8dcdb09d.
//
// Solidity: function setStages((uint80,uint32,bytes32,uint24,uint64,uint64)[] newStages) returns()
func (_Alphegate *AlphegateSession) SetStages(newStages []IERC721MMintStageInfo) (*types.Transaction, error) {
	return _Alphegate.Contract.SetStages(&_Alphegate.TransactOpts, newStages)
}

// SetStages is a paid mutator transaction binding the contract method 0x8dcdb09d.
//
// Solidity: function setStages((uint80,uint32,bytes32,uint24,uint64,uint64)[] newStages) returns()
func (_Alphegate *AlphegateTransactorSession) SetStages(newStages []IERC721MMintStageInfo) (*types.Transaction, error) {
	return _Alphegate.Contract.SetStages(&_Alphegate.TransactOpts, newStages)
}

// SetTimestampExpirySeconds is a paid mutator transaction binding the contract method 0xce2b0ec0.
//
// Solidity: function setTimestampExpirySeconds(uint64 expiry) returns()
func (_Alphegate *AlphegateTransactor) SetTimestampExpirySeconds(opts *bind.TransactOpts, expiry uint64) (*types.Transaction, error) {
	return _Alphegate.contract.Transact(opts, "setTimestampExpirySeconds", expiry)
}

// SetTimestampExpirySeconds is a paid mutator transaction binding the contract method 0xce2b0ec0.
//
// Solidity: function setTimestampExpirySeconds(uint64 expiry) returns()
func (_Alphegate *AlphegateSession) SetTimestampExpirySeconds(expiry uint64) (*types.Transaction, error) {
	return _Alphegate.Contract.SetTimestampExpirySeconds(&_Alphegate.TransactOpts, expiry)
}

// SetTimestampExpirySeconds is a paid mutator transaction binding the contract method 0xce2b0ec0.
//
// Solidity: function setTimestampExpirySeconds(uint64 expiry) returns()
func (_Alphegate *AlphegateTransactorSession) SetTimestampExpirySeconds(expiry uint64) (*types.Transaction, error) {
	return _Alphegate.Contract.SetTimestampExpirySeconds(&_Alphegate.TransactOpts, expiry)
}

// SetTokenURISuffix is a paid mutator transaction binding the contract method 0xa9852bfb.
//
// Solidity: function setTokenURISuffix(string suffix) returns()
func (_Alphegate *AlphegateTransactor) SetTokenURISuffix(opts *bind.TransactOpts, suffix string) (*types.Transaction, error) {
	return _Alphegate.contract.Transact(opts, "setTokenURISuffix", suffix)
}

// SetTokenURISuffix is a paid mutator transaction binding the contract method 0xa9852bfb.
//
// Solidity: function setTokenURISuffix(string suffix) returns()
func (_Alphegate *AlphegateSession) SetTokenURISuffix(suffix string) (*types.Transaction, error) {
	return _Alphegate.Contract.SetTokenURISuffix(&_Alphegate.TransactOpts, suffix)
}

// SetTokenURISuffix is a paid mutator transaction binding the contract method 0xa9852bfb.
//
// Solidity: function setTokenURISuffix(string suffix) returns()
func (_Alphegate *AlphegateTransactorSession) SetTokenURISuffix(suffix string) (*types.Transaction, error) {
	return _Alphegate.Contract.SetTokenURISuffix(&_Alphegate.TransactOpts, suffix)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Alphegate *AlphegateTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Alphegate.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Alphegate *AlphegateSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Alphegate.Contract.TransferFrom(&_Alphegate.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Alphegate *AlphegateTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Alphegate.Contract.TransferFrom(&_Alphegate.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Alphegate *AlphegateTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Alphegate.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Alphegate *AlphegateSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Alphegate.Contract.TransferOwnership(&_Alphegate.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Alphegate *AlphegateTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Alphegate.Contract.TransferOwnership(&_Alphegate.TransactOpts, newOwner)
}

// UpdateStage is a paid mutator transaction binding the contract method 0x73e1607e.
//
// Solidity: function updateStage(uint256 index, uint80 price, uint32 walletLimit, bytes32 merkleRoot, uint24 maxStageSupply, uint64 startTimeUnixSeconds, uint64 endTimeUnixSeconds) returns()
func (_Alphegate *AlphegateTransactor) UpdateStage(opts *bind.TransactOpts, index *big.Int, price *big.Int, walletLimit uint32, merkleRoot [32]byte, maxStageSupply *big.Int, startTimeUnixSeconds uint64, endTimeUnixSeconds uint64) (*types.Transaction, error) {
	return _Alphegate.contract.Transact(opts, "updateStage", index, price, walletLimit, merkleRoot, maxStageSupply, startTimeUnixSeconds, endTimeUnixSeconds)
}

// UpdateStage is a paid mutator transaction binding the contract method 0x73e1607e.
//
// Solidity: function updateStage(uint256 index, uint80 price, uint32 walletLimit, bytes32 merkleRoot, uint24 maxStageSupply, uint64 startTimeUnixSeconds, uint64 endTimeUnixSeconds) returns()
func (_Alphegate *AlphegateSession) UpdateStage(index *big.Int, price *big.Int, walletLimit uint32, merkleRoot [32]byte, maxStageSupply *big.Int, startTimeUnixSeconds uint64, endTimeUnixSeconds uint64) (*types.Transaction, error) {
	return _Alphegate.Contract.UpdateStage(&_Alphegate.TransactOpts, index, price, walletLimit, merkleRoot, maxStageSupply, startTimeUnixSeconds, endTimeUnixSeconds)
}

// UpdateStage is a paid mutator transaction binding the contract method 0x73e1607e.
//
// Solidity: function updateStage(uint256 index, uint80 price, uint32 walletLimit, bytes32 merkleRoot, uint24 maxStageSupply, uint64 startTimeUnixSeconds, uint64 endTimeUnixSeconds) returns()
func (_Alphegate *AlphegateTransactorSession) UpdateStage(index *big.Int, price *big.Int, walletLimit uint32, merkleRoot [32]byte, maxStageSupply *big.Int, startTimeUnixSeconds uint64, endTimeUnixSeconds uint64) (*types.Transaction, error) {
	return _Alphegate.Contract.UpdateStage(&_Alphegate.TransactOpts, index, price, walletLimit, merkleRoot, maxStageSupply, startTimeUnixSeconds, endTimeUnixSeconds)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Alphegate *AlphegateTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Alphegate.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Alphegate *AlphegateSession) Withdraw() (*types.Transaction, error) {
	return _Alphegate.Contract.Withdraw(&_Alphegate.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Alphegate *AlphegateTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Alphegate.Contract.Withdraw(&_Alphegate.TransactOpts)
}

// AlphegateApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Alphegate contract.
type AlphegateApprovalIterator struct {
	Event *AlphegateApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AlphegateApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphegateApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AlphegateApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AlphegateApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphegateApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphegateApproval represents a Approval event raised by the Alphegate contract.
type AlphegateApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Alphegate *AlphegateFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*AlphegateApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Alphegate.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &AlphegateApprovalIterator{contract: _Alphegate.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Alphegate *AlphegateFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AlphegateApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Alphegate.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphegateApproval)
				if err := _Alphegate.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Alphegate *AlphegateFilterer) ParseApproval(log types.Log) (*AlphegateApproval, error) {
	event := new(AlphegateApproval)
	if err := _Alphegate.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlphegateApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Alphegate contract.
type AlphegateApprovalForAllIterator struct {
	Event *AlphegateApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AlphegateApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphegateApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AlphegateApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AlphegateApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphegateApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphegateApprovalForAll represents a ApprovalForAll event raised by the Alphegate contract.
type AlphegateApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Alphegate *AlphegateFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*AlphegateApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Alphegate.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &AlphegateApprovalForAllIterator{contract: _Alphegate.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Alphegate *AlphegateFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *AlphegateApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Alphegate.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphegateApprovalForAll)
				if err := _Alphegate.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Alphegate *AlphegateFilterer) ParseApprovalForAll(log types.Log) (*AlphegateApprovalForAll, error) {
	event := new(AlphegateApprovalForAll)
	if err := _Alphegate.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlphegateConsecutiveTransferIterator is returned from FilterConsecutiveTransfer and is used to iterate over the raw logs and unpacked data for ConsecutiveTransfer events raised by the Alphegate contract.
type AlphegateConsecutiveTransferIterator struct {
	Event *AlphegateConsecutiveTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AlphegateConsecutiveTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphegateConsecutiveTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AlphegateConsecutiveTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AlphegateConsecutiveTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphegateConsecutiveTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphegateConsecutiveTransfer represents a ConsecutiveTransfer event raised by the Alphegate contract.
type AlphegateConsecutiveTransfer struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	From        common.Address
	To          common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConsecutiveTransfer is a free log retrieval operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Alphegate *AlphegateFilterer) FilterConsecutiveTransfer(opts *bind.FilterOpts, fromTokenId []*big.Int, from []common.Address, to []common.Address) (*AlphegateConsecutiveTransferIterator, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Alphegate.contract.FilterLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AlphegateConsecutiveTransferIterator{contract: _Alphegate.contract, event: "ConsecutiveTransfer", logs: logs, sub: sub}, nil
}

// WatchConsecutiveTransfer is a free log subscription operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Alphegate *AlphegateFilterer) WatchConsecutiveTransfer(opts *bind.WatchOpts, sink chan<- *AlphegateConsecutiveTransfer, fromTokenId []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Alphegate.contract.WatchLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphegateConsecutiveTransfer)
				if err := _Alphegate.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConsecutiveTransfer is a log parse operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Alphegate *AlphegateFilterer) ParseConsecutiveTransfer(log types.Log) (*AlphegateConsecutiveTransfer, error) {
	event := new(AlphegateConsecutiveTransfer)
	if err := _Alphegate.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlphegateOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Alphegate contract.
type AlphegateOwnershipTransferredIterator struct {
	Event *AlphegateOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AlphegateOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphegateOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AlphegateOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AlphegateOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphegateOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphegateOwnershipTransferred represents a OwnershipTransferred event raised by the Alphegate contract.
type AlphegateOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Alphegate *AlphegateFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AlphegateOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Alphegate.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AlphegateOwnershipTransferredIterator{contract: _Alphegate.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Alphegate *AlphegateFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AlphegateOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Alphegate.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphegateOwnershipTransferred)
				if err := _Alphegate.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Alphegate *AlphegateFilterer) ParseOwnershipTransferred(log types.Log) (*AlphegateOwnershipTransferred, error) {
	event := new(AlphegateOwnershipTransferred)
	if err := _Alphegate.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlphegatePermanentBaseURIIterator is returned from FilterPermanentBaseURI and is used to iterate over the raw logs and unpacked data for PermanentBaseURI events raised by the Alphegate contract.
type AlphegatePermanentBaseURIIterator struct {
	Event *AlphegatePermanentBaseURI // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AlphegatePermanentBaseURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphegatePermanentBaseURI)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AlphegatePermanentBaseURI)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AlphegatePermanentBaseURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphegatePermanentBaseURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphegatePermanentBaseURI represents a PermanentBaseURI event raised by the Alphegate contract.
type AlphegatePermanentBaseURI struct {
	BaseURI string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPermanentBaseURI is a free log retrieval operation binding the contract event 0xc6a6c2b165e62c9d37fc51a18ed76e5be22304bc1d337877c98f31c23e40b0f5.
//
// Solidity: event PermanentBaseURI(string baseURI)
func (_Alphegate *AlphegateFilterer) FilterPermanentBaseURI(opts *bind.FilterOpts) (*AlphegatePermanentBaseURIIterator, error) {

	logs, sub, err := _Alphegate.contract.FilterLogs(opts, "PermanentBaseURI")
	if err != nil {
		return nil, err
	}
	return &AlphegatePermanentBaseURIIterator{contract: _Alphegate.contract, event: "PermanentBaseURI", logs: logs, sub: sub}, nil
}

// WatchPermanentBaseURI is a free log subscription operation binding the contract event 0xc6a6c2b165e62c9d37fc51a18ed76e5be22304bc1d337877c98f31c23e40b0f5.
//
// Solidity: event PermanentBaseURI(string baseURI)
func (_Alphegate *AlphegateFilterer) WatchPermanentBaseURI(opts *bind.WatchOpts, sink chan<- *AlphegatePermanentBaseURI) (event.Subscription, error) {

	logs, sub, err := _Alphegate.contract.WatchLogs(opts, "PermanentBaseURI")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphegatePermanentBaseURI)
				if err := _Alphegate.contract.UnpackLog(event, "PermanentBaseURI", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePermanentBaseURI is a log parse operation binding the contract event 0xc6a6c2b165e62c9d37fc51a18ed76e5be22304bc1d337877c98f31c23e40b0f5.
//
// Solidity: event PermanentBaseURI(string baseURI)
func (_Alphegate *AlphegateFilterer) ParsePermanentBaseURI(log types.Log) (*AlphegatePermanentBaseURI, error) {
	event := new(AlphegatePermanentBaseURI)
	if err := _Alphegate.contract.UnpackLog(event, "PermanentBaseURI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlphegateSetActiveStageIterator is returned from FilterSetActiveStage and is used to iterate over the raw logs and unpacked data for SetActiveStage events raised by the Alphegate contract.
type AlphegateSetActiveStageIterator struct {
	Event *AlphegateSetActiveStage // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AlphegateSetActiveStageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphegateSetActiveStage)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AlphegateSetActiveStage)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AlphegateSetActiveStageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphegateSetActiveStageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphegateSetActiveStage represents a SetActiveStage event raised by the Alphegate contract.
type AlphegateSetActiveStage struct {
	ActiveStage *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetActiveStage is a free log retrieval operation binding the contract event 0x160d6de2c069c3adf7f4c1252236d0b325c0e3eb963ddb10c67a81aadf9a3058.
//
// Solidity: event SetActiveStage(uint256 activeStage)
func (_Alphegate *AlphegateFilterer) FilterSetActiveStage(opts *bind.FilterOpts) (*AlphegateSetActiveStageIterator, error) {

	logs, sub, err := _Alphegate.contract.FilterLogs(opts, "SetActiveStage")
	if err != nil {
		return nil, err
	}
	return &AlphegateSetActiveStageIterator{contract: _Alphegate.contract, event: "SetActiveStage", logs: logs, sub: sub}, nil
}

// WatchSetActiveStage is a free log subscription operation binding the contract event 0x160d6de2c069c3adf7f4c1252236d0b325c0e3eb963ddb10c67a81aadf9a3058.
//
// Solidity: event SetActiveStage(uint256 activeStage)
func (_Alphegate *AlphegateFilterer) WatchSetActiveStage(opts *bind.WatchOpts, sink chan<- *AlphegateSetActiveStage) (event.Subscription, error) {

	logs, sub, err := _Alphegate.contract.WatchLogs(opts, "SetActiveStage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphegateSetActiveStage)
				if err := _Alphegate.contract.UnpackLog(event, "SetActiveStage", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetActiveStage is a log parse operation binding the contract event 0x160d6de2c069c3adf7f4c1252236d0b325c0e3eb963ddb10c67a81aadf9a3058.
//
// Solidity: event SetActiveStage(uint256 activeStage)
func (_Alphegate *AlphegateFilterer) ParseSetActiveStage(log types.Log) (*AlphegateSetActiveStage, error) {
	event := new(AlphegateSetActiveStage)
	if err := _Alphegate.contract.UnpackLog(event, "SetActiveStage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlphegateSetBaseURIIterator is returned from FilterSetBaseURI and is used to iterate over the raw logs and unpacked data for SetBaseURI events raised by the Alphegate contract.
type AlphegateSetBaseURIIterator struct {
	Event *AlphegateSetBaseURI // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AlphegateSetBaseURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphegateSetBaseURI)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AlphegateSetBaseURI)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AlphegateSetBaseURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphegateSetBaseURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphegateSetBaseURI represents a SetBaseURI event raised by the Alphegate contract.
type AlphegateSetBaseURI struct {
	BaseURI string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetBaseURI is a free log retrieval operation binding the contract event 0x23c8c9488efebfd474e85a7956de6f39b17c7ab88502d42a623db2d8e382bbaa.
//
// Solidity: event SetBaseURI(string baseURI)
func (_Alphegate *AlphegateFilterer) FilterSetBaseURI(opts *bind.FilterOpts) (*AlphegateSetBaseURIIterator, error) {

	logs, sub, err := _Alphegate.contract.FilterLogs(opts, "SetBaseURI")
	if err != nil {
		return nil, err
	}
	return &AlphegateSetBaseURIIterator{contract: _Alphegate.contract, event: "SetBaseURI", logs: logs, sub: sub}, nil
}

// WatchSetBaseURI is a free log subscription operation binding the contract event 0x23c8c9488efebfd474e85a7956de6f39b17c7ab88502d42a623db2d8e382bbaa.
//
// Solidity: event SetBaseURI(string baseURI)
func (_Alphegate *AlphegateFilterer) WatchSetBaseURI(opts *bind.WatchOpts, sink chan<- *AlphegateSetBaseURI) (event.Subscription, error) {

	logs, sub, err := _Alphegate.contract.WatchLogs(opts, "SetBaseURI")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphegateSetBaseURI)
				if err := _Alphegate.contract.UnpackLog(event, "SetBaseURI", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetBaseURI is a log parse operation binding the contract event 0x23c8c9488efebfd474e85a7956de6f39b17c7ab88502d42a623db2d8e382bbaa.
//
// Solidity: event SetBaseURI(string baseURI)
func (_Alphegate *AlphegateFilterer) ParseSetBaseURI(log types.Log) (*AlphegateSetBaseURI, error) {
	event := new(AlphegateSetBaseURI)
	if err := _Alphegate.contract.UnpackLog(event, "SetBaseURI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlphegateSetCosignerIterator is returned from FilterSetCosigner and is used to iterate over the raw logs and unpacked data for SetCosigner events raised by the Alphegate contract.
type AlphegateSetCosignerIterator struct {
	Event *AlphegateSetCosigner // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AlphegateSetCosignerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphegateSetCosigner)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AlphegateSetCosigner)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AlphegateSetCosignerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphegateSetCosignerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphegateSetCosigner represents a SetCosigner event raised by the Alphegate contract.
type AlphegateSetCosigner struct {
	Cosigner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetCosigner is a free log retrieval operation binding the contract event 0xaea1573caf7b4fdd079b947d86c1be6c725642c47582f8f9bd2c7d2a30bf0bd9.
//
// Solidity: event SetCosigner(address cosigner)
func (_Alphegate *AlphegateFilterer) FilterSetCosigner(opts *bind.FilterOpts) (*AlphegateSetCosignerIterator, error) {

	logs, sub, err := _Alphegate.contract.FilterLogs(opts, "SetCosigner")
	if err != nil {
		return nil, err
	}
	return &AlphegateSetCosignerIterator{contract: _Alphegate.contract, event: "SetCosigner", logs: logs, sub: sub}, nil
}

// WatchSetCosigner is a free log subscription operation binding the contract event 0xaea1573caf7b4fdd079b947d86c1be6c725642c47582f8f9bd2c7d2a30bf0bd9.
//
// Solidity: event SetCosigner(address cosigner)
func (_Alphegate *AlphegateFilterer) WatchSetCosigner(opts *bind.WatchOpts, sink chan<- *AlphegateSetCosigner) (event.Subscription, error) {

	logs, sub, err := _Alphegate.contract.WatchLogs(opts, "SetCosigner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphegateSetCosigner)
				if err := _Alphegate.contract.UnpackLog(event, "SetCosigner", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetCosigner is a log parse operation binding the contract event 0xaea1573caf7b4fdd079b947d86c1be6c725642c47582f8f9bd2c7d2a30bf0bd9.
//
// Solidity: event SetCosigner(address cosigner)
func (_Alphegate *AlphegateFilterer) ParseSetCosigner(log types.Log) (*AlphegateSetCosigner, error) {
	event := new(AlphegateSetCosigner)
	if err := _Alphegate.contract.UnpackLog(event, "SetCosigner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlphegateSetCrossmintAddressIterator is returned from FilterSetCrossmintAddress and is used to iterate over the raw logs and unpacked data for SetCrossmintAddress events raised by the Alphegate contract.
type AlphegateSetCrossmintAddressIterator struct {
	Event *AlphegateSetCrossmintAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AlphegateSetCrossmintAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphegateSetCrossmintAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AlphegateSetCrossmintAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AlphegateSetCrossmintAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphegateSetCrossmintAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphegateSetCrossmintAddress represents a SetCrossmintAddress event raised by the Alphegate contract.
type AlphegateSetCrossmintAddress struct {
	CrossmintAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetCrossmintAddress is a free log retrieval operation binding the contract event 0xf477d93c015f2a73c2ccc5ed37078d12123b80fc5d12e0014c60b913bc1a1ec4.
//
// Solidity: event SetCrossmintAddress(address crossmintAddress)
func (_Alphegate *AlphegateFilterer) FilterSetCrossmintAddress(opts *bind.FilterOpts) (*AlphegateSetCrossmintAddressIterator, error) {

	logs, sub, err := _Alphegate.contract.FilterLogs(opts, "SetCrossmintAddress")
	if err != nil {
		return nil, err
	}
	return &AlphegateSetCrossmintAddressIterator{contract: _Alphegate.contract, event: "SetCrossmintAddress", logs: logs, sub: sub}, nil
}

// WatchSetCrossmintAddress is a free log subscription operation binding the contract event 0xf477d93c015f2a73c2ccc5ed37078d12123b80fc5d12e0014c60b913bc1a1ec4.
//
// Solidity: event SetCrossmintAddress(address crossmintAddress)
func (_Alphegate *AlphegateFilterer) WatchSetCrossmintAddress(opts *bind.WatchOpts, sink chan<- *AlphegateSetCrossmintAddress) (event.Subscription, error) {

	logs, sub, err := _Alphegate.contract.WatchLogs(opts, "SetCrossmintAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphegateSetCrossmintAddress)
				if err := _Alphegate.contract.UnpackLog(event, "SetCrossmintAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetCrossmintAddress is a log parse operation binding the contract event 0xf477d93c015f2a73c2ccc5ed37078d12123b80fc5d12e0014c60b913bc1a1ec4.
//
// Solidity: event SetCrossmintAddress(address crossmintAddress)
func (_Alphegate *AlphegateFilterer) ParseSetCrossmintAddress(log types.Log) (*AlphegateSetCrossmintAddress, error) {
	event := new(AlphegateSetCrossmintAddress)
	if err := _Alphegate.contract.UnpackLog(event, "SetCrossmintAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlphegateSetGlobalWalletLimitIterator is returned from FilterSetGlobalWalletLimit and is used to iterate over the raw logs and unpacked data for SetGlobalWalletLimit events raised by the Alphegate contract.
type AlphegateSetGlobalWalletLimitIterator struct {
	Event *AlphegateSetGlobalWalletLimit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AlphegateSetGlobalWalletLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphegateSetGlobalWalletLimit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AlphegateSetGlobalWalletLimit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AlphegateSetGlobalWalletLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphegateSetGlobalWalletLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphegateSetGlobalWalletLimit represents a SetGlobalWalletLimit event raised by the Alphegate contract.
type AlphegateSetGlobalWalletLimit struct {
	GlobalWalletLimit *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSetGlobalWalletLimit is a free log retrieval operation binding the contract event 0x5307de8ad7d34d5ddfd5171435c143bdc645493980f453eb5d7cdb3e494a1b35.
//
// Solidity: event SetGlobalWalletLimit(uint256 globalWalletLimit)
func (_Alphegate *AlphegateFilterer) FilterSetGlobalWalletLimit(opts *bind.FilterOpts) (*AlphegateSetGlobalWalletLimitIterator, error) {

	logs, sub, err := _Alphegate.contract.FilterLogs(opts, "SetGlobalWalletLimit")
	if err != nil {
		return nil, err
	}
	return &AlphegateSetGlobalWalletLimitIterator{contract: _Alphegate.contract, event: "SetGlobalWalletLimit", logs: logs, sub: sub}, nil
}

// WatchSetGlobalWalletLimit is a free log subscription operation binding the contract event 0x5307de8ad7d34d5ddfd5171435c143bdc645493980f453eb5d7cdb3e494a1b35.
//
// Solidity: event SetGlobalWalletLimit(uint256 globalWalletLimit)
func (_Alphegate *AlphegateFilterer) WatchSetGlobalWalletLimit(opts *bind.WatchOpts, sink chan<- *AlphegateSetGlobalWalletLimit) (event.Subscription, error) {

	logs, sub, err := _Alphegate.contract.WatchLogs(opts, "SetGlobalWalletLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphegateSetGlobalWalletLimit)
				if err := _Alphegate.contract.UnpackLog(event, "SetGlobalWalletLimit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetGlobalWalletLimit is a log parse operation binding the contract event 0x5307de8ad7d34d5ddfd5171435c143bdc645493980f453eb5d7cdb3e494a1b35.
//
// Solidity: event SetGlobalWalletLimit(uint256 globalWalletLimit)
func (_Alphegate *AlphegateFilterer) ParseSetGlobalWalletLimit(log types.Log) (*AlphegateSetGlobalWalletLimit, error) {
	event := new(AlphegateSetGlobalWalletLimit)
	if err := _Alphegate.contract.UnpackLog(event, "SetGlobalWalletLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlphegateSetMaxMintableSupplyIterator is returned from FilterSetMaxMintableSupply and is used to iterate over the raw logs and unpacked data for SetMaxMintableSupply events raised by the Alphegate contract.
type AlphegateSetMaxMintableSupplyIterator struct {
	Event *AlphegateSetMaxMintableSupply // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AlphegateSetMaxMintableSupplyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphegateSetMaxMintableSupply)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AlphegateSetMaxMintableSupply)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AlphegateSetMaxMintableSupplyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphegateSetMaxMintableSupplyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphegateSetMaxMintableSupply represents a SetMaxMintableSupply event raised by the Alphegate contract.
type AlphegateSetMaxMintableSupply struct {
	MaxMintableSupply *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSetMaxMintableSupply is a free log retrieval operation binding the contract event 0xc7bbc2b288fc13314546ea4aa51f6bcf71b7ba4740beeb3d32e9acef57b6668a.
//
// Solidity: event SetMaxMintableSupply(uint256 maxMintableSupply)
func (_Alphegate *AlphegateFilterer) FilterSetMaxMintableSupply(opts *bind.FilterOpts) (*AlphegateSetMaxMintableSupplyIterator, error) {

	logs, sub, err := _Alphegate.contract.FilterLogs(opts, "SetMaxMintableSupply")
	if err != nil {
		return nil, err
	}
	return &AlphegateSetMaxMintableSupplyIterator{contract: _Alphegate.contract, event: "SetMaxMintableSupply", logs: logs, sub: sub}, nil
}

// WatchSetMaxMintableSupply is a free log subscription operation binding the contract event 0xc7bbc2b288fc13314546ea4aa51f6bcf71b7ba4740beeb3d32e9acef57b6668a.
//
// Solidity: event SetMaxMintableSupply(uint256 maxMintableSupply)
func (_Alphegate *AlphegateFilterer) WatchSetMaxMintableSupply(opts *bind.WatchOpts, sink chan<- *AlphegateSetMaxMintableSupply) (event.Subscription, error) {

	logs, sub, err := _Alphegate.contract.WatchLogs(opts, "SetMaxMintableSupply")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphegateSetMaxMintableSupply)
				if err := _Alphegate.contract.UnpackLog(event, "SetMaxMintableSupply", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetMaxMintableSupply is a log parse operation binding the contract event 0xc7bbc2b288fc13314546ea4aa51f6bcf71b7ba4740beeb3d32e9acef57b6668a.
//
// Solidity: event SetMaxMintableSupply(uint256 maxMintableSupply)
func (_Alphegate *AlphegateFilterer) ParseSetMaxMintableSupply(log types.Log) (*AlphegateSetMaxMintableSupply, error) {
	event := new(AlphegateSetMaxMintableSupply)
	if err := _Alphegate.contract.UnpackLog(event, "SetMaxMintableSupply", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlphegateSetMintableIterator is returned from FilterSetMintable and is used to iterate over the raw logs and unpacked data for SetMintable events raised by the Alphegate contract.
type AlphegateSetMintableIterator struct {
	Event *AlphegateSetMintable // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AlphegateSetMintableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphegateSetMintable)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AlphegateSetMintable)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AlphegateSetMintableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphegateSetMintableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphegateSetMintable represents a SetMintable event raised by the Alphegate contract.
type AlphegateSetMintable struct {
	Mintable bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetMintable is a free log retrieval operation binding the contract event 0xe717a2bfc51e250b028aaac5eb448e76f4df26b9609956782bff49097bb792cf.
//
// Solidity: event SetMintable(bool mintable)
func (_Alphegate *AlphegateFilterer) FilterSetMintable(opts *bind.FilterOpts) (*AlphegateSetMintableIterator, error) {

	logs, sub, err := _Alphegate.contract.FilterLogs(opts, "SetMintable")
	if err != nil {
		return nil, err
	}
	return &AlphegateSetMintableIterator{contract: _Alphegate.contract, event: "SetMintable", logs: logs, sub: sub}, nil
}

// WatchSetMintable is a free log subscription operation binding the contract event 0xe717a2bfc51e250b028aaac5eb448e76f4df26b9609956782bff49097bb792cf.
//
// Solidity: event SetMintable(bool mintable)
func (_Alphegate *AlphegateFilterer) WatchSetMintable(opts *bind.WatchOpts, sink chan<- *AlphegateSetMintable) (event.Subscription, error) {

	logs, sub, err := _Alphegate.contract.WatchLogs(opts, "SetMintable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphegateSetMintable)
				if err := _Alphegate.contract.UnpackLog(event, "SetMintable", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetMintable is a log parse operation binding the contract event 0xe717a2bfc51e250b028aaac5eb448e76f4df26b9609956782bff49097bb792cf.
//
// Solidity: event SetMintable(bool mintable)
func (_Alphegate *AlphegateFilterer) ParseSetMintable(log types.Log) (*AlphegateSetMintable, error) {
	event := new(AlphegateSetMintable)
	if err := _Alphegate.contract.UnpackLog(event, "SetMintable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlphegateSetTimestampExpirySecondsIterator is returned from FilterSetTimestampExpirySeconds and is used to iterate over the raw logs and unpacked data for SetTimestampExpirySeconds events raised by the Alphegate contract.
type AlphegateSetTimestampExpirySecondsIterator struct {
	Event *AlphegateSetTimestampExpirySeconds // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AlphegateSetTimestampExpirySecondsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphegateSetTimestampExpirySeconds)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AlphegateSetTimestampExpirySeconds)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AlphegateSetTimestampExpirySecondsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphegateSetTimestampExpirySecondsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphegateSetTimestampExpirySeconds represents a SetTimestampExpirySeconds event raised by the Alphegate contract.
type AlphegateSetTimestampExpirySeconds struct {
	Expiry uint64
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetTimestampExpirySeconds is a free log retrieval operation binding the contract event 0x41b9126ccd8cb4505310c40a376055b5ef246bd4c9214de02af31ef4f26b1b5f.
//
// Solidity: event SetTimestampExpirySeconds(uint64 expiry)
func (_Alphegate *AlphegateFilterer) FilterSetTimestampExpirySeconds(opts *bind.FilterOpts) (*AlphegateSetTimestampExpirySecondsIterator, error) {

	logs, sub, err := _Alphegate.contract.FilterLogs(opts, "SetTimestampExpirySeconds")
	if err != nil {
		return nil, err
	}
	return &AlphegateSetTimestampExpirySecondsIterator{contract: _Alphegate.contract, event: "SetTimestampExpirySeconds", logs: logs, sub: sub}, nil
}

// WatchSetTimestampExpirySeconds is a free log subscription operation binding the contract event 0x41b9126ccd8cb4505310c40a376055b5ef246bd4c9214de02af31ef4f26b1b5f.
//
// Solidity: event SetTimestampExpirySeconds(uint64 expiry)
func (_Alphegate *AlphegateFilterer) WatchSetTimestampExpirySeconds(opts *bind.WatchOpts, sink chan<- *AlphegateSetTimestampExpirySeconds) (event.Subscription, error) {

	logs, sub, err := _Alphegate.contract.WatchLogs(opts, "SetTimestampExpirySeconds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphegateSetTimestampExpirySeconds)
				if err := _Alphegate.contract.UnpackLog(event, "SetTimestampExpirySeconds", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetTimestampExpirySeconds is a log parse operation binding the contract event 0x41b9126ccd8cb4505310c40a376055b5ef246bd4c9214de02af31ef4f26b1b5f.
//
// Solidity: event SetTimestampExpirySeconds(uint64 expiry)
func (_Alphegate *AlphegateFilterer) ParseSetTimestampExpirySeconds(log types.Log) (*AlphegateSetTimestampExpirySeconds, error) {
	event := new(AlphegateSetTimestampExpirySeconds)
	if err := _Alphegate.contract.UnpackLog(event, "SetTimestampExpirySeconds", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlphegateTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Alphegate contract.
type AlphegateTransferIterator struct {
	Event *AlphegateTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AlphegateTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphegateTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AlphegateTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AlphegateTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphegateTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphegateTransfer represents a Transfer event raised by the Alphegate contract.
type AlphegateTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Alphegate *AlphegateFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*AlphegateTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Alphegate.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &AlphegateTransferIterator{contract: _Alphegate.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Alphegate *AlphegateFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AlphegateTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Alphegate.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphegateTransfer)
				if err := _Alphegate.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Alphegate *AlphegateFilterer) ParseTransfer(log types.Log) (*AlphegateTransfer, error) {
	event := new(AlphegateTransfer)
	if err := _Alphegate.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlphegateUpdateStageIterator is returned from FilterUpdateStage and is used to iterate over the raw logs and unpacked data for UpdateStage events raised by the Alphegate contract.
type AlphegateUpdateStageIterator struct {
	Event *AlphegateUpdateStage // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AlphegateUpdateStageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphegateUpdateStage)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AlphegateUpdateStage)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AlphegateUpdateStageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphegateUpdateStageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphegateUpdateStage represents a UpdateStage event raised by the Alphegate contract.
type AlphegateUpdateStage struct {
	Stage                *big.Int
	Price                *big.Int
	WalletLimit          uint32
	MerkleRoot           [32]byte
	MaxStageSupply       *big.Int
	StartTimeUnixSeconds uint64
	EndTimeUnixSeconds   uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdateStage is a free log retrieval operation binding the contract event 0xb3268648542a1bb1b2dd12e3b14aeb5a3ab22c592de96bdd3e842154a5b394fa.
//
// Solidity: event UpdateStage(uint256 stage, uint80 price, uint32 walletLimit, bytes32 merkleRoot, uint24 maxStageSupply, uint64 startTimeUnixSeconds, uint64 endTimeUnixSeconds)
func (_Alphegate *AlphegateFilterer) FilterUpdateStage(opts *bind.FilterOpts) (*AlphegateUpdateStageIterator, error) {

	logs, sub, err := _Alphegate.contract.FilterLogs(opts, "UpdateStage")
	if err != nil {
		return nil, err
	}
	return &AlphegateUpdateStageIterator{contract: _Alphegate.contract, event: "UpdateStage", logs: logs, sub: sub}, nil
}

// WatchUpdateStage is a free log subscription operation binding the contract event 0xb3268648542a1bb1b2dd12e3b14aeb5a3ab22c592de96bdd3e842154a5b394fa.
//
// Solidity: event UpdateStage(uint256 stage, uint80 price, uint32 walletLimit, bytes32 merkleRoot, uint24 maxStageSupply, uint64 startTimeUnixSeconds, uint64 endTimeUnixSeconds)
func (_Alphegate *AlphegateFilterer) WatchUpdateStage(opts *bind.WatchOpts, sink chan<- *AlphegateUpdateStage) (event.Subscription, error) {

	logs, sub, err := _Alphegate.contract.WatchLogs(opts, "UpdateStage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphegateUpdateStage)
				if err := _Alphegate.contract.UnpackLog(event, "UpdateStage", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateStage is a log parse operation binding the contract event 0xb3268648542a1bb1b2dd12e3b14aeb5a3ab22c592de96bdd3e842154a5b394fa.
//
// Solidity: event UpdateStage(uint256 stage, uint80 price, uint32 walletLimit, bytes32 merkleRoot, uint24 maxStageSupply, uint64 startTimeUnixSeconds, uint64 endTimeUnixSeconds)
func (_Alphegate *AlphegateFilterer) ParseUpdateStage(log types.Log) (*AlphegateUpdateStage, error) {
	event := new(AlphegateUpdateStage)
	if err := _Alphegate.contract.UnpackLog(event, "UpdateStage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlphegateWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Alphegate contract.
type AlphegateWithdrawIterator struct {
	Event *AlphegateWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AlphegateWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphegateWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AlphegateWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AlphegateWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphegateWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphegateWithdraw represents a Withdraw event raised by the Alphegate contract.
type AlphegateWithdraw struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x5b6b431d4476a211bb7d41c20d1aab9ae2321deee0d20be3d9fc9b1093fa6e3d.
//
// Solidity: event Withdraw(uint256 value)
func (_Alphegate *AlphegateFilterer) FilterWithdraw(opts *bind.FilterOpts) (*AlphegateWithdrawIterator, error) {

	logs, sub, err := _Alphegate.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &AlphegateWithdrawIterator{contract: _Alphegate.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x5b6b431d4476a211bb7d41c20d1aab9ae2321deee0d20be3d9fc9b1093fa6e3d.
//
// Solidity: event Withdraw(uint256 value)
func (_Alphegate *AlphegateFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *AlphegateWithdraw) (event.Subscription, error) {

	logs, sub, err := _Alphegate.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphegateWithdraw)
				if err := _Alphegate.contract.UnpackLog(event, "Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdraw is a log parse operation binding the contract event 0x5b6b431d4476a211bb7d41c20d1aab9ae2321deee0d20be3d9fc9b1093fa6e3d.
//
// Solidity: event Withdraw(uint256 value)
func (_Alphegate *AlphegateFilterer) ParseWithdraw(log types.Log) (*AlphegateWithdraw, error) {
	event := new(AlphegateWithdraw)
	if err := _Alphegate.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
