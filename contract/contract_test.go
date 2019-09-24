package contract

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// TestETH_CheckExist ...
func TestGetHostList(t *testing.T) {
	lists, e := getHotList("D:\\EventSummary.csv")
	for _, list := range lists {
		log.Infof("%+v", list)
	}
	t.Log(lists, e)
}

func TestVideoInfo(t *testing.T) {
	conn, err := ethclient.Dial(defaultGatewayAddress)
	if err != nil {
		return
	}

	addr := common.HexToAddress("0xc15266BB7C1FC17d5aD614dfdbA7316038F35725")
	info, err := NewVideoInfo(addr, conn)
	if err != nil {
		return
	}

	struct0s, err := info.GetHotVideos(&bind.CallOpts{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(struct0s)

}
