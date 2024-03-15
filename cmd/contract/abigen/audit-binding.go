// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package audit

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

// AuditMetaData contains all meta data concerning the Audit contract.
var AuditMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"retVal\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPKS\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPKT\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getParamsGU\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSKH\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"keyArr\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"storedPKT\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"storedSKH\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"keys\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"storedPKT\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"storedSKH\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queryAudit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"sendAudit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"input2\",\"type\":\"bytes\"}],\"name\":\"setOwnerKeys\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"setPKO\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"setPKS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"input2\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"input3\",\"type\":\"bytes\"}],\"name\":\"setParamsGU\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"input\",\"type\":\"string\"}],\"name\":\"setSign\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506001805460ff19169055610cad806100275f395ff3fe608060405234801561000f575f80fd5b50600436106100f0575f3560e01c80634ddffb4b116100935780636d4ce63c116100635780636d4ce63c146101d057806377c4e362146101e05780638ce47f78146101f3578063cb426cb114610206575f80fd5b80634ddffb4b1461018557806350a6cab31461019857806356d0022f146101ab57806360fe47b1146101be575f80fd5b8063307540f6116100ce578063307540f61461013057806331fbbae2146101465780633e0610341461015b5780633f47283e14610172575f80fd5b8063076ab3e5146100f45780631027d316146101125780631afbca031461011a575b5f80fd5b6100fc61020e565b604051610109919061078a565b60405180910390f35b6100fc61029e565b60015460ff166040519015158152602001610109565b6101386102b0565b6040516101099291906107a3565b610159610154366004610857565b6103cc565b005b6101636103dc565b604051610109939291906108a4565b610159610180366004610904565b610590565b610138610193366004610986565b6105bc565b6101596101a636600461099d565b6105ea565b6101596101b93660046109fd565b61067f565b6101596101cc366004610986565b5f55565b5f54604051908152602001610109565b6101596101ee3660046109fd565b6106f7565b6101596102013660046109fd565b610703565b6100fc61070f565b60606006805461021d90610a2f565b80601f016020809104026020016040519081016040528092919081815260200182805461024990610a2f565b80156102945780601f1061026b57610100808354040283529160200191610294565b820191905f5260205f20905b81548152906001019060200180831161027757829003601f168201915b5050505050905090565b60606008600101805461021d90610a2f565b6008805481906102bf90610a2f565b80601f01602080910402602001604051908101604052809291908181526020018280546102eb90610a2f565b80156103365780601f1061030d57610100808354040283529160200191610336565b820191905f5260205f20905b81548152906001019060200180831161031957829003601f168201915b50505050509080600101805461034b90610a2f565b80601f016020809104026020016040519081016040528092919081815260200182805461037790610a2f565b80156103c25780601f10610399576101008083540402835291602001916103c2565b820191905f5260205f20905b8154815290600101906020018083116103a557829003601f168201915b5050505050905082565b60076103d88282610ab5565b5050565b60608060606002600360048280546103f390610a2f565b80601f016020809104026020016040519081016040528092919081815260200182805461041f90610a2f565b801561046a5780601f106104415761010080835404028352916020019161046a565b820191905f5260205f20905b81548152906001019060200180831161044d57829003601f168201915b5050505050925081805461047d90610a2f565b80601f01602080910402602001604051908101604052809291908181526020018280546104a990610a2f565b80156104f45780601f106104cb576101008083540402835291602001916104f4565b820191905f5260205f20905b8154815290600101906020018083116104d757829003601f168201915b5050505050915080805461050790610a2f565b80601f016020809104026020016040519081016040528092919081815260200182805461053390610a2f565b801561057e5780601f106105555761010080835404028352916020019161057e565b820191905f5260205f20905b81548152906001019060200180831161056157829003601f168201915b50505050509050925092509250909192565b600261059c8482610ab5565b5060036105a98382610ab5565b5060046105b68282610ab5565b50505050565b600a81815481106105cb575f80fd5b905f5260205f2090600202015f91509050805f0180546102bf90610a2f565b604080518082019091528281526020810182905260088061060b8582610ab5565b50602082015160018201906106209082610ab5565b5050600a80546001810182555f91909152600891506002027fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a801806106658382610b71565b5060018181019061067890840182610b71565b5050505050565b610687610720565b5f80835160206106979190610c3e565b905060018382865f60095f19f115915081156106b5576106b5610c63565b825180516020909101207f37610255ab3f0df385209ed77d20f6af0a56ae9c81fcf83234b398d0d674743a016105b6576001805460ff19168117905550505050565b60066103d88282610ab5565b60056103d88282610ab5565b606060085f01805461021d90610a2f565b60405180602001604052806001905b606081526020019060019003908161072f5790505090565b5f81518084525f5b8181101561076b5760208185018101518683018201520161074f565b505f602082860101526020601f19601f83011685010191505092915050565b602081525f61079c6020830184610747565b9392505050565b604081525f6107b56040830185610747565b82810360208401526107c78185610747565b95945050505050565b634e487b7160e01b5f52604160045260245ffd5b5f67ffffffffffffffff808411156107fe576107fe6107d0565b604051601f8501601f19908116603f01168101908282118183101715610826576108266107d0565b8160405280935085815286868601111561083e575f80fd5b858560208301375f602087830101525050509392505050565b5f60208284031215610867575f80fd5b813567ffffffffffffffff81111561087d575f80fd5b8201601f8101841361088d575f80fd5b61089c848235602084016107e4565b949350505050565b606081525f6108b66060830186610747565b82810360208401526108c88186610747565b905082810360408401526108dc8185610747565b9695505050505050565b5f82601f8301126108f5575f80fd5b61079c838335602085016107e4565b5f805f60608486031215610916575f80fd5b833567ffffffffffffffff8082111561092d575f80fd5b610939878388016108e6565b9450602086013591508082111561094e575f80fd5b61095a878388016108e6565b9350604086013591508082111561096f575f80fd5b5061097c868287016108e6565b9150509250925092565b5f60208284031215610996575f80fd5b5035919050565b5f80604083850312156109ae575f80fd5b823567ffffffffffffffff808211156109c5575f80fd5b6109d1868387016108e6565b935060208501359150808211156109e6575f80fd5b506109f3858286016108e6565b9150509250929050565b5f60208284031215610a0d575f80fd5b813567ffffffffffffffff811115610a23575f80fd5b61089c848285016108e6565b600181811c90821680610a4357607f821691505b602082108103610a6157634e487b7160e01b5f52602260045260245ffd5b50919050565b601f821115610ab0575f81815260208120601f850160051c81016020861015610a8d5750805b601f850160051c820191505b81811015610aac57828155600101610a99565b5050505b505050565b815167ffffffffffffffff811115610acf57610acf6107d0565b610ae381610add8454610a2f565b84610a67565b602080601f831160018114610b16575f8415610aff5750858301515b5f19600386901b1c1916600185901b178555610aac565b5f85815260208120601f198616915b82811015610b4457888601518255948401946001909101908401610b25565b5085821015610b6157878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b818103610b7c575050565b610b868254610a2f565b67ffffffffffffffff811115610b9e57610b9e6107d0565b610bac81610add8454610a2f565b5f601f821160018114610bdd575f8315610bc65750848201545b5f19600385901b1c1916600184901b178455610678565b5f85815260209020601f198416905f86815260209020845b83811015610c155782860154825560019586019590910190602001610bf5565b5085831015610b61579301545f1960f8600387901b161c19169092555050600190811b01905550565b80820180821115610c5d57634e487b7160e01b5f52601160045260245ffd5b92915050565b634e487b7160e01b5f52600160045260245ffdfea264697066735822122002a7b07d8827579ca80d5960d159adf9965a5eeb72ffd79c1851a60e3653fb9a64736f6c63430008140033",
}

// AuditABI is the input ABI used to generate the binding from.
// Deprecated: Use AuditMetaData.ABI instead.
var AuditABI = AuditMetaData.ABI

// AuditBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AuditMetaData.Bin instead.
var AuditBin = AuditMetaData.Bin

// DeployAudit deploys a new Ethereum contract, binding an instance of Audit to it.
func DeployAudit(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Audit, error) {
	parsed, err := AuditMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AuditBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Audit{AuditCaller: AuditCaller{contract: contract}, AuditTransactor: AuditTransactor{contract: contract}, AuditFilterer: AuditFilterer{contract: contract}}, nil
}

// Audit is an auto generated Go binding around an Ethereum contract.
type Audit struct {
	AuditCaller     // Read-only binding to the contract
	AuditTransactor // Write-only binding to the contract
	AuditFilterer   // Log filterer for contract events
}

// AuditCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuditCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuditTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuditTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuditFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuditFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuditSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuditSession struct {
	Contract     *Audit            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuditCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuditCallerSession struct {
	Contract *AuditCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AuditTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuditTransactorSession struct {
	Contract     *AuditTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuditRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuditRaw struct {
	Contract *Audit // Generic contract binding to access the raw methods on
}

// AuditCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuditCallerRaw struct {
	Contract *AuditCaller // Generic read-only contract binding to access the raw methods on
}

// AuditTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuditTransactorRaw struct {
	Contract *AuditTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAudit creates a new instance of Audit, bound to a specific deployed contract.
func NewAudit(address common.Address, backend bind.ContractBackend) (*Audit, error) {
	contract, err := bindAudit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Audit{AuditCaller: AuditCaller{contract: contract}, AuditTransactor: AuditTransactor{contract: contract}, AuditFilterer: AuditFilterer{contract: contract}}, nil
}

// NewAuditCaller creates a new read-only instance of Audit, bound to a specific deployed contract.
func NewAuditCaller(address common.Address, caller bind.ContractCaller) (*AuditCaller, error) {
	contract, err := bindAudit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuditCaller{contract: contract}, nil
}

// NewAuditTransactor creates a new write-only instance of Audit, bound to a specific deployed contract.
func NewAuditTransactor(address common.Address, transactor bind.ContractTransactor) (*AuditTransactor, error) {
	contract, err := bindAudit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuditTransactor{contract: contract}, nil
}

// NewAuditFilterer creates a new log filterer instance of Audit, bound to a specific deployed contract.
func NewAuditFilterer(address common.Address, filterer bind.ContractFilterer) (*AuditFilterer, error) {
	contract, err := bindAudit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuditFilterer{contract: contract}, nil
}

// bindAudit binds a generic wrapper to an already deployed contract.
func bindAudit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AuditMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Audit *AuditRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Audit.Contract.AuditCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Audit *AuditRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Audit.Contract.AuditTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Audit *AuditRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Audit.Contract.AuditTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Audit *AuditCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Audit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Audit *AuditTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Audit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Audit *AuditTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Audit.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint256 retVal)
func (_Audit *AuditCaller) Get(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Audit.contract.Call(opts, &out, "get")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint256 retVal)
func (_Audit *AuditSession) Get() (*big.Int, error) {
	return _Audit.Contract.Get(&_Audit.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint256 retVal)
func (_Audit *AuditCallerSession) Get() (*big.Int, error) {
	return _Audit.Contract.Get(&_Audit.CallOpts)
}

// GetPKS is a free data retrieval call binding the contract method 0x076ab3e5.
//
// Solidity: function getPKS() view returns(bytes)
func (_Audit *AuditCaller) GetPKS(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Audit.contract.Call(opts, &out, "getPKS")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetPKS is a free data retrieval call binding the contract method 0x076ab3e5.
//
// Solidity: function getPKS() view returns(bytes)
func (_Audit *AuditSession) GetPKS() ([]byte, error) {
	return _Audit.Contract.GetPKS(&_Audit.CallOpts)
}

// GetPKS is a free data retrieval call binding the contract method 0x076ab3e5.
//
// Solidity: function getPKS() view returns(bytes)
func (_Audit *AuditCallerSession) GetPKS() ([]byte, error) {
	return _Audit.Contract.GetPKS(&_Audit.CallOpts)
}

// GetPKT is a free data retrieval call binding the contract method 0xcb426cb1.
//
// Solidity: function getPKT() view returns(bytes)
func (_Audit *AuditCaller) GetPKT(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Audit.contract.Call(opts, &out, "getPKT")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetPKT is a free data retrieval call binding the contract method 0xcb426cb1.
//
// Solidity: function getPKT() view returns(bytes)
func (_Audit *AuditSession) GetPKT() ([]byte, error) {
	return _Audit.Contract.GetPKT(&_Audit.CallOpts)
}

// GetPKT is a free data retrieval call binding the contract method 0xcb426cb1.
//
// Solidity: function getPKT() view returns(bytes)
func (_Audit *AuditCallerSession) GetPKT() ([]byte, error) {
	return _Audit.Contract.GetPKT(&_Audit.CallOpts)
}

// GetParamsGU is a free data retrieval call binding the contract method 0x3e061034.
//
// Solidity: function getParamsGU() view returns(bytes, bytes, bytes)
func (_Audit *AuditCaller) GetParamsGU(opts *bind.CallOpts) ([]byte, []byte, []byte, error) {
	var out []interface{}
	err := _Audit.contract.Call(opts, &out, "getParamsGU")

	if err != nil {
		return *new([]byte), *new([]byte), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	out2 := *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return out0, out1, out2, err

}

// GetParamsGU is a free data retrieval call binding the contract method 0x3e061034.
//
// Solidity: function getParamsGU() view returns(bytes, bytes, bytes)
func (_Audit *AuditSession) GetParamsGU() ([]byte, []byte, []byte, error) {
	return _Audit.Contract.GetParamsGU(&_Audit.CallOpts)
}

// GetParamsGU is a free data retrieval call binding the contract method 0x3e061034.
//
// Solidity: function getParamsGU() view returns(bytes, bytes, bytes)
func (_Audit *AuditCallerSession) GetParamsGU() ([]byte, []byte, []byte, error) {
	return _Audit.Contract.GetParamsGU(&_Audit.CallOpts)
}

// GetSKH is a free data retrieval call binding the contract method 0x1027d316.
//
// Solidity: function getSKH() view returns(bytes)
func (_Audit *AuditCaller) GetSKH(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Audit.contract.Call(opts, &out, "getSKH")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetSKH is a free data retrieval call binding the contract method 0x1027d316.
//
// Solidity: function getSKH() view returns(bytes)
func (_Audit *AuditSession) GetSKH() ([]byte, error) {
	return _Audit.Contract.GetSKH(&_Audit.CallOpts)
}

// GetSKH is a free data retrieval call binding the contract method 0x1027d316.
//
// Solidity: function getSKH() view returns(bytes)
func (_Audit *AuditCallerSession) GetSKH() ([]byte, error) {
	return _Audit.Contract.GetSKH(&_Audit.CallOpts)
}

// KeyArr is a free data retrieval call binding the contract method 0x4ddffb4b.
//
// Solidity: function keyArr(uint256 ) view returns(bytes storedPKT, bytes storedSKH)
func (_Audit *AuditCaller) KeyArr(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StoredPKT []byte
	StoredSKH []byte
}, error) {
	var out []interface{}
	err := _Audit.contract.Call(opts, &out, "keyArr", arg0)

	outstruct := new(struct {
		StoredPKT []byte
		StoredSKH []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StoredPKT = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.StoredSKH = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// KeyArr is a free data retrieval call binding the contract method 0x4ddffb4b.
//
// Solidity: function keyArr(uint256 ) view returns(bytes storedPKT, bytes storedSKH)
func (_Audit *AuditSession) KeyArr(arg0 *big.Int) (struct {
	StoredPKT []byte
	StoredSKH []byte
}, error) {
	return _Audit.Contract.KeyArr(&_Audit.CallOpts, arg0)
}

// KeyArr is a free data retrieval call binding the contract method 0x4ddffb4b.
//
// Solidity: function keyArr(uint256 ) view returns(bytes storedPKT, bytes storedSKH)
func (_Audit *AuditCallerSession) KeyArr(arg0 *big.Int) (struct {
	StoredPKT []byte
	StoredSKH []byte
}, error) {
	return _Audit.Contract.KeyArr(&_Audit.CallOpts, arg0)
}

// Keys is a free data retrieval call binding the contract method 0x307540f6.
//
// Solidity: function keys() view returns(bytes storedPKT, bytes storedSKH)
func (_Audit *AuditCaller) Keys(opts *bind.CallOpts) (struct {
	StoredPKT []byte
	StoredSKH []byte
}, error) {
	var out []interface{}
	err := _Audit.contract.Call(opts, &out, "keys")

	outstruct := new(struct {
		StoredPKT []byte
		StoredSKH []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StoredPKT = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.StoredSKH = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// Keys is a free data retrieval call binding the contract method 0x307540f6.
//
// Solidity: function keys() view returns(bytes storedPKT, bytes storedSKH)
func (_Audit *AuditSession) Keys() (struct {
	StoredPKT []byte
	StoredSKH []byte
}, error) {
	return _Audit.Contract.Keys(&_Audit.CallOpts)
}

// Keys is a free data retrieval call binding the contract method 0x307540f6.
//
// Solidity: function keys() view returns(bytes storedPKT, bytes storedSKH)
func (_Audit *AuditCallerSession) Keys() (struct {
	StoredPKT []byte
	StoredSKH []byte
}, error) {
	return _Audit.Contract.Keys(&_Audit.CallOpts)
}

// QueryAudit is a free data retrieval call binding the contract method 0x1afbca03.
//
// Solidity: function queryAudit() view returns(bool)
func (_Audit *AuditCaller) QueryAudit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Audit.contract.Call(opts, &out, "queryAudit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// QueryAudit is a free data retrieval call binding the contract method 0x1afbca03.
//
// Solidity: function queryAudit() view returns(bool)
func (_Audit *AuditSession) QueryAudit() (bool, error) {
	return _Audit.Contract.QueryAudit(&_Audit.CallOpts)
}

// QueryAudit is a free data retrieval call binding the contract method 0x1afbca03.
//
// Solidity: function queryAudit() view returns(bool)
func (_Audit *AuditCallerSession) QueryAudit() (bool, error) {
	return _Audit.Contract.QueryAudit(&_Audit.CallOpts)
}

// SendAudit is a paid mutator transaction binding the contract method 0x56d0022f.
//
// Solidity: function sendAudit(bytes input) returns()
func (_Audit *AuditTransactor) SendAudit(opts *bind.TransactOpts, input []byte) (*types.Transaction, error) {
	return _Audit.contract.Transact(opts, "sendAudit", input)
}

// SendAudit is a paid mutator transaction binding the contract method 0x56d0022f.
//
// Solidity: function sendAudit(bytes input) returns()
func (_Audit *AuditSession) SendAudit(input []byte) (*types.Transaction, error) {
	return _Audit.Contract.SendAudit(&_Audit.TransactOpts, input)
}

// SendAudit is a paid mutator transaction binding the contract method 0x56d0022f.
//
// Solidity: function sendAudit(bytes input) returns()
func (_Audit *AuditTransactorSession) SendAudit(input []byte) (*types.Transaction, error) {
	return _Audit.Contract.SendAudit(&_Audit.TransactOpts, input)
}

// Set is a paid mutator transaction binding the contract method 0x60fe47b1.
//
// Solidity: function set(uint256 x) returns()
func (_Audit *AuditTransactor) Set(opts *bind.TransactOpts, x *big.Int) (*types.Transaction, error) {
	return _Audit.contract.Transact(opts, "set", x)
}

// Set is a paid mutator transaction binding the contract method 0x60fe47b1.
//
// Solidity: function set(uint256 x) returns()
func (_Audit *AuditSession) Set(x *big.Int) (*types.Transaction, error) {
	return _Audit.Contract.Set(&_Audit.TransactOpts, x)
}

// Set is a paid mutator transaction binding the contract method 0x60fe47b1.
//
// Solidity: function set(uint256 x) returns()
func (_Audit *AuditTransactorSession) Set(x *big.Int) (*types.Transaction, error) {
	return _Audit.Contract.Set(&_Audit.TransactOpts, x)
}

// SetOwnerKeys is a paid mutator transaction binding the contract method 0x50a6cab3.
//
// Solidity: function setOwnerKeys(bytes input1, bytes input2) returns()
func (_Audit *AuditTransactor) SetOwnerKeys(opts *bind.TransactOpts, input1 []byte, input2 []byte) (*types.Transaction, error) {
	return _Audit.contract.Transact(opts, "setOwnerKeys", input1, input2)
}

// SetOwnerKeys is a paid mutator transaction binding the contract method 0x50a6cab3.
//
// Solidity: function setOwnerKeys(bytes input1, bytes input2) returns()
func (_Audit *AuditSession) SetOwnerKeys(input1 []byte, input2 []byte) (*types.Transaction, error) {
	return _Audit.Contract.SetOwnerKeys(&_Audit.TransactOpts, input1, input2)
}

// SetOwnerKeys is a paid mutator transaction binding the contract method 0x50a6cab3.
//
// Solidity: function setOwnerKeys(bytes input1, bytes input2) returns()
func (_Audit *AuditTransactorSession) SetOwnerKeys(input1 []byte, input2 []byte) (*types.Transaction, error) {
	return _Audit.Contract.SetOwnerKeys(&_Audit.TransactOpts, input1, input2)
}

// SetPKO is a paid mutator transaction binding the contract method 0x8ce47f78.
//
// Solidity: function setPKO(bytes input) returns()
func (_Audit *AuditTransactor) SetPKO(opts *bind.TransactOpts, input []byte) (*types.Transaction, error) {
	return _Audit.contract.Transact(opts, "setPKO", input)
}

// SetPKO is a paid mutator transaction binding the contract method 0x8ce47f78.
//
// Solidity: function setPKO(bytes input) returns()
func (_Audit *AuditSession) SetPKO(input []byte) (*types.Transaction, error) {
	return _Audit.Contract.SetPKO(&_Audit.TransactOpts, input)
}

// SetPKO is a paid mutator transaction binding the contract method 0x8ce47f78.
//
// Solidity: function setPKO(bytes input) returns()
func (_Audit *AuditTransactorSession) SetPKO(input []byte) (*types.Transaction, error) {
	return _Audit.Contract.SetPKO(&_Audit.TransactOpts, input)
}

// SetPKS is a paid mutator transaction binding the contract method 0x77c4e362.
//
// Solidity: function setPKS(bytes input) returns()
func (_Audit *AuditTransactor) SetPKS(opts *bind.TransactOpts, input []byte) (*types.Transaction, error) {
	return _Audit.contract.Transact(opts, "setPKS", input)
}

// SetPKS is a paid mutator transaction binding the contract method 0x77c4e362.
//
// Solidity: function setPKS(bytes input) returns()
func (_Audit *AuditSession) SetPKS(input []byte) (*types.Transaction, error) {
	return _Audit.Contract.SetPKS(&_Audit.TransactOpts, input)
}

// SetPKS is a paid mutator transaction binding the contract method 0x77c4e362.
//
// Solidity: function setPKS(bytes input) returns()
func (_Audit *AuditTransactorSession) SetPKS(input []byte) (*types.Transaction, error) {
	return _Audit.Contract.SetPKS(&_Audit.TransactOpts, input)
}

// SetParamsGU is a paid mutator transaction binding the contract method 0x3f47283e.
//
// Solidity: function setParamsGU(bytes input1, bytes input2, bytes input3) returns()
func (_Audit *AuditTransactor) SetParamsGU(opts *bind.TransactOpts, input1 []byte, input2 []byte, input3 []byte) (*types.Transaction, error) {
	return _Audit.contract.Transact(opts, "setParamsGU", input1, input2, input3)
}

// SetParamsGU is a paid mutator transaction binding the contract method 0x3f47283e.
//
// Solidity: function setParamsGU(bytes input1, bytes input2, bytes input3) returns()
func (_Audit *AuditSession) SetParamsGU(input1 []byte, input2 []byte, input3 []byte) (*types.Transaction, error) {
	return _Audit.Contract.SetParamsGU(&_Audit.TransactOpts, input1, input2, input3)
}

// SetParamsGU is a paid mutator transaction binding the contract method 0x3f47283e.
//
// Solidity: function setParamsGU(bytes input1, bytes input2, bytes input3) returns()
func (_Audit *AuditTransactorSession) SetParamsGU(input1 []byte, input2 []byte, input3 []byte) (*types.Transaction, error) {
	return _Audit.Contract.SetParamsGU(&_Audit.TransactOpts, input1, input2, input3)
}

// SetSign is a paid mutator transaction binding the contract method 0x31fbbae2.
//
// Solidity: function setSign(string input) returns()
func (_Audit *AuditTransactor) SetSign(opts *bind.TransactOpts, input string) (*types.Transaction, error) {
	return _Audit.contract.Transact(opts, "setSign", input)
}

// SetSign is a paid mutator transaction binding the contract method 0x31fbbae2.
//
// Solidity: function setSign(string input) returns()
func (_Audit *AuditSession) SetSign(input string) (*types.Transaction, error) {
	return _Audit.Contract.SetSign(&_Audit.TransactOpts, input)
}

// SetSign is a paid mutator transaction binding the contract method 0x31fbbae2.
//
// Solidity: function setSign(string input) returns()
func (_Audit *AuditTransactorSession) SetSign(input string) (*types.Transaction, error) {
	return _Audit.Contract.SetSign(&_Audit.TransactOpts, input)
}
