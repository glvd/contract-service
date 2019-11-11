package contract

import (
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
	DefaultNodeAddress = "0x1bEE31E960E05ff8009E651033df6a851B7D0815"
	DefaultGatway = "http://139.196.215.224:8545"
	DefaultTagAddress = "0x8858038b42a788499efC1DCCFe34123f07288D6f"
	DefaultMessageAddress = "0x2429F978deE1D931dcFC15C2A40a5986002745Bd"

	testContract = NewContract(ETHClient(DefaultGatway),
		FileKey("945d35cd4a6549213e8d37feb5d708ec98906902", "123"),
		//HexKey("20841a230fc4ace1ff70f616aeba24af85648ba7914c9c03492d15be5082a827"),
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

// TestContract_GetVideo ...
func TestContract_GetVideo(t *testing.T) {
	messages, err := testContract.GetVideo("abp874")
	if err != nil {
		t.Log(err)
	}
	t.Log(messages)
}

// TestContract_AddVideo ...
func TestContract_AddVideo(t *testing.T) {
	//e := testContract.OpenMessageAuthority()
	//if e != nil {
	//	t.Log(e)
	//	return
	//}

	e := testContract.AddVideo("abp874", "randomid", "{}", "v0.0.1")
	if e != nil {
		t.Log(e)
	}
}

// TestContract_DeployMessage ...
func TestContract_DeployMessage(t *testing.T) {
	addr, e := testContract.DeployMessage()
	t.Log(addr, e)
}

// TestContract_DeployTag ...
func TestContract_DeployTag(t *testing.T) {
	msgAddr := common.HexToAddress(DefaultMessageAddress)
	addr, e := testContract.DeployTag(&msgAddr)
	t.Log(addr, e)
}
