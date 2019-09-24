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
	infos, e := os.Open("D:\\ipfstest\\accnode\\54c0fa4a3d982656c51fe7dfbdcc21923a7678cb")
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
		FindNo:     "abp-721",
		Bangumi:    "abp-721",
		Intro:      "",
		Alias:      nil,
		Role:       nil,
		ThumbHash:  "thumbxxxxxxxxxxxxx",
		PosterHash: "posterxxxxxxxxxxx",
		Uncensored: false,
	}, Struct1{
		Bangumi:   "abp-721",
		Visit:     0,
		Series:    "",
		Tags:      nil,
		Length:    "",
		Sharpness: "",
		Format:    "",
	}, Struct2{
		Bangumi:    "abp-721",
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
