package contract

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/goextension/log"
	"go.uber.org/zap"
)

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
}

// TestContract_AddNodes ...
func TestContract_AddNodes(t *testing.T) {
	c := NewContract(ETHClient(DefaultGatway), HexKey("8fce5cf38a3c20d35a59c6a025cb0fd8a3c1201a27bbb778d1899dd2f4692ca4"), Node(DefaultNodeAddress))
	tm := time.Now()
	//idx := tm.UnixNano()
	//key, e := ioutil.ReadFile("./test_key/rsa.crt")
	//if e != nil {
	//	t.Error(e)
	//}
	//_ = key
	//log.Info("sec", idx)
	//
	//enc := dhcrypto.NewCipherEncoder(key, int(idx), tm)
	//bytes, e := enc.Encode("/ip4/13.124.213.107/tcp/4001/ipfs/QmS9knxyQkdiFrGbzEb2FNqnq2yEzvBKx2pEGNBnXTWi7h")
	//if e != nil {
	//	t.Error(e)
	//}
	source := "/ip4/13.124.213.107/tcp/4001/ipfs/QmS9knxyQkdiFrGbzEb2FNqnq2yEzvBKx2pEGNBnXTWi7r"

	t.Log("source:", source)
	e := c.AddNodes(true, tm, source)
	if e != nil {
		t.Error(e)
	}
}

// TestContract_GetNodes ...
func TestContract_GetNodes(t *testing.T) {
	c := NewContract(ETHClient(DefaultGatway), HexKey("2ed78769ad77af7fa01734e5f3302f03e7e40b94c4bdb1abaa5f54615b9ea0b1"), Node(DefaultNodeAddress))
	key, e := ioutil.ReadFile("./test_key/private.pem")
	if e != nil {
		t.Error(e)
	}
	_ = key

	strings, i, e := c.GetNodes(time.Time{})
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
