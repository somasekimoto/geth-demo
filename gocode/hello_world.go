// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package helloworld

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

// HelloworldMetaData contains all meta data concerning the Helloworld contract.
var HelloworldMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"GetBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetData\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"Greet\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Hello\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"SetData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506104e18061001d5f395ff3fe608060405234801561000f575f80fd5b506004361061007a575f3560e01c8063b7e8260c11610058578063b7e8260c146100ae578063bcdfe0d5146100c1578063efdeaaf5146100e8578063f8f8a912146100f9575f80fd5b806357ea89b61461007e57806373d4a13a1461008857806376b8e528146100a6575b5f80fd5b610086610107565b005b610090610133565b60405161009d919061025c565b60405180910390f35b6100906101be565b6100866100bc3660046102bc565b61024d565b60408051808201909152600b81526a12195b1b1bc815dbdc9b1960aa1b6020820152610090565b6100906100f63660046102bc565b90565b60405147815260200161009d565b60405133904780156108fc02915f818181858888f19350505050158015610130573d5f803e3d5ffd5b50565b5f805461013f90610367565b80601f016020809104026020016040519081016040528092919081815260200182805461016b90610367565b80156101b65780601f1061018d576101008083540402835291602001916101b6565b820191905f5260205f20905b81548152906001019060200180831161019957829003601f168201915b505050505081565b60605f80546101cc90610367565b80601f01602080910402602001604051908101604052809291908181526020018280546101f890610367565b80156102435780601f1061021a57610100808354040283529160200191610243565b820191905f5260205f20905b81548152906001019060200180831161022657829003601f168201915b5050505050905090565b5f61025882826103eb565b5050565b5f602080835283518060208501525f5b818110156102885785810183015185820160400152820161026c565b505f604082860101526040601f19601f8301168501019250505092915050565b634e487b7160e01b5f52604160045260245ffd5b5f602082840312156102cc575f80fd5b813567ffffffffffffffff808211156102e3575f80fd5b818401915084601f8301126102f6575f80fd5b813581811115610308576103086102a8565b604051601f8201601f19908116603f01168101908382118183101715610330576103306102a8565b81604052828152876020848701011115610348575f80fd5b826020860160208301375f928101602001929092525095945050505050565b600181811c9082168061037b57607f821691505b60208210810361039957634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156103e657805f5260205f20601f840160051c810160208510156103c45750805b601f840160051c820191505b818110156103e3575f81556001016103d0565b50505b505050565b815167ffffffffffffffff811115610405576104056102a8565b610419816104138454610367565b8461039f565b602080601f83116001811461044c575f84156104355750858301515b5f19600386901b1c1916600185901b1785556104a3565b5f85815260208120601f198616915b8281101561047a5788860151825594840194600190910190840161045b565b508582101561049757878501515f19600388901b60f8161c191681555b505060018460011b0185555b50505050505056fea26469706673582212203957cd4bf277b5a196201496304d3ddee001125df463cd39769ff5f5fb9943e064736f6c63430008170033",
}

// HelloworldABI is the input ABI used to generate the binding from.
// Deprecated: Use HelloworldMetaData.ABI instead.
var HelloworldABI = HelloworldMetaData.ABI

// HelloworldBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HelloworldMetaData.Bin instead.
var HelloworldBin = HelloworldMetaData.Bin

// DeployHelloworld deploys a new Ethereum contract, binding an instance of Helloworld to it.
func DeployHelloworld(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Helloworld, error) {
	parsed, err := HelloworldMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HelloworldBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Helloworld{HelloworldCaller: HelloworldCaller{contract: contract}, HelloworldTransactor: HelloworldTransactor{contract: contract}, HelloworldFilterer: HelloworldFilterer{contract: contract}}, nil
}

// Helloworld is an auto generated Go binding around an Ethereum contract.
type Helloworld struct {
	HelloworldCaller     // Read-only binding to the contract
	HelloworldTransactor // Write-only binding to the contract
	HelloworldFilterer   // Log filterer for contract events
}

// HelloworldCaller is an auto generated read-only Go binding around an Ethereum contract.
type HelloworldCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloworldTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HelloworldTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloworldFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HelloworldFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloworldSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HelloworldSession struct {
	Contract     *Helloworld       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HelloworldCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HelloworldCallerSession struct {
	Contract *HelloworldCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// HelloworldTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HelloworldTransactorSession struct {
	Contract     *HelloworldTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// HelloworldRaw is an auto generated low-level Go binding around an Ethereum contract.
type HelloworldRaw struct {
	Contract *Helloworld // Generic contract binding to access the raw methods on
}

// HelloworldCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HelloworldCallerRaw struct {
	Contract *HelloworldCaller // Generic read-only contract binding to access the raw methods on
}

// HelloworldTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HelloworldTransactorRaw struct {
	Contract *HelloworldTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHelloworld creates a new instance of Helloworld, bound to a specific deployed contract.
func NewHelloworld(address common.Address, backend bind.ContractBackend) (*Helloworld, error) {
	contract, err := bindHelloworld(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Helloworld{HelloworldCaller: HelloworldCaller{contract: contract}, HelloworldTransactor: HelloworldTransactor{contract: contract}, HelloworldFilterer: HelloworldFilterer{contract: contract}}, nil
}

// NewHelloworldCaller creates a new read-only instance of Helloworld, bound to a specific deployed contract.
func NewHelloworldCaller(address common.Address, caller bind.ContractCaller) (*HelloworldCaller, error) {
	contract, err := bindHelloworld(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HelloworldCaller{contract: contract}, nil
}

// NewHelloworldTransactor creates a new write-only instance of Helloworld, bound to a specific deployed contract.
func NewHelloworldTransactor(address common.Address, transactor bind.ContractTransactor) (*HelloworldTransactor, error) {
	contract, err := bindHelloworld(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HelloworldTransactor{contract: contract}, nil
}

// NewHelloworldFilterer creates a new log filterer instance of Helloworld, bound to a specific deployed contract.
func NewHelloworldFilterer(address common.Address, filterer bind.ContractFilterer) (*HelloworldFilterer, error) {
	contract, err := bindHelloworld(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HelloworldFilterer{contract: contract}, nil
}

// bindHelloworld binds a generic wrapper to an already deployed contract.
func bindHelloworld(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HelloworldMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Helloworld *HelloworldRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Helloworld.Contract.HelloworldCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Helloworld *HelloworldRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Helloworld.Contract.HelloworldTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Helloworld *HelloworldRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Helloworld.Contract.HelloworldTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Helloworld *HelloworldCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Helloworld.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Helloworld *HelloworldTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Helloworld.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Helloworld *HelloworldTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Helloworld.Contract.contract.Transact(opts, method, params...)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8f8a912.
//
// Solidity: function GetBalance() view returns(uint256)
func (_Helloworld *HelloworldCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Helloworld.contract.Call(opts, &out, "GetBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xf8f8a912.
//
// Solidity: function GetBalance() view returns(uint256)
func (_Helloworld *HelloworldSession) GetBalance() (*big.Int, error) {
	return _Helloworld.Contract.GetBalance(&_Helloworld.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8f8a912.
//
// Solidity: function GetBalance() view returns(uint256)
func (_Helloworld *HelloworldCallerSession) GetBalance() (*big.Int, error) {
	return _Helloworld.Contract.GetBalance(&_Helloworld.CallOpts)
}

// GetData is a free data retrieval call binding the contract method 0x76b8e528.
//
// Solidity: function GetData() view returns(string)
func (_Helloworld *HelloworldCaller) GetData(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Helloworld.contract.Call(opts, &out, "GetData")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetData is a free data retrieval call binding the contract method 0x76b8e528.
//
// Solidity: function GetData() view returns(string)
func (_Helloworld *HelloworldSession) GetData() (string, error) {
	return _Helloworld.Contract.GetData(&_Helloworld.CallOpts)
}

// GetData is a free data retrieval call binding the contract method 0x76b8e528.
//
// Solidity: function GetData() view returns(string)
func (_Helloworld *HelloworldCallerSession) GetData() (string, error) {
	return _Helloworld.Contract.GetData(&_Helloworld.CallOpts)
}

// Greet is a free data retrieval call binding the contract method 0xefdeaaf5.
//
// Solidity: function Greet(string str) pure returns(string)
func (_Helloworld *HelloworldCaller) Greet(opts *bind.CallOpts, str string) (string, error) {
	var out []interface{}
	err := _Helloworld.contract.Call(opts, &out, "Greet", str)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Greet is a free data retrieval call binding the contract method 0xefdeaaf5.
//
// Solidity: function Greet(string str) pure returns(string)
func (_Helloworld *HelloworldSession) Greet(str string) (string, error) {
	return _Helloworld.Contract.Greet(&_Helloworld.CallOpts, str)
}

// Greet is a free data retrieval call binding the contract method 0xefdeaaf5.
//
// Solidity: function Greet(string str) pure returns(string)
func (_Helloworld *HelloworldCallerSession) Greet(str string) (string, error) {
	return _Helloworld.Contract.Greet(&_Helloworld.CallOpts, str)
}

// Hello is a free data retrieval call binding the contract method 0xbcdfe0d5.
//
// Solidity: function Hello() pure returns(string)
func (_Helloworld *HelloworldCaller) Hello(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Helloworld.contract.Call(opts, &out, "Hello")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Hello is a free data retrieval call binding the contract method 0xbcdfe0d5.
//
// Solidity: function Hello() pure returns(string)
func (_Helloworld *HelloworldSession) Hello() (string, error) {
	return _Helloworld.Contract.Hello(&_Helloworld.CallOpts)
}

// Hello is a free data retrieval call binding the contract method 0xbcdfe0d5.
//
// Solidity: function Hello() pure returns(string)
func (_Helloworld *HelloworldCallerSession) Hello() (string, error) {
	return _Helloworld.Contract.Hello(&_Helloworld.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() view returns(string)
func (_Helloworld *HelloworldCaller) Data(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Helloworld.contract.Call(opts, &out, "data")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() view returns(string)
func (_Helloworld *HelloworldSession) Data() (string, error) {
	return _Helloworld.Contract.Data(&_Helloworld.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() view returns(string)
func (_Helloworld *HelloworldCallerSession) Data() (string, error) {
	return _Helloworld.Contract.Data(&_Helloworld.CallOpts)
}

// SetData is a paid mutator transaction binding the contract method 0xb7e8260c.
//
// Solidity: function SetData(string str) returns()
func (_Helloworld *HelloworldTransactor) SetData(opts *bind.TransactOpts, str string) (*types.Transaction, error) {
	return _Helloworld.contract.Transact(opts, "SetData", str)
}

// SetData is a paid mutator transaction binding the contract method 0xb7e8260c.
//
// Solidity: function SetData(string str) returns()
func (_Helloworld *HelloworldSession) SetData(str string) (*types.Transaction, error) {
	return _Helloworld.Contract.SetData(&_Helloworld.TransactOpts, str)
}

// SetData is a paid mutator transaction binding the contract method 0xb7e8260c.
//
// Solidity: function SetData(string str) returns()
func (_Helloworld *HelloworldTransactorSession) SetData(str string) (*types.Transaction, error) {
	return _Helloworld.Contract.SetData(&_Helloworld.TransactOpts, str)
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() returns()
func (_Helloworld *HelloworldTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Helloworld.contract.Transact(opts, "Withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() returns()
func (_Helloworld *HelloworldSession) Withdraw() (*types.Transaction, error) {
	return _Helloworld.Contract.Withdraw(&_Helloworld.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() returns()
func (_Helloworld *HelloworldTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Helloworld.Contract.Withdraw(&_Helloworld.TransactOpts)
}
