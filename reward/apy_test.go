package reward

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/math"
	"testing"
	"trojan/distribute/wallet"
)

func TestAPY(t *testing.T) {
	reward := []int{5000, 3000, 1000, 500, 300, 100, 50, 30, 10}
	sum := 0
	for _, v := range reward {
		sum = sum + v
	}
	fmt.Println("sum ", sum)
	apy := 50
	// n * 1000 000 *0.5 / 365 = n * 5 * 10 + 9990
	fmt.Println(" daily count ", 100000*apy/100/365)
	fmt.Println(" people num ", sum/(100000*apy/100/365-5*10))

	reward = []int{2000, 1000, 800, 500, 300, 100, 50, 30, 10}
	sum = 0
	for _, v := range reward {
		sum = sum + v
	}
	fmt.Println("sum ", sum)
	apy = 50
	// n * 1000 000 *0.5 / 365 = n * 5 * 10 + 9990
	fmt.Println(" daily count ", 100000*apy/100/365)
	fmt.Println(" people apy ", sum/(100000*apy/100/365-5*10))
}

func TestNormalAPY(t *testing.T) {
	apy := 36
	daily36 := 100000 * apy / 100 / 365
	fmt.Println(" daily count ", daily36)
	apy = 48
	daily48 := 100000 * apy / 100 / 365
	fmt.Println(" daily count ", daily48)
	fmt.Println(" daily count ", daily36*6/10, " ", daily48*6/10)

	// 6 4
	v, _ := math.ParseBig256("300000000000000000000000")

	fmt.Println(" ", wallet.ToEth(v))
	//pool := 0

}
