// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package efes

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

// ISRCNFTInputMintStageInfo is an auto generated low-level Go binding around an user-defined struct.
type ISRCNFTInputMintStageInfo struct {
	WhiteSalePrice       *big.Int
	PublicSalePrice      *big.Int
	WhiteSaleHour        *big.Int
	PublicSaleHour       *big.Int
	MaxStageSupply       *big.Int
	StartTimeUnixSeconds uint64
}

// EfesMetaData contains all meta data concerning the Efes contract.
var EfesMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_tokenURISuffix\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_currentBaseURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_cosigner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minGasToTransferAndStore\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_lzEndpoint\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_expire\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_maxSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ApprovalCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceQueryForZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CosignerNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HasStopped\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCosignSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStage\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStartAndEndTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintERC2309QuantityExceedsLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintZeroQuantity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Mintable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoStageSupplyLeft\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotMintable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnershipNotInitializedForExtraData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimestampExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromIncorrectOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToNonERC721ReceiverImplementer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"URIQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WMStopped\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"notSatifiedSig\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"ConsecutiveTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hashedPayload\",\"type\":\"bytes32\"}],\"name\":\"CreditCleared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hashedPayload\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"CreditStored\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_reason\",\"type\":\"bytes\"}],\"name\":\"MessageFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_toAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"ReceiveFromChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_payloadHash\",\"type\":\"bytes32\"}],\"name\":\"RetryMessageSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"_toAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"SendToChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_dstChainIdToBatchLimit\",\"type\":\"uint256\"}],\"name\":\"SetDstChainIdToBatchLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_dstChainIdToTransferGas\",\"type\":\"uint256\"}],\"name\":\"SetDstChainIdToTransferGas\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_type\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_minDstGas\",\"type\":\"uint256\"}],\"name\":\"SetMinDstGas\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_minGasToTransferAndStore\",\"type\":\"uint256\"}],\"name\":\"SetMinGasToTransferAndStore\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"precrime\",\"type\":\"address\"}],\"name\":\"SetPrecrime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_remoteChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_path\",\"type\":\"bytes\"}],\"name\":\"SetTrustedRemote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_remoteChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_remoteAddress\",\"type\":\"bytes\"}],\"name\":\"SetTrustedRemoteAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_PAYLOAD_SIZE_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FUNCTION_TYPE_SEND\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_chainID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beneficiary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"clearCredits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cosigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentBaseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"decode\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"qty\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"requestId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"dstChainIdToBatchLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"dstChainIdToTransferGas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_toAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"bool\",\"name\":\"_useZro\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"_adapterParams\",\"type\":\"bytes\"}],\"name\":\"estimateSendBatchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zroFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_toAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_useZro\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"_adapterParams\",\"type\":\"bytes\"}],\"name\":\"estimateSendFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zroFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expireTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"failedMessages\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"}],\"name\":\"forceResumeReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"getActiveStageFromTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_version\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_chainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_configType\",\"type\":\"uint256\"}],\"name\":\"getConfig\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumberStages\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_remoteChainId\",\"type\":\"uint16\"}],\"name\":\"getTrustedRemoteAddress\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gettimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"}],\"name\":\"isTrustedRemote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lzEndpoint\",\"outputs\":[{\"internalType\":\"contractILayerZeroEndpoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"lzReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"minDstGasLookup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minGasToTransferAndStore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mintStages\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"whiteSalePrice\",\"type\":\"uint80\"},{\"internalType\":\"uint80\",\"name\":\"publicSalePrice\",\"type\":\"uint80\"},{\"internalType\":\"uint24\",\"name\":\"whiteSaleHour\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"publicSaleHour\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"maxStageSupply\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"startTimeUnixSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"endWhiteTimeUnixSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTimeUnixSeconds\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mint_to\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"nonblockingLzReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"qty\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"ownermint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"payloadSizeLimitLookup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"precrime\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"retryMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_toAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"addresspayable\",\"name\":\"_refundAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_zroPaymentAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_adapterParams\",\"type\":\"bytes\"}],\"name\":\"sendBatchFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_toAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_refundAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_zroPaymentAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_adapterParams\",\"type\":\"bytes\"}],\"name\":\"sendFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_version\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_chainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_configType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_config\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainIdToBatchLimit\",\"type\":\"uint256\"}],\"name\":\"setDstChainIdToBatchLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainIdToTransferGas\",\"type\":\"uint256\"}],\"name\":\"setDstChainIdToTransferGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_packetType\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_minGas\",\"type\":\"uint256\"}],\"name\":\"setMinDstGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minGasToTransferAndStore\",\"type\":\"uint256\"}],\"name\":\"setMinGasToTransferAndStore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_mintable\",\"type\":\"bool\"}],\"name\":\"setMintable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_size\",\"type\":\"uint256\"}],\"name\":\"setPayloadSizeLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_precrime\",\"type\":\"address\"}],\"name\":\"setPrecrime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_version\",\"type\":\"uint16\"}],\"name\":\"setReceiveVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_version\",\"type\":\"uint16\"}],\"name\":\"setSendVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint80\",\"name\":\"whiteSalePrice\",\"type\":\"uint80\"},{\"internalType\":\"uint80\",\"name\":\"publicSalePrice\",\"type\":\"uint80\"},{\"internalType\":\"uint24\",\"name\":\"whiteSaleHour\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"publicSaleHour\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"maxStageSupply\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"startTimeUnixSeconds\",\"type\":\"uint64\"}],\"internalType\":\"structISRCNFT.InputMintStageInfo[]\",\"name\":\"newStages\",\"type\":\"tuple[]\"}],\"name\":\"setStages\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"suffix\",\"type\":\"string\"}],\"name\":\"setTokenURISuffix\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_remoteChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_path\",\"type\":\"bytes\"}],\"name\":\"setTrustedRemote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_remoteChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_remoteAddress\",\"type\":\"bytes\"}],\"name\":\"setTrustedRemoteAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stageMintedCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint24\",\"name\":\"_days\",\"type\":\"uint24\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"storedCredits\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"creditsRemain\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenIdStakeDays\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenURISuffix\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalMinted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"trustedRemoteLookup\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_cosigner\",\"type\":\"address\"}],\"name\":\"updateCosigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_expire\",\"type\":\"uint32\"}],\"name\":\"updateExpire\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"whiteSalePrice\",\"type\":\"uint80\"},{\"internalType\":\"uint80\",\"name\":\"publicSalePrice\",\"type\":\"uint80\"},{\"internalType\":\"uint24\",\"name\":\"maxStageSupply\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"startTimeUnixSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint24\",\"name\":\"whiteSaleHour\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"publicSaleHour\",\"type\":\"uint24\"}],\"name\":\"updateStage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405260405162006548380380620065488339810160408190526200002691620002d0565b89898585808033806200005457604051631e4fbdf760e01b8152600060048201526024015b60405180910390fd5b6200005f8162000189565b506001600160a01b031660805250600160065581620000cd5760405162461bcd60e51b8152602060048201526024808201527f6d696e476173546f5472616e73666572416e6453746f7265206d7573742062656044820152630203e20360e41b60648201526084016200004b565b50600755600d620000df838262000478565b50600e620000ee828262000478565b506000600b5550506013805460ff19166001179055601662000111898262000478565b50601562000120888262000478565b5060138054610100600160a81b0319166101006001600160a01b038981169190910291909117909155601480546001600160a01b03191691871691909117905562000173603c63ffffffff841662000544565b601755601c555062000570975050505050505050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126200020157600080fd5b81516001600160401b03808211156200021e576200021e620001d9565b604051601f8301601f19908116603f01168101908282118183101715620002495762000249620001d9565b816040528381526020925086838588010111156200026657600080fd5b600091505b838210156200028a57858201830151818301840152908201906200026b565b600093810190920192909252949350505050565b80516001600160a01b0381168114620002b657600080fd5b919050565b805163ffffffff81168114620002b657600080fd5b6000806000806000806000806000806101408b8d031215620002f157600080fd5b8a516001600160401b03808211156200030957600080fd5b620003178e838f01620001ef565b9b5060208d01519150808211156200032e57600080fd5b6200033c8e838f01620001ef565b9a5060408d01519150808211156200035357600080fd5b620003618e838f01620001ef565b995060608d01519150808211156200037857600080fd5b50620003878d828e01620001ef565b9750506200039860808c016200029e565b9550620003a860a08c016200029e565b945060c08b01519350620003bf60e08c016200029e565b9250620003d06101008c01620002bb565b91506101208b015190509295989b9194979a5092959850565b600181811c90821680620003fe57607f821691505b6020821081036200041f57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200047357600081815260208120601f850160051c810160208610156200044e5750805b601f850160051c820191505b818110156200046f578281556001016200045a565b5050505b505050565b81516001600160401b03811115620004945762000494620001d9565b620004ac81620004a58454620003e9565b8462000425565b602080601f831160018114620004e45760008415620004cb5750858301515b600019600386901b1c1916600185901b1785556200046f565b600085815260208120601f198616915b828110156200051557888601518255948401946001909101908401620004f4565b5085821015620005345787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b80820281158282048414176200056a57634e487b7160e01b600052601160045260246000fd5b92915050565b608051615f84620005c460003960008181610c7701528181610fa1015281816112ad015281816114d901528181611afe01528181612a4f01528181612f8e015281816130ab01526143b20152615f846000f3fe6080604052600436106104885760003560e01c80637ba0e2e711610255578063beab713111610144578063dbbc853b116100c1578063f235364111610085578063f235364114610e84578063f2fde38b14610ea4578063f41aa65514610ec4578063f5ecbdbc14610ed7578063fa25f9b614610ef7578063fd0f745614610f2457600080fd5b8063dbbc853b14610ddf578063df2a5b3b14610df4578063e5c5e9a314610e14578063e985e9c514610e44578063eb8d72b714610e6457600080fd5b8063d12473a511610108578063d12473a514610d63578063d1deba1f14610d83578063d325a79414610d96578063d5abeb0114610db6578063dabc0e1d14610dcc57600080fd5b8063beab713114610ccc578063c446183414610ced578063c5b2d40414610d03578063c87b56dd14610d23578063cbed8b9c14610d4357600080fd5b8063a22cb465116101d2578063af3fb21c11610196578063af3fb21c14610c1d578063afa2d19014610c45578063b353aaa714610c65578063b88d4fde14610c99578063baf3292d14610cac57600080fd5b8063a22cb46514610b95578063a2309ff814610bb5578063a6c3d16514610bca578063a9852bfb14610bea578063ab3ffb9314610c0a57600080fd5b8063950c8a7411610219578063950c8a7414610b0b57806395d89b4114610b2b57806397d4ccd214610b405780639ea5d6b114610b555780639f38369a14610b7557600080fd5b80637ba0e2e714610a6f5780637c15dd6014610a825780638cfd8f5c14610a955780638da5cb5b14610acd5780638ffa1f2a14610aeb57600080fd5b80633ccfd60b1161037c57806351905636116102f957806367808a34116102bd57806367808a34146109c0578063681232ad146109e057806370a0823114610a0557806370da24ee14610a25578063715018a614610a3a5780637533d78814610a4f57600080fd5b806351905636146108fe57806355f804b3146109115780635b8c41e6146109315780636352211e1461098057806366ad5c8a146109a057600080fd5b806342d65a8d1161034057806342d65a8d1461085457806348288190146108745780634ac3f4ff1461088a5780634bf365df146108b75780634ca6d84d146108d157600080fd5b80633ccfd60b146107bf5780633d8b38f6146107d45780633f1f4fa4146107f457806342842e0e1461082157806342966c681461083457600080fd5b8063150b7a021161040a578063285d70d4116103ce578063285d70d4146107145780632a205e3d146107345780632c1cc7f914610769578063310a1ee31461078957806338af3eed1461079f57600080fd5b8063150b7a021461060c57806318160ddd146106455780631824c4de1461065e57806322a3ecf91461067e57806323b872dd1461070157600080fd5b8063095ea7b311610451578063095ea7b31461055e5780630b4cad4c146105715780630df37483146105915780630f461d5e146105b157806310ddb137146105ec57600080fd5b80621d35671461048d57806301ffc9a7146104af57806306fdde03146104e457806307e0db1714610506578063081812fc14610526575b600080fd5b34801561049957600080fd5b506104ad6104a8366004614a22565b610f9e565b005b3480156104bb57600080fd5b506104cf6104ca366004614acc565b6111cf565b60405190151581526020015b60405180910390f35b3480156104f057600080fd5b506104f96111fa565b6040516104db9190614b39565b34801561051257600080fd5b506104ad610521366004614b4c565b61128c565b34801561053257600080fd5b50610546610541366004614b67565b611315565b6040516001600160a01b0390911681526020016104db565b6104ad61056c366004614ba0565b611359565b34801561057d57600080fd5b506104ad61058c366004614b67565b6113f9565b34801561059d57600080fd5b506104ad6105ac366004614bcc565b611499565b3480156105bd57600080fd5b506105de6105cc366004614b67565b60196020526000908152604090205481565b6040519081526020016104db565b3480156105f857600080fd5b506104ad610607366004614b4c565b6114b8565b34801561061857600080fd5b5061062c610627366004614cab565b611510565b6040516001600160e01b031990911681526020016104db565b34801561065157600080fd5b50600c54600b54036105de565b34801561066a57600080fd5b506104ad610679366004614d29565b611521565b34801561068a57600080fd5b506106d2610699366004614b67565b600a6020526000908152604090208054600182015460029092015461ffff821692620100009092046001600160a01b0316919060ff1684565b6040805161ffff90951685526001600160a01b03909316602085015291830152151560608201526080016104db565b6104ad61070f366004614d55565b6115cb565b34801561072057600080fd5b506104ad61072f366004614da6565b611609565b34801561074057600080fd5b5061075461074f366004614dc1565b611624565b604080519283526020830191909152016104db565b34801561077557600080fd5b506104ad610784366004614e66565b61164a565b34801561079557600080fd5b506105de60175481565b3480156107ab57600080fd5b50601454610546906001600160a01b031681565b3480156107cb57600080fd5b506104ad611902565b3480156107e057600080fd5b506104cf6107ef366004614ee7565b6119bc565b34801561080057600080fd5b506105de61080f366004614b4c565b60036020526000908152604090205481565b6104ad61082f366004614d55565b611a88565b34801561084057600080fd5b506104ad61084f366004614b67565b611aa3565b34801561086057600080fd5b506104ad61086f366004614ee7565b611adf565b34801561088057600080fd5b506105de60075481565b34801561089657600080fd5b506105de6108a5366004614b4c565b60086020526000908152604090205481565b3480156108c357600080fd5b506013546104cf9060ff1681565b3480156108dd57600080fd5b506105de6108ec366004614b67565b601a6020526000908152604090205481565b6104ad61090c366004614f39565b611b65565b34801561091d57600080fd5b506104ad61092c366004614ff2565b611b7c565b34801561093d57600080fd5b506105de61094c366004615033565b6005602090815260009384526040808520845180860184018051928152908401958401959095209452929052825290205481565b34801561098c57600080fd5b5061054661099b366004614b67565b611b91565b3480156109ac57600080fd5b506104ad6109bb366004614a22565b611b9c565b3480156109cc57600080fd5b506105de6109db366004615094565b611c78565b3480156109ec57600080fd5b506013546105469061010090046001600160a01b031681565b348015610a1157600080fd5b506105de610a203660046150b1565b611d2c565b348015610a3157600080fd5b506018546105de565b348015610a4657600080fd5b506104ad611d7a565b348015610a5b57600080fd5b506104f9610a6a366004614b4c565b611d8e565b6104ad610a7d366004614ff2565b611e28565b348015610a8e57600080fd5b50426105de565b348015610aa157600080fd5b506105de610ab03660046150ce565b600260209081526000928352604080842090915290825290205481565b348015610ad957600080fd5b506000546001600160a01b0316610546565b348015610af757600080fd5b506104ad610b063660046150f8565b611e91565b348015610b1757600080fd5b50600454610546906001600160a01b031681565b348015610b3757600080fd5b506104f96120d9565b348015610b4c57600080fd5b506104f96120e8565b348015610b6157600080fd5b506104ad610b70366004614bcc565b6120f5565b348015610b8157600080fd5b506104f9610b90366004614b4c565b6121a4565b348015610ba157600080fd5b506104ad610bb036600461512c565b6122ba565b348015610bc157600080fd5b506105de612326565b348015610bd657600080fd5b506104ad610be5366004614ee7565b612336565b348015610bf657600080fd5b506104ad610c05366004614ff2565b6123bf565b6104ad610c183660046151e1565b6123d4565b348015610c2957600080fd5b50610c32600181565b60405161ffff90911681526020016104db565b348015610c5157600080fd5b506104ad610c60366004615296565b6123e3565b348015610c7157600080fd5b506105467f000000000000000000000000000000000000000000000000000000000000000081565b6104ad610ca7366004614cab565b61286e565b348015610cb857600080fd5b506104ad610cc73660046150b1565b6128b2565b348015610cd857600080fd5b5060405163ffffffff461681526020016104db565b348015610cf957600080fd5b506105de61271081565b348015610d0f57600080fd5b506104ad610d1e36600461531c565b612908565b348015610d2f57600080fd5b506104f9610d3e366004614b67565b612927565b348015610d4f57600080fd5b506104ad610d5e366004615339565b612a30565b348015610d6f57600080fd5b506104ad610d7e366004614bcc565b612ac5565b6104ad610d91366004614a22565b612b75565b348015610da257600080fd5b506104ad610db13660046150b1565b612d8b565b348015610dc257600080fd5b506105de601c5481565b6104ad610dda3660046153a7565b612dbb565b348015610deb57600080fd5b506104f9612e20565b348015610e0057600080fd5b506104ad610e0f3660046153c7565b612e2d565b348015610e2057600080fd5b50610e34610e2f3660046150f8565b612e97565b6040516104db9493929190615403565b348015610e5057600080fd5b506104cf610e5f366004615443565b612ec7565b348015610e7057600080fd5b506104ad610e7f366004614ee7565b612ef5565b348015610e9057600080fd5b50610754610e9f36600461547c565b612f4f565b348015610eb057600080fd5b506104ad610ebf3660046150b1565b61301a565b6104ad610ed23660046154f5565b613055565b348015610ee357600080fd5b506104f9610ef2366004615535565b61307a565b348015610f0357600080fd5b506105de610f12366004614b4c565b60096020526000908152604090205481565b348015610f3057600080fd5b50610f44610f3f366004614b67565b61312b565b604080516001600160501b03998a16815298909716602089015262ffffff958616968801969096529284166060870152921660808501526001600160401b0390911660a084015260c083015260e0820152610100016104db565b337f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03161461101b5760405162461bcd60e51b815260206004820152601e60248201527f4c7a4170703a20696e76616c696420656e64706f696e742063616c6c6572000060448201526064015b60405180910390fd5b61ffff86166000908152600160205260408120805461103990615582565b80601f016020809104026020016040519081016040528092919081815260200182805461106590615582565b80156110b25780601f10611087576101008083540402835291602001916110b2565b820191906000526020600020905b81548152906001019060200180831161109557829003601f168201915b505050505090508051868690501480156110cd575060008151115b80156110f55750805160208201206040516110eb90889088906155bc565b6040518091039020145b6111505760405162461bcd60e51b815260206004820152602660248201527f4c7a4170703a20696e76616c696420736f757263652073656e64696e6720636f6044820152651b9d1c9858dd60d21b6064820152608401611012565b6111c68787878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8a018190048102820181019092528881528a9350915088908890819084018382808284376000920191909152506131a792505050565b50505050505050565b60006001600160e01b031982166322bac5d960e01b14806111f457506111f482613220565b92915050565b6060600d805461120990615582565b80601f016020809104026020016040519081016040528092919081815260200182805461123590615582565b80156112825780601f1061125757610100808354040283529160200191611282565b820191906000526020600020905b81548152906001019060200180831161126557829003601f168201915b5050505050905090565b61129461326e565b6040516307e0db1760e01b815261ffff821660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906307e0db17906024015b600060405180830381600087803b1580156112fa57600080fd5b505af115801561130e573d6000803e3d6000fd5b5050505050565b60006113208261329b565b61133d576040516333d1c03960e21b815260040160405180910390fd5b506000908152601160205260409020546001600160a01b031690565b600061136482611b91565b9050336001600160a01b0382161461139d576113808133612ec7565b61139d576040516367d9dca160e11b815260040160405180910390fd5b60008281526011602052604080822080546001600160a01b0319166001600160a01b0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b61140161326e565b6000811161145d5760405162461bcd60e51b8152602060048201526024808201527f6d696e476173546f5472616e73666572416e6453746f7265206d7573742062656044820152630203e20360e41b6064820152608401611012565b60078190556040518181527ffebbc4f8bb9ec2313950c718d43123124b15778efda4c1f1d529de2995b4f34d906020015b60405180910390a150565b6114a161326e565b61ffff909116600090815260036020526040902055565b6114c061326e565b6040516310ddb13760e01b815261ffff821660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906310ddb137906024016112e0565b630a85bd0160e11b5b949350505050565b3361152b83611b91565b6001600160a01b03161461156d5760405162461bcd60e51b81526020600482015260096024820152682727aa2fa7aba722a960b91b6044820152606401611012565b6000828152601a6020526040902054156115995760405162461bcd60e51b8152600401611012906155cc565b6115ab6201518062ffffff8316615604565b6115b5904261561b565b6000928352601a60205260409092209190915550565b6000818152601a60205260409020544210156115f95760405162461bcd60e51b8152600401611012906155cc565b6116048383836132c3565b505050565b61161161326e565b6013805460ff1916911515919091179055565b60008061163c878761163588613450565b8787612f4f565b915091509550959350505050565b61165261326e565b60185487106116745760405163e82a532960e01b815260040160405180910390fd5b600187106116bd576116bd6001600160401b038416601861169660018b61562e565b815481106116a6576116a6615641565b90600052602060002090600402016003015461349b565b85601888815481106116d1576116d1615641565b906000526020600020906004020160000160006101000a8154816001600160501b0302191690836001600160501b03160217905550846018888154811061171a5761171a615641565b9060005260206000209060040201600001600a6101000a8154816001600160501b0302191690836001600160501b03160217905550836018888154811061176357611763615641565b9060005260206000209060040201600001601a6101000a81548162ffffff021916908362ffffff16021790555082601888815481106117a4576117a4615641565b906000526020600020906004020160010160006101000a8154816001600160401b0302191690836001600160401b0316021790555081601888815481106117ed576117ed615641565b906000526020600020906004020160000160146101000a81548162ffffff021916908362ffffff160217905550806018888154811061182e5761182e615641565b60009182526020909120600490910201805462ffffff60b81b1916600160b81b62ffffff9384160217905561186890610e10908416615604565b61187b906001600160401b03851661561b565b6018888154811061188e5761188e615641565b6000918252602090912060026004909202010155610e106118af8383615657565b62ffffff166118be9190615604565b6118d1906001600160401b03851661561b565b601888815481106118e4576118e4615641565b90600052602060002090600402016003018190555050505050505050565b6014546001600160a01b0316331461191957600080fd5b6040514790600090339083908381818185875af1925050503d806000811461195d576040519150601f19603f3d011682016040523d82523d6000602084013e611962565b606091505b505090508061198457604051631d42c86760e21b815260040160405180910390fd5b6040518281527f5b6b431d4476a211bb7d41c20d1aab9ae2321deee0d20be3d9fc9b1093fa6e3d906020015b60405180910390a15050565b61ffff8316600090815260016020526040812080548291906119dd90615582565b80601f0160208091040260200160405190810160405280929190818152602001828054611a0990615582565b8015611a565780601f10611a2b57610100808354040283529160200191611a56565b820191906000526020600020905b815481529060010190602001808311611a3957829003601f168201915b505050505090508383604051611a6d9291906155bc565b60405180910390208180519060200120149150509392505050565b6116048383836040518060200160405280600081525061286e565b6000818152601a6020526040902054421015611ad15760405162461bcd60e51b8152600401611012906155cc565b611adc8160016134bc565b50565b611ae761326e565b6040516342d65a8d60e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906342d65a8d90611b37908690869086906004016156a3565b600060405180830381600087803b158015611b5157600080fd5b505af11580156111c6573d6000803e3d6000fd5b6111c6878787611b7488613450565b8787876135f5565b611b8461326e565b6015611604828483615707565b60006111f4826137c9565b333014611bfa5760405162461bcd60e51b815260206004820152602660248201527f4e6f6e626c6f636b696e674c7a4170703a2063616c6c6572206d7573742062656044820152650204c7a4170760d41b6064820152608401611012565b611c708686868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f89018190048102820181019092528781528993509150879087908190840183828082843760009201919091525061383092505050565b505050505050565b6000805b601854811015611d125760188181548110611c9957611c99615641565b60009182526020909120600160049092020101546001600160401b0390811690841610801590611cf5575060188181548110611cd757611cd7615641565b906000526020600020906004020160030154836001600160401b0316105b15611d005792915050565b80611d0a816157c6565b915050611c7c565b5060405163e82a532960e01b815260040160405180910390fd5b60006001600160a01b038216611d55576040516323d3ad8160e21b815260040160405180910390fd5b506001600160a01b03166000908152601060205260409020546001600160401b031690565b611d8261326e565b611d8c6000613987565b565b60016020526000908152604090208054611da790615582565b80601f0160208091040260200160405190810160405280929190818152602001828054611dd390615582565b8015611e205780601f10611df557610100808354040283529160200191611e20565b820191906000526020600020905b815481529060010190602001808311611e0357829003601f168201915b505050505081565b611e306139d7565b600080611e7284848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250613a3092505050565b91509150611e81823383613b28565b5050611e8d6001600655565b5050565b611e996139d7565b80516020808301919091206000818152600a90925260409091206002015460ff16611efa5760405162461bcd60e51b81526020600482015260116024820152701b9bc818dc99591a5d1cc81cdd1bdc9959607a1b6044820152606401611012565b600082806020019051810190611f109190615824565b6000848152600a602052604081208054600190910154929450909250611f4c9161ffff8216916201000090046001600160a01b03169085613dcf565b6000848152600a60205260409020600101549091508111611fc15760405162461bcd60e51b815260206004820152602960248201527f6e6f7420656e6f7567682067617320746f2070726f6365737320637265646974604482015268103a3930b739b332b960b91b6064820152608401611012565b81518103612038576000838152600a602052604080822080546001600160b01b031916815560018101929092556002909101805460ff19169055517fd7be02b8dd0d27bd0517a9cb4d7469ce27df4313821ae5ec1ff69acc594ba2339061202b9085815260200190565b60405180910390a16120cc565b604080516080810182526000858152600a6020818152848320805461ffff80821687526001600160a01b03620100008084048216868a019081529989018b8152600160608b01818152998f90529790965297519851169096026001600160b01b03199091169690951695909517939093178455915191830191909155516002909101805491151560ff199092169190911790555b505050611adc6001600655565b6060600e805461120990615582565b60158054611da790615582565b6120fd61326e565b600081116121585760405162461bcd60e51b815260206004820152602260248201527f647374436861696e4964546f42617463684c696d6974206d757374206265203e604482015261020360f41b6064820152608401611012565b61ffff8216600081815260086020908152604091829020849055815192835282018390527f7315f7654d594ead24a30160ed9ba2d23247f543016b918343591e93d7afdb6d91016119b0565b61ffff81166000908152600160205260408120805460609291906121c790615582565b80601f01602080910402602001604051908101604052809291908181526020018280546121f390615582565b80156122405780601f1061221557610100808354040283529160200191612240565b820191906000526020600020905b81548152906001019060200180831161222357829003601f168201915b5050505050905080516000036122985760405162461bcd60e51b815260206004820152601d60248201527f4c7a4170703a206e6f20747275737465642070617468207265636f72640000006044820152606401611012565b6122b36000601483516122ab919061562e565b839190613e1b565b9392505050565b3360008181526012602090815260408083206001600160a01b03871680855290835292819020805460ff191686151590811790915590519081529192917f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a35050565b6000612331600b5490565b905090565b61233e61326e565b818130604051602001612353939291906158de565b60408051601f1981840301815291815261ffff851660009081526001602052209061237e9082615904565b507f8c0400cfe2d1199b1a725c78960bcc2a344d869b80590d0f2bd005db15a572ce8383836040516123b2939291906156a3565b60405180910390a1505050565b6123c761326e565b6016611604828483615707565b6111c6878787878787876135f5565b6123eb61326e565b8060008190036124285760405162461bcd60e51b81526020600482015260086024820152674e4f5f535441474560c01b6044820152606401611012565b60005b60185481101561249b576018805480612446576124466159c3565b60008281526020812060046000199093019283020180546001600160e81b031916815560018101805467ffffffffffffffff191690556002810182905560030155905580612493816157c6565b91505061242b565b5060005b8281101561286857600181106124f2576124f28484838181106124c4576124c4615641565b905060c0020160a00160208101906124dc9190615094565b6001600160401b0316601861169660018561562e565b601860405180610100016040528086868581811061251257612512615641565b61252892602060c09092020190810191506159d9565b6001600160501b0316815260200186868581811061254857612548615641565b905060c00201602001602081019061256091906159d9565b6001600160501b0316815260200186868581811061258057612580615641565b905060c00201604001602081019061259891906159f4565b62ffffff1681526020018686858181106125b4576125b4615641565b905060c0020160600160208101906125cc91906159f4565b62ffffff1681526020018686858181106125e8576125e8615641565b905060c00201608001602081019061260091906159f4565b62ffffff16815260200186868581811061261c5761261c615641565b905060c0020160a00160208101906126349190615094565b6001600160401b03168152602001610e1087878681811061265757612657615641565b905060c00201604001602081019061266f91906159f4565b62ffffff1661267e9190615604565b87878681811061269057612690615641565b905060c0020160a00160208101906126a89190615094565b6001600160401b03166126bb919061561b565b8152602001610e108787868181106126d5576126d5615641565b905060c0020160400160208101906126ed91906159f4565b8888878181106126ff576126ff615641565b905060c00201606001602081019061271791906159f4565b6127219190615657565b62ffffff166127309190615604565b87878681811061274257612742615641565b905060c0020160a001602081019061275a9190615094565b6001600160401b031661276d919061561b565b9052815460018082018455600093845260209384902083516004909302018054948401516040850151606086015160808701516001600160501b039687166001600160a01b031990991698909817600160501b96909316959095029190911765ffffffffffff60a01b1916600160a01b62ffffff9283160262ffffff60b81b191617600160b81b948216949094029390931762ffffff60d01b1916600160d01b939095169290920293909317815560a0820151928101805467ffffffffffffffff19166001600160401b039094169390931790925560c0810151600283015560e0015160039091015580612860816157c6565b91505061249f565b50505050565b6128798484846115cb565b6001600160a01b0383163b156128685761289584848484613f28565b612868576040516368d2bf6b60e11b815260040160405180910390fd5b6128ba61326e565b600480546001600160a01b0319166001600160a01b0383169081179091556040519081527f5db758e995a17ec1ad84bdef7e8c3293a0bd6179bcce400dff5d4c3d87db726b9060200161148e565b61291061326e565b612921603c63ffffffff8316615604565b60175550565b60606129328261329b565b61294f57604051630a14c4b560e41b815260040160405180910390fd5b60006015805461295e90615582565b80601f016020809104026020016040519081016040528092919081815260200182805461298a90615582565b80156129d75780601f106129ac576101008083540402835291602001916129d7565b820191906000526020600020905b8154815290600101906020018083116129ba57829003601f168201915b5050505050905080516000036129fc57604051806020016040528060008152506122b3565b80612a0684614010565b6016604051602001612a1a93929190615a0f565b6040516020818303038152906040529392505050565b612a3861326e565b6040516332fb62e760e21b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063cbed8b9c90612a8c9088908890889088908890600401615aaf565b600060405180830381600087803b158015612aa657600080fd5b505af1158015612aba573d6000803e3d6000fd5b505050505050505050565b612acd61326e565b60008111612b295760405162461bcd60e51b815260206004820152602360248201527f647374436861696e4964546f5472616e73666572476173206d7573742062652060448201526203e20360ec1b6064820152608401611012565b61ffff8216600081815260096020908152604091829020849055815192835282018390527fc46df2983228ac2d9754e94a0d565e6671665dc8ad38602bc8e544f0685a29fb91016119b0565b61ffff86166000908152600560205260408082209051612b9890889088906155bc565b90815260408051602092819003830190206001600160401b03871660009081529252902054905080612c185760405162461bcd60e51b815260206004820152602360248201527f4e6f6e626c6f636b696e674c7a4170703a206e6f2073746f726564206d65737360448201526261676560e81b6064820152608401611012565b808383604051612c299291906155bc565b604051809103902014612c885760405162461bcd60e51b815260206004820152602160248201527f4e6f6e626c6f636b696e674c7a4170703a20696e76616c6964207061796c6f616044820152601960fa1b6064820152608401611012565b61ffff87166000908152600560205260408082209051612cab90899089906155bc565b90815260408051602092819003830181206001600160401b038916600090815290845282902093909355601f88018290048202830182019052868252612d43918991899089908190840183828082843760009201919091525050604080516020601f8a018190048102820181019092528881528a93509150889088908190840183828082843760009201919091525061383092505050565b7fc264d91f3adc5588250e1551f547752ca0cfa8f6b530d243b9f9f4cab10ea8e58787878785604051612d7a959493929190615add565b60405180910390a150505050505050565b612d9361326e565b601380546001600160a01b0390921661010002610100600160a81b0319909216919091179055565b612dc36139d7565b600080612e0584848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250613a3092505050565b91509150612e14828683613b28565b50506116046001600655565b60168054611da790615582565b612e3561326e565b61ffff83811660008181526002602090815260408083209487168084529482529182902085905581519283528201929092529081018290527f9d5c7c0b934da8fefa9c7760c98383778a12dfbfc0c3b3106518f43fb9508ac0906060016123b2565b600060606000606084806020019051810190612eb39190615b23565b939c919b5099509197509095505050505050565b6001600160a01b03918216600090815260126020908152604080832093909416825291909152205460ff1690565b612efd61326e565b61ffff83166000908152600160205260409020612f1b828483615707565b507ffa41487ad5d6728f0b19276fa1eddc16558578f5109fc39d2dc33c3230470dab8383836040516123b2939291906156a3565b60008060008686604051602001612f67929190615c34565b60408051601f198184030181529082905263040a7bb160e41b825291506001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906340a7bb1090612fcb908b90309086908b908b90600401615c59565b6040805180830381865afa158015612fe7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061300b9190615cad565b92509250509550959350505050565b61302261326e565b6001600160a01b03811661304c57604051631e4fbdf760e01b815260006004820152602401611012565b611adc81613987565b61305d61326e565b6130656139d7565b613070828483613b28565b6116046001600655565b604051633d7b2f6f60e21b815261ffff808616600483015284166024820152306044820152606481018290526060907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063f5ecbdbc90608401600060405180830381865afa1580156130fa573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526131229190810190615cd1565b95945050505050565b6018818154811061313b57600080fd5b600091825260209091206004909102018054600182015460028301546003909301546001600160501b038084169550600160501b8404169362ffffff600160a01b8504811694600160b81b8104821694600160d01b909104909116926001600160401b03909116919088565b60008061320a5a60966366ad5c8a60e01b898989896040516024016131cf9493929190615d05565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b03199093169290921790915230929190614054565b9150915081611c7057611c7086868686856140de565b60006301ffc9a760e01b6001600160e01b03198316148061325157506380ac58cd60e01b6001600160e01b03198316145b806111f45750506001600160e01b031916635b5e139f60e01b1490565b6000546001600160a01b03163314611d8c5760405163118cdaa760e01b8152336004820152602401611012565b6000600b54821080156111f45750506000908152600f6020526040902054600160e01b161590565b60006132ce826137c9565b9050836001600160a01b0316816001600160a01b0316146133015760405162a1148160e81b815260040160405180910390fd5b6000828152601160205260409020805461332d8187335b6001600160a01b039081169116811491141790565b6133585761333b8633612ec7565b61335857604051632ce44b5f60e11b815260040160405180910390fd5b6001600160a01b03851661337f57604051633a954ecd60e21b815260040160405180910390fd5b801561338a57600082555b6001600160a01b038681166000908152601060205260408082208054600019019055918716808252919020805460010190554260a01b17600160e11b176000858152600f6020526040812091909155600160e11b8416900361341c57600184016000818152600f6020526040812054900361341a57600b54811461341a576000818152600f602052604090208490555b505b83856001600160a01b0316876001600160a01b0316600080516020615f2f83398151915260405160405180910390a4611c70565b6040805160018082528183019092526060916000919060208083019080368337019050509050828160008151811061348a5761348a615641565b602090810291909101015292915050565b80821015611e8d57604051631750215560e11b815260040160405180910390fd5b60006134c7836137c9565b9050806000806134e586600090815260116020526040902080549091565b915091508415613525576134fa818433613318565b613525576135088333612ec7565b61352557604051632ce44b5f60e11b815260040160405180910390fd5b801561353057600082555b6001600160a01b038316600081815260106020526040902080546fffffffffffffffffffffffffffffffff0190554260a01b17600360e01b176000878152600f6020526040812091909155600160e11b851690036135be57600186016000818152600f602052604081205490036135bc57600b5481146135bc576000818152600f602052604090208590555b505b60405186906000906001600160a01b03861690600080516020615f2f833981519152908390a45050600c8054600101905550505050565b600084511161363c5760405162461bcd60e51b8152602060048201526013602482015272746f6b656e4964735b5d20697320656d70747960681b6044820152606401611012565b835160011480613660575061ffff8616600090815260086020526040902054845111155b6136b75760405162461bcd60e51b815260206004820152602260248201527f62617463682073697a65206578636565647320647374206261746368206c696d6044820152611a5d60f21b6064820152608401611012565b60005b84518110156136fa576136e88888888885815181106136db576136db615641565b602002602001015161417b565b806136f2816157c6565b9150506136ba565b5060008585604051602001613710929190615c34565b6040516020818303038152906040529050613755876001848851600960008d61ffff1661ffff168152602001908152602001600020546137509190615604565b6141b4565b613763878286868634614289565b856040516137719190615d22565b6040518091039020886001600160a01b03168861ffff167fe1b87c47fdeb4f9cbadbca9df3af7aba453bb6e501075d0440d88125b711522a886040516137b79190615d3e565b60405180910390a45050505050505050565b600081600b54811015613817576000818152600f602052604081205490600160e01b82169003613815575b806000036122b35750600019016000818152600f60205260409020546137f4565b505b604051636f96cda160e11b815260040160405180910390fd5b600080828060200190518101906138479190615824565b60148201519193509150600061385f88838386613dcf565b905082518110156139335784516020808701919091206040805160808101825261ffff808d1682526001600160a01b038088168387019081528385018881526001606086018181526000898152600a909a529887902095518654935190941662010000026001600160b01b03199093169390941692909217178355519082015592516002909301805493151560ff199094169390931790925590517f10e0b70d256bccc84b7027506978bd8b68984a870788b93b479def144c839ad7906139299083908990615d51565b60405180910390a1505b816001600160a01b03168760405161394b9190615d22565b60405180910390208961ffff167f5b821db8a46f8ecbe1941ba2f51cfeea9643268b56631f70d45e2a745d990265866040516137b79190615d3e565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600260065403613a295760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401611012565b6002600655565b600080600080600080613a4287612e97565b9350935093509350601b83604051613a5a9190615d22565b9081526040519081900360200190205460ff1615613aa35760405162461bcd60e51b815260040161101290602080825260049082015263155cd95960e21b604082015260600190565b6001601b84604051613ab59190615d22565b908152604051602091819003919091019020805460ff1916911515919091179055601354613afe9061010090046001600160a01b0316613af8338746888861442e565b836144ce565b613b1b576040516302129fe960e11b815260040160405180910390fd5b5091959194509092505050565b601c548363ffffffff16613b3a612326565b613b44919061561b565b1115613b7d5760405162461bcd60e51b815260206004820152600860248201526709eeccae4be9ac2f60c31b6044820152606401611012565b60135460ff16613ba057604051630952c8a960e11b815260040160405180910390fd5b6000613bab82611c78565b90504260188281548110613bc157613bc1615641565b9060005260206000209060040201600301541015613bf257604051633402a54760e11b815260040160405180910390fd5b60188181548110613c0557613c05615641565b60009182526020808320600492909202909101548383526019909152604090912054600160d01b90910462ffffff1690613c469063ffffffff87169061561b565b1115613c65576040516371dd1daf60e11b815260040160405180910390fd5b600060188281548110613c7a57613c7a615641565b9060005260206000209060040201600001600a9054906101000a90046001600160501b03166001600160501b0316905060188281548110613cbd57613cbd615641565b906000526020600020906004020160020154836001600160401b031611613d53574260188381548110613cf257613cf2615641565b9060005260206000209060040201600201541015613d235760405163ab7a6d3b60e01b815260040160405180910390fd5b60188281548110613d3657613d36615641565b60009182526020909120600490910201546001600160501b031690505b6000546001600160a01b03163314613d9557613d7563ffffffff861682615604565b341015613d9557604051630717c22560e51b815260040160405180910390fd5b6000828152601960205260408120805463ffffffff88169290613db990849061561b565b9091555061130e90508463ffffffff8716614530565b6000825b8251811015613122576007545a1061312257613e098686858481518110613dfc57613dfc615641565b602002602001015161454a565b80613e13816157c6565b915050613dd3565b606081613e2981601f61561b565b1015613e685760405162461bcd60e51b815260206004820152600e60248201526d736c6963655f6f766572666c6f7760901b6044820152606401611012565b613e72828461561b565b84511015613eb65760405162461bcd60e51b8152602060048201526011602482015270736c6963655f6f75744f66426f756e647360781b6044820152606401611012565b606082158015613ed55760405191506000825260208201604052613f1f565b6040519150601f8416801560200281840101858101878315602002848b0101015b81831015613f0e578051835260209283019201613ef6565b5050858452601f01601f1916604052505b50949350505050565b604051630a85bd0160e11b81526000906001600160a01b0385169063150b7a0290613f5d903390899088908890600401615d6a565b6020604051808303816000875af1925050508015613f98575060408051601f3d908101601f19168201909252613f9591810190615d9d565b60015b613ff6573d808015613fc6576040519150601f19603f3d011682016040523d82523d6000602084013e613fcb565b606091505b508051600003613fee576040516368d2bf6b60e11b815260040160405180910390fd5b805181602001fd5b6001600160e01b031916630a85bd0160e11b149050611519565b606060a06040510180604052602081039150506000815280825b600183039250600a81066030018353600a90048061402a5750819003601f19909101908152919050565b6000606060008060008661ffff166001600160401b0381111561407957614079614be8565b6040519080825280601f01601f1916602001820160405280156140a3576020820181803683370190505b50905060008087516020890160008d8df191503d9250868311156140c5578692505b828152826000602083013e909890975095505050505050565b8180519060200120600560008761ffff1661ffff1681526020019081526020016000208560405161410f9190615d22565b9081526040805191829003602090810183206001600160401b0388166000908152915220919091557fe183f33de2837795525b4792ca4cd60535bd77c53b7e7030060bfcf5734d6b0c9061416c9087908790879087908790615dba565b60405180910390a15050505050565b6000818152601a60205260409020544210156141a95760405162461bcd60e51b8152600401611012906155cc565b612868843083611a88565b60006141bf836145aa565b61ffff808716600090815260026020908152604080832093891683529290522054909150806142305760405162461bcd60e51b815260206004820152601a60248201527f4c7a4170703a206d696e4761734c696d6974206e6f74207365740000000000006044820152606401611012565b61423a838261561b565b821015611c705760405162461bcd60e51b815260206004820152601b60248201527f4c7a4170703a20676173206c696d697420697320746f6f206c6f7700000000006044820152606401611012565b61ffff8616600090815260016020526040812080546142a790615582565b80601f01602080910402602001604051908101604052809291908181526020018280546142d390615582565b80156143205780601f106142f557610100808354040283529160200191614320565b820191906000526020600020905b81548152906001019060200180831161430357829003601f168201915b5050505050905080516000036143915760405162461bcd60e51b815260206004820152603060248201527f4c7a4170703a2064657374696e6174696f6e20636861696e206973206e6f742060448201526f61207472757374656420736f7572636560801b6064820152608401611012565b61439c878751614606565b60405162c5803160e81b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063c58031009084906143f3908b9086908c908c908c908c90600401615e0c565b6000604051808303818588803b15801561440c57600080fd5b505af1158015614420573d6000803e3d6000fd5b505050505050505050505050565b60135460405160009182916144609130918a9161010090046001600160a01b0316908a908a908a908a90602001615e73565b60408051601f1981840301815282825280516020918201207f19457468657265756d205369676e6564204d6573736167653a0a33320000000084830152603c80850182905283518086039091018152605c90940190925282519201919091209091505b979650505050505050565b60008060006144dd8585614677565b50909250905060008160038111156144f7576144f7615eff565b1480156145155750856001600160a01b0316826001600160a01b0316145b8061452657506145268686866146c4565b9695505050505050565b611e8d82826040518060200160405280600081525061479f565b6145538161329b565b158061457f57506145638161329b565b801561457f57503061457482611b91565b6001600160a01b0316145b61458857600080fd5b6145918161329b565b61459f576116048282614530565b611604308383611a88565b60006022825110156145fe5760405162461bcd60e51b815260206004820152601c60248201527f4c7a4170703a20696e76616c69642061646170746572506172616d73000000006044820152606401611012565b506022015190565b61ffff82166000908152600360205260408120549081900361462757506127105b808211156116045760405162461bcd60e51b815260206004820181905260248201527f4c7a4170703a207061796c6f61642073697a6520697320746f6f206c617267656044820152606401611012565b600080600083516041036146b15760208401516040850151606086015160001a6146a388828585614805565b9550955095505050506146bd565b50508151600091506002905b9250925092565b6000806000856001600160a01b031685856040516024016146e6929190615d51565b60408051601f198184030181529181526020820180516001600160e01b0316630b135d3f60e11b1790525161471b9190615d22565b600060405180830381855afa9150503d8060008114614756576040519150601f19603f3d011682016040523d82523d6000602084013e61475b565b606091505b509150915081801561476f57506020815110155b801561452657508051630b135d3f60e11b906147949083016020908101908401615f15565b149695505050505050565b6147a983836148d4565b6001600160a01b0383163b1561160457600b548281035b6147d36000868380600101945086613f28565b6147f0576040516368d2bf6b60e11b815260040160405180910390fd5b8181106147c05781600b541461130e57600080fd5b600080807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084111561484057506000915060039050826148ca565b604080516000808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015614894573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166148c0575060009250600191508290506148ca565b9250600091508190505b9450945094915050565b600b5460008290036148f95760405163b562e8dd60e01b815260040160405180910390fd5b6001600160a01b03831660008181526010602090815260408083208054680100000000000000018802019055848352600f90915281206001851460e11b4260a01b17831790558284019083908390600080516020615f2f8339815191528180a4600183015b8181146149845780836000600080516020615f2f833981519152600080a460010161495e565b50816000036149a557604051622e076360e81b815260040160405180910390fd5b600b5550505050565b803561ffff811681146149c057600080fd5b919050565b60008083601f8401126149d757600080fd5b5081356001600160401b038111156149ee57600080fd5b602083019150836020828501011115614a0657600080fd5b9250929050565b6001600160401b0381168114611adc57600080fd5b60008060008060008060808789031215614a3b57600080fd5b614a44876149ae565b955060208701356001600160401b0380821115614a6057600080fd5b614a6c8a838b016149c5565b909750955060408901359150614a8182614a0d565b90935060608801359080821115614a9757600080fd5b50614aa489828a016149c5565b979a9699509497509295939492505050565b6001600160e01b031981168114611adc57600080fd5b600060208284031215614ade57600080fd5b81356122b381614ab6565b60005b83811015614b04578181015183820152602001614aec565b50506000910152565b60008151808452614b25816020860160208601614ae9565b601f01601f19169290920160200192915050565b6020815260006122b36020830184614b0d565b600060208284031215614b5e57600080fd5b6122b3826149ae565b600060208284031215614b7957600080fd5b5035919050565b6001600160a01b0381168114611adc57600080fd5b80356149c081614b80565b60008060408385031215614bb357600080fd5b8235614bbe81614b80565b946020939093013593505050565b60008060408385031215614bdf57600080fd5b614bbe836149ae565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715614c2657614c26614be8565b604052919050565b60006001600160401b03821115614c4757614c47614be8565b50601f01601f191660200190565b600082601f830112614c6657600080fd5b8135614c79614c7482614c2e565b614bfe565b818152846020838601011115614c8e57600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008060808587031215614cc157600080fd5b8435614ccc81614b80565b93506020850135614cdc81614b80565b92506040850135915060608501356001600160401b03811115614cfe57600080fd5b614d0a87828801614c55565b91505092959194509250565b803562ffffff811681146149c057600080fd5b60008060408385031215614d3c57600080fd5b82359150614d4c60208401614d16565b90509250929050565b600080600060608486031215614d6a57600080fd5b8335614d7581614b80565b92506020840135614d8581614b80565b929592945050506040919091013590565b803580151581146149c057600080fd5b600060208284031215614db857600080fd5b6122b382614d96565b600080600080600060a08688031215614dd957600080fd5b614de2866149ae565b945060208601356001600160401b0380821115614dfe57600080fd5b614e0a89838a01614c55565b955060408801359450614e1f60608901614d96565b93506080880135915080821115614e3557600080fd5b50614e4288828901614c55565b9150509295509295909350565b80356001600160501b03811681146149c057600080fd5b600080600080600080600060e0888a031215614e8157600080fd5b87359650614e9160208901614e4f565b9550614e9f60408901614e4f565b9450614ead60608901614d16565b93506080880135614ebd81614a0d565b9250614ecb60a08901614d16565b9150614ed960c08901614d16565b905092959891949750929550565b600080600060408486031215614efc57600080fd5b614f05846149ae565b925060208401356001600160401b03811115614f2057600080fd5b614f2c868287016149c5565b9497909650939450505050565b600080600080600080600060e0888a031215614f5457600080fd5b8735614f5f81614b80565b9650614f6d602089016149ae565b955060408801356001600160401b0380821115614f8957600080fd5b614f958b838c01614c55565b965060608a0135955060808a01359150614fae82614b80565b90935060a089013590614fc082614b80565b90925060c08901359080821115614fd657600080fd5b50614fe38a828b01614c55565b91505092959891949750929550565b6000806020838503121561500557600080fd5b82356001600160401b0381111561501b57600080fd5b615027858286016149c5565b90969095509350505050565b60008060006060848603121561504857600080fd5b615051846149ae565b925060208401356001600160401b0381111561506c57600080fd5b61507886828701614c55565b925050604084013561508981614a0d565b809150509250925092565b6000602082840312156150a657600080fd5b81356122b381614a0d565b6000602082840312156150c357600080fd5b81356122b381614b80565b600080604083850312156150e157600080fd5b6150ea836149ae565b9150614d4c602084016149ae565b60006020828403121561510a57600080fd5b81356001600160401b0381111561512057600080fd5b61151984828501614c55565b6000806040838503121561513f57600080fd5b823561514a81614b80565b9150614d4c60208401614d96565b60006001600160401b0382111561517157615171614be8565b5060051b60200190565b600082601f83011261518c57600080fd5b8135602061519c614c7483615158565b82815260059290921b840181019181810190868411156151bb57600080fd5b8286015b848110156151d657803583529183019183016151bf565b509695505050505050565b600080600080600080600060e0888a0312156151fc57600080fd5b873561520781614b80565b9650615215602089016149ae565b955060408801356001600160401b038082111561523157600080fd5b61523d8b838c01614c55565b965060608a013591508082111561525357600080fd5b61525f8b838c0161517b565b955060808a0135915061527182614b80565b81945061528060a08b01614b95565b935060c08a0135915080821115614fd657600080fd5b600080602083850312156152a957600080fd5b82356001600160401b03808211156152c057600080fd5b818501915085601f8301126152d457600080fd5b8135818111156152e357600080fd5b86602060c0830285010111156152f857600080fd5b60209290920196919550909350505050565b63ffffffff81168114611adc57600080fd5b60006020828403121561532e57600080fd5b81356122b38161530a565b60008060008060006080868803121561535157600080fd5b61535a866149ae565b9450615368602087016149ae565b93506040860135925060608601356001600160401b0381111561538a57600080fd5b615396888289016149c5565b969995985093965092949392505050565b6000806000604084860312156153bc57600080fd5b8335614f0581614b80565b6000806000606084860312156153dc57600080fd5b6153e5846149ae565b92506153f3602085016149ae565b9150604084013590509250925092565b63ffffffff851681526080602082015260006154226080830186614b0d565b6001600160401b038516604084015282810360608401526144c38185614b0d565b6000806040838503121561545657600080fd5b823561546181614b80565b9150602083013561547181614b80565b809150509250929050565b600080600080600060a0868803121561549457600080fd5b61549d866149ae565b945060208601356001600160401b03808211156154b957600080fd5b6154c589838a01614c55565b955060408801359150808211156154db57600080fd5b6154e789838a0161517b565b9450614e1f60608901614d96565b60008060006060848603121561550a57600080fd5b833561551581614b80565b925060208401356155258161530a565b9150604084013561508981614a0d565b6000806000806080858703121561554b57600080fd5b615554856149ae565b9350615562602086016149ae565b9250604085013561557281614b80565b9396929550929360600135925050565b600181811c9082168061559657607f821691505b6020821081036155b657634e487b7160e01b600052602260045260246000fd5b50919050565b8183823760009101908152919050565b602080825260089082015267494e5f5354414b4560c01b604082015260600190565b634e487b7160e01b600052601160045260246000fd5b80820281158282048414176111f4576111f46155ee565b808201808211156111f4576111f46155ee565b818103818111156111f4576111f46155ee565b634e487b7160e01b600052603260045260246000fd5b62ffffff818116838216019080821115615673576156736155ee565b5092915050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b61ffff8416815260406020820152600061312260408301848661567a565b601f82111561160457600081815260208120601f850160051c810160208610156156e85750805b601f850160051c820191505b81811015611c70578281556001016156f4565b6001600160401b0383111561571e5761571e614be8565b6157328361572c8354615582565b836156c1565b6000601f841160018114615766576000851561574e5750838201355b600019600387901b1c1916600186901b17835561130e565b600083815260209020601f19861690835b828110156157975786850135825560209485019460019092019101615777565b50868210156157b45760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b6000600182016157d8576157d86155ee565b5060010190565b600082601f8301126157f057600080fd5b81516157fe614c7482614c2e565b81815284602083860101111561581357600080fd5b611519826020830160208701614ae9565b6000806040838503121561583757600080fd5b82516001600160401b038082111561584e57600080fd5b61585a868387016157df565b935060209150818501518181111561587157600080fd5b85019050601f8101861361588457600080fd5b8051615892614c7482615158565b81815260059190911b820183019083810190888311156158b157600080fd5b928401925b828410156158cf578351825292840192908401906158b6565b80955050505050509250929050565b8284823760609190911b6bffffffffffffffffffffffff19169101908152601401919050565b81516001600160401b0381111561591d5761591d614be8565b6159318161592b8454615582565b846156c1565b602080601f831160018114615966576000841561594e5750858301515b600019600386901b1c1916600185901b178555611c70565b600085815260208120601f198616915b8281101561599557888601518255948401946001909101908401615976565b50858210156159b35787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052603160045260246000fd5b6000602082840312156159eb57600080fd5b6122b382614e4f565b600060208284031215615a0657600080fd5b6122b382614d16565b600084516020615a228285838a01614ae9565b855191840191615a358184848a01614ae9565b8554920191600090615a4681615582565b60018281168015615a5e5760018114615a7357615a9f565b60ff1984168752821515830287019450615a9f565b896000528560002060005b84811015615a9757815489820152908301908701615a7e565b505082870194505b50929a9950505050505050505050565b600061ffff8088168352808716602084015250846040830152608060608301526144c360808301848661567a565b61ffff86168152608060208201526000615afb60808301868861567a565b6001600160401b0394909416604083015250606001529392505050565b80516149c081614a0d565b600080600080600080600080610100898b031215615b4057600080fd5b8851615b4b81614b80565b60208a0151909850615b5c81614b80565b60408a0151909750615b6d81614b80565b60608a0151909650615b7e8161530a565b60808a0151909550615b8f8161530a565b60a08a01519094506001600160401b0380821115615bac57600080fd5b615bb88c838d016157df565b9450615bc660c08c01615b18565b935060e08b0151915080821115615bdc57600080fd5b50615be98b828c016157df565b9150509295985092959890939650565b600081518084526020808501945080840160005b83811015615c2957815187529582019590820190600101615c0d565b509495945050505050565b604081526000615c476040830185614b0d565b82810360208401526131228185615bf9565b61ffff861681526001600160a01b038516602082015260a060408201819052600090615c8790830186614b0d565b84151560608401528281036080840152615ca18185614b0d565b98975050505050505050565b60008060408385031215615cc057600080fd5b505080516020909101519092909150565b600060208284031215615ce357600080fd5b81516001600160401b03811115615cf957600080fd5b611519848285016157df565b61ffff851681526080602082015260006154226080830186614b0d565b60008251615d34818460208701614ae9565b9190910192915050565b6020815260006122b36020830184615bf9565b8281526040602082015260006115196040830184614b0d565b6001600160a01b038581168252841660208201526040810183905260806060820181905260009061452690830184614b0d565b600060208284031215615daf57600080fd5b81516122b381614ab6565b61ffff8616815260a060208201526000615dd760a0830187614b0d565b6001600160401b03861660408401528281036060840152615df88186614b0d565b90508281036080840152615ca18185614b0d565b61ffff8716815260c060208201526000615e2960c0830188614b0d565b8281036040840152615e3b8188614b0d565b6001600160a01b0387811660608601528616608085015283810360a08501529050615e668185614b0d565b9998505050505050505050565b60006bffffffffffffffffffffffff19808a60601b168352808960601b166014840152808860601b1660288401525063ffffffff60e01b808760e01b16603c840152808660e01b166040840152508351615ed4816044850160208801614ae9565b60c09390931b6001600160c01b03191660449290930191820192909252604c01979650505050505050565b634e487b7160e01b600052602160045260246000fd5b600060208284031215615f2757600080fd5b505191905056feddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa264697066735822122085457782e025ea537b1fca5b30ba288776df33e7943aaf36d3669a5d7e95da9664736f6c63430008140033",
}

// EfesABI is the input ABI used to generate the binding from.
// Deprecated: Use EfesMetaData.ABI instead.
var EfesABI = EfesMetaData.ABI

// EfesBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EfesMetaData.Bin instead.
var EfesBin = EfesMetaData.Bin

// DeployEfes deploys a new Ethereum contract, binding an instance of Efes to it.
func DeployEfes(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string, _tokenURISuffix string, _currentBaseURI string, _cosigner common.Address, _beneficiary common.Address, _minGasToTransferAndStore *big.Int, _lzEndpoint common.Address, _expire uint32, _maxSupply *big.Int) (common.Address, *types.Transaction, *Efes, error) {
	parsed, err := EfesMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EfesBin), backend, _name, _symbol, _tokenURISuffix, _currentBaseURI, _cosigner, _beneficiary, _minGasToTransferAndStore, _lzEndpoint, _expire, _maxSupply)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Efes{EfesCaller: EfesCaller{contract: contract}, EfesTransactor: EfesTransactor{contract: contract}, EfesFilterer: EfesFilterer{contract: contract}}, nil
}

// Efes is an auto generated Go binding around an Ethereum contract.
type Efes struct {
	EfesCaller     // Read-only binding to the contract
	EfesTransactor // Write-only binding to the contract
	EfesFilterer   // Log filterer for contract events
}

// EfesCaller is an auto generated read-only Go binding around an Ethereum contract.
type EfesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EfesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EfesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EfesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EfesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EfesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EfesSession struct {
	Contract     *Efes             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EfesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EfesCallerSession struct {
	Contract *EfesCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EfesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EfesTransactorSession struct {
	Contract     *EfesTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EfesRaw is an auto generated low-level Go binding around an Ethereum contract.
type EfesRaw struct {
	Contract *Efes // Generic contract binding to access the raw methods on
}

// EfesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EfesCallerRaw struct {
	Contract *EfesCaller // Generic read-only contract binding to access the raw methods on
}

// EfesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EfesTransactorRaw struct {
	Contract *EfesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEfes creates a new instance of Efes, bound to a specific deployed contract.
func NewEfes(address common.Address, backend bind.ContractBackend) (*Efes, error) {
	contract, err := bindEfes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Efes{EfesCaller: EfesCaller{contract: contract}, EfesTransactor: EfesTransactor{contract: contract}, EfesFilterer: EfesFilterer{contract: contract}}, nil
}

// NewEfesCaller creates a new read-only instance of Efes, bound to a specific deployed contract.
func NewEfesCaller(address common.Address, caller bind.ContractCaller) (*EfesCaller, error) {
	contract, err := bindEfes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EfesCaller{contract: contract}, nil
}

// NewEfesTransactor creates a new write-only instance of Efes, bound to a specific deployed contract.
func NewEfesTransactor(address common.Address, transactor bind.ContractTransactor) (*EfesTransactor, error) {
	contract, err := bindEfes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EfesTransactor{contract: contract}, nil
}

// NewEfesFilterer creates a new log filterer instance of Efes, bound to a specific deployed contract.
func NewEfesFilterer(address common.Address, filterer bind.ContractFilterer) (*EfesFilterer, error) {
	contract, err := bindEfes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EfesFilterer{contract: contract}, nil
}

// bindEfes binds a generic wrapper to an already deployed contract.
func bindEfes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EfesMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Efes *EfesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Efes.Contract.EfesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Efes *EfesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Efes.Contract.EfesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Efes *EfesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Efes.Contract.EfesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Efes *EfesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Efes.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Efes *EfesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Efes.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Efes *EfesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Efes.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTPAYLOADSIZELIMIT is a free data retrieval call binding the contract method 0xc4461834.
//
// Solidity: function DEFAULT_PAYLOAD_SIZE_LIMIT() view returns(uint256)
func (_Efes *EfesCaller) DEFAULTPAYLOADSIZELIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "DEFAULT_PAYLOAD_SIZE_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTPAYLOADSIZELIMIT is a free data retrieval call binding the contract method 0xc4461834.
//
// Solidity: function DEFAULT_PAYLOAD_SIZE_LIMIT() view returns(uint256)
func (_Efes *EfesSession) DEFAULTPAYLOADSIZELIMIT() (*big.Int, error) {
	return _Efes.Contract.DEFAULTPAYLOADSIZELIMIT(&_Efes.CallOpts)
}

// DEFAULTPAYLOADSIZELIMIT is a free data retrieval call binding the contract method 0xc4461834.
//
// Solidity: function DEFAULT_PAYLOAD_SIZE_LIMIT() view returns(uint256)
func (_Efes *EfesCallerSession) DEFAULTPAYLOADSIZELIMIT() (*big.Int, error) {
	return _Efes.Contract.DEFAULTPAYLOADSIZELIMIT(&_Efes.CallOpts)
}

// FUNCTIONTYPESEND is a free data retrieval call binding the contract method 0xaf3fb21c.
//
// Solidity: function FUNCTION_TYPE_SEND() view returns(uint16)
func (_Efes *EfesCaller) FUNCTIONTYPESEND(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "FUNCTION_TYPE_SEND")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// FUNCTIONTYPESEND is a free data retrieval call binding the contract method 0xaf3fb21c.
//
// Solidity: function FUNCTION_TYPE_SEND() view returns(uint16)
func (_Efes *EfesSession) FUNCTIONTYPESEND() (uint16, error) {
	return _Efes.Contract.FUNCTIONTYPESEND(&_Efes.CallOpts)
}

// FUNCTIONTYPESEND is a free data retrieval call binding the contract method 0xaf3fb21c.
//
// Solidity: function FUNCTION_TYPE_SEND() view returns(uint16)
func (_Efes *EfesCallerSession) FUNCTIONTYPESEND() (uint16, error) {
	return _Efes.Contract.FUNCTIONTYPESEND(&_Efes.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xbeab7131.
//
// Solidity: function _chainID() view returns(uint32)
func (_Efes *EfesCaller) ChainID(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "_chainID")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ChainID is a free data retrieval call binding the contract method 0xbeab7131.
//
// Solidity: function _chainID() view returns(uint32)
func (_Efes *EfesSession) ChainID() (uint32, error) {
	return _Efes.Contract.ChainID(&_Efes.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xbeab7131.
//
// Solidity: function _chainID() view returns(uint32)
func (_Efes *EfesCallerSession) ChainID() (uint32, error) {
	return _Efes.Contract.ChainID(&_Efes.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Efes *EfesCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Efes *EfesSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Efes.Contract.BalanceOf(&_Efes.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Efes *EfesCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Efes.Contract.BalanceOf(&_Efes.CallOpts, owner)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_Efes *EfesCaller) Beneficiary(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "beneficiary")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_Efes *EfesSession) Beneficiary() (common.Address, error) {
	return _Efes.Contract.Beneficiary(&_Efes.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_Efes *EfesCallerSession) Beneficiary() (common.Address, error) {
	return _Efes.Contract.Beneficiary(&_Efes.CallOpts)
}

// Cosigner is a free data retrieval call binding the contract method 0x681232ad.
//
// Solidity: function cosigner() view returns(address)
func (_Efes *EfesCaller) Cosigner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "cosigner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Cosigner is a free data retrieval call binding the contract method 0x681232ad.
//
// Solidity: function cosigner() view returns(address)
func (_Efes *EfesSession) Cosigner() (common.Address, error) {
	return _Efes.Contract.Cosigner(&_Efes.CallOpts)
}

// Cosigner is a free data retrieval call binding the contract method 0x681232ad.
//
// Solidity: function cosigner() view returns(address)
func (_Efes *EfesCallerSession) Cosigner() (common.Address, error) {
	return _Efes.Contract.Cosigner(&_Efes.CallOpts)
}

// CurrentBaseURI is a free data retrieval call binding the contract method 0x97d4ccd2.
//
// Solidity: function currentBaseURI() view returns(string)
func (_Efes *EfesCaller) CurrentBaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "currentBaseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CurrentBaseURI is a free data retrieval call binding the contract method 0x97d4ccd2.
//
// Solidity: function currentBaseURI() view returns(string)
func (_Efes *EfesSession) CurrentBaseURI() (string, error) {
	return _Efes.Contract.CurrentBaseURI(&_Efes.CallOpts)
}

// CurrentBaseURI is a free data retrieval call binding the contract method 0x97d4ccd2.
//
// Solidity: function currentBaseURI() view returns(string)
func (_Efes *EfesCallerSession) CurrentBaseURI() (string, error) {
	return _Efes.Contract.CurrentBaseURI(&_Efes.CallOpts)
}

// Decode is a free data retrieval call binding the contract method 0xe5c5e9a3.
//
// Solidity: function decode(bytes data) pure returns(uint32 qty, string requestId, uint64 timestamp, bytes sig)
func (_Efes *EfesCaller) Decode(opts *bind.CallOpts, data []byte) (struct {
	Qty       uint32
	RequestId string
	Timestamp uint64
	Sig       []byte
}, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "decode", data)

	outstruct := new(struct {
		Qty       uint32
		RequestId string
		Timestamp uint64
		Sig       []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Qty = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.RequestId = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.Sig = *abi.ConvertType(out[3], new([]byte)).(*[]byte)

	return *outstruct, err

}

// Decode is a free data retrieval call binding the contract method 0xe5c5e9a3.
//
// Solidity: function decode(bytes data) pure returns(uint32 qty, string requestId, uint64 timestamp, bytes sig)
func (_Efes *EfesSession) Decode(data []byte) (struct {
	Qty       uint32
	RequestId string
	Timestamp uint64
	Sig       []byte
}, error) {
	return _Efes.Contract.Decode(&_Efes.CallOpts, data)
}

// Decode is a free data retrieval call binding the contract method 0xe5c5e9a3.
//
// Solidity: function decode(bytes data) pure returns(uint32 qty, string requestId, uint64 timestamp, bytes sig)
func (_Efes *EfesCallerSession) Decode(data []byte) (struct {
	Qty       uint32
	RequestId string
	Timestamp uint64
	Sig       []byte
}, error) {
	return _Efes.Contract.Decode(&_Efes.CallOpts, data)
}

// DstChainIdToBatchLimit is a free data retrieval call binding the contract method 0x4ac3f4ff.
//
// Solidity: function dstChainIdToBatchLimit(uint16 ) view returns(uint256)
func (_Efes *EfesCaller) DstChainIdToBatchLimit(opts *bind.CallOpts, arg0 uint16) (*big.Int, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "dstChainIdToBatchLimit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DstChainIdToBatchLimit is a free data retrieval call binding the contract method 0x4ac3f4ff.
//
// Solidity: function dstChainIdToBatchLimit(uint16 ) view returns(uint256)
func (_Efes *EfesSession) DstChainIdToBatchLimit(arg0 uint16) (*big.Int, error) {
	return _Efes.Contract.DstChainIdToBatchLimit(&_Efes.CallOpts, arg0)
}

// DstChainIdToBatchLimit is a free data retrieval call binding the contract method 0x4ac3f4ff.
//
// Solidity: function dstChainIdToBatchLimit(uint16 ) view returns(uint256)
func (_Efes *EfesCallerSession) DstChainIdToBatchLimit(arg0 uint16) (*big.Int, error) {
	return _Efes.Contract.DstChainIdToBatchLimit(&_Efes.CallOpts, arg0)
}

// DstChainIdToTransferGas is a free data retrieval call binding the contract method 0xfa25f9b6.
//
// Solidity: function dstChainIdToTransferGas(uint16 ) view returns(uint256)
func (_Efes *EfesCaller) DstChainIdToTransferGas(opts *bind.CallOpts, arg0 uint16) (*big.Int, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "dstChainIdToTransferGas", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DstChainIdToTransferGas is a free data retrieval call binding the contract method 0xfa25f9b6.
//
// Solidity: function dstChainIdToTransferGas(uint16 ) view returns(uint256)
func (_Efes *EfesSession) DstChainIdToTransferGas(arg0 uint16) (*big.Int, error) {
	return _Efes.Contract.DstChainIdToTransferGas(&_Efes.CallOpts, arg0)
}

// DstChainIdToTransferGas is a free data retrieval call binding the contract method 0xfa25f9b6.
//
// Solidity: function dstChainIdToTransferGas(uint16 ) view returns(uint256)
func (_Efes *EfesCallerSession) DstChainIdToTransferGas(arg0 uint16) (*big.Int, error) {
	return _Efes.Contract.DstChainIdToTransferGas(&_Efes.CallOpts, arg0)
}

// EstimateSendBatchFee is a free data retrieval call binding the contract method 0xf2353641.
//
// Solidity: function estimateSendBatchFee(uint16 _dstChainId, bytes _toAddress, uint256[] _tokenIds, bool _useZro, bytes _adapterParams) view returns(uint256 nativeFee, uint256 zroFee)
func (_Efes *EfesCaller) EstimateSendBatchFee(opts *bind.CallOpts, _dstChainId uint16, _toAddress []byte, _tokenIds []*big.Int, _useZro bool, _adapterParams []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "estimateSendBatchFee", _dstChainId, _toAddress, _tokenIds, _useZro, _adapterParams)

	outstruct := new(struct {
		NativeFee *big.Int
		ZroFee    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NativeFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ZroFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// EstimateSendBatchFee is a free data retrieval call binding the contract method 0xf2353641.
//
// Solidity: function estimateSendBatchFee(uint16 _dstChainId, bytes _toAddress, uint256[] _tokenIds, bool _useZro, bytes _adapterParams) view returns(uint256 nativeFee, uint256 zroFee)
func (_Efes *EfesSession) EstimateSendBatchFee(_dstChainId uint16, _toAddress []byte, _tokenIds []*big.Int, _useZro bool, _adapterParams []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	return _Efes.Contract.EstimateSendBatchFee(&_Efes.CallOpts, _dstChainId, _toAddress, _tokenIds, _useZro, _adapterParams)
}

// EstimateSendBatchFee is a free data retrieval call binding the contract method 0xf2353641.
//
// Solidity: function estimateSendBatchFee(uint16 _dstChainId, bytes _toAddress, uint256[] _tokenIds, bool _useZro, bytes _adapterParams) view returns(uint256 nativeFee, uint256 zroFee)
func (_Efes *EfesCallerSession) EstimateSendBatchFee(_dstChainId uint16, _toAddress []byte, _tokenIds []*big.Int, _useZro bool, _adapterParams []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	return _Efes.Contract.EstimateSendBatchFee(&_Efes.CallOpts, _dstChainId, _toAddress, _tokenIds, _useZro, _adapterParams)
}

// EstimateSendFee is a free data retrieval call binding the contract method 0x2a205e3d.
//
// Solidity: function estimateSendFee(uint16 _dstChainId, bytes _toAddress, uint256 _tokenId, bool _useZro, bytes _adapterParams) view returns(uint256 nativeFee, uint256 zroFee)
func (_Efes *EfesCaller) EstimateSendFee(opts *bind.CallOpts, _dstChainId uint16, _toAddress []byte, _tokenId *big.Int, _useZro bool, _adapterParams []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "estimateSendFee", _dstChainId, _toAddress, _tokenId, _useZro, _adapterParams)

	outstruct := new(struct {
		NativeFee *big.Int
		ZroFee    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NativeFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ZroFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// EstimateSendFee is a free data retrieval call binding the contract method 0x2a205e3d.
//
// Solidity: function estimateSendFee(uint16 _dstChainId, bytes _toAddress, uint256 _tokenId, bool _useZro, bytes _adapterParams) view returns(uint256 nativeFee, uint256 zroFee)
func (_Efes *EfesSession) EstimateSendFee(_dstChainId uint16, _toAddress []byte, _tokenId *big.Int, _useZro bool, _adapterParams []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	return _Efes.Contract.EstimateSendFee(&_Efes.CallOpts, _dstChainId, _toAddress, _tokenId, _useZro, _adapterParams)
}

// EstimateSendFee is a free data retrieval call binding the contract method 0x2a205e3d.
//
// Solidity: function estimateSendFee(uint16 _dstChainId, bytes _toAddress, uint256 _tokenId, bool _useZro, bytes _adapterParams) view returns(uint256 nativeFee, uint256 zroFee)
func (_Efes *EfesCallerSession) EstimateSendFee(_dstChainId uint16, _toAddress []byte, _tokenId *big.Int, _useZro bool, _adapterParams []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	return _Efes.Contract.EstimateSendFee(&_Efes.CallOpts, _dstChainId, _toAddress, _tokenId, _useZro, _adapterParams)
}

// ExpireTime is a free data retrieval call binding the contract method 0x310a1ee3.
//
// Solidity: function expireTime() view returns(uint256)
func (_Efes *EfesCaller) ExpireTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "expireTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpireTime is a free data retrieval call binding the contract method 0x310a1ee3.
//
// Solidity: function expireTime() view returns(uint256)
func (_Efes *EfesSession) ExpireTime() (*big.Int, error) {
	return _Efes.Contract.ExpireTime(&_Efes.CallOpts)
}

// ExpireTime is a free data retrieval call binding the contract method 0x310a1ee3.
//
// Solidity: function expireTime() view returns(uint256)
func (_Efes *EfesCallerSession) ExpireTime() (*big.Int, error) {
	return _Efes.Contract.ExpireTime(&_Efes.CallOpts)
}

// FailedMessages is a free data retrieval call binding the contract method 0x5b8c41e6.
//
// Solidity: function failedMessages(uint16 , bytes , uint64 ) view returns(bytes32)
func (_Efes *EfesCaller) FailedMessages(opts *bind.CallOpts, arg0 uint16, arg1 []byte, arg2 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "failedMessages", arg0, arg1, arg2)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FailedMessages is a free data retrieval call binding the contract method 0x5b8c41e6.
//
// Solidity: function failedMessages(uint16 , bytes , uint64 ) view returns(bytes32)
func (_Efes *EfesSession) FailedMessages(arg0 uint16, arg1 []byte, arg2 uint64) ([32]byte, error) {
	return _Efes.Contract.FailedMessages(&_Efes.CallOpts, arg0, arg1, arg2)
}

// FailedMessages is a free data retrieval call binding the contract method 0x5b8c41e6.
//
// Solidity: function failedMessages(uint16 , bytes , uint64 ) view returns(bytes32)
func (_Efes *EfesCallerSession) FailedMessages(arg0 uint16, arg1 []byte, arg2 uint64) ([32]byte, error) {
	return _Efes.Contract.FailedMessages(&_Efes.CallOpts, arg0, arg1, arg2)
}

// GetActiveStageFromTimestamp is a free data retrieval call binding the contract method 0x67808a34.
//
// Solidity: function getActiveStageFromTimestamp(uint64 timestamp) view returns(uint256)
func (_Efes *EfesCaller) GetActiveStageFromTimestamp(opts *bind.CallOpts, timestamp uint64) (*big.Int, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "getActiveStageFromTimestamp", timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveStageFromTimestamp is a free data retrieval call binding the contract method 0x67808a34.
//
// Solidity: function getActiveStageFromTimestamp(uint64 timestamp) view returns(uint256)
func (_Efes *EfesSession) GetActiveStageFromTimestamp(timestamp uint64) (*big.Int, error) {
	return _Efes.Contract.GetActiveStageFromTimestamp(&_Efes.CallOpts, timestamp)
}

// GetActiveStageFromTimestamp is a free data retrieval call binding the contract method 0x67808a34.
//
// Solidity: function getActiveStageFromTimestamp(uint64 timestamp) view returns(uint256)
func (_Efes *EfesCallerSession) GetActiveStageFromTimestamp(timestamp uint64) (*big.Int, error) {
	return _Efes.Contract.GetActiveStageFromTimestamp(&_Efes.CallOpts, timestamp)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Efes *EfesCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Efes *EfesSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Efes.Contract.GetApproved(&_Efes.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Efes *EfesCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Efes.Contract.GetApproved(&_Efes.CallOpts, tokenId)
}

// GetConfig is a free data retrieval call binding the contract method 0xf5ecbdbc.
//
// Solidity: function getConfig(uint16 _version, uint16 _chainId, address , uint256 _configType) view returns(bytes)
func (_Efes *EfesCaller) GetConfig(opts *bind.CallOpts, _version uint16, _chainId uint16, arg2 common.Address, _configType *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "getConfig", _version, _chainId, arg2, _configType)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetConfig is a free data retrieval call binding the contract method 0xf5ecbdbc.
//
// Solidity: function getConfig(uint16 _version, uint16 _chainId, address , uint256 _configType) view returns(bytes)
func (_Efes *EfesSession) GetConfig(_version uint16, _chainId uint16, arg2 common.Address, _configType *big.Int) ([]byte, error) {
	return _Efes.Contract.GetConfig(&_Efes.CallOpts, _version, _chainId, arg2, _configType)
}

// GetConfig is a free data retrieval call binding the contract method 0xf5ecbdbc.
//
// Solidity: function getConfig(uint16 _version, uint16 _chainId, address , uint256 _configType) view returns(bytes)
func (_Efes *EfesCallerSession) GetConfig(_version uint16, _chainId uint16, arg2 common.Address, _configType *big.Int) ([]byte, error) {
	return _Efes.Contract.GetConfig(&_Efes.CallOpts, _version, _chainId, arg2, _configType)
}

// GetNumberStages is a free data retrieval call binding the contract method 0x70da24ee.
//
// Solidity: function getNumberStages() view returns(uint256)
func (_Efes *EfesCaller) GetNumberStages(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "getNumberStages")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumberStages is a free data retrieval call binding the contract method 0x70da24ee.
//
// Solidity: function getNumberStages() view returns(uint256)
func (_Efes *EfesSession) GetNumberStages() (*big.Int, error) {
	return _Efes.Contract.GetNumberStages(&_Efes.CallOpts)
}

// GetNumberStages is a free data retrieval call binding the contract method 0x70da24ee.
//
// Solidity: function getNumberStages() view returns(uint256)
func (_Efes *EfesCallerSession) GetNumberStages() (*big.Int, error) {
	return _Efes.Contract.GetNumberStages(&_Efes.CallOpts)
}

// GetTrustedRemoteAddress is a free data retrieval call binding the contract method 0x9f38369a.
//
// Solidity: function getTrustedRemoteAddress(uint16 _remoteChainId) view returns(bytes)
func (_Efes *EfesCaller) GetTrustedRemoteAddress(opts *bind.CallOpts, _remoteChainId uint16) ([]byte, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "getTrustedRemoteAddress", _remoteChainId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetTrustedRemoteAddress is a free data retrieval call binding the contract method 0x9f38369a.
//
// Solidity: function getTrustedRemoteAddress(uint16 _remoteChainId) view returns(bytes)
func (_Efes *EfesSession) GetTrustedRemoteAddress(_remoteChainId uint16) ([]byte, error) {
	return _Efes.Contract.GetTrustedRemoteAddress(&_Efes.CallOpts, _remoteChainId)
}

// GetTrustedRemoteAddress is a free data retrieval call binding the contract method 0x9f38369a.
//
// Solidity: function getTrustedRemoteAddress(uint16 _remoteChainId) view returns(bytes)
func (_Efes *EfesCallerSession) GetTrustedRemoteAddress(_remoteChainId uint16) ([]byte, error) {
	return _Efes.Contract.GetTrustedRemoteAddress(&_Efes.CallOpts, _remoteChainId)
}

// Gettimestamp is a free data retrieval call binding the contract method 0x7c15dd60.
//
// Solidity: function gettimestamp() view returns(uint256)
func (_Efes *EfesCaller) Gettimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "gettimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Gettimestamp is a free data retrieval call binding the contract method 0x7c15dd60.
//
// Solidity: function gettimestamp() view returns(uint256)
func (_Efes *EfesSession) Gettimestamp() (*big.Int, error) {
	return _Efes.Contract.Gettimestamp(&_Efes.CallOpts)
}

// Gettimestamp is a free data retrieval call binding the contract method 0x7c15dd60.
//
// Solidity: function gettimestamp() view returns(uint256)
func (_Efes *EfesCallerSession) Gettimestamp() (*big.Int, error) {
	return _Efes.Contract.Gettimestamp(&_Efes.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Efes *EfesCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Efes *EfesSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Efes.Contract.IsApprovedForAll(&_Efes.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Efes *EfesCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Efes.Contract.IsApprovedForAll(&_Efes.CallOpts, owner, operator)
}

// IsTrustedRemote is a free data retrieval call binding the contract method 0x3d8b38f6.
//
// Solidity: function isTrustedRemote(uint16 _srcChainId, bytes _srcAddress) view returns(bool)
func (_Efes *EfesCaller) IsTrustedRemote(opts *bind.CallOpts, _srcChainId uint16, _srcAddress []byte) (bool, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "isTrustedRemote", _srcChainId, _srcAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedRemote is a free data retrieval call binding the contract method 0x3d8b38f6.
//
// Solidity: function isTrustedRemote(uint16 _srcChainId, bytes _srcAddress) view returns(bool)
func (_Efes *EfesSession) IsTrustedRemote(_srcChainId uint16, _srcAddress []byte) (bool, error) {
	return _Efes.Contract.IsTrustedRemote(&_Efes.CallOpts, _srcChainId, _srcAddress)
}

// IsTrustedRemote is a free data retrieval call binding the contract method 0x3d8b38f6.
//
// Solidity: function isTrustedRemote(uint16 _srcChainId, bytes _srcAddress) view returns(bool)
func (_Efes *EfesCallerSession) IsTrustedRemote(_srcChainId uint16, _srcAddress []byte) (bool, error) {
	return _Efes.Contract.IsTrustedRemote(&_Efes.CallOpts, _srcChainId, _srcAddress)
}

// LzEndpoint is a free data retrieval call binding the contract method 0xb353aaa7.
//
// Solidity: function lzEndpoint() view returns(address)
func (_Efes *EfesCaller) LzEndpoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "lzEndpoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LzEndpoint is a free data retrieval call binding the contract method 0xb353aaa7.
//
// Solidity: function lzEndpoint() view returns(address)
func (_Efes *EfesSession) LzEndpoint() (common.Address, error) {
	return _Efes.Contract.LzEndpoint(&_Efes.CallOpts)
}

// LzEndpoint is a free data retrieval call binding the contract method 0xb353aaa7.
//
// Solidity: function lzEndpoint() view returns(address)
func (_Efes *EfesCallerSession) LzEndpoint() (common.Address, error) {
	return _Efes.Contract.LzEndpoint(&_Efes.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Efes *EfesCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Efes *EfesSession) MaxSupply() (*big.Int, error) {
	return _Efes.Contract.MaxSupply(&_Efes.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Efes *EfesCallerSession) MaxSupply() (*big.Int, error) {
	return _Efes.Contract.MaxSupply(&_Efes.CallOpts)
}

// MinDstGasLookup is a free data retrieval call binding the contract method 0x8cfd8f5c.
//
// Solidity: function minDstGasLookup(uint16 , uint16 ) view returns(uint256)
func (_Efes *EfesCaller) MinDstGasLookup(opts *bind.CallOpts, arg0 uint16, arg1 uint16) (*big.Int, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "minDstGasLookup", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDstGasLookup is a free data retrieval call binding the contract method 0x8cfd8f5c.
//
// Solidity: function minDstGasLookup(uint16 , uint16 ) view returns(uint256)
func (_Efes *EfesSession) MinDstGasLookup(arg0 uint16, arg1 uint16) (*big.Int, error) {
	return _Efes.Contract.MinDstGasLookup(&_Efes.CallOpts, arg0, arg1)
}

// MinDstGasLookup is a free data retrieval call binding the contract method 0x8cfd8f5c.
//
// Solidity: function minDstGasLookup(uint16 , uint16 ) view returns(uint256)
func (_Efes *EfesCallerSession) MinDstGasLookup(arg0 uint16, arg1 uint16) (*big.Int, error) {
	return _Efes.Contract.MinDstGasLookup(&_Efes.CallOpts, arg0, arg1)
}

// MinGasToTransferAndStore is a free data retrieval call binding the contract method 0x48288190.
//
// Solidity: function minGasToTransferAndStore() view returns(uint256)
func (_Efes *EfesCaller) MinGasToTransferAndStore(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "minGasToTransferAndStore")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinGasToTransferAndStore is a free data retrieval call binding the contract method 0x48288190.
//
// Solidity: function minGasToTransferAndStore() view returns(uint256)
func (_Efes *EfesSession) MinGasToTransferAndStore() (*big.Int, error) {
	return _Efes.Contract.MinGasToTransferAndStore(&_Efes.CallOpts)
}

// MinGasToTransferAndStore is a free data retrieval call binding the contract method 0x48288190.
//
// Solidity: function minGasToTransferAndStore() view returns(uint256)
func (_Efes *EfesCallerSession) MinGasToTransferAndStore() (*big.Int, error) {
	return _Efes.Contract.MinGasToTransferAndStore(&_Efes.CallOpts)
}

// MintStages is a free data retrieval call binding the contract method 0xfd0f7456.
//
// Solidity: function mintStages(uint256 ) view returns(uint80 whiteSalePrice, uint80 publicSalePrice, uint24 whiteSaleHour, uint24 publicSaleHour, uint24 maxStageSupply, uint64 startTimeUnixSeconds, uint256 endWhiteTimeUnixSeconds, uint256 endTimeUnixSeconds)
func (_Efes *EfesCaller) MintStages(opts *bind.CallOpts, arg0 *big.Int) (struct {
	WhiteSalePrice          *big.Int
	PublicSalePrice         *big.Int
	WhiteSaleHour           *big.Int
	PublicSaleHour          *big.Int
	MaxStageSupply          *big.Int
	StartTimeUnixSeconds    uint64
	EndWhiteTimeUnixSeconds *big.Int
	EndTimeUnixSeconds      *big.Int
}, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "mintStages", arg0)

	outstruct := new(struct {
		WhiteSalePrice          *big.Int
		PublicSalePrice         *big.Int
		WhiteSaleHour           *big.Int
		PublicSaleHour          *big.Int
		MaxStageSupply          *big.Int
		StartTimeUnixSeconds    uint64
		EndWhiteTimeUnixSeconds *big.Int
		EndTimeUnixSeconds      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.WhiteSalePrice = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PublicSalePrice = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.WhiteSaleHour = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.PublicSaleHour = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.MaxStageSupply = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.StartTimeUnixSeconds = *abi.ConvertType(out[5], new(uint64)).(*uint64)
	outstruct.EndWhiteTimeUnixSeconds = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.EndTimeUnixSeconds = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// MintStages is a free data retrieval call binding the contract method 0xfd0f7456.
//
// Solidity: function mintStages(uint256 ) view returns(uint80 whiteSalePrice, uint80 publicSalePrice, uint24 whiteSaleHour, uint24 publicSaleHour, uint24 maxStageSupply, uint64 startTimeUnixSeconds, uint256 endWhiteTimeUnixSeconds, uint256 endTimeUnixSeconds)
func (_Efes *EfesSession) MintStages(arg0 *big.Int) (struct {
	WhiteSalePrice          *big.Int
	PublicSalePrice         *big.Int
	WhiteSaleHour           *big.Int
	PublicSaleHour          *big.Int
	MaxStageSupply          *big.Int
	StartTimeUnixSeconds    uint64
	EndWhiteTimeUnixSeconds *big.Int
	EndTimeUnixSeconds      *big.Int
}, error) {
	return _Efes.Contract.MintStages(&_Efes.CallOpts, arg0)
}

// MintStages is a free data retrieval call binding the contract method 0xfd0f7456.
//
// Solidity: function mintStages(uint256 ) view returns(uint80 whiteSalePrice, uint80 publicSalePrice, uint24 whiteSaleHour, uint24 publicSaleHour, uint24 maxStageSupply, uint64 startTimeUnixSeconds, uint256 endWhiteTimeUnixSeconds, uint256 endTimeUnixSeconds)
func (_Efes *EfesCallerSession) MintStages(arg0 *big.Int) (struct {
	WhiteSalePrice          *big.Int
	PublicSalePrice         *big.Int
	WhiteSaleHour           *big.Int
	PublicSaleHour          *big.Int
	MaxStageSupply          *big.Int
	StartTimeUnixSeconds    uint64
	EndWhiteTimeUnixSeconds *big.Int
	EndTimeUnixSeconds      *big.Int
}, error) {
	return _Efes.Contract.MintStages(&_Efes.CallOpts, arg0)
}

// Mintable is a free data retrieval call binding the contract method 0x4bf365df.
//
// Solidity: function mintable() view returns(bool)
func (_Efes *EfesCaller) Mintable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "mintable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Mintable is a free data retrieval call binding the contract method 0x4bf365df.
//
// Solidity: function mintable() view returns(bool)
func (_Efes *EfesSession) Mintable() (bool, error) {
	return _Efes.Contract.Mintable(&_Efes.CallOpts)
}

// Mintable is a free data retrieval call binding the contract method 0x4bf365df.
//
// Solidity: function mintable() view returns(bool)
func (_Efes *EfesCallerSession) Mintable() (bool, error) {
	return _Efes.Contract.Mintable(&_Efes.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Efes *EfesCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Efes *EfesSession) Name() (string, error) {
	return _Efes.Contract.Name(&_Efes.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Efes *EfesCallerSession) Name() (string, error) {
	return _Efes.Contract.Name(&_Efes.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Efes *EfesCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Efes *EfesSession) Owner() (common.Address, error) {
	return _Efes.Contract.Owner(&_Efes.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Efes *EfesCallerSession) Owner() (common.Address, error) {
	return _Efes.Contract.Owner(&_Efes.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Efes *EfesCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Efes *EfesSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Efes.Contract.OwnerOf(&_Efes.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Efes *EfesCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Efes.Contract.OwnerOf(&_Efes.CallOpts, tokenId)
}

// PayloadSizeLimitLookup is a free data retrieval call binding the contract method 0x3f1f4fa4.
//
// Solidity: function payloadSizeLimitLookup(uint16 ) view returns(uint256)
func (_Efes *EfesCaller) PayloadSizeLimitLookup(opts *bind.CallOpts, arg0 uint16) (*big.Int, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "payloadSizeLimitLookup", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PayloadSizeLimitLookup is a free data retrieval call binding the contract method 0x3f1f4fa4.
//
// Solidity: function payloadSizeLimitLookup(uint16 ) view returns(uint256)
func (_Efes *EfesSession) PayloadSizeLimitLookup(arg0 uint16) (*big.Int, error) {
	return _Efes.Contract.PayloadSizeLimitLookup(&_Efes.CallOpts, arg0)
}

// PayloadSizeLimitLookup is a free data retrieval call binding the contract method 0x3f1f4fa4.
//
// Solidity: function payloadSizeLimitLookup(uint16 ) view returns(uint256)
func (_Efes *EfesCallerSession) PayloadSizeLimitLookup(arg0 uint16) (*big.Int, error) {
	return _Efes.Contract.PayloadSizeLimitLookup(&_Efes.CallOpts, arg0)
}

// Precrime is a free data retrieval call binding the contract method 0x950c8a74.
//
// Solidity: function precrime() view returns(address)
func (_Efes *EfesCaller) Precrime(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "precrime")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Precrime is a free data retrieval call binding the contract method 0x950c8a74.
//
// Solidity: function precrime() view returns(address)
func (_Efes *EfesSession) Precrime() (common.Address, error) {
	return _Efes.Contract.Precrime(&_Efes.CallOpts)
}

// Precrime is a free data retrieval call binding the contract method 0x950c8a74.
//
// Solidity: function precrime() view returns(address)
func (_Efes *EfesCallerSession) Precrime() (common.Address, error) {
	return _Efes.Contract.Precrime(&_Efes.CallOpts)
}

// StageMintedCounts is a free data retrieval call binding the contract method 0x0f461d5e.
//
// Solidity: function stageMintedCounts(uint256 ) view returns(uint256)
func (_Efes *EfesCaller) StageMintedCounts(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "stageMintedCounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StageMintedCounts is a free data retrieval call binding the contract method 0x0f461d5e.
//
// Solidity: function stageMintedCounts(uint256 ) view returns(uint256)
func (_Efes *EfesSession) StageMintedCounts(arg0 *big.Int) (*big.Int, error) {
	return _Efes.Contract.StageMintedCounts(&_Efes.CallOpts, arg0)
}

// StageMintedCounts is a free data retrieval call binding the contract method 0x0f461d5e.
//
// Solidity: function stageMintedCounts(uint256 ) view returns(uint256)
func (_Efes *EfesCallerSession) StageMintedCounts(arg0 *big.Int) (*big.Int, error) {
	return _Efes.Contract.StageMintedCounts(&_Efes.CallOpts, arg0)
}

// StoredCredits is a free data retrieval call binding the contract method 0x22a3ecf9.
//
// Solidity: function storedCredits(bytes32 ) view returns(uint16 srcChainId, address toAddress, uint256 index, bool creditsRemain)
func (_Efes *EfesCaller) StoredCredits(opts *bind.CallOpts, arg0 [32]byte) (struct {
	SrcChainId    uint16
	ToAddress     common.Address
	Index         *big.Int
	CreditsRemain bool
}, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "storedCredits", arg0)

	outstruct := new(struct {
		SrcChainId    uint16
		ToAddress     common.Address
		Index         *big.Int
		CreditsRemain bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SrcChainId = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.ToAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Index = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CreditsRemain = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// StoredCredits is a free data retrieval call binding the contract method 0x22a3ecf9.
//
// Solidity: function storedCredits(bytes32 ) view returns(uint16 srcChainId, address toAddress, uint256 index, bool creditsRemain)
func (_Efes *EfesSession) StoredCredits(arg0 [32]byte) (struct {
	SrcChainId    uint16
	ToAddress     common.Address
	Index         *big.Int
	CreditsRemain bool
}, error) {
	return _Efes.Contract.StoredCredits(&_Efes.CallOpts, arg0)
}

// StoredCredits is a free data retrieval call binding the contract method 0x22a3ecf9.
//
// Solidity: function storedCredits(bytes32 ) view returns(uint16 srcChainId, address toAddress, uint256 index, bool creditsRemain)
func (_Efes *EfesCallerSession) StoredCredits(arg0 [32]byte) (struct {
	SrcChainId    uint16
	ToAddress     common.Address
	Index         *big.Int
	CreditsRemain bool
}, error) {
	return _Efes.Contract.StoredCredits(&_Efes.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Efes *EfesCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Efes *EfesSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Efes.Contract.SupportsInterface(&_Efes.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Efes *EfesCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Efes.Contract.SupportsInterface(&_Efes.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Efes *EfesCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Efes *EfesSession) Symbol() (string, error) {
	return _Efes.Contract.Symbol(&_Efes.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Efes *EfesCallerSession) Symbol() (string, error) {
	return _Efes.Contract.Symbol(&_Efes.CallOpts)
}

// TokenIdStakeDays is a free data retrieval call binding the contract method 0x4ca6d84d.
//
// Solidity: function tokenIdStakeDays(uint256 ) view returns(uint256)
func (_Efes *EfesCaller) TokenIdStakeDays(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "tokenIdStakeDays", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIdStakeDays is a free data retrieval call binding the contract method 0x4ca6d84d.
//
// Solidity: function tokenIdStakeDays(uint256 ) view returns(uint256)
func (_Efes *EfesSession) TokenIdStakeDays(arg0 *big.Int) (*big.Int, error) {
	return _Efes.Contract.TokenIdStakeDays(&_Efes.CallOpts, arg0)
}

// TokenIdStakeDays is a free data retrieval call binding the contract method 0x4ca6d84d.
//
// Solidity: function tokenIdStakeDays(uint256 ) view returns(uint256)
func (_Efes *EfesCallerSession) TokenIdStakeDays(arg0 *big.Int) (*big.Int, error) {
	return _Efes.Contract.TokenIdStakeDays(&_Efes.CallOpts, arg0)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Efes *EfesCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Efes *EfesSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Efes.Contract.TokenURI(&_Efes.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Efes *EfesCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Efes.Contract.TokenURI(&_Efes.CallOpts, tokenId)
}

// TokenURISuffix is a free data retrieval call binding the contract method 0xdbbc853b.
//
// Solidity: function tokenURISuffix() view returns(string)
func (_Efes *EfesCaller) TokenURISuffix(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "tokenURISuffix")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURISuffix is a free data retrieval call binding the contract method 0xdbbc853b.
//
// Solidity: function tokenURISuffix() view returns(string)
func (_Efes *EfesSession) TokenURISuffix() (string, error) {
	return _Efes.Contract.TokenURISuffix(&_Efes.CallOpts)
}

// TokenURISuffix is a free data retrieval call binding the contract method 0xdbbc853b.
//
// Solidity: function tokenURISuffix() view returns(string)
func (_Efes *EfesCallerSession) TokenURISuffix() (string, error) {
	return _Efes.Contract.TokenURISuffix(&_Efes.CallOpts)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_Efes *EfesCaller) TotalMinted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "totalMinted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_Efes *EfesSession) TotalMinted() (*big.Int, error) {
	return _Efes.Contract.TotalMinted(&_Efes.CallOpts)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_Efes *EfesCallerSession) TotalMinted() (*big.Int, error) {
	return _Efes.Contract.TotalMinted(&_Efes.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Efes *EfesCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Efes *EfesSession) TotalSupply() (*big.Int, error) {
	return _Efes.Contract.TotalSupply(&_Efes.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Efes *EfesCallerSession) TotalSupply() (*big.Int, error) {
	return _Efes.Contract.TotalSupply(&_Efes.CallOpts)
}

// TrustedRemoteLookup is a free data retrieval call binding the contract method 0x7533d788.
//
// Solidity: function trustedRemoteLookup(uint16 ) view returns(bytes)
func (_Efes *EfesCaller) TrustedRemoteLookup(opts *bind.CallOpts, arg0 uint16) ([]byte, error) {
	var out []interface{}
	err := _Efes.contract.Call(opts, &out, "trustedRemoteLookup", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// TrustedRemoteLookup is a free data retrieval call binding the contract method 0x7533d788.
//
// Solidity: function trustedRemoteLookup(uint16 ) view returns(bytes)
func (_Efes *EfesSession) TrustedRemoteLookup(arg0 uint16) ([]byte, error) {
	return _Efes.Contract.TrustedRemoteLookup(&_Efes.CallOpts, arg0)
}

// TrustedRemoteLookup is a free data retrieval call binding the contract method 0x7533d788.
//
// Solidity: function trustedRemoteLookup(uint16 ) view returns(bytes)
func (_Efes *EfesCallerSession) TrustedRemoteLookup(arg0 uint16) ([]byte, error) {
	return _Efes.Contract.TrustedRemoteLookup(&_Efes.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_Efes *EfesTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Efes.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_Efes *EfesSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Efes.Contract.Approve(&_Efes.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_Efes *EfesTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Efes.Contract.Approve(&_Efes.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Efes *EfesTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Efes.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Efes *EfesSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Efes.Contract.Burn(&_Efes.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Efes *EfesTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Efes.Contract.Burn(&_Efes.TransactOpts, tokenId)
}

// ClearCredits is a paid mutator transaction binding the contract method 0x8ffa1f2a.
//
// Solidity: function clearCredits(bytes _payload) returns()
func (_Efes *EfesTransactor) ClearCredits(opts *bind.TransactOpts, _payload []byte) (*types.Transaction, error) {
	return _Efes.contract.Transact(opts, "clearCredits", _payload)
}

// ClearCredits is a paid mutator transaction binding the contract method 0x8ffa1f2a.
//
// Solidity: function clearCredits(bytes _payload) returns()
func (_Efes *EfesSession) ClearCredits(_payload []byte) (*types.Transaction, error) {
	return _Efes.Contract.ClearCredits(&_Efes.TransactOpts, _payload)
}

// ClearCredits is a paid mutator transaction binding the contract method 0x8ffa1f2a.
//
// Solidity: function clearCredits(bytes _payload) returns()
func (_Efes *EfesTransactorSession) ClearCredits(_payload []byte) (*types.Transaction, error) {
	return _Efes.Contract.ClearCredits(&_Efes.TransactOpts, _payload)
}

// ForceResumeReceive is a paid mutator transaction binding the contract method 0x42d65a8d.
//
// Solidity: function forceResumeReceive(uint16 _srcChainId, bytes _srcAddress) returns()
func (_Efes *EfesTransactor) ForceResumeReceive(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte) (*types.Transaction, error) {
	return _Efes.contract.Transact(opts, "forceResumeReceive", _srcChainId, _srcAddress)
}

// ForceResumeReceive is a paid mutator transaction binding the contract method 0x42d65a8d.
//
// Solidity: function forceResumeReceive(uint16 _srcChainId, bytes _srcAddress) returns()
func (_Efes *EfesSession) ForceResumeReceive(_srcChainId uint16, _srcAddress []byte) (*types.Transaction, error) {
	return _Efes.Contract.ForceResumeReceive(&_Efes.TransactOpts, _srcChainId, _srcAddress)
}

// ForceResumeReceive is a paid mutator transaction binding the contract method 0x42d65a8d.
//
// Solidity: function forceResumeReceive(uint16 _srcChainId, bytes _srcAddress) returns()
func (_Efes *EfesTransactorSession) ForceResumeReceive(_srcChainId uint16, _srcAddress []byte) (*types.Transaction, error) {
	return _Efes.Contract.ForceResumeReceive(&_Efes.TransactOpts, _srcChainId, _srcAddress)
}

// LzReceive is a paid mutator transaction binding the contract method 0x001d3567.
//
// Solidity: function lzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Efes *EfesTransactor) LzReceive(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Efes.contract.Transact(opts, "lzReceive", _srcChainId, _srcAddress, _nonce, _payload)
}

// LzReceive is a paid mutator transaction binding the contract method 0x001d3567.
//
// Solidity: function lzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Efes *EfesSession) LzReceive(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Efes.Contract.LzReceive(&_Efes.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// LzReceive is a paid mutator transaction binding the contract method 0x001d3567.
//
// Solidity: function lzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Efes *EfesTransactorSession) LzReceive(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Efes.Contract.LzReceive(&_Efes.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// Mint is a paid mutator transaction binding the contract method 0x7ba0e2e7.
//
// Solidity: function mint(bytes data) payable returns()
func (_Efes *EfesTransactor) Mint(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _Efes.contract.Transact(opts, "mint", data)
}

// Mint is a paid mutator transaction binding the contract method 0x7ba0e2e7.
//
// Solidity: function mint(bytes data) payable returns()
func (_Efes *EfesSession) Mint(data []byte) (*types.Transaction, error) {
	return _Efes.Contract.Mint(&_Efes.TransactOpts, data)
}

// Mint is a paid mutator transaction binding the contract method 0x7ba0e2e7.
//
// Solidity: function mint(bytes data) payable returns()
func (_Efes *EfesTransactorSession) Mint(data []byte) (*types.Transaction, error) {
	return _Efes.Contract.Mint(&_Efes.TransactOpts, data)
}

// MintTo is a paid mutator transaction binding the contract method 0xdabc0e1d.
//
// Solidity: function mint_to(address to, bytes data) payable returns()
func (_Efes *EfesTransactor) MintTo(opts *bind.TransactOpts, to common.Address, data []byte) (*types.Transaction, error) {
	return _Efes.contract.Transact(opts, "mint_to", to, data)
}

// MintTo is a paid mutator transaction binding the contract method 0xdabc0e1d.
//
// Solidity: function mint_to(address to, bytes data) payable returns()
func (_Efes *EfesSession) MintTo(to common.Address, data []byte) (*types.Transaction, error) {
	return _Efes.Contract.MintTo(&_Efes.TransactOpts, to, data)
}

// MintTo is a paid mutator transaction binding the contract method 0xdabc0e1d.
//
// Solidity: function mint_to(address to, bytes data) payable returns()
func (_Efes *EfesTransactorSession) MintTo(to common.Address, data []byte) (*types.Transaction, error) {
	return _Efes.Contract.MintTo(&_Efes.TransactOpts, to, data)
}

// NonblockingLzReceive is a paid mutator transaction binding the contract method 0x66ad5c8a.
//
// Solidity: function nonblockingLzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Efes *EfesTransactor) NonblockingLzReceive(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Efes.contract.Transact(opts, "nonblockingLzReceive", _srcChainId, _srcAddress, _nonce, _payload)
}

// NonblockingLzReceive is a paid mutator transaction binding the contract method 0x66ad5c8a.
//
// Solidity: function nonblockingLzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Efes *EfesSession) NonblockingLzReceive(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Efes.Contract.NonblockingLzReceive(&_Efes.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// NonblockingLzReceive is a paid mutator transaction binding the contract method 0x66ad5c8a.
//
// Solidity: function nonblockingLzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Efes *EfesTransactorSession) NonblockingLzReceive(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Efes.Contract.NonblockingLzReceive(&_Efes.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Efes *EfesTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Efes.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Efes *EfesSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Efes.Contract.OnERC721Received(&_Efes.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Efes *EfesTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Efes.Contract.OnERC721Received(&_Efes.TransactOpts, arg0, arg1, arg2, arg3)
}

// Ownermint is a paid mutator transaction binding the contract method 0xf41aa655.
//
// Solidity: function ownermint(address to, uint32 qty, uint64 timestamp) payable returns()
func (_Efes *EfesTransactor) Ownermint(opts *bind.TransactOpts, to common.Address, qty uint32, timestamp uint64) (*types.Transaction, error) {
	return _Efes.contract.Transact(opts, "ownermint", to, qty, timestamp)
}

// Ownermint is a paid mutator transaction binding the contract method 0xf41aa655.
//
// Solidity: function ownermint(address to, uint32 qty, uint64 timestamp) payable returns()
func (_Efes *EfesSession) Ownermint(to common.Address, qty uint32, timestamp uint64) (*types.Transaction, error) {
	return _Efes.Contract.Ownermint(&_Efes.TransactOpts, to, qty, timestamp)
}

// Ownermint is a paid mutator transaction binding the contract method 0xf41aa655.
//
// Solidity: function ownermint(address to, uint32 qty, uint64 timestamp) payable returns()
func (_Efes *EfesTransactorSession) Ownermint(to common.Address, qty uint32, timestamp uint64) (*types.Transaction, error) {
	return _Efes.Contract.Ownermint(&_Efes.TransactOpts, to, qty, timestamp)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Efes *EfesTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Efes.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Efes *EfesSession) RenounceOwnership() (*types.Transaction, error) {
	return _Efes.Contract.RenounceOwnership(&_Efes.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Efes *EfesTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Efes.Contract.RenounceOwnership(&_Efes.TransactOpts)
}

// RetryMessage is a paid mutator transaction binding the contract method 0xd1deba1f.
//
// Solidity: function retryMessage(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) payable returns()
func (_Efes *EfesTransactor) RetryMessage(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Efes.contract.Transact(opts, "retryMessage", _srcChainId, _srcAddress, _nonce, _payload)
}

// RetryMessage is a paid mutator transaction binding the contract method 0xd1deba1f.
//
// Solidity: function retryMessage(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) payable returns()
func (_Efes *EfesSession) RetryMessage(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Efes.Contract.RetryMessage(&_Efes.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// RetryMessage is a paid mutator transaction binding the contract method 0xd1deba1f.
//
// Solidity: function retryMessage(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) payable returns()
func (_Efes *EfesTransactorSession) RetryMessage(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Efes.Contract.RetryMessage(&_Efes.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Efes *EfesTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Efes.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Efes *EfesSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Efes.Contract.SafeTransferFrom(&_Efes.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Efes *EfesTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Efes.Contract.SafeTransferFrom(&_Efes.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_Efes *EfesTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Efes.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_Efes *EfesSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Efes.Contract.SafeTransferFrom0(&_Efes.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_Efes *EfesTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Efes.Contract.SafeTransferFrom0(&_Efes.TransactOpts, from, to, tokenId, _data)
}

// SendBatchFrom is a paid mutator transaction binding the contract method 0xab3ffb93.
//
// Solidity: function sendBatchFrom(address _from, uint16 _dstChainId, bytes _toAddress, uint256[] _tokenIds, address _refundAddress, address _zroPaymentAddress, bytes _adapterParams) payable returns()
func (_Efes *EfesTransactor) SendBatchFrom(opts *bind.TransactOpts, _from common.Address, _dstChainId uint16, _toAddress []byte, _tokenIds []*big.Int, _refundAddress common.Address, _zroPaymentAddress common.Address, _adapterParams []byte) (*types.Transaction, error) {
	return _Efes.contract.Transact(opts, "sendBatchFrom", _from, _dstChainId, _toAddress, _tokenIds, _refundAddress, _zroPaymentAddress, _adapterParams)
}

// SendBatchFrom is a paid mutator transaction binding the contract method 0xab3ffb93.
//
// Solidity: function sendBatchFrom(address _from, uint16 _dstChainId, bytes _toAddress, uint256[] _tokenIds, address _refundAddress, address _zroPaymentAddress, bytes _adapterParams) payable returns()
func (_Efes *EfesSession) SendBatchFrom(_from common.Address, _dstChainId uint16, _toAddress []byte, _tokenIds []*big.Int, _refundAddress common.Address, _zroPaymentAddress common.Address, _adapterParams []byte) (*types.Transaction, error) {
	return _Efes.Contract.SendBatchFrom(&_Efes.TransactOpts, _from, _dstChainId, _toAddress, _tokenIds, _refundAddress, _zroPaymentAddress, _adapterParams)
}

// SendBatchFrom is a paid mutator transaction binding the contract method 0xab3ffb93.
//
// Solidity: function sendBatchFrom(address _from, uint16 _dstChainId, bytes _toAddress, uint256[] _tokenIds, address _refundAddress, address _zroPaymentAddress, bytes _adapterParams) payable returns()
func (_Efes *EfesTransactorSession) SendBatchFrom(_from common.Address, _dstChainId uint16, _toAddress []byte, _tokenIds []*big.Int, _refundAddress common.Address, _zroPaymentAddress common.Address, _adapterParams []byte) (*types.Transaction, error) {
	return _Efes.Contract.SendBatchFrom(&_Efes.TransactOpts, _from, _dstChainId, _toAddress, _tokenIds, _refundAddress, _zroPaymentAddress, _adapterParams)
}

// SendFrom is a paid mutator transaction binding the contract method 0x51905636.
//
// Solidity: function sendFrom(address _from, uint16 _dstChainId, bytes _toAddress, uint256 _tokenId, address _refundAddress, address _zroPaymentAddress, bytes _adapterParams) payable returns()
func (_Efes *EfesTransactor) SendFrom(opts *bind.TransactOpts, _from common.Address, _dstChainId uint16, _toAddress []byte, _tokenId *big.Int, _refundAddress common.Address, _zroPaymentAddress common.Address, _adapterParams []byte) (*types.Transaction, error) {
	return _Efes.contract.Transact(opts, "sendFrom", _from, _dstChainId, _toAddress, _tokenId, _refundAddress, _zroPaymentAddress, _adapterParams)
}

// SendFrom is a paid mutator transaction binding the contract method 0x51905636.
//
// Solidity: function sendFrom(address _from, uint16 _dstChainId, bytes _toAddress, uint256 _tokenId, address _refundAddress, address _zroPaymentAddress, bytes _adapterParams) payable returns()
func (_Efes *EfesSession) SendFrom(_from common.Address, _dstChainId uint16, _toAddress []byte, _tokenId *big.Int, _refundAddress common.Address, _zroPaymentAddress common.Address, _adapterParams []byte) (*types.Transaction, error) {
	return _Efes.Contract.SendFrom(&_Efes.TransactOpts, _from, _dstChainId, _toAddress, _tokenId, _refundAddress, _zroPaymentAddress, _adapterParams)
}

// SendFrom is a paid mutator transaction binding the contract method 0x51905636.
//
// Solidity: function sendFrom(address _from, uint16 _dstChainId, bytes _toAddress, uint256 _tokenId, address _refundAddress, address _zroPaymentAddress, bytes _adapterParams) payable returns()
func (_Efes *EfesTransactorSession) SendFrom(_from common.Address, _dstChainId uint16, _toAddress []byte, _tokenId *big.Int, _refundAddress common.Address, _zroPaymentAddress common.Address, _adapterParams []byte) (*types.Transaction, error) {
	return _Efes.Contract.SendFrom(&_Efes.TransactOpts, _from, _dstChainId, _toAddress, _tokenId, _refundAddress, _zroPaymentAddress, _adapterParams)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Efes *EfesTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Efes.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Efes *EfesSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Efes.Contract.SetApprovalForAll(&_Efes.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Efes *EfesTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Efes.Contract.SetApprovalForAll(&_Efes.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_Efes *EfesTransactor) SetBaseURI(opts *bind.TransactOpts, baseURI string) (*types.Transaction, error) {
	return _Efes.contract.Transact(opts, "setBaseURI", baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_Efes *EfesSession) SetBaseURI(baseURI string) (*types.Transaction, error) {
	return _Efes.Contract.SetBaseURI(&_Efes.TransactOpts, baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_Efes *EfesTransactorSession) SetBaseURI(baseURI string) (*types.Transaction, error) {
	return _Efes.Contract.SetBaseURI(&_Efes.TransactOpts, baseURI)
}

// SetConfig is a paid mutator transaction binding the contract method 0xcbed8b9c.
//
// Solidity: function setConfig(uint16 _version, uint16 _chainId, uint256 _configType, bytes _config) returns()
func (_Efes *EfesTransactor) SetConfig(opts *bind.TransactOpts, _version uint16, _chainId uint16, _configType *big.Int, _config []byte) (*types.Transaction, error) {
	return _Efes.contract.Transact(opts, "setConfig", _version, _chainId, _configType, _config)
}

// SetConfig is a paid mutator transaction binding the contract method 0xcbed8b9c.
//
// Solidity: function setConfig(uint16 _version, uint16 _chainId, uint256 _configType, bytes _config) returns()
func (_Efes *EfesSession) SetConfig(_version uint16, _chainId uint16, _configType *big.Int, _config []byte) (*types.Transaction, error) {
	return _Efes.Contract.SetConfig(&_Efes.TransactOpts, _version, _chainId, _configType, _config)
}

// SetConfig is a paid mutator transaction binding the contract method 0xcbed8b9c.
//
// Solidity: function setConfig(uint16 _version, uint16 _chainId, uint256 _configType, bytes _config) returns()
func (_Efes *EfesTransactorSession) SetConfig(_version uint16, _chainId uint16, _configType *big.Int, _config []byte) (*types.Transaction, error) {
	return _Efes.Contract.SetConfig(&_Efes.TransactOpts, _version, _chainId, _configType, _config)
}

// SetDstChainIdToBatchLimit is a paid mutator transaction binding the contract method 0x9ea5d6b1.
//
// Solidity: function setDstChainIdToBatchLimit(uint16 _dstChainId, uint256 _dstChainIdToBatchLimit) returns()
func (_Efes *EfesTransactor) SetDstChainIdToBatchLimit(opts *bind.TransactOpts, _dstChainId uint16, _dstChainIdToBatchLimit *big.Int) (*types.Transaction, error) {
	return _Efes.contract.Transact(opts, "setDstChainIdToBatchLimit", _dstChainId, _dstChainIdToBatchLimit)
}

// SetDstChainIdToBatchLimit is a paid mutator transaction binding the contract method 0x9ea5d6b1.
//
// Solidity: function setDstChainIdToBatchLimit(uint16 _dstChainId, uint256 _dstChainIdToBatchLimit) returns()
func (_Efes *EfesSession) SetDstChainIdToBatchLimit(_dstChainId uint16, _dstChainIdToBatchLimit *big.Int) (*types.Transaction, error) {
	return _Efes.Contract.SetDstChainIdToBatchLimit(&_Efes.TransactOpts, _dstChainId, _dstChainIdToBatchLimit)
}

// SetDstChainIdToBatchLimit is a paid mutator transaction binding the contract method 0x9ea5d6b1.
//
// Solidity: function setDstChainIdToBatchLimit(uint16 _dstChainId, uint256 _dstChainIdToBatchLimit) returns()
func (_Efes *EfesTransactorSession) SetDstChainIdToBatchLimit(_dstChainId uint16, _dstChainIdToBatchLimit *big.Int) (*types.Transaction, error) {
	return _Efes.Contract.SetDstChainIdToBatchLimit(&_Efes.TransactOpts, _dstChainId, _dstChainIdToBatchLimit)
}

// SetDstChainIdToTransferGas is a paid mutator transaction binding the contract method 0xd12473a5.
//
// Solidity: function setDstChainIdToTransferGas(uint16 _dstChainId, uint256 _dstChainIdToTransferGas) returns()
func (_Efes *EfesTransactor) SetDstChainIdToTransferGas(opts *bind.TransactOpts, _dstChainId uint16, _dstChainIdToTransferGas *big.Int) (*types.Transaction, error) {
	return _Efes.contract.Transact(opts, "setDstChainIdToTransferGas", _dstChainId, _dstChainIdToTransferGas)
}

// SetDstChainIdToTransferGas is a paid mutator transaction binding the contract method 0xd12473a5.
//
// Solidity: function setDstChainIdToTransferGas(uint16 _dstChainId, uint256 _dstChainIdToTransferGas) returns()
func (_Efes *EfesSession) SetDstChainIdToTransferGas(_dstChainId uint16, _dstChainIdToTransferGas *big.Int) (*types.Transaction, error) {
	return _Efes.Contract.SetDstChainIdToTransferGas(&_Efes.TransactOpts, _dstChainId, _dstChainIdToTransferGas)
}

// SetDstChainIdToTransferGas is a paid mutator transaction binding the contract method 0xd12473a5.
//
// Solidity: function setDstChainIdToTransferGas(uint16 _dstChainId, uint256 _dstChainIdToTransferGas) returns()
func (_Efes *EfesTransactorSession) SetDstChainIdToTransferGas(_dstChainId uint16, _dstChainIdToTransferGas *big.Int) (*types.Transaction, error) {
	return _Efes.Contract.SetDstChainIdToTransferGas(&_Efes.TransactOpts, _dstChainId, _dstChainIdToTransferGas)
}

// SetMinDstGas is a paid mutator transaction binding the contract method 0xdf2a5b3b.
//
// Solidity: function setMinDstGas(uint16 _dstChainId, uint16 _packetType, uint256 _minGas) returns()
func (_Efes *EfesTransactor) SetMinDstGas(opts *bind.TransactOpts, _dstChainId uint16, _packetType uint16, _minGas *big.Int) (*types.Transaction, error) {
	return _Efes.contract.Transact(opts, "setMinDstGas", _dstChainId, _packetType, _minGas)
}

// SetMinDstGas is a paid mutator transaction binding the contract method 0xdf2a5b3b.
//
// Solidity: function setMinDstGas(uint16 _dstChainId, uint16 _packetType, uint256 _minGas) returns()
func (_Efes *EfesSession) SetMinDstGas(_dstChainId uint16, _packetType uint16, _minGas *big.Int) (*types.Transaction, error) {
	return _Efes.Contract.SetMinDstGas(&_Efes.TransactOpts, _dstChainId, _packetType, _minGas)
}

// SetMinDstGas is a paid mutator transaction binding the contract method 0xdf2a5b3b.
//
// Solidity: function setMinDstGas(uint16 _dstChainId, uint16 _packetType, uint256 _minGas) returns()
func (_Efes *EfesTransactorSession) SetMinDstGas(_dstChainId uint16, _packetType uint16, _minGas *big.Int) (*types.Transaction, error) {
	return _Efes.Contract.SetMinDstGas(&_Efes.TransactOpts, _dstChainId, _packetType, _minGas)
}

// SetMinGasToTransferAndStore is a paid mutator transaction binding the contract method 0x0b4cad4c.
//
// Solidity: function setMinGasToTransferAndStore(uint256 _minGasToTransferAndStore) returns()
func (_Efes *EfesTransactor) SetMinGasToTransferAndStore(opts *bind.TransactOpts, _minGasToTransferAndStore *big.Int) (*types.Transaction, error) {
	return _Efes.contract.Transact(opts, "setMinGasToTransferAndStore", _minGasToTransferAndStore)
}

// SetMinGasToTransferAndStore is a paid mutator transaction binding the contract method 0x0b4cad4c.
//
// Solidity: function setMinGasToTransferAndStore(uint256 _minGasToTransferAndStore) returns()
func (_Efes *EfesSession) SetMinGasToTransferAndStore(_minGasToTransferAndStore *big.Int) (*types.Transaction, error) {
	return _Efes.Contract.SetMinGasToTransferAndStore(&_Efes.TransactOpts, _minGasToTransferAndStore)
}

// SetMinGasToTransferAndStore is a paid mutator transaction binding the contract method 0x0b4cad4c.
//
// Solidity: function setMinGasToTransferAndStore(uint256 _minGasToTransferAndStore) returns()
func (_Efes *EfesTransactorSession) SetMinGasToTransferAndStore(_minGasToTransferAndStore *big.Int) (*types.Transaction, error) {
	return _Efes.Contract.SetMinGasToTransferAndStore(&_Efes.TransactOpts, _minGasToTransferAndStore)
}

// SetMintable is a paid mutator transaction binding the contract method 0x285d70d4.
//
// Solidity: function setMintable(bool _mintable) returns()
func (_Efes *EfesTransactor) SetMintable(opts *bind.TransactOpts, _mintable bool) (*types.Transaction, error) {
	return _Efes.contract.Transact(opts, "setMintable", _mintable)
}

// SetMintable is a paid mutator transaction binding the contract method 0x285d70d4.
//
// Solidity: function setMintable(bool _mintable) returns()
func (_Efes *EfesSession) SetMintable(_mintable bool) (*types.Transaction, error) {
	return _Efes.Contract.SetMintable(&_Efes.TransactOpts, _mintable)
}

// SetMintable is a paid mutator transaction binding the contract method 0x285d70d4.
//
// Solidity: function setMintable(bool _mintable) returns()
func (_Efes *EfesTransactorSession) SetMintable(_mintable bool) (*types.Transaction, error) {
	return _Efes.Contract.SetMintable(&_Efes.TransactOpts, _mintable)
}

// SetPayloadSizeLimit is a paid mutator transaction binding the contract method 0x0df37483.
//
// Solidity: function setPayloadSizeLimit(uint16 _dstChainId, uint256 _size) returns()
func (_Efes *EfesTransactor) SetPayloadSizeLimit(opts *bind.TransactOpts, _dstChainId uint16, _size *big.Int) (*types.Transaction, error) {
	return _Efes.contract.Transact(opts, "setPayloadSizeLimit", _dstChainId, _size)
}

// SetPayloadSizeLimit is a paid mutator transaction binding the contract method 0x0df37483.
//
// Solidity: function setPayloadSizeLimit(uint16 _dstChainId, uint256 _size) returns()
func (_Efes *EfesSession) SetPayloadSizeLimit(_dstChainId uint16, _size *big.Int) (*types.Transaction, error) {
	return _Efes.Contract.SetPayloadSizeLimit(&_Efes.TransactOpts, _dstChainId, _size)
}

// SetPayloadSizeLimit is a paid mutator transaction binding the contract method 0x0df37483.
//
// Solidity: function setPayloadSizeLimit(uint16 _dstChainId, uint256 _size) returns()
func (_Efes *EfesTransactorSession) SetPayloadSizeLimit(_dstChainId uint16, _size *big.Int) (*types.Transaction, error) {
	return _Efes.Contract.SetPayloadSizeLimit(&_Efes.TransactOpts, _dstChainId, _size)
}

// SetPrecrime is a paid mutator transaction binding the contract method 0xbaf3292d.
//
// Solidity: function setPrecrime(address _precrime) returns()
func (_Efes *EfesTransactor) SetPrecrime(opts *bind.TransactOpts, _precrime common.Address) (*types.Transaction, error) {
	return _Efes.contract.Transact(opts, "setPrecrime", _precrime)
}

// SetPrecrime is a paid mutator transaction binding the contract method 0xbaf3292d.
//
// Solidity: function setPrecrime(address _precrime) returns()
func (_Efes *EfesSession) SetPrecrime(_precrime common.Address) (*types.Transaction, error) {
	return _Efes.Contract.SetPrecrime(&_Efes.TransactOpts, _precrime)
}

// SetPrecrime is a paid mutator transaction binding the contract method 0xbaf3292d.
//
// Solidity: function setPrecrime(address _precrime) returns()
func (_Efes *EfesTransactorSession) SetPrecrime(_precrime common.Address) (*types.Transaction, error) {
	return _Efes.Contract.SetPrecrime(&_Efes.TransactOpts, _precrime)
}

// SetReceiveVersion is a paid mutator transaction binding the contract method 0x10ddb137.
//
// Solidity: function setReceiveVersion(uint16 _version) returns()
func (_Efes *EfesTransactor) SetReceiveVersion(opts *bind.TransactOpts, _version uint16) (*types.Transaction, error) {
	return _Efes.contract.Transact(opts, "setReceiveVersion", _version)
}

// SetReceiveVersion is a paid mutator transaction binding the contract method 0x10ddb137.
//
// Solidity: function setReceiveVersion(uint16 _version) returns()
func (_Efes *EfesSession) SetReceiveVersion(_version uint16) (*types.Transaction, error) {
	return _Efes.Contract.SetReceiveVersion(&_Efes.TransactOpts, _version)
}

// SetReceiveVersion is a paid mutator transaction binding the contract method 0x10ddb137.
//
// Solidity: function setReceiveVersion(uint16 _version) returns()
func (_Efes *EfesTransactorSession) SetReceiveVersion(_version uint16) (*types.Transaction, error) {
	return _Efes.Contract.SetReceiveVersion(&_Efes.TransactOpts, _version)
}

// SetSendVersion is a paid mutator transaction binding the contract method 0x07e0db17.
//
// Solidity: function setSendVersion(uint16 _version) returns()
func (_Efes *EfesTransactor) SetSendVersion(opts *bind.TransactOpts, _version uint16) (*types.Transaction, error) {
	return _Efes.contract.Transact(opts, "setSendVersion", _version)
}

// SetSendVersion is a paid mutator transaction binding the contract method 0x07e0db17.
//
// Solidity: function setSendVersion(uint16 _version) returns()
func (_Efes *EfesSession) SetSendVersion(_version uint16) (*types.Transaction, error) {
	return _Efes.Contract.SetSendVersion(&_Efes.TransactOpts, _version)
}

// SetSendVersion is a paid mutator transaction binding the contract method 0x07e0db17.
//
// Solidity: function setSendVersion(uint16 _version) returns()
func (_Efes *EfesTransactorSession) SetSendVersion(_version uint16) (*types.Transaction, error) {
	return _Efes.Contract.SetSendVersion(&_Efes.TransactOpts, _version)
}

// SetStages is a paid mutator transaction binding the contract method 0xafa2d190.
//
// Solidity: function setStages((uint80,uint80,uint24,uint24,uint24,uint64)[] newStages) returns()
func (_Efes *EfesTransactor) SetStages(opts *bind.TransactOpts, newStages []ISRCNFTInputMintStageInfo) (*types.Transaction, error) {
	return _Efes.contract.Transact(opts, "setStages", newStages)
}

// SetStages is a paid mutator transaction binding the contract method 0xafa2d190.
//
// Solidity: function setStages((uint80,uint80,uint24,uint24,uint24,uint64)[] newStages) returns()
func (_Efes *EfesSession) SetStages(newStages []ISRCNFTInputMintStageInfo) (*types.Transaction, error) {
	return _Efes.Contract.SetStages(&_Efes.TransactOpts, newStages)
}

// SetStages is a paid mutator transaction binding the contract method 0xafa2d190.
//
// Solidity: function setStages((uint80,uint80,uint24,uint24,uint24,uint64)[] newStages) returns()
func (_Efes *EfesTransactorSession) SetStages(newStages []ISRCNFTInputMintStageInfo) (*types.Transaction, error) {
	return _Efes.Contract.SetStages(&_Efes.TransactOpts, newStages)
}

// SetTokenURISuffix is a paid mutator transaction binding the contract method 0xa9852bfb.
//
// Solidity: function setTokenURISuffix(string suffix) returns()
func (_Efes *EfesTransactor) SetTokenURISuffix(opts *bind.TransactOpts, suffix string) (*types.Transaction, error) {
	return _Efes.contract.Transact(opts, "setTokenURISuffix", suffix)
}

// SetTokenURISuffix is a paid mutator transaction binding the contract method 0xa9852bfb.
//
// Solidity: function setTokenURISuffix(string suffix) returns()
func (_Efes *EfesSession) SetTokenURISuffix(suffix string) (*types.Transaction, error) {
	return _Efes.Contract.SetTokenURISuffix(&_Efes.TransactOpts, suffix)
}

// SetTokenURISuffix is a paid mutator transaction binding the contract method 0xa9852bfb.
//
// Solidity: function setTokenURISuffix(string suffix) returns()
func (_Efes *EfesTransactorSession) SetTokenURISuffix(suffix string) (*types.Transaction, error) {
	return _Efes.Contract.SetTokenURISuffix(&_Efes.TransactOpts, suffix)
}

// SetTrustedRemote is a paid mutator transaction binding the contract method 0xeb8d72b7.
//
// Solidity: function setTrustedRemote(uint16 _remoteChainId, bytes _path) returns()
func (_Efes *EfesTransactor) SetTrustedRemote(opts *bind.TransactOpts, _remoteChainId uint16, _path []byte) (*types.Transaction, error) {
	return _Efes.contract.Transact(opts, "setTrustedRemote", _remoteChainId, _path)
}

// SetTrustedRemote is a paid mutator transaction binding the contract method 0xeb8d72b7.
//
// Solidity: function setTrustedRemote(uint16 _remoteChainId, bytes _path) returns()
func (_Efes *EfesSession) SetTrustedRemote(_remoteChainId uint16, _path []byte) (*types.Transaction, error) {
	return _Efes.Contract.SetTrustedRemote(&_Efes.TransactOpts, _remoteChainId, _path)
}

// SetTrustedRemote is a paid mutator transaction binding the contract method 0xeb8d72b7.
//
// Solidity: function setTrustedRemote(uint16 _remoteChainId, bytes _path) returns()
func (_Efes *EfesTransactorSession) SetTrustedRemote(_remoteChainId uint16, _path []byte) (*types.Transaction, error) {
	return _Efes.Contract.SetTrustedRemote(&_Efes.TransactOpts, _remoteChainId, _path)
}

// SetTrustedRemoteAddress is a paid mutator transaction binding the contract method 0xa6c3d165.
//
// Solidity: function setTrustedRemoteAddress(uint16 _remoteChainId, bytes _remoteAddress) returns()
func (_Efes *EfesTransactor) SetTrustedRemoteAddress(opts *bind.TransactOpts, _remoteChainId uint16, _remoteAddress []byte) (*types.Transaction, error) {
	return _Efes.contract.Transact(opts, "setTrustedRemoteAddress", _remoteChainId, _remoteAddress)
}

// SetTrustedRemoteAddress is a paid mutator transaction binding the contract method 0xa6c3d165.
//
// Solidity: function setTrustedRemoteAddress(uint16 _remoteChainId, bytes _remoteAddress) returns()
func (_Efes *EfesSession) SetTrustedRemoteAddress(_remoteChainId uint16, _remoteAddress []byte) (*types.Transaction, error) {
	return _Efes.Contract.SetTrustedRemoteAddress(&_Efes.TransactOpts, _remoteChainId, _remoteAddress)
}

// SetTrustedRemoteAddress is a paid mutator transaction binding the contract method 0xa6c3d165.
//
// Solidity: function setTrustedRemoteAddress(uint16 _remoteChainId, bytes _remoteAddress) returns()
func (_Efes *EfesTransactorSession) SetTrustedRemoteAddress(_remoteChainId uint16, _remoteAddress []byte) (*types.Transaction, error) {
	return _Efes.Contract.SetTrustedRemoteAddress(&_Efes.TransactOpts, _remoteChainId, _remoteAddress)
}

// Stake is a paid mutator transaction binding the contract method 0x1824c4de.
//
// Solidity: function stake(uint256 tokenId, uint24 _days) returns()
func (_Efes *EfesTransactor) Stake(opts *bind.TransactOpts, tokenId *big.Int, _days *big.Int) (*types.Transaction, error) {
	return _Efes.contract.Transact(opts, "stake", tokenId, _days)
}

// Stake is a paid mutator transaction binding the contract method 0x1824c4de.
//
// Solidity: function stake(uint256 tokenId, uint24 _days) returns()
func (_Efes *EfesSession) Stake(tokenId *big.Int, _days *big.Int) (*types.Transaction, error) {
	return _Efes.Contract.Stake(&_Efes.TransactOpts, tokenId, _days)
}

// Stake is a paid mutator transaction binding the contract method 0x1824c4de.
//
// Solidity: function stake(uint256 tokenId, uint24 _days) returns()
func (_Efes *EfesTransactorSession) Stake(tokenId *big.Int, _days *big.Int) (*types.Transaction, error) {
	return _Efes.Contract.Stake(&_Efes.TransactOpts, tokenId, _days)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 _tokenId) payable returns()
func (_Efes *EfesTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Efes.contract.Transact(opts, "transferFrom", from, to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 _tokenId) payable returns()
func (_Efes *EfesSession) TransferFrom(from common.Address, to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Efes.Contract.TransferFrom(&_Efes.TransactOpts, from, to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 _tokenId) payable returns()
func (_Efes *EfesTransactorSession) TransferFrom(from common.Address, to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Efes.Contract.TransferFrom(&_Efes.TransactOpts, from, to, _tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Efes *EfesTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Efes.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Efes *EfesSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Efes.Contract.TransferOwnership(&_Efes.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Efes *EfesTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Efes.Contract.TransferOwnership(&_Efes.TransactOpts, newOwner)
}

// UpdateCosigner is a paid mutator transaction binding the contract method 0xd325a794.
//
// Solidity: function updateCosigner(address _cosigner) returns()
func (_Efes *EfesTransactor) UpdateCosigner(opts *bind.TransactOpts, _cosigner common.Address) (*types.Transaction, error) {
	return _Efes.contract.Transact(opts, "updateCosigner", _cosigner)
}

// UpdateCosigner is a paid mutator transaction binding the contract method 0xd325a794.
//
// Solidity: function updateCosigner(address _cosigner) returns()
func (_Efes *EfesSession) UpdateCosigner(_cosigner common.Address) (*types.Transaction, error) {
	return _Efes.Contract.UpdateCosigner(&_Efes.TransactOpts, _cosigner)
}

// UpdateCosigner is a paid mutator transaction binding the contract method 0xd325a794.
//
// Solidity: function updateCosigner(address _cosigner) returns()
func (_Efes *EfesTransactorSession) UpdateCosigner(_cosigner common.Address) (*types.Transaction, error) {
	return _Efes.Contract.UpdateCosigner(&_Efes.TransactOpts, _cosigner)
}

// UpdateExpire is a paid mutator transaction binding the contract method 0xc5b2d404.
//
// Solidity: function updateExpire(uint32 _expire) returns()
func (_Efes *EfesTransactor) UpdateExpire(opts *bind.TransactOpts, _expire uint32) (*types.Transaction, error) {
	return _Efes.contract.Transact(opts, "updateExpire", _expire)
}

// UpdateExpire is a paid mutator transaction binding the contract method 0xc5b2d404.
//
// Solidity: function updateExpire(uint32 _expire) returns()
func (_Efes *EfesSession) UpdateExpire(_expire uint32) (*types.Transaction, error) {
	return _Efes.Contract.UpdateExpire(&_Efes.TransactOpts, _expire)
}

// UpdateExpire is a paid mutator transaction binding the contract method 0xc5b2d404.
//
// Solidity: function updateExpire(uint32 _expire) returns()
func (_Efes *EfesTransactorSession) UpdateExpire(_expire uint32) (*types.Transaction, error) {
	return _Efes.Contract.UpdateExpire(&_Efes.TransactOpts, _expire)
}

// UpdateStage is a paid mutator transaction binding the contract method 0x2c1cc7f9.
//
// Solidity: function updateStage(uint256 index, uint80 whiteSalePrice, uint80 publicSalePrice, uint24 maxStageSupply, uint64 startTimeUnixSeconds, uint24 whiteSaleHour, uint24 publicSaleHour) returns()
func (_Efes *EfesTransactor) UpdateStage(opts *bind.TransactOpts, index *big.Int, whiteSalePrice *big.Int, publicSalePrice *big.Int, maxStageSupply *big.Int, startTimeUnixSeconds uint64, whiteSaleHour *big.Int, publicSaleHour *big.Int) (*types.Transaction, error) {
	return _Efes.contract.Transact(opts, "updateStage", index, whiteSalePrice, publicSalePrice, maxStageSupply, startTimeUnixSeconds, whiteSaleHour, publicSaleHour)
}

// UpdateStage is a paid mutator transaction binding the contract method 0x2c1cc7f9.
//
// Solidity: function updateStage(uint256 index, uint80 whiteSalePrice, uint80 publicSalePrice, uint24 maxStageSupply, uint64 startTimeUnixSeconds, uint24 whiteSaleHour, uint24 publicSaleHour) returns()
func (_Efes *EfesSession) UpdateStage(index *big.Int, whiteSalePrice *big.Int, publicSalePrice *big.Int, maxStageSupply *big.Int, startTimeUnixSeconds uint64, whiteSaleHour *big.Int, publicSaleHour *big.Int) (*types.Transaction, error) {
	return _Efes.Contract.UpdateStage(&_Efes.TransactOpts, index, whiteSalePrice, publicSalePrice, maxStageSupply, startTimeUnixSeconds, whiteSaleHour, publicSaleHour)
}

// UpdateStage is a paid mutator transaction binding the contract method 0x2c1cc7f9.
//
// Solidity: function updateStage(uint256 index, uint80 whiteSalePrice, uint80 publicSalePrice, uint24 maxStageSupply, uint64 startTimeUnixSeconds, uint24 whiteSaleHour, uint24 publicSaleHour) returns()
func (_Efes *EfesTransactorSession) UpdateStage(index *big.Int, whiteSalePrice *big.Int, publicSalePrice *big.Int, maxStageSupply *big.Int, startTimeUnixSeconds uint64, whiteSaleHour *big.Int, publicSaleHour *big.Int) (*types.Transaction, error) {
	return _Efes.Contract.UpdateStage(&_Efes.TransactOpts, index, whiteSalePrice, publicSalePrice, maxStageSupply, startTimeUnixSeconds, whiteSaleHour, publicSaleHour)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Efes *EfesTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Efes.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Efes *EfesSession) Withdraw() (*types.Transaction, error) {
	return _Efes.Contract.Withdraw(&_Efes.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Efes *EfesTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Efes.Contract.Withdraw(&_Efes.TransactOpts)
}

// EfesApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Efes contract.
type EfesApprovalIterator struct {
	Event *EfesApproval // Event containing the contract specifics and raw log

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
func (it *EfesApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EfesApproval)
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
		it.Event = new(EfesApproval)
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
func (it *EfesApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EfesApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EfesApproval represents a Approval event raised by the Efes contract.
type EfesApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Efes *EfesFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*EfesApprovalIterator, error) {

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

	logs, sub, err := _Efes.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &EfesApprovalIterator{contract: _Efes.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Efes *EfesFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *EfesApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Efes.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EfesApproval)
				if err := _Efes.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Efes *EfesFilterer) ParseApproval(log types.Log) (*EfesApproval, error) {
	event := new(EfesApproval)
	if err := _Efes.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EfesApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Efes contract.
type EfesApprovalForAllIterator struct {
	Event *EfesApprovalForAll // Event containing the contract specifics and raw log

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
func (it *EfesApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EfesApprovalForAll)
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
		it.Event = new(EfesApprovalForAll)
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
func (it *EfesApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EfesApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EfesApprovalForAll represents a ApprovalForAll event raised by the Efes contract.
type EfesApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Efes *EfesFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*EfesApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Efes.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &EfesApprovalForAllIterator{contract: _Efes.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Efes *EfesFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *EfesApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Efes.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EfesApprovalForAll)
				if err := _Efes.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Efes *EfesFilterer) ParseApprovalForAll(log types.Log) (*EfesApprovalForAll, error) {
	event := new(EfesApprovalForAll)
	if err := _Efes.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EfesConsecutiveTransferIterator is returned from FilterConsecutiveTransfer and is used to iterate over the raw logs and unpacked data for ConsecutiveTransfer events raised by the Efes contract.
type EfesConsecutiveTransferIterator struct {
	Event *EfesConsecutiveTransfer // Event containing the contract specifics and raw log

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
func (it *EfesConsecutiveTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EfesConsecutiveTransfer)
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
		it.Event = new(EfesConsecutiveTransfer)
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
func (it *EfesConsecutiveTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EfesConsecutiveTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EfesConsecutiveTransfer represents a ConsecutiveTransfer event raised by the Efes contract.
type EfesConsecutiveTransfer struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	From        common.Address
	To          common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConsecutiveTransfer is a free log retrieval operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Efes *EfesFilterer) FilterConsecutiveTransfer(opts *bind.FilterOpts, fromTokenId []*big.Int, from []common.Address, to []common.Address) (*EfesConsecutiveTransferIterator, error) {

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

	logs, sub, err := _Efes.contract.FilterLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EfesConsecutiveTransferIterator{contract: _Efes.contract, event: "ConsecutiveTransfer", logs: logs, sub: sub}, nil
}

// WatchConsecutiveTransfer is a free log subscription operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Efes *EfesFilterer) WatchConsecutiveTransfer(opts *bind.WatchOpts, sink chan<- *EfesConsecutiveTransfer, fromTokenId []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Efes.contract.WatchLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EfesConsecutiveTransfer)
				if err := _Efes.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
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
func (_Efes *EfesFilterer) ParseConsecutiveTransfer(log types.Log) (*EfesConsecutiveTransfer, error) {
	event := new(EfesConsecutiveTransfer)
	if err := _Efes.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EfesCreditClearedIterator is returned from FilterCreditCleared and is used to iterate over the raw logs and unpacked data for CreditCleared events raised by the Efes contract.
type EfesCreditClearedIterator struct {
	Event *EfesCreditCleared // Event containing the contract specifics and raw log

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
func (it *EfesCreditClearedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EfesCreditCleared)
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
		it.Event = new(EfesCreditCleared)
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
func (it *EfesCreditClearedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EfesCreditClearedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EfesCreditCleared represents a CreditCleared event raised by the Efes contract.
type EfesCreditCleared struct {
	HashedPayload [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCreditCleared is a free log retrieval operation binding the contract event 0xd7be02b8dd0d27bd0517a9cb4d7469ce27df4313821ae5ec1ff69acc594ba233.
//
// Solidity: event CreditCleared(bytes32 _hashedPayload)
func (_Efes *EfesFilterer) FilterCreditCleared(opts *bind.FilterOpts) (*EfesCreditClearedIterator, error) {

	logs, sub, err := _Efes.contract.FilterLogs(opts, "CreditCleared")
	if err != nil {
		return nil, err
	}
	return &EfesCreditClearedIterator{contract: _Efes.contract, event: "CreditCleared", logs: logs, sub: sub}, nil
}

// WatchCreditCleared is a free log subscription operation binding the contract event 0xd7be02b8dd0d27bd0517a9cb4d7469ce27df4313821ae5ec1ff69acc594ba233.
//
// Solidity: event CreditCleared(bytes32 _hashedPayload)
func (_Efes *EfesFilterer) WatchCreditCleared(opts *bind.WatchOpts, sink chan<- *EfesCreditCleared) (event.Subscription, error) {

	logs, sub, err := _Efes.contract.WatchLogs(opts, "CreditCleared")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EfesCreditCleared)
				if err := _Efes.contract.UnpackLog(event, "CreditCleared", log); err != nil {
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

// ParseCreditCleared is a log parse operation binding the contract event 0xd7be02b8dd0d27bd0517a9cb4d7469ce27df4313821ae5ec1ff69acc594ba233.
//
// Solidity: event CreditCleared(bytes32 _hashedPayload)
func (_Efes *EfesFilterer) ParseCreditCleared(log types.Log) (*EfesCreditCleared, error) {
	event := new(EfesCreditCleared)
	if err := _Efes.contract.UnpackLog(event, "CreditCleared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EfesCreditStoredIterator is returned from FilterCreditStored and is used to iterate over the raw logs and unpacked data for CreditStored events raised by the Efes contract.
type EfesCreditStoredIterator struct {
	Event *EfesCreditStored // Event containing the contract specifics and raw log

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
func (it *EfesCreditStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EfesCreditStored)
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
		it.Event = new(EfesCreditStored)
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
func (it *EfesCreditStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EfesCreditStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EfesCreditStored represents a CreditStored event raised by the Efes contract.
type EfesCreditStored struct {
	HashedPayload [32]byte
	Payload       []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCreditStored is a free log retrieval operation binding the contract event 0x10e0b70d256bccc84b7027506978bd8b68984a870788b93b479def144c839ad7.
//
// Solidity: event CreditStored(bytes32 _hashedPayload, bytes _payload)
func (_Efes *EfesFilterer) FilterCreditStored(opts *bind.FilterOpts) (*EfesCreditStoredIterator, error) {

	logs, sub, err := _Efes.contract.FilterLogs(opts, "CreditStored")
	if err != nil {
		return nil, err
	}
	return &EfesCreditStoredIterator{contract: _Efes.contract, event: "CreditStored", logs: logs, sub: sub}, nil
}

// WatchCreditStored is a free log subscription operation binding the contract event 0x10e0b70d256bccc84b7027506978bd8b68984a870788b93b479def144c839ad7.
//
// Solidity: event CreditStored(bytes32 _hashedPayload, bytes _payload)
func (_Efes *EfesFilterer) WatchCreditStored(opts *bind.WatchOpts, sink chan<- *EfesCreditStored) (event.Subscription, error) {

	logs, sub, err := _Efes.contract.WatchLogs(opts, "CreditStored")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EfesCreditStored)
				if err := _Efes.contract.UnpackLog(event, "CreditStored", log); err != nil {
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

// ParseCreditStored is a log parse operation binding the contract event 0x10e0b70d256bccc84b7027506978bd8b68984a870788b93b479def144c839ad7.
//
// Solidity: event CreditStored(bytes32 _hashedPayload, bytes _payload)
func (_Efes *EfesFilterer) ParseCreditStored(log types.Log) (*EfesCreditStored, error) {
	event := new(EfesCreditStored)
	if err := _Efes.contract.UnpackLog(event, "CreditStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EfesMessageFailedIterator is returned from FilterMessageFailed and is used to iterate over the raw logs and unpacked data for MessageFailed events raised by the Efes contract.
type EfesMessageFailedIterator struct {
	Event *EfesMessageFailed // Event containing the contract specifics and raw log

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
func (it *EfesMessageFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EfesMessageFailed)
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
		it.Event = new(EfesMessageFailed)
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
func (it *EfesMessageFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EfesMessageFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EfesMessageFailed represents a MessageFailed event raised by the Efes contract.
type EfesMessageFailed struct {
	SrcChainId uint16
	SrcAddress []byte
	Nonce      uint64
	Payload    []byte
	Reason     []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMessageFailed is a free log retrieval operation binding the contract event 0xe183f33de2837795525b4792ca4cd60535bd77c53b7e7030060bfcf5734d6b0c.
//
// Solidity: event MessageFailed(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload, bytes _reason)
func (_Efes *EfesFilterer) FilterMessageFailed(opts *bind.FilterOpts) (*EfesMessageFailedIterator, error) {

	logs, sub, err := _Efes.contract.FilterLogs(opts, "MessageFailed")
	if err != nil {
		return nil, err
	}
	return &EfesMessageFailedIterator{contract: _Efes.contract, event: "MessageFailed", logs: logs, sub: sub}, nil
}

// WatchMessageFailed is a free log subscription operation binding the contract event 0xe183f33de2837795525b4792ca4cd60535bd77c53b7e7030060bfcf5734d6b0c.
//
// Solidity: event MessageFailed(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload, bytes _reason)
func (_Efes *EfesFilterer) WatchMessageFailed(opts *bind.WatchOpts, sink chan<- *EfesMessageFailed) (event.Subscription, error) {

	logs, sub, err := _Efes.contract.WatchLogs(opts, "MessageFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EfesMessageFailed)
				if err := _Efes.contract.UnpackLog(event, "MessageFailed", log); err != nil {
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

// ParseMessageFailed is a log parse operation binding the contract event 0xe183f33de2837795525b4792ca4cd60535bd77c53b7e7030060bfcf5734d6b0c.
//
// Solidity: event MessageFailed(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload, bytes _reason)
func (_Efes *EfesFilterer) ParseMessageFailed(log types.Log) (*EfesMessageFailed, error) {
	event := new(EfesMessageFailed)
	if err := _Efes.contract.UnpackLog(event, "MessageFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EfesOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Efes contract.
type EfesOwnershipTransferredIterator struct {
	Event *EfesOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EfesOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EfesOwnershipTransferred)
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
		it.Event = new(EfesOwnershipTransferred)
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
func (it *EfesOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EfesOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EfesOwnershipTransferred represents a OwnershipTransferred event raised by the Efes contract.
type EfesOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Efes *EfesFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EfesOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Efes.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EfesOwnershipTransferredIterator{contract: _Efes.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Efes *EfesFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EfesOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Efes.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EfesOwnershipTransferred)
				if err := _Efes.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Efes *EfesFilterer) ParseOwnershipTransferred(log types.Log) (*EfesOwnershipTransferred, error) {
	event := new(EfesOwnershipTransferred)
	if err := _Efes.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EfesReceiveFromChainIterator is returned from FilterReceiveFromChain and is used to iterate over the raw logs and unpacked data for ReceiveFromChain events raised by the Efes contract.
type EfesReceiveFromChainIterator struct {
	Event *EfesReceiveFromChain // Event containing the contract specifics and raw log

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
func (it *EfesReceiveFromChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EfesReceiveFromChain)
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
		it.Event = new(EfesReceiveFromChain)
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
func (it *EfesReceiveFromChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EfesReceiveFromChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EfesReceiveFromChain represents a ReceiveFromChain event raised by the Efes contract.
type EfesReceiveFromChain struct {
	SrcChainId uint16
	SrcAddress common.Hash
	ToAddress  common.Address
	TokenIds   []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterReceiveFromChain is a free log retrieval operation binding the contract event 0x5b821db8a46f8ecbe1941ba2f51cfeea9643268b56631f70d45e2a745d990265.
//
// Solidity: event ReceiveFromChain(uint16 indexed _srcChainId, bytes indexed _srcAddress, address indexed _toAddress, uint256[] _tokenIds)
func (_Efes *EfesFilterer) FilterReceiveFromChain(opts *bind.FilterOpts, _srcChainId []uint16, _srcAddress [][]byte, _toAddress []common.Address) (*EfesReceiveFromChainIterator, error) {

	var _srcChainIdRule []interface{}
	for _, _srcChainIdItem := range _srcChainId {
		_srcChainIdRule = append(_srcChainIdRule, _srcChainIdItem)
	}
	var _srcAddressRule []interface{}
	for _, _srcAddressItem := range _srcAddress {
		_srcAddressRule = append(_srcAddressRule, _srcAddressItem)
	}
	var _toAddressRule []interface{}
	for _, _toAddressItem := range _toAddress {
		_toAddressRule = append(_toAddressRule, _toAddressItem)
	}

	logs, sub, err := _Efes.contract.FilterLogs(opts, "ReceiveFromChain", _srcChainIdRule, _srcAddressRule, _toAddressRule)
	if err != nil {
		return nil, err
	}
	return &EfesReceiveFromChainIterator{contract: _Efes.contract, event: "ReceiveFromChain", logs: logs, sub: sub}, nil
}

// WatchReceiveFromChain is a free log subscription operation binding the contract event 0x5b821db8a46f8ecbe1941ba2f51cfeea9643268b56631f70d45e2a745d990265.
//
// Solidity: event ReceiveFromChain(uint16 indexed _srcChainId, bytes indexed _srcAddress, address indexed _toAddress, uint256[] _tokenIds)
func (_Efes *EfesFilterer) WatchReceiveFromChain(opts *bind.WatchOpts, sink chan<- *EfesReceiveFromChain, _srcChainId []uint16, _srcAddress [][]byte, _toAddress []common.Address) (event.Subscription, error) {

	var _srcChainIdRule []interface{}
	for _, _srcChainIdItem := range _srcChainId {
		_srcChainIdRule = append(_srcChainIdRule, _srcChainIdItem)
	}
	var _srcAddressRule []interface{}
	for _, _srcAddressItem := range _srcAddress {
		_srcAddressRule = append(_srcAddressRule, _srcAddressItem)
	}
	var _toAddressRule []interface{}
	for _, _toAddressItem := range _toAddress {
		_toAddressRule = append(_toAddressRule, _toAddressItem)
	}

	logs, sub, err := _Efes.contract.WatchLogs(opts, "ReceiveFromChain", _srcChainIdRule, _srcAddressRule, _toAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EfesReceiveFromChain)
				if err := _Efes.contract.UnpackLog(event, "ReceiveFromChain", log); err != nil {
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

// ParseReceiveFromChain is a log parse operation binding the contract event 0x5b821db8a46f8ecbe1941ba2f51cfeea9643268b56631f70d45e2a745d990265.
//
// Solidity: event ReceiveFromChain(uint16 indexed _srcChainId, bytes indexed _srcAddress, address indexed _toAddress, uint256[] _tokenIds)
func (_Efes *EfesFilterer) ParseReceiveFromChain(log types.Log) (*EfesReceiveFromChain, error) {
	event := new(EfesReceiveFromChain)
	if err := _Efes.contract.UnpackLog(event, "ReceiveFromChain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EfesRetryMessageSuccessIterator is returned from FilterRetryMessageSuccess and is used to iterate over the raw logs and unpacked data for RetryMessageSuccess events raised by the Efes contract.
type EfesRetryMessageSuccessIterator struct {
	Event *EfesRetryMessageSuccess // Event containing the contract specifics and raw log

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
func (it *EfesRetryMessageSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EfesRetryMessageSuccess)
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
		it.Event = new(EfesRetryMessageSuccess)
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
func (it *EfesRetryMessageSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EfesRetryMessageSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EfesRetryMessageSuccess represents a RetryMessageSuccess event raised by the Efes contract.
type EfesRetryMessageSuccess struct {
	SrcChainId  uint16
	SrcAddress  []byte
	Nonce       uint64
	PayloadHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRetryMessageSuccess is a free log retrieval operation binding the contract event 0xc264d91f3adc5588250e1551f547752ca0cfa8f6b530d243b9f9f4cab10ea8e5.
//
// Solidity: event RetryMessageSuccess(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes32 _payloadHash)
func (_Efes *EfesFilterer) FilterRetryMessageSuccess(opts *bind.FilterOpts) (*EfesRetryMessageSuccessIterator, error) {

	logs, sub, err := _Efes.contract.FilterLogs(opts, "RetryMessageSuccess")
	if err != nil {
		return nil, err
	}
	return &EfesRetryMessageSuccessIterator{contract: _Efes.contract, event: "RetryMessageSuccess", logs: logs, sub: sub}, nil
}

// WatchRetryMessageSuccess is a free log subscription operation binding the contract event 0xc264d91f3adc5588250e1551f547752ca0cfa8f6b530d243b9f9f4cab10ea8e5.
//
// Solidity: event RetryMessageSuccess(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes32 _payloadHash)
func (_Efes *EfesFilterer) WatchRetryMessageSuccess(opts *bind.WatchOpts, sink chan<- *EfesRetryMessageSuccess) (event.Subscription, error) {

	logs, sub, err := _Efes.contract.WatchLogs(opts, "RetryMessageSuccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EfesRetryMessageSuccess)
				if err := _Efes.contract.UnpackLog(event, "RetryMessageSuccess", log); err != nil {
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

// ParseRetryMessageSuccess is a log parse operation binding the contract event 0xc264d91f3adc5588250e1551f547752ca0cfa8f6b530d243b9f9f4cab10ea8e5.
//
// Solidity: event RetryMessageSuccess(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes32 _payloadHash)
func (_Efes *EfesFilterer) ParseRetryMessageSuccess(log types.Log) (*EfesRetryMessageSuccess, error) {
	event := new(EfesRetryMessageSuccess)
	if err := _Efes.contract.UnpackLog(event, "RetryMessageSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EfesSendToChainIterator is returned from FilterSendToChain and is used to iterate over the raw logs and unpacked data for SendToChain events raised by the Efes contract.
type EfesSendToChainIterator struct {
	Event *EfesSendToChain // Event containing the contract specifics and raw log

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
func (it *EfesSendToChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EfesSendToChain)
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
		it.Event = new(EfesSendToChain)
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
func (it *EfesSendToChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EfesSendToChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EfesSendToChain represents a SendToChain event raised by the Efes contract.
type EfesSendToChain struct {
	DstChainId uint16
	From       common.Address
	ToAddress  common.Hash
	TokenIds   []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSendToChain is a free log retrieval operation binding the contract event 0xe1b87c47fdeb4f9cbadbca9df3af7aba453bb6e501075d0440d88125b711522a.
//
// Solidity: event SendToChain(uint16 indexed _dstChainId, address indexed _from, bytes indexed _toAddress, uint256[] _tokenIds)
func (_Efes *EfesFilterer) FilterSendToChain(opts *bind.FilterOpts, _dstChainId []uint16, _from []common.Address, _toAddress [][]byte) (*EfesSendToChainIterator, error) {

	var _dstChainIdRule []interface{}
	for _, _dstChainIdItem := range _dstChainId {
		_dstChainIdRule = append(_dstChainIdRule, _dstChainIdItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toAddressRule []interface{}
	for _, _toAddressItem := range _toAddress {
		_toAddressRule = append(_toAddressRule, _toAddressItem)
	}

	logs, sub, err := _Efes.contract.FilterLogs(opts, "SendToChain", _dstChainIdRule, _fromRule, _toAddressRule)
	if err != nil {
		return nil, err
	}
	return &EfesSendToChainIterator{contract: _Efes.contract, event: "SendToChain", logs: logs, sub: sub}, nil
}

// WatchSendToChain is a free log subscription operation binding the contract event 0xe1b87c47fdeb4f9cbadbca9df3af7aba453bb6e501075d0440d88125b711522a.
//
// Solidity: event SendToChain(uint16 indexed _dstChainId, address indexed _from, bytes indexed _toAddress, uint256[] _tokenIds)
func (_Efes *EfesFilterer) WatchSendToChain(opts *bind.WatchOpts, sink chan<- *EfesSendToChain, _dstChainId []uint16, _from []common.Address, _toAddress [][]byte) (event.Subscription, error) {

	var _dstChainIdRule []interface{}
	for _, _dstChainIdItem := range _dstChainId {
		_dstChainIdRule = append(_dstChainIdRule, _dstChainIdItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toAddressRule []interface{}
	for _, _toAddressItem := range _toAddress {
		_toAddressRule = append(_toAddressRule, _toAddressItem)
	}

	logs, sub, err := _Efes.contract.WatchLogs(opts, "SendToChain", _dstChainIdRule, _fromRule, _toAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EfesSendToChain)
				if err := _Efes.contract.UnpackLog(event, "SendToChain", log); err != nil {
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

// ParseSendToChain is a log parse operation binding the contract event 0xe1b87c47fdeb4f9cbadbca9df3af7aba453bb6e501075d0440d88125b711522a.
//
// Solidity: event SendToChain(uint16 indexed _dstChainId, address indexed _from, bytes indexed _toAddress, uint256[] _tokenIds)
func (_Efes *EfesFilterer) ParseSendToChain(log types.Log) (*EfesSendToChain, error) {
	event := new(EfesSendToChain)
	if err := _Efes.contract.UnpackLog(event, "SendToChain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EfesSetDstChainIdToBatchLimitIterator is returned from FilterSetDstChainIdToBatchLimit and is used to iterate over the raw logs and unpacked data for SetDstChainIdToBatchLimit events raised by the Efes contract.
type EfesSetDstChainIdToBatchLimitIterator struct {
	Event *EfesSetDstChainIdToBatchLimit // Event containing the contract specifics and raw log

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
func (it *EfesSetDstChainIdToBatchLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EfesSetDstChainIdToBatchLimit)
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
		it.Event = new(EfesSetDstChainIdToBatchLimit)
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
func (it *EfesSetDstChainIdToBatchLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EfesSetDstChainIdToBatchLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EfesSetDstChainIdToBatchLimit represents a SetDstChainIdToBatchLimit event raised by the Efes contract.
type EfesSetDstChainIdToBatchLimit struct {
	DstChainId             uint16
	DstChainIdToBatchLimit *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetDstChainIdToBatchLimit is a free log retrieval operation binding the contract event 0x7315f7654d594ead24a30160ed9ba2d23247f543016b918343591e93d7afdb6d.
//
// Solidity: event SetDstChainIdToBatchLimit(uint16 _dstChainId, uint256 _dstChainIdToBatchLimit)
func (_Efes *EfesFilterer) FilterSetDstChainIdToBatchLimit(opts *bind.FilterOpts) (*EfesSetDstChainIdToBatchLimitIterator, error) {

	logs, sub, err := _Efes.contract.FilterLogs(opts, "SetDstChainIdToBatchLimit")
	if err != nil {
		return nil, err
	}
	return &EfesSetDstChainIdToBatchLimitIterator{contract: _Efes.contract, event: "SetDstChainIdToBatchLimit", logs: logs, sub: sub}, nil
}

// WatchSetDstChainIdToBatchLimit is a free log subscription operation binding the contract event 0x7315f7654d594ead24a30160ed9ba2d23247f543016b918343591e93d7afdb6d.
//
// Solidity: event SetDstChainIdToBatchLimit(uint16 _dstChainId, uint256 _dstChainIdToBatchLimit)
func (_Efes *EfesFilterer) WatchSetDstChainIdToBatchLimit(opts *bind.WatchOpts, sink chan<- *EfesSetDstChainIdToBatchLimit) (event.Subscription, error) {

	logs, sub, err := _Efes.contract.WatchLogs(opts, "SetDstChainIdToBatchLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EfesSetDstChainIdToBatchLimit)
				if err := _Efes.contract.UnpackLog(event, "SetDstChainIdToBatchLimit", log); err != nil {
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

// ParseSetDstChainIdToBatchLimit is a log parse operation binding the contract event 0x7315f7654d594ead24a30160ed9ba2d23247f543016b918343591e93d7afdb6d.
//
// Solidity: event SetDstChainIdToBatchLimit(uint16 _dstChainId, uint256 _dstChainIdToBatchLimit)
func (_Efes *EfesFilterer) ParseSetDstChainIdToBatchLimit(log types.Log) (*EfesSetDstChainIdToBatchLimit, error) {
	event := new(EfesSetDstChainIdToBatchLimit)
	if err := _Efes.contract.UnpackLog(event, "SetDstChainIdToBatchLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EfesSetDstChainIdToTransferGasIterator is returned from FilterSetDstChainIdToTransferGas and is used to iterate over the raw logs and unpacked data for SetDstChainIdToTransferGas events raised by the Efes contract.
type EfesSetDstChainIdToTransferGasIterator struct {
	Event *EfesSetDstChainIdToTransferGas // Event containing the contract specifics and raw log

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
func (it *EfesSetDstChainIdToTransferGasIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EfesSetDstChainIdToTransferGas)
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
		it.Event = new(EfesSetDstChainIdToTransferGas)
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
func (it *EfesSetDstChainIdToTransferGasIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EfesSetDstChainIdToTransferGasIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EfesSetDstChainIdToTransferGas represents a SetDstChainIdToTransferGas event raised by the Efes contract.
type EfesSetDstChainIdToTransferGas struct {
	DstChainId              uint16
	DstChainIdToTransferGas *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterSetDstChainIdToTransferGas is a free log retrieval operation binding the contract event 0xc46df2983228ac2d9754e94a0d565e6671665dc8ad38602bc8e544f0685a29fb.
//
// Solidity: event SetDstChainIdToTransferGas(uint16 _dstChainId, uint256 _dstChainIdToTransferGas)
func (_Efes *EfesFilterer) FilterSetDstChainIdToTransferGas(opts *bind.FilterOpts) (*EfesSetDstChainIdToTransferGasIterator, error) {

	logs, sub, err := _Efes.contract.FilterLogs(opts, "SetDstChainIdToTransferGas")
	if err != nil {
		return nil, err
	}
	return &EfesSetDstChainIdToTransferGasIterator{contract: _Efes.contract, event: "SetDstChainIdToTransferGas", logs: logs, sub: sub}, nil
}

// WatchSetDstChainIdToTransferGas is a free log subscription operation binding the contract event 0xc46df2983228ac2d9754e94a0d565e6671665dc8ad38602bc8e544f0685a29fb.
//
// Solidity: event SetDstChainIdToTransferGas(uint16 _dstChainId, uint256 _dstChainIdToTransferGas)
func (_Efes *EfesFilterer) WatchSetDstChainIdToTransferGas(opts *bind.WatchOpts, sink chan<- *EfesSetDstChainIdToTransferGas) (event.Subscription, error) {

	logs, sub, err := _Efes.contract.WatchLogs(opts, "SetDstChainIdToTransferGas")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EfesSetDstChainIdToTransferGas)
				if err := _Efes.contract.UnpackLog(event, "SetDstChainIdToTransferGas", log); err != nil {
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

// ParseSetDstChainIdToTransferGas is a log parse operation binding the contract event 0xc46df2983228ac2d9754e94a0d565e6671665dc8ad38602bc8e544f0685a29fb.
//
// Solidity: event SetDstChainIdToTransferGas(uint16 _dstChainId, uint256 _dstChainIdToTransferGas)
func (_Efes *EfesFilterer) ParseSetDstChainIdToTransferGas(log types.Log) (*EfesSetDstChainIdToTransferGas, error) {
	event := new(EfesSetDstChainIdToTransferGas)
	if err := _Efes.contract.UnpackLog(event, "SetDstChainIdToTransferGas", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EfesSetMinDstGasIterator is returned from FilterSetMinDstGas and is used to iterate over the raw logs and unpacked data for SetMinDstGas events raised by the Efes contract.
type EfesSetMinDstGasIterator struct {
	Event *EfesSetMinDstGas // Event containing the contract specifics and raw log

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
func (it *EfesSetMinDstGasIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EfesSetMinDstGas)
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
		it.Event = new(EfesSetMinDstGas)
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
func (it *EfesSetMinDstGasIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EfesSetMinDstGasIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EfesSetMinDstGas represents a SetMinDstGas event raised by the Efes contract.
type EfesSetMinDstGas struct {
	DstChainId uint16
	Type       uint16
	MinDstGas  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetMinDstGas is a free log retrieval operation binding the contract event 0x9d5c7c0b934da8fefa9c7760c98383778a12dfbfc0c3b3106518f43fb9508ac0.
//
// Solidity: event SetMinDstGas(uint16 _dstChainId, uint16 _type, uint256 _minDstGas)
func (_Efes *EfesFilterer) FilterSetMinDstGas(opts *bind.FilterOpts) (*EfesSetMinDstGasIterator, error) {

	logs, sub, err := _Efes.contract.FilterLogs(opts, "SetMinDstGas")
	if err != nil {
		return nil, err
	}
	return &EfesSetMinDstGasIterator{contract: _Efes.contract, event: "SetMinDstGas", logs: logs, sub: sub}, nil
}

// WatchSetMinDstGas is a free log subscription operation binding the contract event 0x9d5c7c0b934da8fefa9c7760c98383778a12dfbfc0c3b3106518f43fb9508ac0.
//
// Solidity: event SetMinDstGas(uint16 _dstChainId, uint16 _type, uint256 _minDstGas)
func (_Efes *EfesFilterer) WatchSetMinDstGas(opts *bind.WatchOpts, sink chan<- *EfesSetMinDstGas) (event.Subscription, error) {

	logs, sub, err := _Efes.contract.WatchLogs(opts, "SetMinDstGas")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EfesSetMinDstGas)
				if err := _Efes.contract.UnpackLog(event, "SetMinDstGas", log); err != nil {
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

// ParseSetMinDstGas is a log parse operation binding the contract event 0x9d5c7c0b934da8fefa9c7760c98383778a12dfbfc0c3b3106518f43fb9508ac0.
//
// Solidity: event SetMinDstGas(uint16 _dstChainId, uint16 _type, uint256 _minDstGas)
func (_Efes *EfesFilterer) ParseSetMinDstGas(log types.Log) (*EfesSetMinDstGas, error) {
	event := new(EfesSetMinDstGas)
	if err := _Efes.contract.UnpackLog(event, "SetMinDstGas", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EfesSetMinGasToTransferAndStoreIterator is returned from FilterSetMinGasToTransferAndStore and is used to iterate over the raw logs and unpacked data for SetMinGasToTransferAndStore events raised by the Efes contract.
type EfesSetMinGasToTransferAndStoreIterator struct {
	Event *EfesSetMinGasToTransferAndStore // Event containing the contract specifics and raw log

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
func (it *EfesSetMinGasToTransferAndStoreIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EfesSetMinGasToTransferAndStore)
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
		it.Event = new(EfesSetMinGasToTransferAndStore)
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
func (it *EfesSetMinGasToTransferAndStoreIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EfesSetMinGasToTransferAndStoreIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EfesSetMinGasToTransferAndStore represents a SetMinGasToTransferAndStore event raised by the Efes contract.
type EfesSetMinGasToTransferAndStore struct {
	MinGasToTransferAndStore *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSetMinGasToTransferAndStore is a free log retrieval operation binding the contract event 0xfebbc4f8bb9ec2313950c718d43123124b15778efda4c1f1d529de2995b4f34d.
//
// Solidity: event SetMinGasToTransferAndStore(uint256 _minGasToTransferAndStore)
func (_Efes *EfesFilterer) FilterSetMinGasToTransferAndStore(opts *bind.FilterOpts) (*EfesSetMinGasToTransferAndStoreIterator, error) {

	logs, sub, err := _Efes.contract.FilterLogs(opts, "SetMinGasToTransferAndStore")
	if err != nil {
		return nil, err
	}
	return &EfesSetMinGasToTransferAndStoreIterator{contract: _Efes.contract, event: "SetMinGasToTransferAndStore", logs: logs, sub: sub}, nil
}

// WatchSetMinGasToTransferAndStore is a free log subscription operation binding the contract event 0xfebbc4f8bb9ec2313950c718d43123124b15778efda4c1f1d529de2995b4f34d.
//
// Solidity: event SetMinGasToTransferAndStore(uint256 _minGasToTransferAndStore)
func (_Efes *EfesFilterer) WatchSetMinGasToTransferAndStore(opts *bind.WatchOpts, sink chan<- *EfesSetMinGasToTransferAndStore) (event.Subscription, error) {

	logs, sub, err := _Efes.contract.WatchLogs(opts, "SetMinGasToTransferAndStore")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EfesSetMinGasToTransferAndStore)
				if err := _Efes.contract.UnpackLog(event, "SetMinGasToTransferAndStore", log); err != nil {
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

// ParseSetMinGasToTransferAndStore is a log parse operation binding the contract event 0xfebbc4f8bb9ec2313950c718d43123124b15778efda4c1f1d529de2995b4f34d.
//
// Solidity: event SetMinGasToTransferAndStore(uint256 _minGasToTransferAndStore)
func (_Efes *EfesFilterer) ParseSetMinGasToTransferAndStore(log types.Log) (*EfesSetMinGasToTransferAndStore, error) {
	event := new(EfesSetMinGasToTransferAndStore)
	if err := _Efes.contract.UnpackLog(event, "SetMinGasToTransferAndStore", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EfesSetPrecrimeIterator is returned from FilterSetPrecrime and is used to iterate over the raw logs and unpacked data for SetPrecrime events raised by the Efes contract.
type EfesSetPrecrimeIterator struct {
	Event *EfesSetPrecrime // Event containing the contract specifics and raw log

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
func (it *EfesSetPrecrimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EfesSetPrecrime)
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
		it.Event = new(EfesSetPrecrime)
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
func (it *EfesSetPrecrimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EfesSetPrecrimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EfesSetPrecrime represents a SetPrecrime event raised by the Efes contract.
type EfesSetPrecrime struct {
	Precrime common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetPrecrime is a free log retrieval operation binding the contract event 0x5db758e995a17ec1ad84bdef7e8c3293a0bd6179bcce400dff5d4c3d87db726b.
//
// Solidity: event SetPrecrime(address precrime)
func (_Efes *EfesFilterer) FilterSetPrecrime(opts *bind.FilterOpts) (*EfesSetPrecrimeIterator, error) {

	logs, sub, err := _Efes.contract.FilterLogs(opts, "SetPrecrime")
	if err != nil {
		return nil, err
	}
	return &EfesSetPrecrimeIterator{contract: _Efes.contract, event: "SetPrecrime", logs: logs, sub: sub}, nil
}

// WatchSetPrecrime is a free log subscription operation binding the contract event 0x5db758e995a17ec1ad84bdef7e8c3293a0bd6179bcce400dff5d4c3d87db726b.
//
// Solidity: event SetPrecrime(address precrime)
func (_Efes *EfesFilterer) WatchSetPrecrime(opts *bind.WatchOpts, sink chan<- *EfesSetPrecrime) (event.Subscription, error) {

	logs, sub, err := _Efes.contract.WatchLogs(opts, "SetPrecrime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EfesSetPrecrime)
				if err := _Efes.contract.UnpackLog(event, "SetPrecrime", log); err != nil {
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

// ParseSetPrecrime is a log parse operation binding the contract event 0x5db758e995a17ec1ad84bdef7e8c3293a0bd6179bcce400dff5d4c3d87db726b.
//
// Solidity: event SetPrecrime(address precrime)
func (_Efes *EfesFilterer) ParseSetPrecrime(log types.Log) (*EfesSetPrecrime, error) {
	event := new(EfesSetPrecrime)
	if err := _Efes.contract.UnpackLog(event, "SetPrecrime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EfesSetTrustedRemoteIterator is returned from FilterSetTrustedRemote and is used to iterate over the raw logs and unpacked data for SetTrustedRemote events raised by the Efes contract.
type EfesSetTrustedRemoteIterator struct {
	Event *EfesSetTrustedRemote // Event containing the contract specifics and raw log

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
func (it *EfesSetTrustedRemoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EfesSetTrustedRemote)
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
		it.Event = new(EfesSetTrustedRemote)
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
func (it *EfesSetTrustedRemoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EfesSetTrustedRemoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EfesSetTrustedRemote represents a SetTrustedRemote event raised by the Efes contract.
type EfesSetTrustedRemote struct {
	RemoteChainId uint16
	Path          []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedRemote is a free log retrieval operation binding the contract event 0xfa41487ad5d6728f0b19276fa1eddc16558578f5109fc39d2dc33c3230470dab.
//
// Solidity: event SetTrustedRemote(uint16 _remoteChainId, bytes _path)
func (_Efes *EfesFilterer) FilterSetTrustedRemote(opts *bind.FilterOpts) (*EfesSetTrustedRemoteIterator, error) {

	logs, sub, err := _Efes.contract.FilterLogs(opts, "SetTrustedRemote")
	if err != nil {
		return nil, err
	}
	return &EfesSetTrustedRemoteIterator{contract: _Efes.contract, event: "SetTrustedRemote", logs: logs, sub: sub}, nil
}

// WatchSetTrustedRemote is a free log subscription operation binding the contract event 0xfa41487ad5d6728f0b19276fa1eddc16558578f5109fc39d2dc33c3230470dab.
//
// Solidity: event SetTrustedRemote(uint16 _remoteChainId, bytes _path)
func (_Efes *EfesFilterer) WatchSetTrustedRemote(opts *bind.WatchOpts, sink chan<- *EfesSetTrustedRemote) (event.Subscription, error) {

	logs, sub, err := _Efes.contract.WatchLogs(opts, "SetTrustedRemote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EfesSetTrustedRemote)
				if err := _Efes.contract.UnpackLog(event, "SetTrustedRemote", log); err != nil {
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

// ParseSetTrustedRemote is a log parse operation binding the contract event 0xfa41487ad5d6728f0b19276fa1eddc16558578f5109fc39d2dc33c3230470dab.
//
// Solidity: event SetTrustedRemote(uint16 _remoteChainId, bytes _path)
func (_Efes *EfesFilterer) ParseSetTrustedRemote(log types.Log) (*EfesSetTrustedRemote, error) {
	event := new(EfesSetTrustedRemote)
	if err := _Efes.contract.UnpackLog(event, "SetTrustedRemote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EfesSetTrustedRemoteAddressIterator is returned from FilterSetTrustedRemoteAddress and is used to iterate over the raw logs and unpacked data for SetTrustedRemoteAddress events raised by the Efes contract.
type EfesSetTrustedRemoteAddressIterator struct {
	Event *EfesSetTrustedRemoteAddress // Event containing the contract specifics and raw log

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
func (it *EfesSetTrustedRemoteAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EfesSetTrustedRemoteAddress)
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
		it.Event = new(EfesSetTrustedRemoteAddress)
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
func (it *EfesSetTrustedRemoteAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EfesSetTrustedRemoteAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EfesSetTrustedRemoteAddress represents a SetTrustedRemoteAddress event raised by the Efes contract.
type EfesSetTrustedRemoteAddress struct {
	RemoteChainId uint16
	RemoteAddress []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedRemoteAddress is a free log retrieval operation binding the contract event 0x8c0400cfe2d1199b1a725c78960bcc2a344d869b80590d0f2bd005db15a572ce.
//
// Solidity: event SetTrustedRemoteAddress(uint16 _remoteChainId, bytes _remoteAddress)
func (_Efes *EfesFilterer) FilterSetTrustedRemoteAddress(opts *bind.FilterOpts) (*EfesSetTrustedRemoteAddressIterator, error) {

	logs, sub, err := _Efes.contract.FilterLogs(opts, "SetTrustedRemoteAddress")
	if err != nil {
		return nil, err
	}
	return &EfesSetTrustedRemoteAddressIterator{contract: _Efes.contract, event: "SetTrustedRemoteAddress", logs: logs, sub: sub}, nil
}

// WatchSetTrustedRemoteAddress is a free log subscription operation binding the contract event 0x8c0400cfe2d1199b1a725c78960bcc2a344d869b80590d0f2bd005db15a572ce.
//
// Solidity: event SetTrustedRemoteAddress(uint16 _remoteChainId, bytes _remoteAddress)
func (_Efes *EfesFilterer) WatchSetTrustedRemoteAddress(opts *bind.WatchOpts, sink chan<- *EfesSetTrustedRemoteAddress) (event.Subscription, error) {

	logs, sub, err := _Efes.contract.WatchLogs(opts, "SetTrustedRemoteAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EfesSetTrustedRemoteAddress)
				if err := _Efes.contract.UnpackLog(event, "SetTrustedRemoteAddress", log); err != nil {
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

// ParseSetTrustedRemoteAddress is a log parse operation binding the contract event 0x8c0400cfe2d1199b1a725c78960bcc2a344d869b80590d0f2bd005db15a572ce.
//
// Solidity: event SetTrustedRemoteAddress(uint16 _remoteChainId, bytes _remoteAddress)
func (_Efes *EfesFilterer) ParseSetTrustedRemoteAddress(log types.Log) (*EfesSetTrustedRemoteAddress, error) {
	event := new(EfesSetTrustedRemoteAddress)
	if err := _Efes.contract.UnpackLog(event, "SetTrustedRemoteAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EfesTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Efes contract.
type EfesTransferIterator struct {
	Event *EfesTransfer // Event containing the contract specifics and raw log

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
func (it *EfesTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EfesTransfer)
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
		it.Event = new(EfesTransfer)
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
func (it *EfesTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EfesTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EfesTransfer represents a Transfer event raised by the Efes contract.
type EfesTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Efes *EfesFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*EfesTransferIterator, error) {

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

	logs, sub, err := _Efes.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &EfesTransferIterator{contract: _Efes.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Efes *EfesFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *EfesTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Efes.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EfesTransfer)
				if err := _Efes.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Efes *EfesFilterer) ParseTransfer(log types.Log) (*EfesTransfer, error) {
	event := new(EfesTransfer)
	if err := _Efes.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EfesWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Efes contract.
type EfesWithdrawIterator struct {
	Event *EfesWithdraw // Event containing the contract specifics and raw log

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
func (it *EfesWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EfesWithdraw)
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
		it.Event = new(EfesWithdraw)
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
func (it *EfesWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EfesWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EfesWithdraw represents a Withdraw event raised by the Efes contract.
type EfesWithdraw struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x5b6b431d4476a211bb7d41c20d1aab9ae2321deee0d20be3d9fc9b1093fa6e3d.
//
// Solidity: event Withdraw(uint256 value)
func (_Efes *EfesFilterer) FilterWithdraw(opts *bind.FilterOpts) (*EfesWithdrawIterator, error) {

	logs, sub, err := _Efes.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &EfesWithdrawIterator{contract: _Efes.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x5b6b431d4476a211bb7d41c20d1aab9ae2321deee0d20be3d9fc9b1093fa6e3d.
//
// Solidity: event Withdraw(uint256 value)
func (_Efes *EfesFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *EfesWithdraw) (event.Subscription, error) {

	logs, sub, err := _Efes.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EfesWithdraw)
				if err := _Efes.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_Efes *EfesFilterer) ParseWithdraw(log types.Log) (*EfesWithdraw, error) {
	event := new(EfesWithdraw)
	if err := _Efes.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
