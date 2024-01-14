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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GetBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetData\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"Greet\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Hello\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"SetData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5061060e8061001d5f395ff3fe608060405260043610610071575f3560e01c8063b7e8260c1161004c578063b7e8260c14610108578063bcdfe0d514610127578063efdeaaf51461015a578063f8f8a91214610179575f80fd5b806357ea89b6146100b457806373d4a13a146100ca57806376b8e528146100f4575f80fd5b366100b057604080513381523460208201527f88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874910160405180910390a1005b5f80fd5b3480156100bf575f80fd5b506100c8610193565b005b3480156100d5575f80fd5b506100de6101bf565b6040516100eb9190610389565b60405180910390f35b3480156100ff575f80fd5b506100de61024a565b348015610113575f80fd5b506100c86101223660046103e9565b6102d9565b348015610132575f80fd5b5060408051808201909152600b81526a12195b1b1bc815dbdc9b1960aa1b60208201526100de565b348015610165575f80fd5b506100de6101743660046103e9565b610338565b348015610184575f80fd5b506040514781526020016100eb565b60405133904780156108fc02915f818181858888f193505050501580156101bc573d5f803e3d5ffd5b50565b5f80546101cb90610494565b80601f01602080910402602001604051908101604052809291908181526020018280546101f790610494565b80156102425780601f1061021957610100808354040283529160200191610242565b820191905f5260205f20905b81548152906001019060200180831161022557829003601f168201915b505050505081565b60605f805461025890610494565b80601f016020809104026020016040519081016040528092919081815260200182805461028490610494565b80156102cf5780601f106102a6576101008083540402835291602001916102cf565b820191905f5260205f20905b8154815290600101906020018083116102b257829003601f168201915b5050505050905090565b5f8151116103295760405162461bcd60e51b8152602060048201526018602482015277737472696e67206d757374206e6f7420626520656d70747960401b60448201526064015b60405180910390fd5b5f6103348282610518565b5050565b60605f8251116103855760405162461bcd60e51b8152602060048201526018602482015277737472696e67206d757374206e6f7420626520656d70747960401b6044820152606401610320565b5090565b5f602080835283518060208501525f5b818110156103b557858101830151858201604001528201610399565b505f604082860101526040601f19601f8301168501019250505092915050565b634e487b7160e01b5f52604160045260245ffd5b5f602082840312156103f9575f80fd5b813567ffffffffffffffff80821115610410575f80fd5b818401915084601f830112610423575f80fd5b813581811115610435576104356103d5565b604051601f8201601f19908116603f0116810190838211818310171561045d5761045d6103d5565b81604052828152876020848701011115610475575f80fd5b826020860160208301375f928101602001929092525095945050505050565b600181811c908216806104a857607f821691505b6020821081036104c657634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561051357805f5260205f20601f840160051c810160208510156104f15750805b601f840160051c820191505b81811015610510575f81556001016104fd565b50505b505050565b815167ffffffffffffffff811115610532576105326103d5565b610546816105408454610494565b846104cc565b602080601f831160018114610579575f84156105625750858301515b5f19600386901b1c1916600185901b1785556105d0565b5f85815260208120601f198616915b828110156105a757888601518255948401946001909101908401610588565b50858210156105c457878501515f19600388901b60f8161c191681555b505060018460011b0185555b50505050505056fea2646970667358221220562df10e020e80262899256f3382c5ab5d46ed4157aaf312d0e51cdd1d102ddf64736f6c63430008170033",
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

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Helloworld *HelloworldTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Helloworld.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Helloworld *HelloworldSession) Receive() (*types.Transaction, error) {
	return _Helloworld.Contract.Receive(&_Helloworld.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Helloworld *HelloworldTransactorSession) Receive() (*types.Transaction, error) {
	return _Helloworld.Contract.Receive(&_Helloworld.TransactOpts)
}

// HelloworldReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the Helloworld contract.
type HelloworldReceivedIterator struct {
	Event *HelloworldReceived // Event containing the contract specifics and raw log

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
func (it *HelloworldReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HelloworldReceived)
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
		it.Event = new(HelloworldReceived)
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
func (it *HelloworldReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HelloworldReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HelloworldReceived represents a Received event raised by the Helloworld contract.
type HelloworldReceived struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address sender, uint256 amount)
func (_Helloworld *HelloworldFilterer) FilterReceived(opts *bind.FilterOpts) (*HelloworldReceivedIterator, error) {

	logs, sub, err := _Helloworld.contract.FilterLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return &HelloworldReceivedIterator{contract: _Helloworld.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address sender, uint256 amount)
func (_Helloworld *HelloworldFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *HelloworldReceived) (event.Subscription, error) {

	logs, sub, err := _Helloworld.contract.WatchLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HelloworldReceived)
				if err := _Helloworld.contract.UnpackLog(event, "Received", log); err != nil {
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

// ParseReceived is a log parse operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address sender, uint256 amount)
func (_Helloworld *HelloworldFilterer) ParseReceived(log types.Log) (*HelloworldReceived, error) {
	event := new(HelloworldReceived)
	if err := _Helloworld.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
