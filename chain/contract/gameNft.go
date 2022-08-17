// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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
)

// GameNftMetaData contains all meta data concerning the GameNft contract.
var GameNftMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"AdminEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"expires\",\"type\":\"uint64\"}],\"name\":\"UpdateUser\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admins\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"expires\",\"type\":\"uint64\"}],\"name\":\"setUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"direct\",\"type\":\"bool\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"userExpires\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"userOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_tokenURI\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"feeNumerator\",\"type\":\"uint96\"}],\"name\":\"setDefaultRoyalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"feeNumerator\",\"type\":\"uint96\"}],\"name\":\"setTokenRoyalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_tokenURI\",\"type\":\"string\"}],\"name\":\"setTokenURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI_\",\"type\":\"string\"}],\"name\":\"setBaseTokenURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"enableAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"disableAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// GameNftABI is the input ABI used to generate the binding from.
// Deprecated: Use GameNftMetaData.ABI instead.
var GameNftABI = GameNftMetaData.ABI

// GameNft is an auto generated Go binding around an Ethereum contract.
type GameNft struct {
	GameNftCaller     // Read-only binding to the contract
	GameNftTransactor // Write-only binding to the contract
	GameNftFilterer   // Log filterer for contract events
}

// GameNftCaller is an auto generated read-only Go binding around an Ethereum contract.
type GameNftCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameNftTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GameNftTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameNftFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GameNftFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameNftSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GameNftSession struct {
	Contract     *GameNft          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GameNftCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GameNftCallerSession struct {
	Contract *GameNftCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// GameNftTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GameNftTransactorSession struct {
	Contract     *GameNftTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GameNftRaw is an auto generated low-level Go binding around an Ethereum contract.
type GameNftRaw struct {
	Contract *GameNft // Generic contract binding to access the raw methods on
}

// GameNftCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GameNftCallerRaw struct {
	Contract *GameNftCaller // Generic read-only contract binding to access the raw methods on
}

// GameNftTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GameNftTransactorRaw struct {
	Contract *GameNftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGameNft creates a new instance of GameNft, bound to a specific deployed contract.
func NewGameNft(address common.Address, backend bind.ContractBackend) (*GameNft, error) {
	contract, err := bindGameNft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GameNft{GameNftCaller: GameNftCaller{contract: contract}, GameNftTransactor: GameNftTransactor{contract: contract}, GameNftFilterer: GameNftFilterer{contract: contract}}, nil
}

// NewGameNftCaller creates a new read-only instance of GameNft, bound to a specific deployed contract.
func NewGameNftCaller(address common.Address, caller bind.ContractCaller) (*GameNftCaller, error) {
	contract, err := bindGameNft(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GameNftCaller{contract: contract}, nil
}

// NewGameNftTransactor creates a new write-only instance of GameNft, bound to a specific deployed contract.
func NewGameNftTransactor(address common.Address, transactor bind.ContractTransactor) (*GameNftTransactor, error) {
	contract, err := bindGameNft(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GameNftTransactor{contract: contract}, nil
}

// NewGameNftFilterer creates a new log filterer instance of GameNft, bound to a specific deployed contract.
func NewGameNftFilterer(address common.Address, filterer bind.ContractFilterer) (*GameNftFilterer, error) {
	contract, err := bindGameNft(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GameNftFilterer{contract: contract}, nil
}

// bindGameNft binds a generic wrapper to an already deployed contract.
func bindGameNft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GameNftABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GameNft *GameNftRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GameNft.Contract.GameNftCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GameNft *GameNftRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameNft.Contract.GameNftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GameNft *GameNftRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GameNft.Contract.GameNftTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GameNft *GameNftCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GameNft.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GameNft *GameNftTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameNft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GameNft *GameNftTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GameNft.Contract.contract.Transact(opts, method, params...)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_GameNft *GameNftCaller) Admins(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _GameNft.contract.Call(opts, &out, "admins", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_GameNft *GameNftSession) Admins(arg0 common.Address) (bool, error) {
	return _GameNft.Contract.Admins(&_GameNft.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_GameNft *GameNftCallerSession) Admins(arg0 common.Address) (bool, error) {
	return _GameNft.Contract.Admins(&_GameNft.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_GameNft *GameNftCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GameNft.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_GameNft *GameNftSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _GameNft.Contract.BalanceOf(&_GameNft.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_GameNft *GameNftCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _GameNft.Contract.BalanceOf(&_GameNft.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_GameNft *GameNftCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _GameNft.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_GameNft *GameNftSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _GameNft.Contract.GetApproved(&_GameNft.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_GameNft *GameNftCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _GameNft.Contract.GetApproved(&_GameNft.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_GameNft *GameNftCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _GameNft.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_GameNft *GameNftSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _GameNft.Contract.IsApprovedForAll(&_GameNft.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_GameNft *GameNftCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _GameNft.Contract.IsApprovedForAll(&_GameNft.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GameNft *GameNftCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GameNft.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GameNft *GameNftSession) Name() (string, error) {
	return _GameNft.Contract.Name(&_GameNft.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GameNft *GameNftCallerSession) Name() (string, error) {
	return _GameNft.Contract.Name(&_GameNft.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GameNft *GameNftCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GameNft.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GameNft *GameNftSession) Owner() (common.Address, error) {
	return _GameNft.Contract.Owner(&_GameNft.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GameNft *GameNftCallerSession) Owner() (common.Address, error) {
	return _GameNft.Contract.Owner(&_GameNft.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_GameNft *GameNftCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _GameNft.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_GameNft *GameNftSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _GameNft.Contract.OwnerOf(&_GameNft.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_GameNft *GameNftCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _GameNft.Contract.OwnerOf(&_GameNft.CallOpts, tokenId)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_GameNft *GameNftCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GameNft.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_GameNft *GameNftSession) PendingOwner() (common.Address, error) {
	return _GameNft.Contract.PendingOwner(&_GameNft.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_GameNft *GameNftCallerSession) PendingOwner() (common.Address, error) {
	return _GameNft.Contract.PendingOwner(&_GameNft.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address, uint256)
func (_GameNft *GameNftCaller) RoyaltyInfo(opts *bind.CallOpts, _tokenId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _GameNft.contract.Call(opts, &out, "royaltyInfo", _tokenId, _salePrice)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address, uint256)
func (_GameNft *GameNftSession) RoyaltyInfo(_tokenId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	return _GameNft.Contract.RoyaltyInfo(&_GameNft.CallOpts, _tokenId, _salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address, uint256)
func (_GameNft *GameNftCallerSession) RoyaltyInfo(_tokenId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	return _GameNft.Contract.RoyaltyInfo(&_GameNft.CallOpts, _tokenId, _salePrice)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GameNft *GameNftCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _GameNft.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GameNft *GameNftSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GameNft.Contract.SupportsInterface(&_GameNft.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GameNft *GameNftCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GameNft.Contract.SupportsInterface(&_GameNft.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GameNft *GameNftCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GameNft.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GameNft *GameNftSession) Symbol() (string, error) {
	return _GameNft.Contract.Symbol(&_GameNft.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GameNft *GameNftCallerSession) Symbol() (string, error) {
	return _GameNft.Contract.Symbol(&_GameNft.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_GameNft *GameNftCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _GameNft.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_GameNft *GameNftSession) TokenURI(tokenId *big.Int) (string, error) {
	return _GameNft.Contract.TokenURI(&_GameNft.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_GameNft *GameNftCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _GameNft.Contract.TokenURI(&_GameNft.CallOpts, tokenId)
}

// UserExpires is a free data retrieval call binding the contract method 0x8fc88c48.
//
// Solidity: function userExpires(uint256 tokenId) view returns(uint256)
func (_GameNft *GameNftCaller) UserExpires(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GameNft.contract.Call(opts, &out, "userExpires", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserExpires is a free data retrieval call binding the contract method 0x8fc88c48.
//
// Solidity: function userExpires(uint256 tokenId) view returns(uint256)
func (_GameNft *GameNftSession) UserExpires(tokenId *big.Int) (*big.Int, error) {
	return _GameNft.Contract.UserExpires(&_GameNft.CallOpts, tokenId)
}

// UserExpires is a free data retrieval call binding the contract method 0x8fc88c48.
//
// Solidity: function userExpires(uint256 tokenId) view returns(uint256)
func (_GameNft *GameNftCallerSession) UserExpires(tokenId *big.Int) (*big.Int, error) {
	return _GameNft.Contract.UserExpires(&_GameNft.CallOpts, tokenId)
}

// UserOf is a free data retrieval call binding the contract method 0xc2f1f14a.
//
// Solidity: function userOf(uint256 tokenId) view returns(address)
func (_GameNft *GameNftCaller) UserOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _GameNft.contract.Call(opts, &out, "userOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserOf is a free data retrieval call binding the contract method 0xc2f1f14a.
//
// Solidity: function userOf(uint256 tokenId) view returns(address)
func (_GameNft *GameNftSession) UserOf(tokenId *big.Int) (common.Address, error) {
	return _GameNft.Contract.UserOf(&_GameNft.CallOpts, tokenId)
}

// UserOf is a free data retrieval call binding the contract method 0xc2f1f14a.
//
// Solidity: function userOf(uint256 tokenId) view returns(address)
func (_GameNft *GameNftCallerSession) UserOf(tokenId *big.Int) (common.Address, error) {
	return _GameNft.Contract.UserOf(&_GameNft.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_GameNft *GameNftTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GameNft.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_GameNft *GameNftSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GameNft.Contract.Approve(&_GameNft.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_GameNft *GameNftTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GameNft.Contract.Approve(&_GameNft.TransactOpts, to, tokenId)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_GameNft *GameNftTransactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameNft.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_GameNft *GameNftSession) ClaimOwnership() (*types.Transaction, error) {
	return _GameNft.Contract.ClaimOwnership(&_GameNft.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_GameNft *GameNftTransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _GameNft.Contract.ClaimOwnership(&_GameNft.TransactOpts)
}

// DisableAdmin is a paid mutator transaction binding the contract method 0x751e9a9c.
//
// Solidity: function disableAdmin(address _addr) returns()
func (_GameNft *GameNftTransactor) DisableAdmin(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _GameNft.contract.Transact(opts, "disableAdmin", _addr)
}

// DisableAdmin is a paid mutator transaction binding the contract method 0x751e9a9c.
//
// Solidity: function disableAdmin(address _addr) returns()
func (_GameNft *GameNftSession) DisableAdmin(_addr common.Address) (*types.Transaction, error) {
	return _GameNft.Contract.DisableAdmin(&_GameNft.TransactOpts, _addr)
}

// DisableAdmin is a paid mutator transaction binding the contract method 0x751e9a9c.
//
// Solidity: function disableAdmin(address _addr) returns()
func (_GameNft *GameNftTransactorSession) DisableAdmin(_addr common.Address) (*types.Transaction, error) {
	return _GameNft.Contract.DisableAdmin(&_GameNft.TransactOpts, _addr)
}

// EnableAdmin is a paid mutator transaction binding the contract method 0xbea532ff.
//
// Solidity: function enableAdmin(address _addr) returns()
func (_GameNft *GameNftTransactor) EnableAdmin(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _GameNft.contract.Transact(opts, "enableAdmin", _addr)
}

// EnableAdmin is a paid mutator transaction binding the contract method 0xbea532ff.
//
// Solidity: function enableAdmin(address _addr) returns()
func (_GameNft *GameNftSession) EnableAdmin(_addr common.Address) (*types.Transaction, error) {
	return _GameNft.Contract.EnableAdmin(&_GameNft.TransactOpts, _addr)
}

// EnableAdmin is a paid mutator transaction binding the contract method 0xbea532ff.
//
// Solidity: function enableAdmin(address _addr) returns()
func (_GameNft *GameNftTransactorSession) EnableAdmin(_addr common.Address) (*types.Transaction, error) {
	return _GameNft.Contract.EnableAdmin(&_GameNft.TransactOpts, _addr)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 tokenId, address to) returns()
func (_GameNft *GameNftTransactor) Mint(opts *bind.TransactOpts, tokenId *big.Int, to common.Address) (*types.Transaction, error) {
	return _GameNft.contract.Transact(opts, "mint", tokenId, to)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 tokenId, address to) returns()
func (_GameNft *GameNftSession) Mint(tokenId *big.Int, to common.Address) (*types.Transaction, error) {
	return _GameNft.Contract.Mint(&_GameNft.TransactOpts, tokenId, to)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 tokenId, address to) returns()
func (_GameNft *GameNftTransactorSession) Mint(tokenId *big.Int, to common.Address) (*types.Transaction, error) {
	return _GameNft.Contract.Mint(&_GameNft.TransactOpts, tokenId, to)
}

// Mint0 is a paid mutator transaction binding the contract method 0xe67e402c.
//
// Solidity: function mint(uint256 tokenId, address to, string _tokenURI) returns()
func (_GameNft *GameNftTransactor) Mint0(opts *bind.TransactOpts, tokenId *big.Int, to common.Address, _tokenURI string) (*types.Transaction, error) {
	return _GameNft.contract.Transact(opts, "mint0", tokenId, to, _tokenURI)
}

// Mint0 is a paid mutator transaction binding the contract method 0xe67e402c.
//
// Solidity: function mint(uint256 tokenId, address to, string _tokenURI) returns()
func (_GameNft *GameNftSession) Mint0(tokenId *big.Int, to common.Address, _tokenURI string) (*types.Transaction, error) {
	return _GameNft.Contract.Mint0(&_GameNft.TransactOpts, tokenId, to, _tokenURI)
}

// Mint0 is a paid mutator transaction binding the contract method 0xe67e402c.
//
// Solidity: function mint(uint256 tokenId, address to, string _tokenURI) returns()
func (_GameNft *GameNftTransactorSession) Mint0(tokenId *big.Int, to common.Address, _tokenURI string) (*types.Transaction, error) {
	return _GameNft.Contract.Mint0(&_GameNft.TransactOpts, tokenId, to, _tokenURI)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GameNft *GameNftTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameNft.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GameNft *GameNftSession) RenounceOwnership() (*types.Transaction, error) {
	return _GameNft.Contract.RenounceOwnership(&_GameNft.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GameNft *GameNftTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GameNft.Contract.RenounceOwnership(&_GameNft.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_GameNft *GameNftTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GameNft.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_GameNft *GameNftSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GameNft.Contract.SafeTransferFrom(&_GameNft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_GameNft *GameNftTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GameNft.Contract.SafeTransferFrom(&_GameNft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_GameNft *GameNftTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _GameNft.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_GameNft *GameNftSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _GameNft.Contract.SafeTransferFrom0(&_GameNft.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_GameNft *GameNftTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _GameNft.Contract.SafeTransferFrom0(&_GameNft.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_GameNft *GameNftTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _GameNft.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_GameNft *GameNftSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _GameNft.Contract.SetApprovalForAll(&_GameNft.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_GameNft *GameNftTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _GameNft.Contract.SetApprovalForAll(&_GameNft.TransactOpts, operator, approved)
}

// SetBaseTokenURI is a paid mutator transaction binding the contract method 0x30176e13.
//
// Solidity: function setBaseTokenURI(string baseURI_) returns()
func (_GameNft *GameNftTransactor) SetBaseTokenURI(opts *bind.TransactOpts, baseURI_ string) (*types.Transaction, error) {
	return _GameNft.contract.Transact(opts, "setBaseTokenURI", baseURI_)
}

// SetBaseTokenURI is a paid mutator transaction binding the contract method 0x30176e13.
//
// Solidity: function setBaseTokenURI(string baseURI_) returns()
func (_GameNft *GameNftSession) SetBaseTokenURI(baseURI_ string) (*types.Transaction, error) {
	return _GameNft.Contract.SetBaseTokenURI(&_GameNft.TransactOpts, baseURI_)
}

// SetBaseTokenURI is a paid mutator transaction binding the contract method 0x30176e13.
//
// Solidity: function setBaseTokenURI(string baseURI_) returns()
func (_GameNft *GameNftTransactorSession) SetBaseTokenURI(baseURI_ string) (*types.Transaction, error) {
	return _GameNft.Contract.SetBaseTokenURI(&_GameNft.TransactOpts, baseURI_)
}

// SetDefaultRoyalty is a paid mutator transaction binding the contract method 0x04634d8d.
//
// Solidity: function setDefaultRoyalty(address receiver, uint96 feeNumerator) returns()
func (_GameNft *GameNftTransactor) SetDefaultRoyalty(opts *bind.TransactOpts, receiver common.Address, feeNumerator *big.Int) (*types.Transaction, error) {
	return _GameNft.contract.Transact(opts, "setDefaultRoyalty", receiver, feeNumerator)
}

// SetDefaultRoyalty is a paid mutator transaction binding the contract method 0x04634d8d.
//
// Solidity: function setDefaultRoyalty(address receiver, uint96 feeNumerator) returns()
func (_GameNft *GameNftSession) SetDefaultRoyalty(receiver common.Address, feeNumerator *big.Int) (*types.Transaction, error) {
	return _GameNft.Contract.SetDefaultRoyalty(&_GameNft.TransactOpts, receiver, feeNumerator)
}

// SetDefaultRoyalty is a paid mutator transaction binding the contract method 0x04634d8d.
//
// Solidity: function setDefaultRoyalty(address receiver, uint96 feeNumerator) returns()
func (_GameNft *GameNftTransactorSession) SetDefaultRoyalty(receiver common.Address, feeNumerator *big.Int) (*types.Transaction, error) {
	return _GameNft.Contract.SetDefaultRoyalty(&_GameNft.TransactOpts, receiver, feeNumerator)
}

// SetTokenRoyalty is a paid mutator transaction binding the contract method 0x5944c753.
//
// Solidity: function setTokenRoyalty(uint256 tokenId, address receiver, uint96 feeNumerator) returns()
func (_GameNft *GameNftTransactor) SetTokenRoyalty(opts *bind.TransactOpts, tokenId *big.Int, receiver common.Address, feeNumerator *big.Int) (*types.Transaction, error) {
	return _GameNft.contract.Transact(opts, "setTokenRoyalty", tokenId, receiver, feeNumerator)
}

// SetTokenRoyalty is a paid mutator transaction binding the contract method 0x5944c753.
//
// Solidity: function setTokenRoyalty(uint256 tokenId, address receiver, uint96 feeNumerator) returns()
func (_GameNft *GameNftSession) SetTokenRoyalty(tokenId *big.Int, receiver common.Address, feeNumerator *big.Int) (*types.Transaction, error) {
	return _GameNft.Contract.SetTokenRoyalty(&_GameNft.TransactOpts, tokenId, receiver, feeNumerator)
}

// SetTokenRoyalty is a paid mutator transaction binding the contract method 0x5944c753.
//
// Solidity: function setTokenRoyalty(uint256 tokenId, address receiver, uint96 feeNumerator) returns()
func (_GameNft *GameNftTransactorSession) SetTokenRoyalty(tokenId *big.Int, receiver common.Address, feeNumerator *big.Int) (*types.Transaction, error) {
	return _GameNft.Contract.SetTokenRoyalty(&_GameNft.TransactOpts, tokenId, receiver, feeNumerator)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0x162094c4.
//
// Solidity: function setTokenURI(uint256 tokenId, string _tokenURI) returns()
func (_GameNft *GameNftTransactor) SetTokenURI(opts *bind.TransactOpts, tokenId *big.Int, _tokenURI string) (*types.Transaction, error) {
	return _GameNft.contract.Transact(opts, "setTokenURI", tokenId, _tokenURI)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0x162094c4.
//
// Solidity: function setTokenURI(uint256 tokenId, string _tokenURI) returns()
func (_GameNft *GameNftSession) SetTokenURI(tokenId *big.Int, _tokenURI string) (*types.Transaction, error) {
	return _GameNft.Contract.SetTokenURI(&_GameNft.TransactOpts, tokenId, _tokenURI)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0x162094c4.
//
// Solidity: function setTokenURI(uint256 tokenId, string _tokenURI) returns()
func (_GameNft *GameNftTransactorSession) SetTokenURI(tokenId *big.Int, _tokenURI string) (*types.Transaction, error) {
	return _GameNft.Contract.SetTokenURI(&_GameNft.TransactOpts, tokenId, _tokenURI)
}

// SetUser is a paid mutator transaction binding the contract method 0xe030565e.
//
// Solidity: function setUser(uint256 tokenId, address user, uint64 expires) returns()
func (_GameNft *GameNftTransactor) SetUser(opts *bind.TransactOpts, tokenId *big.Int, user common.Address, expires uint64) (*types.Transaction, error) {
	return _GameNft.contract.Transact(opts, "setUser", tokenId, user, expires)
}

// SetUser is a paid mutator transaction binding the contract method 0xe030565e.
//
// Solidity: function setUser(uint256 tokenId, address user, uint64 expires) returns()
func (_GameNft *GameNftSession) SetUser(tokenId *big.Int, user common.Address, expires uint64) (*types.Transaction, error) {
	return _GameNft.Contract.SetUser(&_GameNft.TransactOpts, tokenId, user, expires)
}

// SetUser is a paid mutator transaction binding the contract method 0xe030565e.
//
// Solidity: function setUser(uint256 tokenId, address user, uint64 expires) returns()
func (_GameNft *GameNftTransactorSession) SetUser(tokenId *big.Int, user common.Address, expires uint64) (*types.Transaction, error) {
	return _GameNft.Contract.SetUser(&_GameNft.TransactOpts, tokenId, user, expires)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_GameNft *GameNftTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GameNft.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_GameNft *GameNftSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GameNft.Contract.TransferFrom(&_GameNft.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_GameNft *GameNftTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GameNft.Contract.TransferFrom(&_GameNft.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xb242e534.
//
// Solidity: function transferOwnership(address newOwner, bool direct) returns()
func (_GameNft *GameNftTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address, direct bool) (*types.Transaction, error) {
	return _GameNft.contract.Transact(opts, "transferOwnership", newOwner, direct)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xb242e534.
//
// Solidity: function transferOwnership(address newOwner, bool direct) returns()
func (_GameNft *GameNftSession) TransferOwnership(newOwner common.Address, direct bool) (*types.Transaction, error) {
	return _GameNft.Contract.TransferOwnership(&_GameNft.TransactOpts, newOwner, direct)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xb242e534.
//
// Solidity: function transferOwnership(address newOwner, bool direct) returns()
func (_GameNft *GameNftTransactorSession) TransferOwnership(newOwner common.Address, direct bool) (*types.Transaction, error) {
	return _GameNft.Contract.TransferOwnership(&_GameNft.TransactOpts, newOwner, direct)
}

// GameNftAdminEnabledIterator is returned from FilterAdminEnabled and is used to iterate over the raw logs and unpacked data for AdminEnabled events raised by the GameNft contract.
type GameNftAdminEnabledIterator struct {
	Event *GameNftAdminEnabled // Event containing the contract specifics and raw log

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
func (it *GameNftAdminEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameNftAdminEnabled)
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
		it.Event = new(GameNftAdminEnabled)
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
func (it *GameNftAdminEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameNftAdminEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameNftAdminEnabled represents a AdminEnabled event raised by the GameNft contract.
type GameNftAdminEnabled struct {
	Admin   common.Address
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminEnabled is a free log retrieval operation binding the contract event 0xb3714ef62726d54fcac8919290b50b999590b382adefc7f78eddf14aafb9ba5e.
//
// Solidity: event AdminEnabled(address admin, bool enabled)
func (_GameNft *GameNftFilterer) FilterAdminEnabled(opts *bind.FilterOpts) (*GameNftAdminEnabledIterator, error) {

	logs, sub, err := _GameNft.contract.FilterLogs(opts, "AdminEnabled")
	if err != nil {
		return nil, err
	}
	return &GameNftAdminEnabledIterator{contract: _GameNft.contract, event: "AdminEnabled", logs: logs, sub: sub}, nil
}

// WatchAdminEnabled is a free log subscription operation binding the contract event 0xb3714ef62726d54fcac8919290b50b999590b382adefc7f78eddf14aafb9ba5e.
//
// Solidity: event AdminEnabled(address admin, bool enabled)
func (_GameNft *GameNftFilterer) WatchAdminEnabled(opts *bind.WatchOpts, sink chan<- *GameNftAdminEnabled) (event.Subscription, error) {

	logs, sub, err := _GameNft.contract.WatchLogs(opts, "AdminEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameNftAdminEnabled)
				if err := _GameNft.contract.UnpackLog(event, "AdminEnabled", log); err != nil {
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

// ParseAdminEnabled is a log parse operation binding the contract event 0xb3714ef62726d54fcac8919290b50b999590b382adefc7f78eddf14aafb9ba5e.
//
// Solidity: event AdminEnabled(address admin, bool enabled)
func (_GameNft *GameNftFilterer) ParseAdminEnabled(log types.Log) (*GameNftAdminEnabled, error) {
	event := new(GameNftAdminEnabled)
	if err := _GameNft.contract.UnpackLog(event, "AdminEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameNftApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the GameNft contract.
type GameNftApprovalIterator struct {
	Event *GameNftApproval // Event containing the contract specifics and raw log

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
func (it *GameNftApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameNftApproval)
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
		it.Event = new(GameNftApproval)
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
func (it *GameNftApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameNftApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameNftApproval represents a Approval event raised by the GameNft contract.
type GameNftApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_GameNft *GameNftFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*GameNftApprovalIterator, error) {

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

	logs, sub, err := _GameNft.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &GameNftApprovalIterator{contract: _GameNft.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_GameNft *GameNftFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *GameNftApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _GameNft.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameNftApproval)
				if err := _GameNft.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_GameNft *GameNftFilterer) ParseApproval(log types.Log) (*GameNftApproval, error) {
	event := new(GameNftApproval)
	if err := _GameNft.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameNftApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the GameNft contract.
type GameNftApprovalForAllIterator struct {
	Event *GameNftApprovalForAll // Event containing the contract specifics and raw log

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
func (it *GameNftApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameNftApprovalForAll)
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
		it.Event = new(GameNftApprovalForAll)
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
func (it *GameNftApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameNftApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameNftApprovalForAll represents a ApprovalForAll event raised by the GameNft contract.
type GameNftApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_GameNft *GameNftFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*GameNftApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _GameNft.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &GameNftApprovalForAllIterator{contract: _GameNft.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_GameNft *GameNftFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *GameNftApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _GameNft.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameNftApprovalForAll)
				if err := _GameNft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_GameNft *GameNftFilterer) ParseApprovalForAll(log types.Log) (*GameNftApprovalForAll, error) {
	event := new(GameNftApprovalForAll)
	if err := _GameNft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameNftOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GameNft contract.
type GameNftOwnershipTransferredIterator struct {
	Event *GameNftOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GameNftOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameNftOwnershipTransferred)
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
		it.Event = new(GameNftOwnershipTransferred)
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
func (it *GameNftOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameNftOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameNftOwnershipTransferred represents a OwnershipTransferred event raised by the GameNft contract.
type GameNftOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GameNft *GameNftFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GameNftOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GameNft.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GameNftOwnershipTransferredIterator{contract: _GameNft.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GameNft *GameNftFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GameNftOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GameNft.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameNftOwnershipTransferred)
				if err := _GameNft.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_GameNft *GameNftFilterer) ParseOwnershipTransferred(log types.Log) (*GameNftOwnershipTransferred, error) {
	event := new(GameNftOwnershipTransferred)
	if err := _GameNft.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameNftTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the GameNft contract.
type GameNftTransferIterator struct {
	Event *GameNftTransfer // Event containing the contract specifics and raw log

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
func (it *GameNftTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameNftTransfer)
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
		it.Event = new(GameNftTransfer)
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
func (it *GameNftTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameNftTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameNftTransfer represents a Transfer event raised by the GameNft contract.
type GameNftTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_GameNft *GameNftFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*GameNftTransferIterator, error) {

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

	logs, sub, err := _GameNft.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &GameNftTransferIterator{contract: _GameNft.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_GameNft *GameNftFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *GameNftTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _GameNft.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameNftTransfer)
				if err := _GameNft.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_GameNft *GameNftFilterer) ParseTransfer(log types.Log) (*GameNftTransfer, error) {
	event := new(GameNftTransfer)
	if err := _GameNft.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameNftUpdateUserIterator is returned from FilterUpdateUser and is used to iterate over the raw logs and unpacked data for UpdateUser events raised by the GameNft contract.
type GameNftUpdateUserIterator struct {
	Event *GameNftUpdateUser // Event containing the contract specifics and raw log

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
func (it *GameNftUpdateUserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameNftUpdateUser)
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
		it.Event = new(GameNftUpdateUser)
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
func (it *GameNftUpdateUserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameNftUpdateUserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameNftUpdateUser represents a UpdateUser event raised by the GameNft contract.
type GameNftUpdateUser struct {
	TokenId *big.Int
	User    common.Address
	Expires uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateUser is a free log retrieval operation binding the contract event 0x4e06b4e7000e659094299b3533b47b6aa8ad048e95e872d23d1f4ee55af89cfe.
//
// Solidity: event UpdateUser(uint256 indexed tokenId, address indexed user, uint64 expires)
func (_GameNft *GameNftFilterer) FilterUpdateUser(opts *bind.FilterOpts, tokenId []*big.Int, user []common.Address) (*GameNftUpdateUserIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _GameNft.contract.FilterLogs(opts, "UpdateUser", tokenIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &GameNftUpdateUserIterator{contract: _GameNft.contract, event: "UpdateUser", logs: logs, sub: sub}, nil
}

// WatchUpdateUser is a free log subscription operation binding the contract event 0x4e06b4e7000e659094299b3533b47b6aa8ad048e95e872d23d1f4ee55af89cfe.
//
// Solidity: event UpdateUser(uint256 indexed tokenId, address indexed user, uint64 expires)
func (_GameNft *GameNftFilterer) WatchUpdateUser(opts *bind.WatchOpts, sink chan<- *GameNftUpdateUser, tokenId []*big.Int, user []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _GameNft.contract.WatchLogs(opts, "UpdateUser", tokenIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameNftUpdateUser)
				if err := _GameNft.contract.UnpackLog(event, "UpdateUser", log); err != nil {
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

// ParseUpdateUser is a log parse operation binding the contract event 0x4e06b4e7000e659094299b3533b47b6aa8ad048e95e872d23d1f4ee55af89cfe.
//
// Solidity: event UpdateUser(uint256 indexed tokenId, address indexed user, uint64 expires)
func (_GameNft *GameNftFilterer) ParseUpdateUser(log types.Log) (*GameNftUpdateUser, error) {
	event := new(GameNftUpdateUser)
	if err := _GameNft.contract.UnpackLog(event, "UpdateUser", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
