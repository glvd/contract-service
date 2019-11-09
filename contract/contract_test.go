package contract

import (
	"testing"
	"time"

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
	DefaultNodeAddress = "0x1bEE31E960E05ff8009E651033df6a851B7D0815"
	DefaultGatway = "http://localhost:8545"
	DefaultTagAddress = "0xE6D0182BDB02641674cb2466A2e5A87049Bd5E54"
	DefaultMessageAddress = "0xbeED867a0B16E27C48A74bfee2728C115892eA8F"
	testContract = NewContract(ETHClient(DefaultGatway),
		HexKey("2ed78769ad77af7fa01734e5f3302f03e7e40b94c4bdb1abaa5f54615b9ea0b1"),
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

// TestContract_AddVideo ...
func TestContract_AddVideo(t *testing.T) {
	e := testContract.AddVideo("abp874", "randomid", "{}", "v0.0.1")
	if e != nil {
		t.Log(e)
	}

}
