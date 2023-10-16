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

// GameInfo is an auto generated low-level Go binding around an user-defined struct.
type GameInfo struct {
	GameID    string
	RoundID   string
	ReplayID  string
	PlayerIDs string
}

// GameReplay is an auto generated low-level Go binding around an user-defined struct.
type GameReplay struct {
	DomainSeparationTag int64
	Address             string
	GameInfo            GameInfo
	VRFHeight           uint64
	VRFProof            []byte
	GameResult          GameResult
}

// GameResult is an auto generated low-level Go binding around an user-defined struct.
type GameResult struct {
	GameID  string
	RoundID string
}

// GameReplayContractMetaData contains all meta data concerning the GameReplayContract contract.
var GameReplayContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_roundID\",\"type\":\"string\"}],\"name\":\"getGameReplay\",\"outputs\":[{\"components\":[{\"internalType\":\"int64\",\"name\":\"DomainSeparationTag\",\"type\":\"int64\"},{\"internalType\":\"string\",\"name\":\"Address\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"GameID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"RoundID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ReplayID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"PlayerIDs\",\"type\":\"string\"}],\"internalType\":\"structGame.Info\",\"name\":\"GameInfo\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"VRFHeight\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"VRFProof\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"GameID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"RoundID\",\"type\":\"string\"}],\"internalType\":\"structGame.Result\",\"name\":\"GameResult\",\"type\":\"tuple\"}],\"internalType\":\"structGame.Replay\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"int64\",\"name\":\"DomainSeparationTag\",\"type\":\"int64\"},{\"internalType\":\"string\",\"name\":\"Address\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"GameID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"RoundID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ReplayID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"PlayerIDs\",\"type\":\"string\"}],\"internalType\":\"structGame.Info\",\"name\":\"GameInfo\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"VRFHeight\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"VRFProof\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"GameID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"RoundID\",\"type\":\"string\"}],\"internalType\":\"structGame.Result\",\"name\":\"GameResult\",\"type\":\"tuple\"}],\"internalType\":\"structGame.Replay[]\",\"name\":\"_replays\",\"type\":\"tuple[]\"}],\"name\":\"saveGameReplay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5061002c61002161003160201b60201c565b61003860201b60201c565b6100f9565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b611ba3806101065f395ff3fe608060405234801561000f575f80fd5b5060043610610055575f3560e01c8063715018a6146100595780638da5cb5b14610063578063b970994814610081578063db7884f2146100b1578063f2fde38b146100cd575b5f80fd5b6100616100e9565b005b61006b6100fc565b6040516100789190610b3c565b60405180910390f35b61009b60048036038101906100969190610ca2565b610123565b6040516100a89190610f43565b60405180910390f35b6100cb60048036038101906100c691906113b6565b6106bc565b005b6100e760048036038101906100e29190611427565b6108a6565b005b6100f1610928565b6100fa5f6109a6565b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61012b610a6e565b5f60018360405161013c919061148c565b90815260200160405180910390206040518060c00160405290815f82015f9054906101000a900460070b60070b60070b8152602001600182018054610180906114cf565b80601f01602080910402602001604051908101604052809291908181526020018280546101ac906114cf565b80156101f75780601f106101ce576101008083540402835291602001916101f7565b820191905f5260205f20905b8154815290600101906020018083116101da57829003601f168201915b50505050508152602001600282016040518060800160405290815f8201805461021f906114cf565b80601f016020809104026020016040519081016040528092919081815260200182805461024b906114cf565b80156102965780601f1061026d57610100808354040283529160200191610296565b820191905f5260205f20905b81548152906001019060200180831161027957829003601f168201915b505050505081526020016001820180546102af906114cf565b80601f01602080910402602001604051908101604052809291908181526020018280546102db906114cf565b80156103265780601f106102fd57610100808354040283529160200191610326565b820191905f5260205f20905b81548152906001019060200180831161030957829003601f168201915b5050505050815260200160028201805461033f906114cf565b80601f016020809104026020016040519081016040528092919081815260200182805461036b906114cf565b80156103b65780601f1061038d576101008083540402835291602001916103b6565b820191905f5260205f20905b81548152906001019060200180831161039957829003601f168201915b505050505081526020016003820180546103cf906114cf565b80601f01602080910402602001604051908101604052809291908181526020018280546103fb906114cf565b80156104465780601f1061041d57610100808354040283529160200191610446565b820191905f5260205f20905b81548152906001019060200180831161042957829003601f168201915b5050505050815250508152602001600682015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600782018054610494906114cf565b80601f01602080910402602001604051908101604052809291908181526020018280546104c0906114cf565b801561050b5780601f106104e25761010080835404028352916020019161050b565b820191905f5260205f20905b8154815290600101906020018083116104ee57829003601f168201915b50505050508152602001600882016040518060400160405290815f82018054610533906114cf565b80601f016020809104026020016040519081016040528092919081815260200182805461055f906114cf565b80156105aa5780601f10610581576101008083540402835291602001916105aa565b820191905f5260205f20905b81548152906001019060200180831161058d57829003601f168201915b505050505081526020016001820180546105c3906114cf565b80601f01602080910402602001604051908101604052809291908181526020018280546105ef906114cf565b801561063a5780601f106106115761010080835404028352916020019161063a565b820191905f5260205f20905b81548152906001019060200180831161061d57829003601f168201915b5050505050815250508152505090505f81602001515111836040516020016106629190611549565b604051602081830303815290604052906106b2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106a991906115b2565b60405180910390fd5b5080915050919050565b6106c4610928565b5f815111610707576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106fe9061161c565b60405180910390fd5b5f5b81518110156108a2578181815181106107255761072461163a565b5b602002602001015160018383815181106107425761074161163a565b5b6020026020010151604001516020015160405161075f919061148c565b90815260200160405180910390205f820151815f015f6101000a81548167ffffffffffffffff021916908360070b67ffffffffffffffff16021790555060208201518160010190816107b1919061180d565b506040820151816002015f820151815f0190816107ce919061180d565b5060208201518160010190816107e4919061180d565b5060408201518160020190816107fa919061180d565b506060820151816003019081610810919061180d565b5050506060820151816006015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160070190816108569190611934565b5060a0820151816008015f820151815f019081610873919061180d565b506020820151816001019081610889919061180d565b505050905050808061089a90611a30565b915050610709565b5050565b6108ae610928565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361091c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161091390611ae7565b60405180910390fd5b610925816109a6565b50565b610930610a67565b73ffffffffffffffffffffffffffffffffffffffff1661094e6100fc565b73ffffffffffffffffffffffffffffffffffffffff16146109a4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161099b90611b4f565b60405180910390fd5b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f33905090565b6040518060c001604052805f60070b815260200160608152602001610a91610abb565b81526020015f67ffffffffffffffff16815260200160608152602001610ab5610ae3565b81525090565b6040518060800160405280606081526020016060815260200160608152602001606081525090565b604051806040016040528060608152602001606081525090565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610b2682610afd565b9050919050565b610b3681610b1c565b82525050565b5f602082019050610b4f5f830184610b2d565b92915050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610bb482610b6e565b810181811067ffffffffffffffff82111715610bd357610bd2610b7e565b5b80604052505050565b5f610be5610b55565b9050610bf18282610bab565b919050565b5f67ffffffffffffffff821115610c1057610c0f610b7e565b5b610c1982610b6e565b9050602081019050919050565b828183375f83830152505050565b5f610c46610c4184610bf6565b610bdc565b905082815260208101848484011115610c6257610c61610b6a565b5b610c6d848285610c26565b509392505050565b5f82601f830112610c8957610c88610b66565b5b8135610c99848260208601610c34565b91505092915050565b5f60208284031215610cb757610cb6610b5e565b5b5f82013567ffffffffffffffff811115610cd457610cd3610b62565b5b610ce084828501610c75565b91505092915050565b5f8160070b9050919050565b610cfe81610ce9565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610d3b578082015181840152602081019050610d20565b5f8484015250505050565b5f610d5082610d04565b610d5a8185610d0e565b9350610d6a818560208601610d1e565b610d7381610b6e565b840191505092915050565b5f608083015f8301518482035f860152610d988282610d46565b91505060208301518482036020860152610db28282610d46565b91505060408301518482036040860152610dcc8282610d46565b91505060608301518482036060860152610de68282610d46565b9150508091505092915050565b5f67ffffffffffffffff82169050919050565b610e0f81610df3565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f610e3982610e15565b610e438185610e1f565b9350610e53818560208601610d1e565b610e5c81610b6e565b840191505092915050565b5f604083015f8301518482035f860152610e818282610d46565b91505060208301518482036020860152610e9b8282610d46565b9150508091505092915050565b5f60c083015f830151610ebd5f860182610cf5565b5060208301518482036020860152610ed58282610d46565b91505060408301518482036040860152610eef8282610d7e565b9150506060830151610f046060860182610e06565b5060808301518482036080860152610f1c8282610e2f565b91505060a083015184820360a0860152610f368282610e67565b9150508091505092915050565b5f6020820190508181035f830152610f5b8184610ea8565b905092915050565b5f67ffffffffffffffff821115610f7d57610f7c610b7e565b5b602082029050602081019050919050565b5f80fd5b5f80fd5b5f80fd5b610fa381610ce9565b8114610fad575f80fd5b50565b5f81359050610fbe81610f9a565b92915050565b5f60808284031215610fd957610fd8610f92565b5b610fe36080610bdc565b90505f82013567ffffffffffffffff81111561100257611001610f96565b5b61100e84828501610c75565b5f83015250602082013567ffffffffffffffff81111561103157611030610f96565b5b61103d84828501610c75565b602083015250604082013567ffffffffffffffff81111561106157611060610f96565b5b61106d84828501610c75565b604083015250606082013567ffffffffffffffff81111561109157611090610f96565b5b61109d84828501610c75565b60608301525092915050565b6110b281610df3565b81146110bc575f80fd5b50565b5f813590506110cd816110a9565b92915050565b5f67ffffffffffffffff8211156110ed576110ec610b7e565b5b6110f682610b6e565b9050602081019050919050565b5f611115611110846110d3565b610bdc565b90508281526020810184848401111561113157611130610b6a565b5b61113c848285610c26565b509392505050565b5f82601f83011261115857611157610b66565b5b8135611168848260208601611103565b91505092915050565b5f6040828403121561118657611185610f92565b5b6111906040610bdc565b90505f82013567ffffffffffffffff8111156111af576111ae610f96565b5b6111bb84828501610c75565b5f83015250602082013567ffffffffffffffff8111156111de576111dd610f96565b5b6111ea84828501610c75565b60208301525092915050565b5f60c0828403121561120b5761120a610f92565b5b61121560c0610bdc565b90505f61122484828501610fb0565b5f83015250602082013567ffffffffffffffff81111561124757611246610f96565b5b61125384828501610c75565b602083015250604082013567ffffffffffffffff81111561127757611276610f96565b5b61128384828501610fc4565b6040830152506060611297848285016110bf565b606083015250608082013567ffffffffffffffff8111156112bb576112ba610f96565b5b6112c784828501611144565b60808301525060a082013567ffffffffffffffff8111156112eb576112ea610f96565b5b6112f784828501611171565b60a08301525092915050565b5f61131561131084610f63565b610bdc565b9050808382526020820190506020840283018581111561133857611337610f8e565b5b835b8181101561137f57803567ffffffffffffffff81111561135d5761135c610b66565b5b80860161136a89826111f6565b8552602085019450505060208101905061133a565b5050509392505050565b5f82601f83011261139d5761139c610b66565b5b81356113ad848260208601611303565b91505092915050565b5f602082840312156113cb576113ca610b5e565b5b5f82013567ffffffffffffffff8111156113e8576113e7610b62565b5b6113f484828501611389565b91505092915050565b61140681610b1c565b8114611410575f80fd5b50565b5f81359050611421816113fd565b92915050565b5f6020828403121561143c5761143b610b5e565b5b5f61144984828501611413565b91505092915050565b5f81905092915050565b5f61146682610d04565b6114708185611452565b9350611480818560208601610d1e565b80840191505092915050565b5f611497828461145c565b915081905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806114e657607f821691505b6020821081036114f9576114f86114a2565b5b50919050565b7f47616d65207265706c6179206e6f7420666f756e643a200000000000000000005f82015250565b5f611533601783611452565b915061153e826114ff565b601782019050919050565b5f61155382611527565b915061155f828461145c565b915081905092915050565b5f82825260208201905092915050565b5f61158482610d04565b61158e818561156a565b935061159e818560208601610d1e565b6115a781610b6e565b840191505092915050565b5f6020820190508181035f8301526115ca818461157a565b905092915050565b7f5f7265706c6179732063616e206e6f7420656d707479000000000000000000005f82015250565b5f61160660168361156a565b9150611611826115d2565b602082019050919050565b5f6020820190508181035f830152611633816115fa565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026116c37fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611688565b6116cd8683611688565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61171161170c611707846116e5565b6116ee565b6116e5565b9050919050565b5f819050919050565b61172a836116f7565b61173e61173682611718565b848454611694565b825550505050565b5f90565b611752611746565b61175d818484611721565b505050565b5b81811015611780576117755f8261174a565b600181019050611763565b5050565b601f8211156117c55761179681611667565b61179f84611679565b810160208510156117ae578190505b6117c26117ba85611679565b830182611762565b50505b505050565b5f82821c905092915050565b5f6117e55f19846008026117ca565b1980831691505092915050565b5f6117fd83836117d6565b9150826002028217905092915050565b61181682610d04565b67ffffffffffffffff81111561182f5761182e610b7e565b5b61183982546114cf565b611844828285611784565b5f60209050601f831160018114611875575f8415611863578287015190505b61186d85826117f2565b8655506118d4565b601f19841661188386611667565b5f5b828110156118aa57848901518255600182019150602085019450602081019050611885565b868310156118c757848901516118c3601f8916826117d6565b8355505b6001600288020188555050505b505050505050565b5f819050815f5260205f209050919050565b601f82111561192f57611900816118dc565b61190984611679565b81016020851015611918578190505b61192c61192485611679565b830182611762565b50505b505050565b61193d82610e15565b67ffffffffffffffff81111561195657611955610b7e565b5b61196082546114cf565b61196b8282856118ee565b5f60209050601f83116001811461199c575f841561198a578287015190505b61199485826117f2565b8655506119fb565b601f1984166119aa866118dc565b5f5b828110156119d1578489015182556001820191506020850194506020810190506119ac565b868310156119ee57848901516119ea601f8916826117d6565b8355505b6001600288020188555050505b505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611a3a826116e5565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611a6c57611a6b611a03565b5b600182019050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f611ad160268361156a565b9150611adc82611a77565b604082019050919050565b5f6020820190508181035f830152611afe81611ac5565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f611b3960208361156a565b9150611b4482611b05565b602082019050919050565b5f6020820190508181035f830152611b6681611b2d565b905091905056fea26469706673582212207f85929d260dc48995bb37d8f0b1e989cc16178238be4aaa0df2ef3b394df43c64736f6c63430008150033",
}

// GameReplayContractABI is the input ABI used to generate the binding from.
// Deprecated: Use GameReplayContractMetaData.ABI instead.
var GameReplayContractABI = GameReplayContractMetaData.ABI

// GameReplayContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GameReplayContractMetaData.Bin instead.
var GameReplayContractBin = GameReplayContractMetaData.Bin

// DeployGameReplayContract deploys a new Ethereum contract, binding an instance of GameReplayContract to it.
func DeployGameReplayContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GameReplayContract, error) {
	parsed, err := GameReplayContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GameReplayContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GameReplayContract{GameReplayContractCaller: GameReplayContractCaller{contract: contract}, GameReplayContractTransactor: GameReplayContractTransactor{contract: contract}, GameReplayContractFilterer: GameReplayContractFilterer{contract: contract}}, nil
}

// GameReplayContract is an auto generated Go binding around an Ethereum contract.
type GameReplayContract struct {
	GameReplayContractCaller     // Read-only binding to the contract
	GameReplayContractTransactor // Write-only binding to the contract
	GameReplayContractFilterer   // Log filterer for contract events
}

// GameReplayContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type GameReplayContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameReplayContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GameReplayContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameReplayContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GameReplayContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameReplayContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GameReplayContractSession struct {
	Contract     *GameReplayContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// GameReplayContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GameReplayContractCallerSession struct {
	Contract *GameReplayContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// GameReplayContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GameReplayContractTransactorSession struct {
	Contract     *GameReplayContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// GameReplayContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type GameReplayContractRaw struct {
	Contract *GameReplayContract // Generic contract binding to access the raw methods on
}

// GameReplayContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GameReplayContractCallerRaw struct {
	Contract *GameReplayContractCaller // Generic read-only contract binding to access the raw methods on
}

// GameReplayContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GameReplayContractTransactorRaw struct {
	Contract *GameReplayContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGameReplayContract creates a new instance of GameReplayContract, bound to a specific deployed contract.
func NewGameReplayContract(address common.Address, backend bind.ContractBackend) (*GameReplayContract, error) {
	contract, err := bindGameReplayContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GameReplayContract{GameReplayContractCaller: GameReplayContractCaller{contract: contract}, GameReplayContractTransactor: GameReplayContractTransactor{contract: contract}, GameReplayContractFilterer: GameReplayContractFilterer{contract: contract}}, nil
}

// NewGameReplayContractCaller creates a new read-only instance of GameReplayContract, bound to a specific deployed contract.
func NewGameReplayContractCaller(address common.Address, caller bind.ContractCaller) (*GameReplayContractCaller, error) {
	contract, err := bindGameReplayContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GameReplayContractCaller{contract: contract}, nil
}

// NewGameReplayContractTransactor creates a new write-only instance of GameReplayContract, bound to a specific deployed contract.
func NewGameReplayContractTransactor(address common.Address, transactor bind.ContractTransactor) (*GameReplayContractTransactor, error) {
	contract, err := bindGameReplayContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GameReplayContractTransactor{contract: contract}, nil
}

// NewGameReplayContractFilterer creates a new log filterer instance of GameReplayContract, bound to a specific deployed contract.
func NewGameReplayContractFilterer(address common.Address, filterer bind.ContractFilterer) (*GameReplayContractFilterer, error) {
	contract, err := bindGameReplayContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GameReplayContractFilterer{contract: contract}, nil
}

// bindGameReplayContract binds a generic wrapper to an already deployed contract.
func bindGameReplayContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GameReplayContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GameReplayContract *GameReplayContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GameReplayContract.Contract.GameReplayContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GameReplayContract *GameReplayContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameReplayContract.Contract.GameReplayContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GameReplayContract *GameReplayContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GameReplayContract.Contract.GameReplayContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GameReplayContract *GameReplayContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GameReplayContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GameReplayContract *GameReplayContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameReplayContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GameReplayContract *GameReplayContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GameReplayContract.Contract.contract.Transact(opts, method, params...)
}

// GetGameReplay is a free data retrieval call binding the contract method 0xb9709948.
//
// Solidity: function getGameReplay(string _roundID) view returns((int64,string,(string,string,string,string),uint64,bytes,(string,string)))
func (_GameReplayContract *GameReplayContractCaller) GetGameReplay(opts *bind.CallOpts, _roundID string) (GameReplay, error) {
	var out []interface{}
	err := _GameReplayContract.contract.Call(opts, &out, "getGameReplay", _roundID)

	if err != nil {
		return *new(GameReplay), err
	}

	out0 := *abi.ConvertType(out[0], new(GameReplay)).(*GameReplay)

	return out0, err

}

// GetGameReplay is a free data retrieval call binding the contract method 0xb9709948.
//
// Solidity: function getGameReplay(string _roundID) view returns((int64,string,(string,string,string,string),uint64,bytes,(string,string)))
func (_GameReplayContract *GameReplayContractSession) GetGameReplay(_roundID string) (GameReplay, error) {
	return _GameReplayContract.Contract.GetGameReplay(&_GameReplayContract.CallOpts, _roundID)
}

// GetGameReplay is a free data retrieval call binding the contract method 0xb9709948.
//
// Solidity: function getGameReplay(string _roundID) view returns((int64,string,(string,string,string,string),uint64,bytes,(string,string)))
func (_GameReplayContract *GameReplayContractCallerSession) GetGameReplay(_roundID string) (GameReplay, error) {
	return _GameReplayContract.Contract.GetGameReplay(&_GameReplayContract.CallOpts, _roundID)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GameReplayContract *GameReplayContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GameReplayContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GameReplayContract *GameReplayContractSession) Owner() (common.Address, error) {
	return _GameReplayContract.Contract.Owner(&_GameReplayContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GameReplayContract *GameReplayContractCallerSession) Owner() (common.Address, error) {
	return _GameReplayContract.Contract.Owner(&_GameReplayContract.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GameReplayContract *GameReplayContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameReplayContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GameReplayContract *GameReplayContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _GameReplayContract.Contract.RenounceOwnership(&_GameReplayContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GameReplayContract *GameReplayContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GameReplayContract.Contract.RenounceOwnership(&_GameReplayContract.TransactOpts)
}

// SaveGameReplay is a paid mutator transaction binding the contract method 0xdb7884f2.
//
// Solidity: function saveGameReplay((int64,string,(string,string,string,string),uint64,bytes,(string,string))[] _replays) returns()
func (_GameReplayContract *GameReplayContractTransactor) SaveGameReplay(opts *bind.TransactOpts, _replays []GameReplay) (*types.Transaction, error) {
	return _GameReplayContract.contract.Transact(opts, "saveGameReplay", _replays)
}

// SaveGameReplay is a paid mutator transaction binding the contract method 0xdb7884f2.
//
// Solidity: function saveGameReplay((int64,string,(string,string,string,string),uint64,bytes,(string,string))[] _replays) returns()
func (_GameReplayContract *GameReplayContractSession) SaveGameReplay(_replays []GameReplay) (*types.Transaction, error) {
	return _GameReplayContract.Contract.SaveGameReplay(&_GameReplayContract.TransactOpts, _replays)
}

// SaveGameReplay is a paid mutator transaction binding the contract method 0xdb7884f2.
//
// Solidity: function saveGameReplay((int64,string,(string,string,string,string),uint64,bytes,(string,string))[] _replays) returns()
func (_GameReplayContract *GameReplayContractTransactorSession) SaveGameReplay(_replays []GameReplay) (*types.Transaction, error) {
	return _GameReplayContract.Contract.SaveGameReplay(&_GameReplayContract.TransactOpts, _replays)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GameReplayContract *GameReplayContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GameReplayContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GameReplayContract *GameReplayContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GameReplayContract.Contract.TransferOwnership(&_GameReplayContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GameReplayContract *GameReplayContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GameReplayContract.Contract.TransferOwnership(&_GameReplayContract.TransactOpts, newOwner)
}

// GameReplayContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GameReplayContract contract.
type GameReplayContractOwnershipTransferredIterator struct {
	Event *GameReplayContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GameReplayContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameReplayContractOwnershipTransferred)
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
		it.Event = new(GameReplayContractOwnershipTransferred)
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
func (it *GameReplayContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameReplayContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameReplayContractOwnershipTransferred represents a OwnershipTransferred event raised by the GameReplayContract contract.
type GameReplayContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GameReplayContract *GameReplayContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GameReplayContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GameReplayContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GameReplayContractOwnershipTransferredIterator{contract: _GameReplayContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GameReplayContract *GameReplayContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GameReplayContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GameReplayContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameReplayContractOwnershipTransferred)
				if err := _GameReplayContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_GameReplayContract *GameReplayContractFilterer) ParseOwnershipTransferred(log types.Log) (*GameReplayContractOwnershipTransferred, error) {
	event := new(GameReplayContractOwnershipTransferred)
	if err := _GameReplayContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
