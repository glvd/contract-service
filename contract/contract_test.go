package contract

import (
	"fmt"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/goextension/log"
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
	DefaultGatway = "http://139.196.215.224:8545"
	DefaultMessageAddress = "0xdfbff0f931cf056b9eb5a8b58616024881215f01"
	DefaultTagAddress = "0x431d6215052cd3f0cec1838289a3d99abc496f5e"
	DefaultNodeAddress = "0x39e427cf40f73d4e09e2addd224fab7bd2ddcefa"

	//var DefaultGasLimit = "0xB71B00"
	//DefaultGatway = "http://localhost:8545"
	//DefaultNodeAddress = "0xe512a2a61814b8de98a52a0dfd7e13627e40014a"
	//DefaultMessageAddress = "0xcD05aAD053f605EE1A76B0B78cDfdfd32bBcfB7b"
	//DefaultTagAddress = "0x696687f41be961A521dD331A1609Edf7E4E7b2b5"

	testContract = NewContract(ETHClient(DefaultGatway),
		FileKey("945d35cd4a6549213e8d37feb5d708ec98906902", "123"),
		//HexKey("9efef8ebc3c51e91fb7f9faf7dbd516cb320ade03108c1568c9cee01a39af311"),
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

// TestContract_GetVideo ...
func TestContract_GetVideo(t *testing.T) {
	messages, _, err := testContract.GetVideos("abp-825")
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

	e = testContract.AddVideo("abp874", "randomid", "{}", "v0.0.1")
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
