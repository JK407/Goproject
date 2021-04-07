// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pkgname

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ApinameABI is the input ABI used to generate the binding from.
const ApinameABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"a\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"b\",\"type\":\"string\"}],\"name\":\"equal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"passsword\",\"type\":\"string\"}],\"name\":\"login\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"passsword\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Apiname is an auto generated Go binding around an Ethereum contract.
type Apiname struct {
	ApinameCaller     // Read-only binding to the contract
	ApinameTransactor // Write-only binding to the contract
	ApinameFilterer   // Log filterer for contract events
}

// ApinameCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApinameCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApinameTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApinameTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApinameFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApinameFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApinameSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApinameSession struct {
	Contract     *Apiname          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApinameCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApinameCallerSession struct {
	Contract *ApinameCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ApinameTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApinameTransactorSession struct {
	Contract     *ApinameTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ApinameRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApinameRaw struct {
	Contract *Apiname // Generic contract binding to access the raw methods on
}

// ApinameCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApinameCallerRaw struct {
	Contract *ApinameCaller // Generic read-only contract binding to access the raw methods on
}

// ApinameTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApinameTransactorRaw struct {
	Contract *ApinameTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApiname creates a new instance of Apiname, bound to a specific deployed contract.
func NewApiname(address common.Address, backend bind.ContractBackend) (*Apiname, error) {
	contract, err := bindApiname(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Apiname{ApinameCaller: ApinameCaller{contract: contract}, ApinameTransactor: ApinameTransactor{contract: contract}, ApinameFilterer: ApinameFilterer{contract: contract}}, nil
}

// NewApinameCaller creates a new read-only instance of Apiname, bound to a specific deployed contract.
func NewApinameCaller(address common.Address, caller bind.ContractCaller) (*ApinameCaller, error) {
	contract, err := bindApiname(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApinameCaller{contract: contract}, nil
}

// NewApinameTransactor creates a new write-only instance of Apiname, bound to a specific deployed contract.
func NewApinameTransactor(address common.Address, transactor bind.ContractTransactor) (*ApinameTransactor, error) {
	contract, err := bindApiname(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApinameTransactor{contract: contract}, nil
}

// NewApinameFilterer creates a new log filterer instance of Apiname, bound to a specific deployed contract.
func NewApinameFilterer(address common.Address, filterer bind.ContractFilterer) (*ApinameFilterer, error) {
	contract, err := bindApiname(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApinameFilterer{contract: contract}, nil
}

// bindApiname binds a generic wrapper to an already deployed contract.
func bindApiname(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ApinameABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Apiname *ApinameRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Apiname.Contract.ApinameCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Apiname *ApinameRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Apiname.Contract.ApinameTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Apiname *ApinameRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Apiname.Contract.ApinameTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Apiname *ApinameCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Apiname.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Apiname *ApinameTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Apiname.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Apiname *ApinameTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Apiname.Contract.contract.Transact(opts, method, params...)
}

// Equal is a paid mutator transaction binding the contract method 0x46bdca9a.
//
// Solidity: function equal(string a, string b) returns(bool)
func (_Apiname *ApinameTransactor) Equal(opts *bind.TransactOpts, a string, b string) (*types.Transaction, error) {
	return _Apiname.contract.Transact(opts, "equal", a, b)
}

// Equal is a paid mutator transaction binding the contract method 0x46bdca9a.
//
// Solidity: function equal(string a, string b) returns(bool)
func (_Apiname *ApinameSession) Equal(a string, b string) (*types.Transaction, error) {
	return _Apiname.Contract.Equal(&_Apiname.TransactOpts, a, b)
}

// Equal is a paid mutator transaction binding the contract method 0x46bdca9a.
//
// Solidity: function equal(string a, string b) returns(bool)
func (_Apiname *ApinameTransactorSession) Equal(a string, b string) (*types.Transaction, error) {
	return _Apiname.Contract.Equal(&_Apiname.TransactOpts, a, b)
}

// Login is a paid mutator transaction binding the contract method 0x2c4f5a17.
//
// Solidity: function login(uint256 id, string passsword) returns(bool)
func (_Apiname *ApinameTransactor) Login(opts *bind.TransactOpts, id *big.Int, passsword string) (*types.Transaction, error) {
	return _Apiname.contract.Transact(opts, "login", id, passsword)
}

// Login is a paid mutator transaction binding the contract method 0x2c4f5a17.
//
// Solidity: function login(uint256 id, string passsword) returns(bool)
func (_Apiname *ApinameSession) Login(id *big.Int, passsword string) (*types.Transaction, error) {
	return _Apiname.Contract.Login(&_Apiname.TransactOpts, id, passsword)
}

// Login is a paid mutator transaction binding the contract method 0x2c4f5a17.
//
// Solidity: function login(uint256 id, string passsword) returns(bool)
func (_Apiname *ApinameTransactorSession) Login(id *big.Int, passsword string) (*types.Transaction, error) {
	return _Apiname.Contract.Login(&_Apiname.TransactOpts, id, passsword)
}

// Register is a paid mutator transaction binding the contract method 0xa00fd3c8.
//
// Solidity: function register(uint256 id, string passsword) returns(bool)
func (_Apiname *ApinameTransactor) Register(opts *bind.TransactOpts, id *big.Int, passsword string) (*types.Transaction, error) {
	return _Apiname.contract.Transact(opts, "register", id, passsword)
}

// Register is a paid mutator transaction binding the contract method 0xa00fd3c8.
//
// Solidity: function register(uint256 id, string passsword) returns(bool)
func (_Apiname *ApinameSession) Register(id *big.Int, passsword string) (*types.Transaction, error) {
	return _Apiname.Contract.Register(&_Apiname.TransactOpts, id, passsword)
}

// Register is a paid mutator transaction binding the contract method 0xa00fd3c8.
//
// Solidity: function register(uint256 id, string passsword) returns(bool)
func (_Apiname *ApinameTransactorSession) Register(id *big.Int, passsword string) (*types.Transaction, error) {
	return _Apiname.Contract.Register(&_Apiname.TransactOpts, id, passsword)
}
