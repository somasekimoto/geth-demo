// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gocode

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

// GocodeMetaData contains all meta data concerning the Gocode contract.
var GocodeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"Greet\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Hello\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506101c38061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610034575f3560e01c8063bcdfe0d514610038578063efdeaaf514610071575b5f80fd5b60408051808201909152600b81526a12195b1b1bc815dbdc9b1960aa1b60208201525b6040516100689190610082565b60405180910390f35b61005b61007f3660046100e2565b90565b5f602080835283518060208501525f5b818110156100ae57858101830151858201604001528201610092565b505f604082860101526040601f19601f8301168501019250505092915050565b634e487b7160e01b5f52604160045260245ffd5b5f602082840312156100f2575f80fd5b813567ffffffffffffffff80821115610109575f80fd5b818401915084601f83011261011c575f80fd5b81358181111561012e5761012e6100ce565b604051601f8201601f19908116603f01168101908382118183101715610156576101566100ce565b8160405282815287602084870101111561016e575f80fd5b826020860160208301375f92810160200192909252509594505050505056fea2646970667358221220641b60d99708a155d2a410ac09e59405c7ee61152a294efeb4e9d0e60002912b64736f6c63430008170033",
}

// GocodeABI is the input ABI used to generate the binding from.
// Deprecated: Use GocodeMetaData.ABI instead.
var GocodeABI = GocodeMetaData.ABI

// GocodeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GocodeMetaData.Bin instead.
var GocodeBin = GocodeMetaData.Bin

// DeployGocode deploys a new Ethereum contract, binding an instance of Gocode to it.
func DeployGocode(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Gocode, error) {
	parsed, err := GocodeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GocodeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Gocode{GocodeCaller: GocodeCaller{contract: contract}, GocodeTransactor: GocodeTransactor{contract: contract}, GocodeFilterer: GocodeFilterer{contract: contract}}, nil
}

// Gocode is an auto generated Go binding around an Ethereum contract.
type Gocode struct {
	GocodeCaller     // Read-only binding to the contract
	GocodeTransactor // Write-only binding to the contract
	GocodeFilterer   // Log filterer for contract events
}

// GocodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type GocodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GocodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GocodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GocodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GocodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GocodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GocodeSession struct {
	Contract     *Gocode           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GocodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GocodeCallerSession struct {
	Contract *GocodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// GocodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GocodeTransactorSession struct {
	Contract     *GocodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GocodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type GocodeRaw struct {
	Contract *Gocode // Generic contract binding to access the raw methods on
}

// GocodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GocodeCallerRaw struct {
	Contract *GocodeCaller // Generic read-only contract binding to access the raw methods on
}

// GocodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GocodeTransactorRaw struct {
	Contract *GocodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGocode creates a new instance of Gocode, bound to a specific deployed contract.
func NewGocode(address common.Address, backend bind.ContractBackend) (*Gocode, error) {
	contract, err := bindGocode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gocode{GocodeCaller: GocodeCaller{contract: contract}, GocodeTransactor: GocodeTransactor{contract: contract}, GocodeFilterer: GocodeFilterer{contract: contract}}, nil
}

// NewGocodeCaller creates a new read-only instance of Gocode, bound to a specific deployed contract.
func NewGocodeCaller(address common.Address, caller bind.ContractCaller) (*GocodeCaller, error) {
	contract, err := bindGocode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GocodeCaller{contract: contract}, nil
}

// NewGocodeTransactor creates a new write-only instance of Gocode, bound to a specific deployed contract.
func NewGocodeTransactor(address common.Address, transactor bind.ContractTransactor) (*GocodeTransactor, error) {
	contract, err := bindGocode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GocodeTransactor{contract: contract}, nil
}

// NewGocodeFilterer creates a new log filterer instance of Gocode, bound to a specific deployed contract.
func NewGocodeFilterer(address common.Address, filterer bind.ContractFilterer) (*GocodeFilterer, error) {
	contract, err := bindGocode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GocodeFilterer{contract: contract}, nil
}

// bindGocode binds a generic wrapper to an already deployed contract.
func bindGocode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GocodeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gocode *GocodeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gocode.Contract.GocodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gocode *GocodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gocode.Contract.GocodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gocode *GocodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gocode.Contract.GocodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gocode *GocodeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gocode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gocode *GocodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gocode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gocode *GocodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gocode.Contract.contract.Transact(opts, method, params...)
}

// Greet is a free data retrieval call binding the contract method 0xefdeaaf5.
//
// Solidity: function Greet(string str) pure returns(string)
func (_Gocode *GocodeCaller) Greet(opts *bind.CallOpts, str string) (string, error) {
	var out []interface{}
	err := _Gocode.contract.Call(opts, &out, "Greet", str)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Greet is a free data retrieval call binding the contract method 0xefdeaaf5.
//
// Solidity: function Greet(string str) pure returns(string)
func (_Gocode *GocodeSession) Greet(str string) (string, error) {
	return _Gocode.Contract.Greet(&_Gocode.CallOpts, str)
}

// Greet is a free data retrieval call binding the contract method 0xefdeaaf5.
//
// Solidity: function Greet(string str) pure returns(string)
func (_Gocode *GocodeCallerSession) Greet(str string) (string, error) {
	return _Gocode.Contract.Greet(&_Gocode.CallOpts, str)
}

// Hello is a free data retrieval call binding the contract method 0xbcdfe0d5.
//
// Solidity: function Hello() pure returns(string)
func (_Gocode *GocodeCaller) Hello(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Gocode.contract.Call(opts, &out, "Hello")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Hello is a free data retrieval call binding the contract method 0xbcdfe0d5.
//
// Solidity: function Hello() pure returns(string)
func (_Gocode *GocodeSession) Hello() (string, error) {
	return _Gocode.Contract.Hello(&_Gocode.CallOpts)
}

// Hello is a free data retrieval call binding the contract method 0xbcdfe0d5.
//
// Solidity: function Hello() pure returns(string)
func (_Gocode *GocodeCallerSession) Hello() (string, error) {
	return _Gocode.Contract.Hello(&_Gocode.CallOpts)
}
