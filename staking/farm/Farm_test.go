package farm

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"testing"
	"trojan/distribute/log"
	"trojan/distribute/util"
	"trojan/distribute/wallet"
)

func TestReadFarm(t *testing.T) {
	url := "https://mainnet.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161"

	client, url := util.DialConn(url)

	farm, err := NewFarm(common.HexToAddress("0xdf1afbc5d532a607329b095e39a013eb672a4eb3"), client)
	if err != nil {
		t.Fatalf("can't NewContract: %v", err)
	}

	totalSupply, err := farm.TotalSupply(nil)
	if err != nil {
		log.Error("Failed to retrieve token ", "name: %v", err)
	}
	fmt.Println("totalSupply ", totalSupply, " ", wallet.ToEth(totalSupply))

	balance, err := farm.BalanceOf(nil, common.HexToAddress("0x21F2ccfD76897C58e0083A3Ab1bbD40A066d1516"))
	if err != nil {
		log.Error("Failed to retrieve token ", "name: %v", err)
	}
	fmt.Println("addr balance BalanceOf", balance, " ", wallet.ToEth(balance))
}
