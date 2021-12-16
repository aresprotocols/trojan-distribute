package erc20

import (
	"context"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"os"
	"strconv"
	"testing"
	"time"
	"trojan/distribute/wallet"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
)

func init() {
	log.Root().SetHandler(log.LvlFilterHandler(log.LvlWarn, log.StreamHandler(os.Stderr, log.TerminalFormat(false))))
}

var (
	key, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	addr   = crypto.PubkeyToAddress(key.PublicKey)

	key2, _  = crypto.HexToECDSA("8a1f9a8f95be41cd7ccb6168179afb4504aefe388d1e14474d32c45c72ce7b7a")
	key3, _  = crypto.HexToECDSA("49a7b37aa6f6645917e7b807e9d1c00d4fa71f18343b0d4122a4d2df64dd6fee")
	testAddr = crypto.PubkeyToAddress(key2.PublicKey)
	add3     = crypto.PubkeyToAddress(key3.PublicKey)
)

func TestErc20(t *testing.T) {
	contractBackend := backends.NewSimulatedBackend(core.GenesisAlloc{
		addr:     {Balance: big.NewInt(1000000000000000000)},
		testAddr: {Balance: big.NewInt(1000000000000000000)}},
		10000000000)
	transactOpts, _ := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	keyOpts, _ := bind.NewKeyedTransactorWithChainID(key2, big.NewInt(1337))
	// Deploy the ENS registry

	ensAddr, _, _, err := DeployToken(transactOpts, contractBackend)
	if err != nil {
		t.Fatalf("can't DeployContract: %v", err)
	}
	ens, err := NewToken(ensAddr, contractBackend)
	if err != nil {
		t.Fatalf("can't NewContract: %v", err)
	}

	contractBackend.Commit()

	// Set ourself as the owner of the name.
	name, err := ens.Name(nil)
	if err != nil {
		log.Error("Failed to retrieve token ", "name: %v", err)
	}
	fmt.Println("Token name:", name)

	// Set ourself as the owner of the name.
	symbol, err := ens.Symbol(nil)
	if err != nil {
		log.Error("Failed to retrieve token ", "name: %v", err)
	}
	fmt.Println("Token symbol:", symbol)

	totalSupply, err := ens.TotalSupply(nil)
	if err != nil {
		log.Error("Failed to retrieve token ", "name: %v", err)
	}
	fmt.Println("totalSupply ", totalSupply)

	balance, err := ens.BalanceOf(nil, addr)
	if err != nil {
		log.Error("Failed to retrieve token ", "name: %v", err)
	}
	fmt.Println("addr balance BalanceOf", balance)

	tx, err := ens.Transfer(transactOpts, testAddr, big.NewInt(50000))
	if err != nil {
		log.Error("Failed to request token transfer: %v", err)
	}
	fmt.Printf("Transfer pending: 0x%x\n", tx.Hash())
	contractBackend.Commit()

	balance, err = ens.BalanceOf(nil, testAddr)
	if err != nil {
		log.Error("Failed to retrieve token ", "name: %v", err)
	}
	fmt.Println("testAddr balance BalanceOf", balance)

	tx, err = ens.Approve(keyOpts, addr, big.NewInt(10000))
	if err != nil {
		log.Error("Failed to retrieve Approve ", "name: %v", err)
	}
	contractBackend.Commit()

	balance, err = ens.Allowance(nil, testAddr, addr)
	if err != nil {
		log.Error("Failed to retrieve token ", "name: %v", err)
	}
	fmt.Println("balance Allowance", balance)

	tx, err = ens.TransferFrom(transactOpts, testAddr, add3, big.NewInt(5000))
	if err != nil {
		log.Error("Failed to request token transfer: %v", err)
	}
	fmt.Printf("Transfer pending: 0x%x\n", tx.Hash())
	contractBackend.Commit()

	balance, err = ens.Allowance(nil, testAddr, addr)
	if err != nil {
		log.Error("Failed to retrieve token ", "name: %v", err)
	}
	fmt.Println("Allowance balance ", balance)

	balance, err = ens.BalanceOf(nil, testAddr)
	if err != nil {
		log.Error("Failed to retrieve token ", "name: %v", err)
	}
	fmt.Println("balance BalanceOf", balance)

	balance, err = ens.BalanceOf(nil, add3)
	if err != nil {
		log.Error("Failed to retrieve token ", "name: %v", err)
	}
	fmt.Println("balance BalanceOf", balance)
	contractBackend.Commit()

	balance, _ = contractBackend.BalanceAt(context.Background(), testAddr, nil)
	fmt.Println("balance", balance)
	keyOpts.Value = new(big.Int).SetUint64(10000000)
	tx, err = ens.Approve(keyOpts, addr, big.NewInt(10000))
	if err != nil {
		log.Error("Failed to retrieve ApproveOne ", "name: %v", err)
	}
	contractBackend.Commit()

	balance, _ = contractBackend.BalanceAt(context.Background(), testAddr, nil)
	fmt.Println("11balance", balance)
}

func TestReadErc20(t *testing.T) {
	url := "https://mainnet.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161"

	client, url := dialConn(url)

	ens, err := NewToken(common.HexToAddress("0x358AA737e033F34df7c54306960a38d09AaBd523"), client)
	if err != nil {
		t.Fatalf("can't NewContract: %v", err)
	}

	// Set ourself as the owner of the name.
	name, err := ens.Name(nil)
	if err != nil {
		log.Error("Failed to retrieve token ", "name: %v", err)
	}
	fmt.Println("Token name:", name)

	// Set ourself as the owner of the name.
	symbol, err := ens.Symbol(nil)
	if err != nil {
		log.Error("Failed to retrieve token ", "name: %v", err)
	}
	fmt.Println("Token symbol:", symbol)

	totalSupply, err := ens.TotalSupply(nil)
	if err != nil {
		log.Error("Failed to retrieve token ", "name: %v", err)
	}
	fmt.Println("totalSupply ", totalSupply)

	balance, err := ens.BalanceOf(nil, common.HexToAddress("0x7a646ee13eb104853c651e1d90d143acc9e72cdb"))
	if err != nil {
		log.Error("Failed to retrieve token ", "name: %v", err)
	}
	fmt.Println("addr balance BalanceOf", balance)
}

func TestStakingErc20(t *testing.T) {
	url := "https://mainnet.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161"

	client, url := dialConn(url)

	ens, err := NewToken(common.HexToAddress("0x358AA737e033F34df7c54306960a38d09AaBd523"), client)
	if err != nil {
		t.Fatalf("can't NewContract: %v", err)
	}
	arrs := []common.Address{
		common.HexToAddress("0xdf1afbc5d532a607329b095e39a013eb672a4eb3"),
		common.HexToAddress("0xa99d9fa06dd1827fd39ab2d6e0d8eb1dae9c4b93"),
		common.HexToAddress("0x4c4f6d9fae70236888c4d613199ea4419ada23e8"),
		common.HexToAddress("0xb31d8eba3f5e2d758b54544e4446b39f9cb769ea"),
	}
	total := new(big.Int)
	for _, arr := range arrs {
		balance, err := ens.BalanceOf(nil, arr)
		if err != nil {
			log.Error("Failed to retrieve token ", "name: %v", err)
		}
		total.Add(total, balance)
	}

	fmt.Println("addr balance BalanceOf", total, " ", wallet.ToEth(total))
}

func dialConn(url string) (*ethclient.Client, string) {
	ip := "165.227.99.131"
	port := 8545

	//url = "https://ethrpc.truescan.network"
	//url = "https://kovan.poa.network/"

	if url == "" {
		url = fmt.Sprintf("http://%s", fmt.Sprintf("%s:%d", ip, port))
	}
	// Create an IPC based RPC connection to a remote node
	// "http://39.100.97.129:8545"
	conn, err := ethclient.Dial(url)
	if err != nil {
		fmt.Printf("Failed to connect to the ethereum client: %v \n", err)
	}
	return conn, url
}

func TestReadEns(t *testing.T) {
	url := "https://mainnet.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161"

	client, url := dialConn(url)

	filter, err := NewMainFilterer(common.HexToAddress("0x283Af0B28c62C092C9727F1Ee09c02CA627EB7F5"), client)
	if err != nil {
		t.Fatalf("can't NewContract: %v", err)
	}

	xlsx1 := excelize.NewFile()
	index := xlsx1.NewSheet("ENS")
	xlsx1.SetCellStr("ENS", "A1", "Name")
	xlsx1.SetCellStr("ENS", "B1", "Label")
	xlsx1.SetCellStr("ENS", "C1", "Owner")
	xlsx1.SetCellStr("ENS", "D1", "Cost")
	xlsx1.SetCellStr("ENS", "E1", "Expires")

	number, err := client.BlockNumber(context.Background())
	if err != nil {
		t.Fatalf("can't BlockNumber: %v", err)
	}

	step := uint64(1000000)
	opts := &bind.FilterOpts{
		Start:   0,
		End:     &step,
		Context: context.Background(),
	}
	find := false
	var event *MainNameRegistered
	count := 1
	for i := 1; i < 10000; i++ {
		registeredIterator, err := filter.FilterNameRegistered(opts, nil, nil)
		fmt.Println("FilterNameRegistered err", err, "opts", opts.Start, " end ", *opts.End)

		for registeredIterator.Next() {
			find = true
			event = registeredIterator.Event
			fmt.Println("name", event.Name, "label", event.Cost, "time", ConvertTime(event.Expires.Uint64()), "height", event.Raw.BlockNumber, "tx", event.Raw.TxHash.String())

			arrs := [5]string{"A", "B", "C", "D", "E"}
			count++
			var enums = []string{event.Name, string(event.Label[:]), event.Owner.String(), wallet.ToEth(event.Cost).String(), ConvertTime(event.Expires.Uint64())}
			for ind, aplat := range arrs {
				axis := aplat + strconv.FormatInt(int64(count), 10)
				xlsx1.SetCellStr("ENS", axis, enums[ind])
			}
		}

		if !registeredIterator.Next() && find {
			opts.Start = *opts.End
			end := uint64(0)
			if opts.Start < 12500000 {
				end = opts.Start + step/10
			} else if opts.Start > 12500000 && opts.Start < 13580000 {
				end = opts.Start + step/100
			} else {
				end = opts.Start + step/1000
			}
			opts.End = &end
			fmt.Println("Next err", err, "opts", opts.Start, " end ", end)
		}
		if !find {
			opts.Start = opts.Start + step
			end := opts.Start + step
			opts.End = &end
			fmt.Println("find err", err, "opts", opts.Start, " end ", end)
		}

		if opts.Start > number {
			break
		}
	}

	xlsx1.SetActiveSheet(index)
	err = xlsx1.SaveAs("./ENS Select.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

func ConvertTime(utime uint64) string {
	format := time.Unix(int64(utime), 0).Format("2006-01-02 15:04:05")
	return format
}
