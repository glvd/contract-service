package contract

import (
	"fmt"
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

// TestContract_AddNodes ...
func TestContract_AddNodes(t *testing.T) {
	tm := time.Now()
	source := "/ip4/13.124.213.107/tcp/4001/ipfs/QmS9knxyQkdiFrGbzEb2FNqnq2yEzvBKx2pEGNBnXTWi7r"

	t.Log("source:", source)
	e := testContract.AddNodes(true, tm, source)
	if e != nil {
		t.Error(e)
	}
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

// TestContract_AddHot ...
func TestContract_AddHot(t *testing.T) {
	e := testContract.AddHot(
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
	)
	if e != nil {
		t.Fatal(e)
	}
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
