package contract

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/goextension/log"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

var testContract *Contract

func init() {
	logger, e := zap.NewProduction(
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	)
	if e != nil {
		panic(e)
	}
	log.Register(logger.Sugar())
	DefaultGatway = "http://127.0.0.1:8545"
	//DefaultMessageAddress = "0x2bc8cdc205b187e90533e98bcda07bc375b99e5f"
	//DefaultTagAddress = "0xf85cbfd1234f7ba4c700223780b4e9b8aea47bee"
	//DefaultNodeAddress = "0x39e427cf40f73d4e09e2addd224fab7bd2ddcefa"

	//var DefaultGasLimit = "0xB71B00"
	//DefaultGatway = "http://localhost:8545"
	//DefaultNodeAddress = "0xe512a2a61814b8de98a52a0dfd7e13627e40014a"
	//DefaultMessageAddress = "0xb884d148051f92e5e3b30fe7fc2c90c28d3f03a5"
	//DefaultTagAddress = "0xdacd53b6476f2d6271d93f42eda736deafdd797f"

	testContract = NewContract(ETHClient(DefaultGatway),
		FileKey("54c0fa4a3d982656c51fe7dfbdcc21923a7678cb", "123"),
		//FileKey("945d35cd4a6549213e8d37feb5d708ec98906902", "123"),
		//HexKey("59c98f25f4ea886cd57dd33eda009c38c33b4aa2e124ffed868fc8c705e2fbae"),
		//HexKey("ffedf6fe27fc7c4e994d56e407778b1a6aebafc8e63c394aebad413f719fc257"),
		//HexKey("87d3724ca3eb89db138fa415c9edfffba4ceb5c71a09c9c3d4cdb08e03e3ee68"),
		Node(DefaultNodeAddress),
		Tag(DefaultTagAddress),
		Message(DefaultMessageAddress),
		HashCoin(DefaultDHCAddress),
	)
}

// TestContract_RemoveNodes ...
func TestContract_RemoveNodes(t *testing.T) {

}

// TestContract_GetNodes ...
func TestContract_GetNodes(t *testing.T) {

	strings, i, e := testContract.GetNodes(time.Time{})
	if e != nil {
		t.Error(e)
		return
	}
	t.Log("ts:", i.Int64())

	//tm := time.Unix(i.Int64(), 0)
	//
	//dec := dhcrypto.NewCipherDecode(key, tm)
	for _, s := range strings {
		t.Log("eth:", s)
	}
	//	bytes, e := dec.Decode(s)
	//	if e != nil {
	//		t.Log(e)
	//		continue
	//	}
	//	t.Log("decoded:", string(bytes))
	//}

}

// TestContract_GetHotVideos ...
func TestContract_GetHotVideos(t *testing.T) {
	v, i, e := testContract.GetTagVideos("hot", "20191122")
	if e != nil {
		t.Fatal(e)
	}
	log.Infow("videos", "v", v, "i", i)
}

// TestContract_AddTagVideos ...
func TestContract_AddTagVideos(t *testing.T) {
	v, i, e := testContract.GetTagVideos("video", "ABP-894")
	if e != nil {
		t.Fatal(e)
	}
	log.Infow("videos", "v", v, "i", i)
}

// TestContract_GetVideos ...
func TestContract_GetVideos(t *testing.T) {

}

// TestContract_GetVideo ...
func TestContract_GetVideo(t *testing.T) {
	messages, _, err := testContract.GetVideos("20191122")
	if err != nil {
		t.Log(err)
	}
	t.Log(messages)
}

// TestContract_AddVideo ...
func TestContract_AddVideo(t *testing.T) {
	e := testContract.OpenMessageAuthority()
	if e != nil {
		t.Log(e)
		return
	}
	m := VideoMessage{
		ID:      uuid.New().String(),
		Content: "{}",
		Version: "v0.0.1",
	}
	e = testContract.AddVideo("abp874", m)
	if e != nil {
		t.Log(e)
	}
}

// TestContract_OpenMessageAuthority ...
func TestContract_OpenMessageAuthority(t *testing.T) {
	e := testContract.OpenMessageAuthority()
	if e != nil {
		t.Log(e)
		return
	}
}

// TestContract_AddOrUpdateVideo ...
func TestContract_AddOrUpdateVideo(t *testing.T) {
	e := testContract.OpenMessageAuthority()
	if e != nil {
		t.Log(e)
		return
	}
	m := VideoMessage{
		ID:      uuid.New().String(),
		Content: "{}",
		Version: "v0.0.1",
	}
	e = testContract.AddOrUpdateVideo("abp874", m, true)
	if e != nil {
		t.Log(e)
	}
}

// TestContract_DeployMessage ...
func TestContract_DeployMessage(t *testing.T) {
	addr, e := testContract.DeployMessage()
	t.Log(fmt.Sprintf("%x", addr), e)
}

// TestContract_DeployTag ...
func TestContract_DeployTag(t *testing.T) {
	msgAddr := common.HexToAddress(DefaultMessageAddress)
	addr, e := testContract.DeployTag(&msgAddr)
	t.Log(fmt.Sprintf("%x", addr), e)
}

// TestContract_DeployNode ...
func TestContract_DeployNode(t *testing.T) {
	addr, e := testContract.DeployNode()
	t.Log(fmt.Sprintf("%x", addr), e)
}

// TestContract_AddHot ...
func TestContract_AddHot(t *testing.T) {
	type fields struct {
		contracts *sync.Map
		conn      *ethclient.Client
		key       *ecdsa.PrivateKey
		gasLimit  *big.Int
	}
	type args struct {
		no []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "addhost1",
			fields: fields{},
			args: args{
				no: []string{
					"ABP-825",
					"ABP-721",
					"ABP-811",
					"ABP-792",
					"ABP-666",
					"ABP-340",
					"ABP-784",
					"ABP-440",
					"ABP-361",
					"ABP-178",
					"ABP-171",
					"ABP-159",
					"ABP-145",
					"ABP-138",
					"ABP-119",
					"ABP-108",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := testContract.AddHot(tt.args.no...); (err != nil) != tt.wantErr {
				t.Errorf("AddHot() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestContract_AddEveryDay ...
func TestContract_AddEveryDay(t *testing.T) {
	type fields struct {
		contracts *sync.Map
		conn      *ethclient.Client
		key       *ecdsa.PrivateKey
		gasLimit  *big.Int
	}
	type args struct {
		no []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "everyday1",
			fields: fields{},
			args: args{
				no: []string{
					"ABP-825",
					"ABP-721",
					"ABP-811",
					"ABP-792",
					"ABP-666",
					"ABP-340",
					"ABP-784",
					"ABP-440",
					"ABP-361",
					"ABP-178",
					"ABP-171",
					"ABP-159",
					"ABP-145",
					"ABP-138",
					"ABP-119",
					"ABP-108",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := testContract.AddEveryDay(tt.args.no...); (err != nil) != tt.wantErr {
				t.Errorf("AddEveryDay() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestContract_AddNodes1 ...
func TestContract_AddNodes(t *testing.T) {
	type fields struct {
		contracts *sync.Map
		conn      *ethclient.Client
		key       *ecdsa.PrivateKey
		gasLimit  *big.Int
	}
	type args struct {
		copyOld bool
		ts      time.Time
		ss      []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "addnodes",
			fields: fields{},
			args: args{
				copyOld: false,
				ts:      time.Now(),
				ss: []string{
					"/ip4/47.101.169.94/tcp/14005/ipfs/QmQLowH9Jd1S3aTSL5QbKVAcaJtD7GHTr3GmyLjebvZ9Rq",
					"/ip4/47.101.169.94/tcp/14001/ipfs/QmXNZRTd54Zvarf4sswVvUUnpb4gPQNAhFViozVgG8uwri",
					"/ip4/47.101.169.94/tcp/14006/ipfs/QmaCidjpHqP2p71fTT6B1gGzeHxR5KFQxVRbpaGv9hGRoA",
					"/ip4/13.124.213.107/tcp/14001/ipfs/Qme6QHB4HSCFgWVpuCHpjQVgzPWVf6ahkX9xhk8VV7BTdH",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := testContract.AddNodes(tt.args.copyOld, tt.args.ts, tt.args.ss...); (err != nil) != tt.wantErr {
				t.Errorf("AddNodes() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestContract_GetTagVideos ...
func TestContract_GetTagVideos(t *testing.T) {
	type fields struct {
		contracts *sync.Map
		conn      *ethclient.Client
		key       *ecdsa.PrivateKey
		gasLimit  *big.Int
	}
	type args struct {
		tag  string
		date string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantV   []VideoMessage
		wantI   int64
		wantErr bool
	}{
		{
			name:    "",
			fields:  fields{},
			args:    args{},
			wantV:   nil,
			wantI:   16,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm, _ := time.Parse(TimeStampFormat, "20191202")

			gotV, gotI, err := testContract.GetTagVideos("everyday", tm.Format(TimeStampFormat))
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTagVideos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotI != tt.wantI {
				t.Errorf("GetTagVideos() gotI = %v, want %v", gotI, tt.wantI)
			}
			if gotV != nil {
				t.Logf("GetTagVideos() gotV = %v", gotV)
			}
		})
	}
}

// TestContract_SetVersion ...
func TestContract_SetVersion(t *testing.T) {
	type fields struct {
		contracts *sync.Map
		conn      *ethclient.Client
		key       *ecdsa.PrivateKey
		gasLimit  *big.Int
	}
	type args struct {
		device  int64
		version string
		hash    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "setversionandroid",
			fields: fields{},
			args: args{
				device:  1,
				version: "v0.1.5",
				hash:    "android",
			},
			wantErr: false,
		},
		{
			name:   "setversionios",
			fields: fields{},
			args: args{
				device:  2,
				version: "v0.1.5",
				hash:    "ios",
			},
			wantErr: false,
		},
		{
			name:   "setversionforerunner",
			fields: fields{},
			args: args{
				device:  3,
				version: "v0.1.5",
				hash:    "forerunner",
			},
			wantErr: false,
		},
		{
			name:   "setversiontv",
			fields: fields{},
			args: args{
				device:  4,
				version: "v0.1.5",
				hash:    "tv",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := testContract.SetVersion(tt.args.device, tt.args.version, tt.args.hash); (err != nil) != tt.wantErr {
				t.Errorf("SetVersion() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestContract_GetVersion ...
func TestContract_GetVersion(t *testing.T) {
	type fields struct {
		contracts *sync.Map
		conn      *ethclient.Client
		key       *ecdsa.PrivateKey
		gasLimit  *big.Int
	}
	type args struct {
		device int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantV   string
		wantS   string
		wantErr bool
	}{
		{
			name:   "getversionandroid",
			fields: fields{},
			args: args{
				device: 1,
			},
			wantV:   "v0.1.5",
			wantS:   "android",
			wantErr: false,
		},
		{
			name:   "getversionios",
			fields: fields{},
			args: args{
				device: 2,
			},
			wantV:   "v0.1.5",
			wantS:   "ios",
			wantErr: false,
		},
		{
			name:   "getversionforerunner",
			fields: fields{},
			args: args{
				device: 3,
			},
			wantV:   "v0.1.5",
			wantS:   "forerunner",
			wantErr: false,
		},
		{
			name:   "getversiontv",
			fields: fields{},
			args: args{
				device: 4,
			},
			wantV:   "v0.1.5",
			wantS:   "tv",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotV, gotS, err := testContract.GetVersion(tt.args.device)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotV != tt.wantV {
				t.Errorf("GetVersion() gotV = %v, want %v", gotV, tt.wantV)
			}
			if gotS != tt.wantS {
				t.Errorf("GetVersion() gotS = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

// TestContract_DeployDHC ...
func TestContract_DeployDHC(t *testing.T) {
	type fields struct {
		contracts *sync.Map
		conn      *ethclient.Client
		key       *ecdsa.PrivateKey
		gasLimit  *big.Int
	}
	tests := []struct {
		name     string
		fields   fields
		wantAddr *common.Address
		wantErr  bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := testContract.UnlockDo(func(contract *Contract) {
				gotAddr, err := testContract.DeployDHC("DHashCoin", "DHC", 2)
				if (err != nil) != tt.wantErr {
					t.Errorf("DeployDHC() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if gotAddr != nil {
					t.Logf("DeployDHC() gotAddr = 0x%x", gotAddr)
				}
			})
			if err != nil {
				t.Log(err)
			}
		})
	}
}

// TestContract_Mint ...
func TestContract_Mint(t *testing.T) {
	type fields struct {
		contracts *sync.Map
		conn      *ethclient.Client
		key       *ecdsa.PrivateKey
		gasLimit  *big.Int
	}
	type args struct {
		addr common.Address
		val  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "",
			fields: fields{},
			args: args{
				addr: common.HexToAddress("0x09e0bc7507d136d586f14bea5bb27822ad31e4e1"),
				val:  20000000,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := testContract
			if err := c.Mint(tt.args.addr, tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("Mint() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestContract_Transfer ...
func TestContract_Transfer(t *testing.T) {
	type fields struct {
		contracts *sync.Map
		conn      *ethclient.Client
		key       *ecdsa.PrivateKey
		gasLimit  *big.Int
	}
	type args struct {
		val int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "",
			fields: fields{},
			args: args{
				val: 100000,
			},
			wantErr: false,
		},
	}
	//DefaultTransferAddress = "0xbb84b28db94415a3c0fb2203efebe4b1d808f53c"
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := testContract
			if err := c.Transfer(common.HexToAddress("0xbb84b28db94415a3c0fb2203efebe4b1d808f53c"), tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("Transfer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestContract_GetBalance ...
func TestContract_GetBalance(t *testing.T) {
	type fields struct {
		contracts *sync.Map
		conn      *ethclient.Client
		key       *ecdsa.PrivateKey
		gasLimit  *big.Int
	}
	type args struct {
		address common.Address
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantB   int64
		wantErr bool
	}{
		{
			name:   "",
			fields: fields{},
			args: args{
				address: common.HexToAddress("0xbb84b28db94415a3c0fb2203efebe4b1d808f53c"),
			},
			wantB:   0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := testContract
			gotB, err := c.GetBalance(tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotB != tt.wantB {
				t.Errorf("GetBalance() gotB = %v, want %v", gotB, tt.wantB)
			}

		})
	}
}

// TestContract_Ethereum ...
func TestContract_Ethereum(t *testing.T) {
	type fields struct {
		contracts *sync.Map
		conn      *ethclient.Client
		key       *ecdsa.PrivateKey
		gasLimit  *big.Int
	}
	type args struct {
		address common.Address
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:   "",
			fields: fields{},
			args: args{
				address: common.HexToAddress("0x1174e90188a9eA9E8eEb57BCb57775E16b9116E4"),
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := testContract
			got, err := c.Ethereum(tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("Ethereum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Ethereum() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestContract_TransferEthereum ...
func TestContract_TransferEthereum(t *testing.T) {
	type fields struct {
		contracts *sync.Map
		conn      *ethclient.Client
		key       *ecdsa.PrivateKey
		gasLimit  *big.Int
	}
	type args struct {
		to  common.Address
		val int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "",
			fields: fields{},
			args: args{
				to:  common.HexToAddress("0x1174e90188a9eA9E8eEb57BCb57775E16b9116E4"),
				val: 100,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := testContract
			if err := c.TransferEthereum(tt.args.to, tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("TransferEthereum() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestNewAccount ...
func TestNewAccount(t *testing.T) {

	// Create an account
	key, err := crypto.GenerateKey()
	if err != nil {
		return
	}
	// Get the address
	address := crypto.PubkeyToAddress(key.PublicKey).Hex()
	// 0x8ee3333cDE801ceE9471ADf23370c48b011f82a6
	t.Log(address)
	// Get the private key
	privateKey := hex.EncodeToString(key.D.Bytes())
	t.Log(privateKey)
	// 05b14254a1d0c77a49eae3bdf080f926a2df17d8e2ebdf7af941ea001481e57f

}

// TestContract_Primary ...
func TestContract_Primary(t *testing.T) {
	type fields struct {
		contracts *sync.Map
		conn      *ethclient.Client
		key       *ecdsa.PrivateKey
		gasLimit  *big.Int
	}
	type args struct {
		addr common.Address
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "",
			fields: fields{},
			args: args{
				addr: common.HexToAddress("0x1174e90188a9eA9E8eEb57BCb57775E16b9116E4"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := testContract
			if err := c.Primary(tt.args.addr); (err != nil) != tt.wantErr {
				t.Errorf("Primary() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
