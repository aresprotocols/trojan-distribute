package farm

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
	"trojan/distribute/backends"
	"trojan/distribute/erc20"
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

func TestFarmStake(t *testing.T) {
	url := "https://mainnet.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161"

	client, url := util.DialConn(url)

	filter, err := NewFarm(common.HexToAddress("0xdf1afbc5d532a607329b095e39a013eb672a4eb3"), client)
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
	var event *FarmStaked
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

func TestParseData(t *testing.T) {
	data := "0x003FF073131fa8B77679425589ab8b2b9d6FE688:9366 0x0108aEB343D74a90299b7e526C629ef6Be858Df4:1087 0x037B5c31c607D573608257Cfd52aeA7a721bd82b:8429 0x051Abc20C0440aF42181C1dC9cA4aBFC74C07fE0:2094 0x0Aa350FB487F75d4f5d6ED920A5ae0923ded06E1:20460 0x0C0cA05D83427e722f412F94a56EB945258A9199:63000 0x0Db33CA4E5A34Bc9f186e284AE135C0BB9F34832:745 0x0E1a852E56a26A849E98cC19150ac565572E39Db:941 0x0f67754C0B0ec2F9883832Ec2E7FFD80d6Cb2CE8:400 0x10D92A4B5513BF766245e48E95EE1183F877bA11:7762 0x10e1AaaA29C46ee2b22e6f8aBa90D1f0399811C2:114739 0x114C86aB12413383cEfd71b7c41ca1C371B30878:4633 0x13F79E5c039F5dC19565a11984BC630DF9d6b75c:1000 0x143C72647E99E523EeED2Dd2e09A32D775110A45:380584 0x16286102d787dC6e8C0C71F2f14F581014DE5fbc:86500 0x16d1663A00d4d1a216E0baa84B0AbC69ba35C156:10000 0x1bb37bdAf2cBcD65E6f185e02f7541EB40706d30:1808115 0x1f5eAdcFf318f2F4BDE9Ec4524C1EEfbbD5Dc3d9:110000 0x1F61746155A9d260D0b8132fF37E0b0b7A0A5aC6:102075 0x2130D266a7012d09df3663A3eAcE40434e50838B:20000 0x21A095fD7Fc95218ca5162878a925d46434F64D0:66091 0x21F2ccfD76897C58e0083A3Ab1bbD40A066d1516:6000000 0x2231bf9D7A7239e0860c1824f79CE94C2A3bB5D0:8339 0x236A40Bc5B6e4594193AF0c67339fE1c1430Df06:9632858 0x2539bcCD727A7A7bE9631eBe37155E5b9F987290:95069 0x266f2f26ec15c43338CfeEf7B96b026f1a8a0894:59794 0x267c3Ca6e21d037D9B311796bd7c5D7B9ddb1AF0:44964 0x26B7B95e6F24F48E01357F84DF492F732A7bc484:102033 0x2A0a5e55ab62BBeBf89CCCcED53DAE7F2439CB04:1466 0x2A6339998D7DE47c5559a039befb4bC1B0692A60:800 0x2D368e72c59E6357aAB26D84Acc5834e6fB53247:2063 0x2eB69E7feB11Cc5f8c875F5e6030E3C38ACeD6CB:8234 0x2EE9711a53eFc23d61883CB7A68b3ccc072A9309:6155 0x33f5446aD8D779735b80578BBBc59CB3E95E0a53:958 0x37227Fd3351F34ccC647B95e43b6334Bb56b2BDE:60353 0x37727b76631Dba594358cF63EAE066355841D231:583 0x3CFc606ac23d8AD919D0322aBf206Bd304C81f88:4872 0x41c3683033046Da53Af82E19388af1b6555fe761:72855 0x43f46BC782Aa13f73264C17b449097ff60B4D759:50040 0x44527F99014eD8aD5c3D962F71EA3C24b49A3Af0:20680 0x4556106cf66203163b1E6D290a35Ffe59b0C6FD2:81003 0x45c15e4C3Bc7A9d14C4035C83B05436027e81268:1188 0x466E7Dd6Cdf7a8e57962299B468C2efDA3ee16C4:68442 0x48Ed758E42767d64b22723cbddcb0D6E263128b4:26478 0x494D443D2bC8a273Ab4D7f58F63e03BB665A382F:6198 0x4A66FdC90311745027e0E60C94204D44CFF68568:3639 0x4D1F51173167aFc7Cc3F8E6690aebA878FbAdea4:117800 0x4D6d64fd35927a265Bc279A475127Af8616b33a2:50000 0x4f9bfCaaa5Ccfe9ae6D7F54377Eb978C008Aaa9c:10075 0x4fD2eA6fc04F90eC6761a15934ddD58FbE95959E:18376 0x4FFbF8b517D5FDc9BbEc3B523C746245E1010EF1:1000 0x518e09FB0204008d5B503D55e23855D3B2668779:43742 0x523F63D157E8455d42b1a9895f62a04Cf466A56e:51000 0x5370B3cA55bb561e7bC050C81848873b9bcebBec:2500 0x5466FE082584f39260E5Ba7F060291Ee61a6bd83:50086 0x54Fe93dD1498F2F76b875EF69a1838AEdDaa318c:72680 0x55589E6cB1d8f48DEE149993754ED379dED2AFf0:3422 0x56373aec74a28117BA5bD85cca8bfCec515453f0:7985 0x57e713Eeeccf6C4c9f93CD8a7E1921858f856e0a:12106 0x5826a5FF97d639fBC9acF877F5545De86266cd0D:6666 0x5b2B8b42C06b9E66A566b3B81Baec13eC6cC43b5:235000 0x5c2b578070b6B9C07Ed94885B33420C8fa2F3a72:29725 0x5cF25606315b277eF7f8E9Df721d67b1607E7A08:2552 0x5D8b736D66fAf4C8602c599Ec1Fcd025c078cC9b:6666 0x5e42aa91DdeD9B93550545B70Bd055966e26211d:1117 0x600082e2F9C808552Be6e21F50224334b82684e8:114415 0x60171E620cBe3DDa2C3B934af377b54acDAa84B6:50000 0x60421Ee88CD7272Def1c4393D05e6c7C6e1E975D:51869 0x62Ab94C14A71E50beB8BDF58D57EBF0326BFd061:9098 0x6391aFD761f039D5459CAc133c27b94511f8C756:7667 0x6458f999E595D3bDA5692db9df011CCD4ab13a96:11666 0x6476B5fdf0f69B03266D0fA31eF5057E52D19A0A:197952 0x69e0aF70893f565A598636bAcA2aE332b8Dd3303:20000 0x6bCb2AcEF180dA4CbF10C5d706a1c091E9e9A42A:6382 0x6c3D26ba8a35167A67f45DF955B68E12399E9D8f:284829 0x6E9010b133B81E77800b57a1aAFF2D26390C9d7B:115986 0x72608c8c712196098FbA0cdE8389dFEF4c578AE9:2343 0x73cBF2c52820612E7960fD130160a1375a1793Dd:50075 0x74948Ed29584f962682f655b5a10a2c7aADD5CC3:2800 0x7681Ed0Bbbcd36b55d98D911F97ccc0AD2EB2174:100 0x7700A4D2CeC52b1eD54f48775F490B1738fb49fd:50000 0x778B18D072843adCB58B6937FfDBf7A6509219Aa:50000 0x7792F8CAeE6C041c937F8138e1D91b001c71cb07:2220 0x7B521c0141401487Bb4B95862f5efF9551316411:55047 0x7bc0CEAfffC2F2322F07c3058FfBcC2fA7C2Af7d:61905 0x7C3815d1fa89ff3640C66d2637895e1b39662b2D:5740 0x7c9b2F97125Dbc08ff6211e0085086E01F149dF4:37046 0x7F23aA8dE141108Dc6a7d4eB81E6ce7dDEB8f897:24226 0x897cc2c2bCCc26AC77a34d5F1DeE54C66d43a3d5:2250 0x89a9171fDCdFb7A557A04a1D3623E118527990e7:6023 0x8B2F481dEA6B92c47eA214Fa767ea4b6b302DD64:28564 0x8D4b951B581dB6010BDAa7Fdf14c1aa01d560EA8:2511 0x8DA57b0407C640964332af6f8Bc1939b4Cc5FA5E:716 0x8fA1cEbE79DAD49E51D0224836eB9915E8a75990:1884 0x8fB27a724E1Cf7C8C9772117F1e050FA64a11DF7:1162 0x9032B3bC5C109eB973Ead776b36038f1f5fe1aDe:16400 0x90c75B28aA5ccB93EFFaCb931194d839182CaC7a:429 0x911f99630a801bf0B78e124fE6A1A394c60A0f27:150000 0x91Fc3295Af66F005B754e0cfba3a8f7424f3E801:4281 0x9dE5DE4C87cE3675732F6b2441c2BCFA9bD5aEEc:30000 0xa0628DF86Fee0FDB325b813c9c3Dd8f0eC67E002:6895 0xa29523ac584B6BB79C710455E8Dd8356f9463FFf:3160 0xa45CF0B416067cBAC4cd76e71029F53c67520d87:74800 0xa513754111849Ca342b9CB0FB3F741BDcE9CF8BC:2529 0xA720aE34a450691FB29e748B07d7b4D584415C9a:3333 0xA9C370D45337Ca76947715Ba1814a4f77E69285B:64831 0xa9E007C34457326458913Bc44a27f0Ec29106702:60000 0xaBAa2f908e1821d830658CfAb616b5Ad903F6af8:919 0xaCb5e79188C8FB1d34f9cE5945aC395fdEa42D97:117945 0xaEe06F3B0ed1F0d595F7E391219e37ec0e391111:100000 0xb32826fA883317649b3FCAb5EAA84161F11eEb9B:51520 0xB37143F867698DBF9f411551f26A091DF0fc3AbE:12189 0xB3f9abE9D41368596201c605d439A5C7fdeeD094:11666 0xb45d7Bc95B1463376273F0883738957BDCB0B2Ba:5282 0xB4df1bD990F2dC8BAaE1dE792D70fBCae795E34E:12557 0xB7767355aa657b3Ed66f0de24dA5018a5E2862bd:16190 0xb78faaB3FC6AB3691aAddF94847cA53481e6410E:11666 0xbc12ae8ba54CC4EC64DaDd2CD18c0c8a55DFdFC4:50920 0xbcCA7b4365aF4B5d56449EEca67110986B47C0FD:10121 0xBd6Da64233cC2bC107bb18dE8DA0c719BaAbcEEE:308288 0xc226F82A48dA95C77072e1E9e142569CA758d495:311926 0xc3812Ae9aFD34b26b7A8AeeEF8798D770505F169:6177 0xc53D829c7cD9fccdF1c3b8B8ABe382fbFc12134d:1800 0xc9920899824FFecE2895Dd3Bb99528347143b1F5:2965 0xCA23572e679239801D2B24Ee142c3C5F0c076801:500 0xce5D647821447cf3Ae3B15fC8F90933094e4D57F:2552 0xceC5BcFc29a42eBAdc4aD2C905d9f74Fb105c2B3:293 0xd1989596De5159Efc28fBeE7061b4fD54b0dc711:4845 0xD23535DdD918EB2E4Ab8C07Eb1AC427147bAAB2F:3333 0xd5397296D45efCeC6395F3717305e921aa2F6694:5139 0xd5B3a3C50eF5cc209109A40039C3f310Fb9e8a4C:913 0xD5fd261bF6672106a9263e15425a0B1965E04378:99112 0xD6fE465FB8bE063aa6b63A68b468aa3Bf0eCB145:12264 0xd70ce0E2c4612a2Ce9411be19029f4704221AB4d:2922 0xD88bd07d9eFe8767E32b7955c30730859c660F14:54913 0xda7a4db7385bC0c1459aD435796f7F3C93929421:172687 0xE0085397759265394566EF1AcaB11fF12a61afC0:185782 0xe153315F0F3ED52AFB62FB87443A5f16229814bd:59821 0xe1d18ae098FFB1AD301e0609180f155B329A710A:2259 0xE355AA2fD7f9A1f7f2b10C413fe38bB78fD3CFDb:58967 0xE69E35BAEF6725240Ad58aE49166161D914fe249:570 0xeA744eFc7c819c5915DA2eeD8677914c5b799739:94330 0xEe511C77205B9D5c12A9452Bcd6d85c522163F4B:253702 0xf0B84d97327b4a27D7FF1eF03fe9Bb37E554628d:1000 0xf462496d1096550528D3F0a0DDb9404f9fae52CC:2299 0xf4A5c796Ad24d8a73dD23c7d74817d997C59ad3a:11000 0xF4E288Fa9401E7D777a7FdA92776E75558dC1E76:6123303 0xFB481C5cFd4F2294B4AF25dcB4F636766463D301:3416 0xfbfd7A91BfCf5C52c7c439759Af184162e0Cd1CC:68614 0xFD7929221a5065581be2a3adf3DCaaCD6558E17c:52520 0xfFC6642e1Fd5a2eB6853DcaaC53bFEF845D4DAa8:4286"
	parts := strings.Split(data, " ")
	var users []uint64
	for i, part := range parts {
		user := strings.Split(part, ":")
		value, _ := strconv.ParseUint(user[1], 10, 32)
		fmt.Println("i", i, "part", part, "value", value, "user", user)
		users = append(users, value)
	}
	count1000 := 0
	count10000 := 0
	count100000 := 0
	count1000000 := 0
	for _, user := range users {
		if user > 1000 {
			count1000++
		}
		if user > 10000 {
			count10000++
		}
		if user > 100000 {
			count100000++
		}
		if user > 1000000 {
			count1000000++
		}
	}
	fmt.Println("count1000", count1000, "count10000", count10000, "count100000", count100000, "count1000000", count1000000)
}

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
	key2Opts, _ := bind.NewKeyedTransactorWithChainID(key2, big.NewInt(1337))
	// Deploy the ENS registry

	ensAddr, _, _, err := erc20.DeployToken(transactOpts, contractBackend)
	if err != nil {
		t.Fatalf("can't DeployContract: %v", err)
	}
	ares, err := erc20.NewToken(ensAddr, contractBackend)
	if err != nil {
		t.Fatalf("can't NewContract: %v", err)
	}
	fmt.Println("ensAddr", ensAddr.String())
	contractBackend.Commit()

	totalSupply, err := ares.TotalSupply(nil)
	if err != nil {
		log.Error("Failed to retrieve token ", "name: %v", err)
	}
	fmt.Println("totalSupply ", totalSupply)

	ares.Transfer(transactOpts, testAddr, wallet.EthToWei(100000))
	contractBackend.Commit()

	block, _ := contractBackend.BlockByNumber(context.Background(), nil)
	farmAddr, _, _, err := DeployFarm(transactOpts, contractBackend, new(big.Int).SetUint64(block.Time()))
	if err != nil {
		t.Fatalf("can't DeployContract: %v", err)
	}
	farm, err := NewFarm(farmAddr, contractBackend)
	if err != nil {
		t.Fatalf("can't NewContract: %v", err)
	}
	contractBackend.Commit()

	starttime, _ := farm.Starttime(nil)
	lastUpdateTime, _ := farm.LastUpdateTime(nil)

	block, _ = contractBackend.BlockByNumber(context.Background(), nil)
	fmt.Println("starttime", starttime, "block", block.Time(), "lastUpdateTime", lastUpdateTime)

	dua, _ := farm.DURATION(nil)
	initreward, _ := farm.Initreward(nil)
	fmt.Println("initreward", initreward, " ", wallet.ToEth(initreward))
	rewardPerTokenStored, _ := farm.RewardPerTokenStored(nil)
	fmt.Println("rewardPerTokenStored", rewardPerTokenStored, " ", wallet.ToEth(rewardPerTokenStored))

	rewardPerToken, _ := farm.RewardPerToken(nil)
	fmt.Println("dua", dua, " ", dua.Int64()/24/60/60, "rewardPerToken", rewardPerToken)

	ares.Approve(transactOpts, farmAddr, wallet.EthToWei(1000000))
	contractBackend.Commit()
	_, err = farm.Stake(transactOpts, wallet.EthToWei(1000000))
	contractBackend.Commit()
	for i := 0; i < 100; i++ {
		contractBackend.Commit()
	}

	totalSupply1, _ := farm.TotalSupply(nil)
	lastUpdateTime, _ = farm.LastUpdateTime(nil)
	lastTimeRewardApplicable, _ := farm.LastTimeRewardApplicable(nil)
	rewardRate, _ := farm.RewardRate(nil)
	rewardPerToken, _ = farm.RewardPerToken(nil)

	fmt.Println("totalSupply stake1", totalSupply1, " ", wallet.ToEth(totalSupply1), "err", err)
	fmt.Println("lastUpdateTime", lastUpdateTime, " lastTimeRewardApplicable ", lastTimeRewardApplicable, "err", err)

	rewardPerTokenStored, _ = farm.RewardPerTokenStored(nil)
	fmt.Println("rewardPerTokenStored", rewardPerTokenStored, " ", wallet.ToEth(rewardPerTokenStored), "rewardRate", rewardRate)
	earned, _ := farm.Earned(nil, addr)

	fmt.Println("earned", earned, " ", wallet.ToEth(earned), "rewardPerToken", rewardPerToken)
	rewardPerTokenStored, _ = farm.RewardPerTokenStored(nil)
	fmt.Println("rewardPerTokenStored", rewardPerTokenStored, " ", wallet.ToEth(rewardPerTokenStored))
	ares.Approve(key2Opts, farmAddr, wallet.EthToWei(100000))
	contractBackend.Commit()
	_, err = farm.Stake(key2Opts, wallet.EthToWei(100000))
	contractBackend.Commit()

	rewardRate, _ = farm.RewardRate(nil)
	rewardPerToken, _ = farm.RewardPerToken(nil)

	rewardPerTokenStored, _ = farm.RewardPerTokenStored(nil)
	fmt.Println("rewardPerTokenStored", rewardPerTokenStored, " ", wallet.ToEth(rewardPerTokenStored))
	earned, _ = farm.Earned(nil, testAddr)
	fmt.Println("earned", earned, " ", wallet.ToEth(earned), "rewardRate", rewardRate)
	rewardPerTokenStored, _ = farm.RewardPerTokenStored(nil)
	fmt.Println("rewardPerTokenStored", rewardPerTokenStored, " ", wallet.ToEth(rewardPerTokenStored), "rewardPerToken", rewardPerToken)
}

func TestTime(t *testing.T) {
	timeS := time.Now().Unix()
	fmt.Println(timeS)

	fmt.Println(wallet.EthToWei(500000))
}
