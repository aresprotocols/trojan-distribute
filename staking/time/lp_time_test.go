package lptime

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
	"trojan/distribute/log"
	"trojan/distribute/util"
	"trojan/distribute/wallet"
)

func TestTimeStake(t *testing.T) {
	url := "https://mainnet.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161"

	client, url := util.DialConn(url)

	filter, err := NewLptime(common.HexToAddress("0xa99d9fa06dd1827fd39ab2d6e0d8eb1dae9c4b93"), client)
	if err != nil {
		t.Fatalf("can't NewContract: %v", err)
	}

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
	var event *LptimeStaked
	count := 0
	addrs := make(map[common.Address]*big.Int)

	for i := 1; i < 10000; i++ {
		registeredIterator, err := filter.FilterStaked(opts, nil)
		fmt.Println("FilterStaked err", err, "opts", opts.Start, " end ", *opts.End)

		for registeredIterator.Next() {
			find = true
			event = registeredIterator.Event
			fmt.Println("name", event.User.String(), "Amount", wallet.ToEth(event.Amount), "height", event.Raw.BlockNumber, "tx", event.Raw.TxHash.String())
			addrs[event.User] = event.Amount
			count++
		}

		if !registeredIterator.Next() && find {
			opts.Start = *opts.End
			end := uint64(0)
			if opts.Start < 12500000 {
				end = opts.Start + step/10
			} else {
				end = opts.Start + step/10
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
	fmt.Println(addrs)
	fmt.Println("count", count, "use", len(addrs))

	count = 0
	vaild := make(map[common.Address]uint64)
	for address := range addrs {
		balance, err := filter.BalanceOf(nil, address)
		if err != nil {
			log.Error("Failed to retrieve token ", "name: %v", err)
		}
		value, _ := wallet.ToEth(balance).Uint64()
		if value > 0 {
			vaild[address] = value
			count++
			fmt.Println("addr balance address", address, " ", wallet.ToEth(balance), "count", count)
		}
	}
	fmt.Println("vaild", vaild)
	count = 0
	for address, u := range vaild {
		if u > 1000 {
			count++
			fmt.Println("addr balance address", address, " ", u, "count", count)
		}
	}
}
