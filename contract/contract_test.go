package contract

import (
	"context"
	"os"
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
	infos, e := os.Open("D:\\ipfstest\\accnode\\UTC--2019-08-20T02-58-22.119842600Z--945d35cd4a6549213e8d37feb5d708ec98906902")
	if e != nil {
		t.Fatal(e)
	}
	conn, err := ethclient.Dial(defaultGatewayAddress)
	if err != nil {
		return
	}

	addr := common.HexToAddress("0xc15266BB7C1FC17d5aD614dfdbA7316038F35725")
	info, err := NewVideoInfo(addr, conn)
	if err != nil {
		t.Fatal(err)
	}
	opt, e := bind.NewTransactor(infos, "123")
	if e != nil {
		t.Fatal(e)
	}
	transaction, err := info.Add(opt, Struct0{
		FindNo:     "abp-723",
		Bangumi:    "abp-723",
		Intro:      "",
		Alias:      nil,
		Role:       nil,
		ThumbHash:  "thumbxxxxxxxxxxxxx",
		PosterHash: "posterxxxxxxxxxxx",
		Uncensored: false,
	}, Struct1{
		Bangumi:   "abp-723",
		Visit:     0,
		Series:    "",
		Tags:      nil,
		Length:    "",
		Sharpness: "",
		Format:    "",
	}, Struct2{
		Bangumi:    "abp-723",
		Key:        "",
		M3u8:       "",
		SourceHash: "sourcexxxxxxxxxx",
		M3u8Hash:   "m3u8xxxxxxxxx",
	})
	if err != nil {
		t.Fatal(err)
	}
	_, err = bind.WaitMined(context.Background(), conn, transaction)
	if err != nil {
		t.Fatal(err)
	}
	struct0s, err := info.GetVideoCount(&bind.CallOpts{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(struct0s)

}
