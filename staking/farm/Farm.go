// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package farm

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

// FarmMetaData contains all meta data concerning the Farm contract.
var FarmMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"earned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"endCalcReward\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"exit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getReward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initreward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastTimeRewardApplicable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastUpdateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"notifyRewardAmount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"periodFinish\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardPerTokenStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rtfCoin\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardDistribution\",\"type\":\"address\"}],\"name\":\"setRewardDistribution\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"setRewardStop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"starttime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewardPerTokenPaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isTrue\",\"type\":\"bool\"}],\"name\":\"withTRC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"y\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052733a220f351252089d385b29beca14e27f204c296a6000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550733a220f351252089d385b29beca14e27f204c296a600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506984595161401484a0000060065560006008556000600955348015620000d257600080fd5b5060405160208062002fea83398101806040526020811015620000f457600080fd5b8101908080519060200190929190505050620001156200022a60201b60201c565b600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600781905550620001f862278d006007546200023260201b6200214a1790919060201c565b6008819055506200021d62278d00600654620002bd60201b620021001790919060201c565b60098190555050620003de565b600033905090565b6000808284019050838110151515620002b3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b60006200030783836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f0000000000008152506200030f60201b60201c565b905092915050565b600080831182901515620003c1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b838110156200038557808201518184015260208101905062000368565b50505050905090810190601f168015620003b35780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385811515620003d057fe5b049050809150509392505050565b612bfc80620003ee6000396000f3fe608060405234801561001057600080fd5b50600436106101ce5760003560e01c80638b87634711610104578063a694fc3a116100a2578063df136d6511610071578063df136d65146106ab578063e9fad8ee146106c9578063ebe2b12b146106d3578063f2fde38b146106f1576101ce565b8063a694fc3a146105f7578063c27a871e14610625578063c8f33c911461066f578063cd3daf9d1461068d576101ce565b80638f32d59b116100de5780638f32d59b146104f3578063934acd5b146105155780639c907b581461058f578063a56dfe4a146105ad576101ce565b80638b876347146104335780638da588971461048b5780638da5cb5b146104a9576101ce565b80633c6b16ab11610171578063715018a61161014b578063715018a6146103e3578063776c54cb146103ed5780637b0a47ee146103f757806380faa57d14610415576101ce565b80633c6b16ab146103535780633d18b9121461038157806370a082311461038b576101ce565b80631087b3d7116101ad5780631087b3d7146102c757806318160ddd146102e95780631be05289146103075780632e1a7d4d14610325576101ce565b80628cc262146101d35780630700037d1461022b5780630d68b76114610283575b600080fd5b610215600480360360208110156101e957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610735565b6040518082815260200191505060405180910390f35b61026d6004803603602081101561024157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061081c565b6040518082815260200191505060405180910390f35b6102c56004803603602081101561029957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610834565b005b6102cf6108f4565b604051808215151515815260200191505060405180910390f35b6102f1610907565b6040518082815260200191505060405180910390f35b61030f610911565b6040518082815260200191505060405180910390f35b6103516004803603602081101561033b57600080fd5b8101908080359060200190929190505050610918565b005b61037f6004803603602081101561036957600080fd5b8101908080359060200190929190505050610b4a565b005b610389610dcd565b005b6103cd600480360360208110156103a157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611360565b6040518082815260200191505060405180910390f35b6103eb6113a9565b005b6103f56114e6565b005b6103ff6115b2565b6040518082815260200191505060405180910390f35b61041d6115b8565b6040518082815260200191505060405180910390f35b6104756004803603602081101561044957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506115cb565b6040518082815260200191505060405180910390f35b6104936115e3565b6040518082815260200191505060405180910390f35b6104b16115e9565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6104fb611613565b604051808215151515815260200191505060405180910390f35b61058d6004803603608081101561052b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803515159060200190929190505050611672565b005b610597611c5c565b6040518082815260200191505060405180910390f35b6105b5611c62565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6106236004803603602081101561060d57600080fd5b8101908080359060200190929190505050611c87565b005b61062d611eb9565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610677611edf565b6040518082815260200191505060405180910390f35b610695611ee5565b6040518082815260200191505060405180910390f35b6106b3611f7d565b6040518082815260200191505060405180910390f35b6106d1611f83565b005b6106db611f9e565b6040518082815260200191505060405180910390f35b6107336004803603602081101561070757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611fa4565b005b6000610815600d60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610807670de0b6b3a76400006107f96107e2600c60008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546107d4611ee5565b61202c90919063ffffffff16565b6107eb88611360565b61207690919063ffffffff16565b61210090919063ffffffff16565b61214a90919063ffffffff16565b9050919050565b600d6020528060005260406000206000915090505481565b61083c611613565b15156108b0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b80600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600e60009054906101000a900460ff1681565b6000600154905090565b62278d0081565b33610921611ee5565b600b8190555061092f6115b8565b600a81905550600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415156109fe5761097481610735565b600d60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600b54600c60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b60075442111515610a77576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260098152602001807f6e6f74207374617274000000000000000000000000000000000000000000000081525060200191505060405180910390fd5b600082111515610aef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f43616e6e6f74207769746864726177203000000000000000000000000000000081525060200191505060405180910390fd5b610af8826121d4565b3373ffffffffffffffffffffffffffffffffffffffff167f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5836040518082815260200191505060405180910390a25050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16610b8b6122d3565b73ffffffffffffffffffffffffffffffffffffffff16141515610bf9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526021815260200180612b866021913960400191505060405180910390fd5b6000610c03611ee5565b600b81905550610c116115b8565b600a81905550600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515610ce057610c5681610735565b600d60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600b54600c60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b60085442101515610d0c57610d0162278d008361210090919063ffffffff16565b600981905550610d6f565b6000610d234260085461202c90919063ffffffff16565b90506000610d3c6009548361207690919063ffffffff16565b9050610d6662278d00610d58838761214a90919063ffffffff16565b61210090919063ffffffff16565b60098190555050505b42600a81905550610d8c62278d004261214a90919063ffffffff16565b6008819055507fde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d826040518082815260200191505060405180910390a15050565b33610dd6611ee5565b600b81905550610de46115b8565b600a81905550600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515610eb357610e2981610735565b600d60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600b54600c60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b60075442111515610f2c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260098152602001807f6e6f74207374617274000000000000000000000000000000000000000000000081525060200191505060405180910390fd5b600e60009054906101000a900460ff1615610f465761135d565b6000610f5133610735565b9050600081118015610f6557506008544211155b1561135b576000600d60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561105057600080fd5b505afa158015611064573d6000803e3d6000fd5b505050506040513d602081101561107a57600080fd5b8101908080519060200190929190505050111515611100576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f6765745265776172643a20746f74616c20525446206973207a65726f0000000081525060200191505060405180910390fd5b80600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b1580156111a057600080fd5b505afa1580156111b4573d6000803e3d6000fd5b505050506040513d60208110156111ca57600080fd5b81019080805190602001909291905050501115156112bf57600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561128157600080fd5b505afa158015611295573d6000803e3d6000fd5b505050506040513d60208110156112ab57600080fd5b810190808051906020019092919050505090505b61130c3382600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166122db9092919063ffffffff16565b3373ffffffffffffffffffffffffffffffffffffffff167fe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486826040518082815260200191505060405180910390a25b505b50565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6113b1611613565b1515611425576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166115276122d3565b73ffffffffffffffffffffffffffffffffffffffff16141515611595576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526021815260200180612b866021913960400191505060405180910390fd5b6001600e60006101000a81548160ff021916908315150217905550565b60095481565b60006115c6426008546123ac565b905090565b600c6020528060005260406000206000915090505481565b60075481565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166116566122d3565b73ffffffffffffffffffffffffffffffffffffffff1614905090565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166116b36122d3565b73ffffffffffffffffffffffffffffffffffffffff16141515611721576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526021815260200180612b866021913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16141515156117c6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f4450416464723a20746f6b656e41646472206973207a65726f0000000000000081525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415151561186b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f4450416464723a20726563697069656e74206973207a65726f0000000000000081525060200191505060405180910390fd5b801561194c57813073ffffffffffffffffffffffffffffffffffffffff163110151515611900576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f6e6f742065676c2062616c616e6365000000000000000000000000000000000081525060200191505060405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f19350505050158015611946573d6000803e3d6000fd5b50611c56565b6000849050828173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b1580156119cf57600080fd5b505afa1580156119e3573d6000803e3d6000fd5b505050506040513d60208110156119f957600080fd5b8101908080519060200190929190505050101515611ad9578073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb85856040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015611a9857600080fd5b505af1158015611aac573d6000803e3d6000fd5b505050506040513d6020811015611ac257600080fd5b810190808051906020019092919050505050611c54565b8073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb858373ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b158015611b7357600080fd5b505afa158015611b87573d6000803e3d6000fd5b505050506040513d6020811015611b9d57600080fd5b81019080805190602001909291905050506040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015611c1757600080fd5b505af1158015611c2b573d6000803e3d6000fd5b505050506040513d6020811015611c4157600080fd5b8101908080519060200190929190505050505b505b50505050565b60065481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b33611c90611ee5565b600b81905550611c9e6115b8565b600a81905550600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515611d6d57611ce381610735565b600d60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600b54600c60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b60075442111515611de6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260098152602001807f6e6f74207374617274000000000000000000000000000000000000000000000081525060200191505060405180910390fd5b600082111515611e5e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f43616e6e6f74207374616b65203000000000000000000000000000000000000081525060200191505060405180910390fd5b611e67826123c5565b3373ffffffffffffffffffffffffffffffffffffffff167f9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d836040518082815260200191505060405180910390a25050565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600a5481565b600080611ef0610907565b1415611f0057600b549050611f7a565b611f77611f66611f0e610907565b611f58670de0b6b3a7640000611f4a600954611f3c600a54611f2e6115b8565b61202c90919063ffffffff16565b61207690919063ffffffff16565b61207690919063ffffffff16565b61210090919063ffffffff16565b600b5461214a90919063ffffffff16565b90505b90565b600b5481565b611f94611f8f33611360565b610918565b611f9c610dcd565b565b60085481565b611fac611613565b1515612020576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b612029816124c6565b50565b600061206e83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525061260e565b905092915050565b60008083141561208957600090506120fa565b6000828402905082848281151561209c57fe5b041415156120f5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526021815260200180612b656021913960400191505060405180910390fd5b809150505b92915050565b600061214283836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f0000000000008152506126d0565b905092915050565b60008082840190508381101515156121ca576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b6121e98160015461202c90919063ffffffff16565b60018190555061224181600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461202c90919063ffffffff16565b600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506122d033826000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166122db9092919063ffffffff16565b50565b600033905090565b6123a7838473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb905060e01b8484604051602401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505061279a565b505050565b60008183106123bb57816123bd565b825b905092915050565b6123da8160015461214a90919063ffffffff16565b60018190555061243281600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461214a90919063ffffffff16565b600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506124c33330836000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166129ed909392919063ffffffff16565b50565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561254e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180612b3f6026913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600083831115829015156126bd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015612682578082015181840152602081019050612667565b50505050905090810190601f1680156126af5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385039050809150509392505050565b60008083118290151561277e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015612743578082015181840152602081019050612728565b50505050905090810190601f1680156127705780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b506000838581151561278c57fe5b049050809150509392505050565b6127b98273ffffffffffffffffffffffffffffffffffffffff16612af3565b151561282d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f5361666545524332303a2063616c6c20746f206e6f6e2d636f6e74726163740081525060200191505060405180910390fd5b600060608373ffffffffffffffffffffffffffffffffffffffff16836040518082805190602001908083835b60208310151561287e5780518252602082019150602081019050602083039250612859565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d80600081146128e0576040519150601f19603f3d011682016040523d82523d6000602084013e6128e5565b606091505b509150915081151561295f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656481525060200191505060405180910390fd5b6000815111156129e75780806020019051602081101561297e57600080fd5b810190808051906020019092919050505015156129e6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a815260200180612ba7602a913960400191505060405180910390fd5b5b50505050565b612aed848573ffffffffffffffffffffffffffffffffffffffff166323b872dd905060e01b858585604051602401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505061279a565b50505050565b60008060007fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47060001b9050833f91506000801b8214158015612b355750808214155b9250505091905056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f7743616c6c6572206973206e6f742072657761726420646973747269627574696f6e5361666545524332303a204552433230206f7065726174696f6e20646964206e6f742073756363656564a165627a7a723058205d354f2ce669cc7ce42350cc0c31f432e87765327a6ca5155f6afc73768ab9b60029",
}

// FarmABI is the input ABI used to generate the binding from.
// Deprecated: Use FarmMetaData.ABI instead.
var FarmABI = FarmMetaData.ABI

// FarmBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FarmMetaData.Bin instead.
var FarmBin = FarmMetaData.Bin

// DeployFarm deploys a new Ethereum contract, binding an instance of Farm to it.
func DeployFarm(auth *bind.TransactOpts, backend bind.ContractBackend, _startTime *big.Int) (common.Address, *types.Transaction, *Farm, error) {
	parsed, err := FarmMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FarmBin), backend, _startTime)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Farm{FarmCaller: FarmCaller{contract: contract}, FarmTransactor: FarmTransactor{contract: contract}, FarmFilterer: FarmFilterer{contract: contract}}, nil
}

// Farm is an auto generated Go binding around an Ethereum contract.
type Farm struct {
	FarmCaller     // Read-only binding to the contract
	FarmTransactor // Write-only binding to the contract
	FarmFilterer   // Log filterer for contract events
}

// FarmCaller is an auto generated read-only Go binding around an Ethereum contract.
type FarmCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FarmTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FarmTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FarmFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FarmFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FarmSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FarmSession struct {
	Contract     *Farm             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FarmCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FarmCallerSession struct {
	Contract *FarmCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// FarmTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FarmTransactorSession struct {
	Contract     *FarmTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FarmRaw is an auto generated low-level Go binding around an Ethereum contract.
type FarmRaw struct {
	Contract *Farm // Generic contract binding to access the raw methods on
}

// FarmCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FarmCallerRaw struct {
	Contract *FarmCaller // Generic read-only contract binding to access the raw methods on
}

// FarmTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FarmTransactorRaw struct {
	Contract *FarmTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFarm creates a new instance of Farm, bound to a specific deployed contract.
func NewFarm(address common.Address, backend bind.ContractBackend) (*Farm, error) {
	contract, err := bindFarm(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Farm{FarmCaller: FarmCaller{contract: contract}, FarmTransactor: FarmTransactor{contract: contract}, FarmFilterer: FarmFilterer{contract: contract}}, nil
}

// NewFarmCaller creates a new read-only instance of Farm, bound to a specific deployed contract.
func NewFarmCaller(address common.Address, caller bind.ContractCaller) (*FarmCaller, error) {
	contract, err := bindFarm(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FarmCaller{contract: contract}, nil
}

// NewFarmTransactor creates a new write-only instance of Farm, bound to a specific deployed contract.
func NewFarmTransactor(address common.Address, transactor bind.ContractTransactor) (*FarmTransactor, error) {
	contract, err := bindFarm(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FarmTransactor{contract: contract}, nil
}

// NewFarmFilterer creates a new log filterer instance of Farm, bound to a specific deployed contract.
func NewFarmFilterer(address common.Address, filterer bind.ContractFilterer) (*FarmFilterer, error) {
	contract, err := bindFarm(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FarmFilterer{contract: contract}, nil
}

// bindFarm binds a generic wrapper to an already deployed contract.
func bindFarm(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FarmABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Farm *FarmRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Farm.Contract.FarmCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Farm *FarmRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Farm.Contract.FarmTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Farm *FarmRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Farm.Contract.FarmTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Farm *FarmCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Farm.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Farm *FarmTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Farm.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Farm *FarmTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Farm.Contract.contract.Transact(opts, method, params...)
}

// DURATION is a free data retrieval call binding the contract method 0x1be05289.
//
// Solidity: function DURATION() view returns(uint256)
func (_Farm *FarmCaller) DURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Farm.contract.Call(opts, &out, "DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DURATION is a free data retrieval call binding the contract method 0x1be05289.
//
// Solidity: function DURATION() view returns(uint256)
func (_Farm *FarmSession) DURATION() (*big.Int, error) {
	return _Farm.Contract.DURATION(&_Farm.CallOpts)
}

// DURATION is a free data retrieval call binding the contract method 0x1be05289.
//
// Solidity: function DURATION() view returns(uint256)
func (_Farm *FarmCallerSession) DURATION() (*big.Int, error) {
	return _Farm.Contract.DURATION(&_Farm.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Farm *FarmCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Farm.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Farm *FarmSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Farm.Contract.BalanceOf(&_Farm.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Farm *FarmCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Farm.Contract.BalanceOf(&_Farm.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_Farm *FarmCaller) Earned(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Farm.contract.Call(opts, &out, "earned", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_Farm *FarmSession) Earned(account common.Address) (*big.Int, error) {
	return _Farm.Contract.Earned(&_Farm.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_Farm *FarmCallerSession) Earned(account common.Address) (*big.Int, error) {
	return _Farm.Contract.Earned(&_Farm.CallOpts, account)
}

// EndCalcReward is a free data retrieval call binding the contract method 0x1087b3d7.
//
// Solidity: function endCalcReward() view returns(bool)
func (_Farm *FarmCaller) EndCalcReward(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Farm.contract.Call(opts, &out, "endCalcReward")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EndCalcReward is a free data retrieval call binding the contract method 0x1087b3d7.
//
// Solidity: function endCalcReward() view returns(bool)
func (_Farm *FarmSession) EndCalcReward() (bool, error) {
	return _Farm.Contract.EndCalcReward(&_Farm.CallOpts)
}

// EndCalcReward is a free data retrieval call binding the contract method 0x1087b3d7.
//
// Solidity: function endCalcReward() view returns(bool)
func (_Farm *FarmCallerSession) EndCalcReward() (bool, error) {
	return _Farm.Contract.EndCalcReward(&_Farm.CallOpts)
}

// Initreward is a free data retrieval call binding the contract method 0x9c907b58.
//
// Solidity: function initreward() view returns(uint256)
func (_Farm *FarmCaller) Initreward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Farm.contract.Call(opts, &out, "initreward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Initreward is a free data retrieval call binding the contract method 0x9c907b58.
//
// Solidity: function initreward() view returns(uint256)
func (_Farm *FarmSession) Initreward() (*big.Int, error) {
	return _Farm.Contract.Initreward(&_Farm.CallOpts)
}

// Initreward is a free data retrieval call binding the contract method 0x9c907b58.
//
// Solidity: function initreward() view returns(uint256)
func (_Farm *FarmCallerSession) Initreward() (*big.Int, error) {
	return _Farm.Contract.Initreward(&_Farm.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Farm *FarmCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Farm.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Farm *FarmSession) IsOwner() (bool, error) {
	return _Farm.Contract.IsOwner(&_Farm.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Farm *FarmCallerSession) IsOwner() (bool, error) {
	return _Farm.Contract.IsOwner(&_Farm.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_Farm *FarmCaller) LastTimeRewardApplicable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Farm.contract.Call(opts, &out, "lastTimeRewardApplicable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_Farm *FarmSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _Farm.Contract.LastTimeRewardApplicable(&_Farm.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_Farm *FarmCallerSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _Farm.Contract.LastTimeRewardApplicable(&_Farm.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_Farm *FarmCaller) LastUpdateTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Farm.contract.Call(opts, &out, "lastUpdateTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_Farm *FarmSession) LastUpdateTime() (*big.Int, error) {
	return _Farm.Contract.LastUpdateTime(&_Farm.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_Farm *FarmCallerSession) LastUpdateTime() (*big.Int, error) {
	return _Farm.Contract.LastUpdateTime(&_Farm.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Farm *FarmCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Farm.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Farm *FarmSession) Owner() (common.Address, error) {
	return _Farm.Contract.Owner(&_Farm.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Farm *FarmCallerSession) Owner() (common.Address, error) {
	return _Farm.Contract.Owner(&_Farm.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_Farm *FarmCaller) PeriodFinish(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Farm.contract.Call(opts, &out, "periodFinish")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_Farm *FarmSession) PeriodFinish() (*big.Int, error) {
	return _Farm.Contract.PeriodFinish(&_Farm.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_Farm *FarmCallerSession) PeriodFinish() (*big.Int, error) {
	return _Farm.Contract.PeriodFinish(&_Farm.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_Farm *FarmCaller) RewardPerToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Farm.contract.Call(opts, &out, "rewardPerToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_Farm *FarmSession) RewardPerToken() (*big.Int, error) {
	return _Farm.Contract.RewardPerToken(&_Farm.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_Farm *FarmCallerSession) RewardPerToken() (*big.Int, error) {
	return _Farm.Contract.RewardPerToken(&_Farm.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_Farm *FarmCaller) RewardPerTokenStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Farm.contract.Call(opts, &out, "rewardPerTokenStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_Farm *FarmSession) RewardPerTokenStored() (*big.Int, error) {
	return _Farm.Contract.RewardPerTokenStored(&_Farm.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_Farm *FarmCallerSession) RewardPerTokenStored() (*big.Int, error) {
	return _Farm.Contract.RewardPerTokenStored(&_Farm.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_Farm *FarmCaller) RewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Farm.contract.Call(opts, &out, "rewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_Farm *FarmSession) RewardRate() (*big.Int, error) {
	return _Farm.Contract.RewardRate(&_Farm.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_Farm *FarmCallerSession) RewardRate() (*big.Int, error) {
	return _Farm.Contract.RewardRate(&_Farm.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_Farm *FarmCaller) Rewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Farm.contract.Call(opts, &out, "rewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_Farm *FarmSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _Farm.Contract.Rewards(&_Farm.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_Farm *FarmCallerSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _Farm.Contract.Rewards(&_Farm.CallOpts, arg0)
}

// RtfCoin is a free data retrieval call binding the contract method 0xc27a871e.
//
// Solidity: function rtfCoin() view returns(address)
func (_Farm *FarmCaller) RtfCoin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Farm.contract.Call(opts, &out, "rtfCoin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RtfCoin is a free data retrieval call binding the contract method 0xc27a871e.
//
// Solidity: function rtfCoin() view returns(address)
func (_Farm *FarmSession) RtfCoin() (common.Address, error) {
	return _Farm.Contract.RtfCoin(&_Farm.CallOpts)
}

// RtfCoin is a free data retrieval call binding the contract method 0xc27a871e.
//
// Solidity: function rtfCoin() view returns(address)
func (_Farm *FarmCallerSession) RtfCoin() (common.Address, error) {
	return _Farm.Contract.RtfCoin(&_Farm.CallOpts)
}

// Starttime is a free data retrieval call binding the contract method 0x8da58897.
//
// Solidity: function starttime() view returns(uint256)
func (_Farm *FarmCaller) Starttime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Farm.contract.Call(opts, &out, "starttime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Starttime is a free data retrieval call binding the contract method 0x8da58897.
//
// Solidity: function starttime() view returns(uint256)
func (_Farm *FarmSession) Starttime() (*big.Int, error) {
	return _Farm.Contract.Starttime(&_Farm.CallOpts)
}

// Starttime is a free data retrieval call binding the contract method 0x8da58897.
//
// Solidity: function starttime() view returns(uint256)
func (_Farm *FarmCallerSession) Starttime() (*big.Int, error) {
	return _Farm.Contract.Starttime(&_Farm.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Farm *FarmCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Farm.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Farm *FarmSession) TotalSupply() (*big.Int, error) {
	return _Farm.Contract.TotalSupply(&_Farm.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Farm *FarmCallerSession) TotalSupply() (*big.Int, error) {
	return _Farm.Contract.TotalSupply(&_Farm.CallOpts)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_Farm *FarmCaller) UserRewardPerTokenPaid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Farm.contract.Call(opts, &out, "userRewardPerTokenPaid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_Farm *FarmSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _Farm.Contract.UserRewardPerTokenPaid(&_Farm.CallOpts, arg0)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_Farm *FarmCallerSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _Farm.Contract.UserRewardPerTokenPaid(&_Farm.CallOpts, arg0)
}

// Y is a free data retrieval call binding the contract method 0xa56dfe4a.
//
// Solidity: function y() view returns(address)
func (_Farm *FarmCaller) Y(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Farm.contract.Call(opts, &out, "y")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Y is a free data retrieval call binding the contract method 0xa56dfe4a.
//
// Solidity: function y() view returns(address)
func (_Farm *FarmSession) Y() (common.Address, error) {
	return _Farm.Contract.Y(&_Farm.CallOpts)
}

// Y is a free data retrieval call binding the contract method 0xa56dfe4a.
//
// Solidity: function y() view returns(address)
func (_Farm *FarmCallerSession) Y() (common.Address, error) {
	return _Farm.Contract.Y(&_Farm.CallOpts)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_Farm *FarmTransactor) Exit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Farm.contract.Transact(opts, "exit")
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_Farm *FarmSession) Exit() (*types.Transaction, error) {
	return _Farm.Contract.Exit(&_Farm.TransactOpts)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_Farm *FarmTransactorSession) Exit() (*types.Transaction, error) {
	return _Farm.Contract.Exit(&_Farm.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_Farm *FarmTransactor) GetReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Farm.contract.Transact(opts, "getReward")
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_Farm *FarmSession) GetReward() (*types.Transaction, error) {
	return _Farm.Contract.GetReward(&_Farm.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_Farm *FarmTransactorSession) GetReward() (*types.Transaction, error) {
	return _Farm.Contract.GetReward(&_Farm.TransactOpts)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 reward) returns()
func (_Farm *FarmTransactor) NotifyRewardAmount(opts *bind.TransactOpts, reward *big.Int) (*types.Transaction, error) {
	return _Farm.contract.Transact(opts, "notifyRewardAmount", reward)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 reward) returns()
func (_Farm *FarmSession) NotifyRewardAmount(reward *big.Int) (*types.Transaction, error) {
	return _Farm.Contract.NotifyRewardAmount(&_Farm.TransactOpts, reward)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 reward) returns()
func (_Farm *FarmTransactorSession) NotifyRewardAmount(reward *big.Int) (*types.Transaction, error) {
	return _Farm.Contract.NotifyRewardAmount(&_Farm.TransactOpts, reward)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Farm *FarmTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Farm.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Farm *FarmSession) RenounceOwnership() (*types.Transaction, error) {
	return _Farm.Contract.RenounceOwnership(&_Farm.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Farm *FarmTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Farm.Contract.RenounceOwnership(&_Farm.TransactOpts)
}

// SetRewardDistribution is a paid mutator transaction binding the contract method 0x0d68b761.
//
// Solidity: function setRewardDistribution(address _rewardDistribution) returns()
func (_Farm *FarmTransactor) SetRewardDistribution(opts *bind.TransactOpts, _rewardDistribution common.Address) (*types.Transaction, error) {
	return _Farm.contract.Transact(opts, "setRewardDistribution", _rewardDistribution)
}

// SetRewardDistribution is a paid mutator transaction binding the contract method 0x0d68b761.
//
// Solidity: function setRewardDistribution(address _rewardDistribution) returns()
func (_Farm *FarmSession) SetRewardDistribution(_rewardDistribution common.Address) (*types.Transaction, error) {
	return _Farm.Contract.SetRewardDistribution(&_Farm.TransactOpts, _rewardDistribution)
}

// SetRewardDistribution is a paid mutator transaction binding the contract method 0x0d68b761.
//
// Solidity: function setRewardDistribution(address _rewardDistribution) returns()
func (_Farm *FarmTransactorSession) SetRewardDistribution(_rewardDistribution common.Address) (*types.Transaction, error) {
	return _Farm.Contract.SetRewardDistribution(&_Farm.TransactOpts, _rewardDistribution)
}

// SetRewardStop is a paid mutator transaction binding the contract method 0x776c54cb.
//
// Solidity: function setRewardStop() returns()
func (_Farm *FarmTransactor) SetRewardStop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Farm.contract.Transact(opts, "setRewardStop")
}

// SetRewardStop is a paid mutator transaction binding the contract method 0x776c54cb.
//
// Solidity: function setRewardStop() returns()
func (_Farm *FarmSession) SetRewardStop() (*types.Transaction, error) {
	return _Farm.Contract.SetRewardStop(&_Farm.TransactOpts)
}

// SetRewardStop is a paid mutator transaction binding the contract method 0x776c54cb.
//
// Solidity: function setRewardStop() returns()
func (_Farm *FarmTransactorSession) SetRewardStop() (*types.Transaction, error) {
	return _Farm.Contract.SetRewardStop(&_Farm.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Farm *FarmTransactor) Stake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Farm.contract.Transact(opts, "stake", amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Farm *FarmSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _Farm.Contract.Stake(&_Farm.TransactOpts, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Farm *FarmTransactorSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _Farm.Contract.Stake(&_Farm.TransactOpts, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Farm *FarmTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Farm.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Farm *FarmSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Farm.Contract.TransferOwnership(&_Farm.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Farm *FarmTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Farm.Contract.TransferOwnership(&_Farm.TransactOpts, newOwner)
}

// WithTRC20 is a paid mutator transaction binding the contract method 0x934acd5b.
//
// Solidity: function withTRC20(address tokenAddr, address recipient, uint256 amount, bool isTrue) returns()
func (_Farm *FarmTransactor) WithTRC20(opts *bind.TransactOpts, tokenAddr common.Address, recipient common.Address, amount *big.Int, isTrue bool) (*types.Transaction, error) {
	return _Farm.contract.Transact(opts, "withTRC20", tokenAddr, recipient, amount, isTrue)
}

// WithTRC20 is a paid mutator transaction binding the contract method 0x934acd5b.
//
// Solidity: function withTRC20(address tokenAddr, address recipient, uint256 amount, bool isTrue) returns()
func (_Farm *FarmSession) WithTRC20(tokenAddr common.Address, recipient common.Address, amount *big.Int, isTrue bool) (*types.Transaction, error) {
	return _Farm.Contract.WithTRC20(&_Farm.TransactOpts, tokenAddr, recipient, amount, isTrue)
}

// WithTRC20 is a paid mutator transaction binding the contract method 0x934acd5b.
//
// Solidity: function withTRC20(address tokenAddr, address recipient, uint256 amount, bool isTrue) returns()
func (_Farm *FarmTransactorSession) WithTRC20(tokenAddr common.Address, recipient common.Address, amount *big.Int, isTrue bool) (*types.Transaction, error) {
	return _Farm.Contract.WithTRC20(&_Farm.TransactOpts, tokenAddr, recipient, amount, isTrue)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Farm *FarmTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Farm.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Farm *FarmSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Farm.Contract.Withdraw(&_Farm.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Farm *FarmTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Farm.Contract.Withdraw(&_Farm.TransactOpts, amount)
}

// FarmOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Farm contract.
type FarmOwnershipTransferredIterator struct {
	Event *FarmOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FarmOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmOwnershipTransferred)
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
		it.Event = new(FarmOwnershipTransferred)
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
func (it *FarmOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmOwnershipTransferred represents a OwnershipTransferred event raised by the Farm contract.
type FarmOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Farm *FarmFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FarmOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Farm.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FarmOwnershipTransferredIterator{contract: _Farm.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Farm *FarmFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FarmOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Farm.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmOwnershipTransferred)
				if err := _Farm.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Farm *FarmFilterer) ParseOwnershipTransferred(log types.Log) (*FarmOwnershipTransferred, error) {
	event := new(FarmOwnershipTransferred)
	if err := _Farm.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmRewardAddedIterator is returned from FilterRewardAdded and is used to iterate over the raw logs and unpacked data for RewardAdded events raised by the Farm contract.
type FarmRewardAddedIterator struct {
	Event *FarmRewardAdded // Event containing the contract specifics and raw log

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
func (it *FarmRewardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmRewardAdded)
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
		it.Event = new(FarmRewardAdded)
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
func (it *FarmRewardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmRewardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmRewardAdded represents a RewardAdded event raised by the Farm contract.
type FarmRewardAdded struct {
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardAdded is a free log retrieval operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_Farm *FarmFilterer) FilterRewardAdded(opts *bind.FilterOpts) (*FarmRewardAddedIterator, error) {

	logs, sub, err := _Farm.contract.FilterLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return &FarmRewardAddedIterator{contract: _Farm.contract, event: "RewardAdded", logs: logs, sub: sub}, nil
}

// WatchRewardAdded is a free log subscription operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_Farm *FarmFilterer) WatchRewardAdded(opts *bind.WatchOpts, sink chan<- *FarmRewardAdded) (event.Subscription, error) {

	logs, sub, err := _Farm.contract.WatchLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmRewardAdded)
				if err := _Farm.contract.UnpackLog(event, "RewardAdded", log); err != nil {
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

// ParseRewardAdded is a log parse operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_Farm *FarmFilterer) ParseRewardAdded(log types.Log) (*FarmRewardAdded, error) {
	event := new(FarmRewardAdded)
	if err := _Farm.contract.UnpackLog(event, "RewardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmRewardPaidIterator is returned from FilterRewardPaid and is used to iterate over the raw logs and unpacked data for RewardPaid events raised by the Farm contract.
type FarmRewardPaidIterator struct {
	Event *FarmRewardPaid // Event containing the contract specifics and raw log

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
func (it *FarmRewardPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmRewardPaid)
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
		it.Event = new(FarmRewardPaid)
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
func (it *FarmRewardPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmRewardPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmRewardPaid represents a RewardPaid event raised by the Farm contract.
type FarmRewardPaid struct {
	User   common.Address
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardPaid is a free log retrieval operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_Farm *FarmFilterer) FilterRewardPaid(opts *bind.FilterOpts, user []common.Address) (*FarmRewardPaidIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Farm.contract.FilterLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return &FarmRewardPaidIterator{contract: _Farm.contract, event: "RewardPaid", logs: logs, sub: sub}, nil
}

// WatchRewardPaid is a free log subscription operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_Farm *FarmFilterer) WatchRewardPaid(opts *bind.WatchOpts, sink chan<- *FarmRewardPaid, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Farm.contract.WatchLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmRewardPaid)
				if err := _Farm.contract.UnpackLog(event, "RewardPaid", log); err != nil {
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

// ParseRewardPaid is a log parse operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_Farm *FarmFilterer) ParseRewardPaid(log types.Log) (*FarmRewardPaid, error) {
	event := new(FarmRewardPaid)
	if err := _Farm.contract.UnpackLog(event, "RewardPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Farm contract.
type FarmStakedIterator struct {
	Event *FarmStaked // Event containing the contract specifics and raw log

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
func (it *FarmStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmStaked)
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
		it.Event = new(FarmStaked)
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
func (it *FarmStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmStaked represents a Staked event raised by the Farm contract.
type FarmStaked struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_Farm *FarmFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*FarmStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Farm.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &FarmStakedIterator{contract: _Farm.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_Farm *FarmFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *FarmStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Farm.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmStaked)
				if err := _Farm.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_Farm *FarmFilterer) ParseStaked(log types.Log) (*FarmStaked, error) {
	event := new(FarmStaked)
	if err := _Farm.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Farm contract.
type FarmWithdrawnIterator struct {
	Event *FarmWithdrawn // Event containing the contract specifics and raw log

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
func (it *FarmWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmWithdrawn)
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
		it.Event = new(FarmWithdrawn)
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
func (it *FarmWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmWithdrawn represents a Withdrawn event raised by the Farm contract.
type FarmWithdrawn struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_Farm *FarmFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*FarmWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Farm.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &FarmWithdrawnIterator{contract: _Farm.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_Farm *FarmFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *FarmWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Farm.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmWithdrawn)
				if err := _Farm.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_Farm *FarmFilterer) ParseWithdrawn(log types.Log) (*FarmWithdrawn, error) {
	event := new(FarmWithdrawn)
	if err := _Farm.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
