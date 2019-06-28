package contract

import (
	"testing"
)

// TestETH_CheckExist ...
func TestGetHostList(t *testing.T) {
	lists, e := getHotList("D:\\EventSummary.csv")
	for _, list := range lists {
		log.Infof("%+v", list)
	}
	t.Log(lists, e)
}
