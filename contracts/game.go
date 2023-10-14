// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// GameInfoContractGameInfo is an auto generated low-level Go binding around an user-defined struct.
type GameInfoContractGameInfo struct {
	DomainSeparationTag int64
	ReplayCID           string
	Entropy             []byte
	VRFHeight           uint64
	VRFProof            []byte
	GameResults         []byte
}

// GameMetaData contains all meta data concerning the Game contract.
var GameMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_roundId\",\"type\":\"string\"}],\"name\":\"getGameInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"int64\",\"name\":\"DomainSeparationTag\",\"type\":\"int64\"},{\"internalType\":\"string\",\"name\":\"ReplayCID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"Entropy\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"VRFHeight\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"VRFProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"GameResults\",\"type\":\"bytes\"}],\"internalType\":\"structGameInfoContract.GameInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_roundId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"DomainSeparationTag\",\"type\":\"int64\"},{\"internalType\":\"string\",\"name\":\"ReplayCID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"Entropy\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"VRFHeight\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"VRFProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"GameResults\",\"type\":\"bytes\"}],\"internalType\":\"structGameInfoContract.GameInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"name\":\"setGameInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5061002c61002161003160201b60201c565b61003860201b60201c565b6100f9565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b61138e806101065f395ff3fe608060405234801561000f575f80fd5b5060043610610055575f3560e01c8063715018a6146100595780638da5cb5b14610063578063c857575114610081578063dcc5dfef146100b1578063f2fde38b146100cd575b5f80fd5b6100616100e9565b005b61006b6100fc565b604051610078919061075e565b60405180910390f35b61009b600480360381019061009691906108c4565b610123565b6040516100a89190610aaf565b60405180910390f35b6100cb60048036038101906100c69190610cd6565b610433565b005b6100e760048036038101906100e29190610d76565b610516565b005b6100f1610598565b6100fa5f610616565b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61012b6106de565b5f60018360405161013c9190610ddb565b90815260200160405180910390206040518060c00160405290815f82015f9054906101000a900460070b60070b60070b815260200160018201805461018090610e1e565b80601f01602080910402602001604051908101604052809291908181526020018280546101ac90610e1e565b80156101f75780601f106101ce576101008083540402835291602001916101f7565b820191905f5260205f20905b8154815290600101906020018083116101da57829003601f168201915b5050505050815260200160028201805461021090610e1e565b80601f016020809104026020016040519081016040528092919081815260200182805461023c90610e1e565b80156102875780601f1061025e57610100808354040283529160200191610287565b820191905f5260205f20905b81548152906001019060200180831161026a57829003601f168201915b50505050508152602001600382015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820180546102d190610e1e565b80601f01602080910402602001604051908101604052809291908181526020018280546102fd90610e1e565b80156103485780601f1061031f57610100808354040283529160200191610348565b820191905f5260205f20905b81548152906001019060200180831161032b57829003601f168201915b5050505050815260200160058201805461036190610e1e565b80601f016020809104026020016040519081016040528092919081815260200182805461038d90610e1e565b80156103d85780601f106103af576101008083540402835291602001916103d8565b820191905f5260205f20905b8154815290600101906020018083116103bb57829003601f168201915b50505050508152505090505f8160200151511161042a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161042190610ea8565b60405180910390fd5b80915050919050565b61043b610598565b8060018360405161044c9190610ddb565b90815260200160405180910390205f820151815f015f6101000a81548167ffffffffffffffff021916908360070b67ffffffffffffffff160217905550602082015181600101908161049e919061106c565b5060408201518160020190816104b49190611193565b506060820151816003015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160040190816104f89190611193565b5060a082015181600501908161050e9190611193565b509050505050565b61051e610598565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361058c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610583906112d2565b60405180910390fd5b61059581610616565b50565b6105a06106d7565b73ffffffffffffffffffffffffffffffffffffffff166105be6100fc565b73ffffffffffffffffffffffffffffffffffffffff1614610614576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161060b9061133a565b60405180910390fd5b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f33905090565b6040518060c001604052805f60070b815260200160608152602001606081526020015f67ffffffffffffffff16815260200160608152602001606081525090565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6107488261071f565b9050919050565b6107588161073e565b82525050565b5f6020820190506107715f83018461074f565b92915050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6107d682610790565b810181811067ffffffffffffffff821117156107f5576107f46107a0565b5b80604052505050565b5f610807610777565b905061081382826107cd565b919050565b5f67ffffffffffffffff821115610832576108316107a0565b5b61083b82610790565b9050602081019050919050565b828183375f83830152505050565b5f61086861086384610818565b6107fe565b9050828152602081018484840111156108845761088361078c565b5b61088f848285610848565b509392505050565b5f82601f8301126108ab576108aa610788565b5b81356108bb848260208601610856565b91505092915050565b5f602082840312156108d9576108d8610780565b5b5f82013567ffffffffffffffff8111156108f6576108f5610784565b5b61090284828501610897565b91505092915050565b5f8160070b9050919050565b6109208161090b565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561095d578082015181840152602081019050610942565b5f8484015250505050565b5f61097282610926565b61097c8185610930565b935061098c818560208601610940565b61099581610790565b840191505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f6109c4826109a0565b6109ce81856109aa565b93506109de818560208601610940565b6109e781610790565b840191505092915050565b5f67ffffffffffffffff82169050919050565b610a0e816109f2565b82525050565b5f60c083015f830151610a295f860182610917565b5060208301518482036020860152610a418282610968565b91505060408301518482036040860152610a5b82826109ba565b9150506060830151610a706060860182610a05565b5060808301518482036080860152610a8882826109ba565b91505060a083015184820360a0860152610aa282826109ba565b9150508091505092915050565b5f6020820190508181035f830152610ac78184610a14565b905092915050565b5f80fd5b5f80fd5b610ae08161090b565b8114610aea575f80fd5b50565b5f81359050610afb81610ad7565b92915050565b5f67ffffffffffffffff821115610b1b57610b1a6107a0565b5b610b2482610790565b9050602081019050919050565b5f610b43610b3e84610b01565b6107fe565b905082815260208101848484011115610b5f57610b5e61078c565b5b610b6a848285610848565b509392505050565b5f82601f830112610b8657610b85610788565b5b8135610b96848260208601610b31565b91505092915050565b610ba8816109f2565b8114610bb2575f80fd5b50565b5f81359050610bc381610b9f565b92915050565b5f60c08284031215610bde57610bdd610acf565b5b610be860c06107fe565b90505f610bf784828501610aed565b5f83015250602082013567ffffffffffffffff811115610c1a57610c19610ad3565b5b610c2684828501610897565b602083015250604082013567ffffffffffffffff811115610c4a57610c49610ad3565b5b610c5684828501610b72565b6040830152506060610c6a84828501610bb5565b606083015250608082013567ffffffffffffffff811115610c8e57610c8d610ad3565b5b610c9a84828501610b72565b60808301525060a082013567ffffffffffffffff811115610cbe57610cbd610ad3565b5b610cca84828501610b72565b60a08301525092915050565b5f8060408385031215610cec57610ceb610780565b5b5f83013567ffffffffffffffff811115610d0957610d08610784565b5b610d1585828601610897565b925050602083013567ffffffffffffffff811115610d3657610d35610784565b5b610d4285828601610bc9565b9150509250929050565b610d558161073e565b8114610d5f575f80fd5b50565b5f81359050610d7081610d4c565b92915050565b5f60208284031215610d8b57610d8a610780565b5b5f610d9884828501610d62565b91505092915050565b5f81905092915050565b5f610db582610926565b610dbf8185610da1565b9350610dcf818560208601610940565b80840191505092915050565b5f610de68284610dab565b915081905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610e3557607f821691505b602082108103610e4857610e47610df1565b5b50919050565b5f82825260208201905092915050565b7f47616d6520696e666f206e6f7420666f756e64000000000000000000000000005f82015250565b5f610e92601383610e4e565b9150610e9d82610e5e565b602082019050919050565b5f6020820190508181035f830152610ebf81610e86565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610f227fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610ee7565b610f2c8683610ee7565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f610f70610f6b610f6684610f44565b610f4d565b610f44565b9050919050565b5f819050919050565b610f8983610f56565b610f9d610f9582610f77565b848454610ef3565b825550505050565b5f90565b610fb1610fa5565b610fbc818484610f80565b505050565b5b81811015610fdf57610fd45f82610fa9565b600181019050610fc2565b5050565b601f82111561102457610ff581610ec6565b610ffe84610ed8565b8101602085101561100d578190505b61102161101985610ed8565b830182610fc1565b50505b505050565b5f82821c905092915050565b5f6110445f1984600802611029565b1980831691505092915050565b5f61105c8383611035565b9150826002028217905092915050565b61107582610926565b67ffffffffffffffff81111561108e5761108d6107a0565b5b6110988254610e1e565b6110a3828285610fe3565b5f60209050601f8311600181146110d4575f84156110c2578287015190505b6110cc8582611051565b865550611133565b601f1984166110e286610ec6565b5f5b82811015611109578489015182556001820191506020850194506020810190506110e4565b868310156111265784890151611122601f891682611035565b8355505b6001600288020188555050505b505050505050565b5f819050815f5260205f209050919050565b601f82111561118e5761115f8161113b565b61116884610ed8565b81016020851015611177578190505b61118b61118385610ed8565b830182610fc1565b50505b505050565b61119c826109a0565b67ffffffffffffffff8111156111b5576111b46107a0565b5b6111bf8254610e1e565b6111ca82828561114d565b5f60209050601f8311600181146111fb575f84156111e9578287015190505b6111f38582611051565b86555061125a565b601f1984166112098661113b565b5f5b828110156112305784890151825560018201915060208501945060208101905061120b565b8683101561124d5784890151611249601f891682611035565b8355505b6001600288020188555050505b505050505050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f6112bc602683610e4e565b91506112c782611262565b604082019050919050565b5f6020820190508181035f8301526112e9816112b0565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f611324602083610e4e565b915061132f826112f0565b602082019050919050565b5f6020820190508181035f83015261135181611318565b905091905056fea2646970667358221220a10e0ffb77087937959cccf5f6059961e099b814c689e8101a9d46cb54cafb3464736f6c63430008150033",
}

// GameABI is the input ABI used to generate the binding from.
// Deprecated: Use GameMetaData.ABI instead.
var GameABI = GameMetaData.ABI

// GameBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GameMetaData.Bin instead.
var GameBin = GameMetaData.Bin

// DeployGame deploys a new Ethereum contract, binding an instance of Game to it.
func DeployGame(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Game, error) {
	parsed, err := GameMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GameBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Game{GameCaller: GameCaller{contract: contract}, GameTransactor: GameTransactor{contract: contract}, GameFilterer: GameFilterer{contract: contract}}, nil
}

// Game is an auto generated Go binding around an Ethereum contract.
type Game struct {
	GameCaller     // Read-only binding to the contract
	GameTransactor // Write-only binding to the contract
	GameFilterer   // Log filterer for contract events
}

// GameCaller is an auto generated read-only Go binding around an Ethereum contract.
type GameCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GameTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GameFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GameSession struct {
	Contract     *Game             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GameCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GameCallerSession struct {
	Contract *GameCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// GameTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GameTransactorSession struct {
	Contract     *GameTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GameRaw is an auto generated low-level Go binding around an Ethereum contract.
type GameRaw struct {
	Contract *Game // Generic contract binding to access the raw methods on
}

// GameCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GameCallerRaw struct {
	Contract *GameCaller // Generic read-only contract binding to access the raw methods on
}

// GameTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GameTransactorRaw struct {
	Contract *GameTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGame creates a new instance of Game, bound to a specific deployed contract.
func NewGame(address common.Address, backend bind.ContractBackend) (*Game, error) {
	contract, err := bindGame(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Game{GameCaller: GameCaller{contract: contract}, GameTransactor: GameTransactor{contract: contract}, GameFilterer: GameFilterer{contract: contract}}, nil
}

// NewGameCaller creates a new read-only instance of Game, bound to a specific deployed contract.
func NewGameCaller(address common.Address, caller bind.ContractCaller) (*GameCaller, error) {
	contract, err := bindGame(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GameCaller{contract: contract}, nil
}

// NewGameTransactor creates a new write-only instance of Game, bound to a specific deployed contract.
func NewGameTransactor(address common.Address, transactor bind.ContractTransactor) (*GameTransactor, error) {
	contract, err := bindGame(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GameTransactor{contract: contract}, nil
}

// NewGameFilterer creates a new log filterer instance of Game, bound to a specific deployed contract.
func NewGameFilterer(address common.Address, filterer bind.ContractFilterer) (*GameFilterer, error) {
	contract, err := bindGame(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GameFilterer{contract: contract}, nil
}

// bindGame binds a generic wrapper to an already deployed contract.
func bindGame(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GameMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Game *GameRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Game.Contract.GameCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Game *GameRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Game.Contract.GameTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Game *GameRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Game.Contract.GameTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Game *GameCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Game.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Game *GameTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Game.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Game *GameTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Game.Contract.contract.Transact(opts, method, params...)
}

// GetGameInfo is a free data retrieval call binding the contract method 0xc8575751.
//
// Solidity: function getGameInfo(string _roundId) view returns((int64,string,bytes,uint64,bytes,bytes))
func (_Game *GameCaller) GetGameInfo(opts *bind.CallOpts, _roundId string) (GameInfoContractGameInfo, error) {
	var out []interface{}
	err := _Game.contract.Call(opts, &out, "getGameInfo", _roundId)

	if err != nil {
		return *new(GameInfoContractGameInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(GameInfoContractGameInfo)).(*GameInfoContractGameInfo)

	return out0, err

}

// GetGameInfo is a free data retrieval call binding the contract method 0xc8575751.
//
// Solidity: function getGameInfo(string _roundId) view returns((int64,string,bytes,uint64,bytes,bytes))
func (_Game *GameSession) GetGameInfo(_roundId string) (GameInfoContractGameInfo, error) {
	return _Game.Contract.GetGameInfo(&_Game.CallOpts, _roundId)
}

// GetGameInfo is a free data retrieval call binding the contract method 0xc8575751.
//
// Solidity: function getGameInfo(string _roundId) view returns((int64,string,bytes,uint64,bytes,bytes))
func (_Game *GameCallerSession) GetGameInfo(_roundId string) (GameInfoContractGameInfo, error) {
	return _Game.Contract.GetGameInfo(&_Game.CallOpts, _roundId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Game *GameCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Game.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Game *GameSession) Owner() (common.Address, error) {
	return _Game.Contract.Owner(&_Game.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Game *GameCallerSession) Owner() (common.Address, error) {
	return _Game.Contract.Owner(&_Game.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Game *GameTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Game.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Game *GameSession) RenounceOwnership() (*types.Transaction, error) {
	return _Game.Contract.RenounceOwnership(&_Game.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Game *GameTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Game.Contract.RenounceOwnership(&_Game.TransactOpts)
}

// SetGameInfo is a paid mutator transaction binding the contract method 0xdcc5dfef.
//
// Solidity: function setGameInfo(string _roundId, (int64,string,bytes,uint64,bytes,bytes) _info) returns()
func (_Game *GameTransactor) SetGameInfo(opts *bind.TransactOpts, _roundId string, _info GameInfoContractGameInfo) (*types.Transaction, error) {
	return _Game.contract.Transact(opts, "setGameInfo", _roundId, _info)
}

// SetGameInfo is a paid mutator transaction binding the contract method 0xdcc5dfef.
//
// Solidity: function setGameInfo(string _roundId, (int64,string,bytes,uint64,bytes,bytes) _info) returns()
func (_Game *GameSession) SetGameInfo(_roundId string, _info GameInfoContractGameInfo) (*types.Transaction, error) {
	return _Game.Contract.SetGameInfo(&_Game.TransactOpts, _roundId, _info)
}

// SetGameInfo is a paid mutator transaction binding the contract method 0xdcc5dfef.
//
// Solidity: function setGameInfo(string _roundId, (int64,string,bytes,uint64,bytes,bytes) _info) returns()
func (_Game *GameTransactorSession) SetGameInfo(_roundId string, _info GameInfoContractGameInfo) (*types.Transaction, error) {
	return _Game.Contract.SetGameInfo(&_Game.TransactOpts, _roundId, _info)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Game *GameTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Game.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Game *GameSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Game.Contract.TransferOwnership(&_Game.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Game *GameTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Game.Contract.TransferOwnership(&_Game.TransactOpts, newOwner)
}

// GameOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Game contract.
type GameOwnershipTransferredIterator struct {
	Event *GameOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GameOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameOwnershipTransferred)
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
		it.Event = new(GameOwnershipTransferred)
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
func (it *GameOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameOwnershipTransferred represents a OwnershipTransferred event raised by the Game contract.
type GameOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Game *GameFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GameOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Game.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GameOwnershipTransferredIterator{contract: _Game.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Game *GameFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GameOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Game.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameOwnershipTransferred)
				if err := _Game.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Game *GameFilterer) ParseOwnershipTransferred(log types.Log) (*GameOwnershipTransferred, error) {
	event := new(GameOwnershipTransferred)
	if err := _Game.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
