// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package punchface

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

// PunchfaceMetaData contains all meta data concerning the Punchface contract.
var PunchfaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amt\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040526000805534801561001457600080fd5b50600180546001600160a01b0319163317905561002f610034565b610073565b34600080828254610045919061004c565b9091555050565b8082018082111561006d57634e487b7160e01b600052601160045260246000fd5b92915050565b6101e2806100826000396000f3fe6080604052600436106100435760003560e01c80630ef67887146100575780634d6ce1e51461007b5780635b6b431d1461009b578063f851a440146100bb57600080fd5b36610052576100506100f3565b005b600080fd5b34801561006357600080fd5b506000545b6040519081526020015b60405180910390f35b34801561008757600080fd5b50610068610096366004610151565b61010b565b3480156100a757600080fd5b506100506100b6366004610151565b610126565b3480156100c757600080fd5b506001546100db906001600160a01b031681565b6040516001600160a01b039091168152602001610072565b346000808282546101049190610180565b9091555050565b60008160005461011b9190610180565b600081905592915050565b6001546001600160a01b0316331461013d57600080fd5b8060005461014b9190610199565b60005550565b60006020828403121561016357600080fd5b5035919050565b634e487b7160e01b600052601160045260246000fd5b808201808211156101935761019361016a565b92915050565b818103818111156101935761019361016a56fea2646970667358221220266082b422878886c09b2636f381172b55224d2eccb6f0c1755d93f10d6a766564736f6c63430008110033",
}

// PunchfaceABI is the input ABI used to generate the binding from.
// Deprecated: Use PunchfaceMetaData.ABI instead.
var PunchfaceABI = PunchfaceMetaData.ABI

// PunchfaceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PunchfaceMetaData.Bin instead.
var PunchfaceBin = PunchfaceMetaData.Bin

// DeployPunchface deploys a new Ethereum contract, binding an instance of Punchface to it.
func DeployPunchface(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Punchface, error) {
	parsed, err := PunchfaceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PunchfaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Punchface{PunchfaceCaller: PunchfaceCaller{contract: contract}, PunchfaceTransactor: PunchfaceTransactor{contract: contract}, PunchfaceFilterer: PunchfaceFilterer{contract: contract}}, nil
}

// Punchface is an auto generated Go binding around an Ethereum contract.
type Punchface struct {
	PunchfaceCaller     // Read-only binding to the contract
	PunchfaceTransactor // Write-only binding to the contract
	PunchfaceFilterer   // Log filterer for contract events
}

// PunchfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type PunchfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PunchfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PunchfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PunchfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PunchfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PunchfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PunchfaceSession struct {
	Contract     *Punchface        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PunchfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PunchfaceCallerSession struct {
	Contract *PunchfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// PunchfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PunchfaceTransactorSession struct {
	Contract     *PunchfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PunchfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type PunchfaceRaw struct {
	Contract *Punchface // Generic contract binding to access the raw methods on
}

// PunchfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PunchfaceCallerRaw struct {
	Contract *PunchfaceCaller // Generic read-only contract binding to access the raw methods on
}

// PunchfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PunchfaceTransactorRaw struct {
	Contract *PunchfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPunchface creates a new instance of Punchface, bound to a specific deployed contract.
func NewPunchface(address common.Address, backend bind.ContractBackend) (*Punchface, error) {
	contract, err := bindPunchface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Punchface{PunchfaceCaller: PunchfaceCaller{contract: contract}, PunchfaceTransactor: PunchfaceTransactor{contract: contract}, PunchfaceFilterer: PunchfaceFilterer{contract: contract}}, nil
}

// NewPunchfaceCaller creates a new read-only instance of Punchface, bound to a specific deployed contract.
func NewPunchfaceCaller(address common.Address, caller bind.ContractCaller) (*PunchfaceCaller, error) {
	contract, err := bindPunchface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PunchfaceCaller{contract: contract}, nil
}

// NewPunchfaceTransactor creates a new write-only instance of Punchface, bound to a specific deployed contract.
func NewPunchfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*PunchfaceTransactor, error) {
	contract, err := bindPunchface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PunchfaceTransactor{contract: contract}, nil
}

// NewPunchfaceFilterer creates a new log filterer instance of Punchface, bound to a specific deployed contract.
func NewPunchfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*PunchfaceFilterer, error) {
	contract, err := bindPunchface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PunchfaceFilterer{contract: contract}, nil
}

// bindPunchface binds a generic wrapper to an already deployed contract.
func bindPunchface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PunchfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Punchface *PunchfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Punchface.Contract.PunchfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Punchface *PunchfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Punchface.Contract.PunchfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Punchface *PunchfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Punchface.Contract.PunchfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Punchface *PunchfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Punchface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Punchface *PunchfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Punchface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Punchface *PunchfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Punchface.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0x0ef67887.
//
// Solidity: function Balance() view returns(uint256)
func (_Punchface *PunchfaceCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Punchface.contract.Call(opts, &out, "Balance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0x0ef67887.
//
// Solidity: function Balance() view returns(uint256)
func (_Punchface *PunchfaceSession) Balance() (*big.Int, error) {
	return _Punchface.Contract.Balance(&_Punchface.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0x0ef67887.
//
// Solidity: function Balance() view returns(uint256)
func (_Punchface *PunchfaceCallerSession) Balance() (*big.Int, error) {
	return _Punchface.Contract.Balance(&_Punchface.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Punchface *PunchfaceCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Punchface.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Punchface *PunchfaceSession) Admin() (common.Address, error) {
	return _Punchface.Contract.Admin(&_Punchface.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Punchface *PunchfaceCallerSession) Admin() (common.Address, error) {
	return _Punchface.Contract.Admin(&_Punchface.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x4d6ce1e5.
//
// Solidity: function Deposit(uint256 amt) returns(uint256)
func (_Punchface *PunchfaceTransactor) Deposit(opts *bind.TransactOpts, amt *big.Int) (*types.Transaction, error) {
	return _Punchface.contract.Transact(opts, "Deposit", amt)
}

// Deposit is a paid mutator transaction binding the contract method 0x4d6ce1e5.
//
// Solidity: function Deposit(uint256 amt) returns(uint256)
func (_Punchface *PunchfaceSession) Deposit(amt *big.Int) (*types.Transaction, error) {
	return _Punchface.Contract.Deposit(&_Punchface.TransactOpts, amt)
}

// Deposit is a paid mutator transaction binding the contract method 0x4d6ce1e5.
//
// Solidity: function Deposit(uint256 amt) returns(uint256)
func (_Punchface *PunchfaceTransactorSession) Deposit(amt *big.Int) (*types.Transaction, error) {
	return _Punchface.Contract.Deposit(&_Punchface.TransactOpts, amt)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5b6b431d.
//
// Solidity: function Withdraw(uint256 _amt) returns()
func (_Punchface *PunchfaceTransactor) Withdraw(opts *bind.TransactOpts, _amt *big.Int) (*types.Transaction, error) {
	return _Punchface.contract.Transact(opts, "Withdraw", _amt)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5b6b431d.
//
// Solidity: function Withdraw(uint256 _amt) returns()
func (_Punchface *PunchfaceSession) Withdraw(_amt *big.Int) (*types.Transaction, error) {
	return _Punchface.Contract.Withdraw(&_Punchface.TransactOpts, _amt)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5b6b431d.
//
// Solidity: function Withdraw(uint256 _amt) returns()
func (_Punchface *PunchfaceTransactorSession) Withdraw(_amt *big.Int) (*types.Transaction, error) {
	return _Punchface.Contract.Withdraw(&_Punchface.TransactOpts, _amt)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Punchface *PunchfaceTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Punchface.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Punchface *PunchfaceSession) Receive() (*types.Transaction, error) {
	return _Punchface.Contract.Receive(&_Punchface.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Punchface *PunchfaceTransactorSession) Receive() (*types.Transaction, error) {
	return _Punchface.Contract.Receive(&_Punchface.TransactOpts)
}
