package contract

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"reflect"
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
	//DefaultGatway = "http://127.0.0.1:8545"
	//DefaultMessageAddress = "0x2bc8cdc205b187e90533e98bcda07bc375b99e5f"
	//DefaultTagAddress = "0xf85cbfd1234f7ba4c700223780b4e9b8aea47bee"
	//DefaultNodeAddress = "0x39e427cf40f73d4e09e2addd224fab7bd2ddcefa"

	//var DefaultGasLimit = "0xB71B00"
	//DefaultGatway = "http://localhost:8545"
	//DefaultNodeAddress = "0xe512a2a61814b8de98a52a0dfd7e13627e40014a"
	//DefaultMessageAddress = "0xb884d148051f92e5e3b30fe7fc2c90c28d3f03a5"
	//DefaultTagAddress = "0xdacd53b6476f2d6271d93f42eda736deafdd797f"

	testContract = NewContract(ETHClient(DefaultGatway),
		FileKey("945d35cd4a6549213e8d37feb5d708ec98906902", "123"),
		//HexKey("87d3724ca3eb89db138fa415c9edfffba4ceb5c71a09c9c3d4cdb08e03e3ee68"),
		Node(DefaultNodeAddress),
		Tag(DefaultTagAddress),
		Message(DefaultMessageAddress))

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
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm, _ := time.Parse(TimeStampFormat, "20191202")

			gotV, gotI, err := testContract.GetTagVideos("everyday", tm.Format(TimeStampFormat))
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTagVideos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("GetTagVideos() gotV = %v, want %v", gotV, tt.wantV)
			}
			if gotI != tt.wantI {
				t.Errorf("GetTagVideos() gotI = %v, want %v", gotI, tt.wantI)
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
