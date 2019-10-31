package service

import (
	"fmt"
	"path/filepath"

	"github.com/glvd/conversion"
)

func initConversion(cfg ConversionConfig) (*conversion.Task, error) {
	eng, err := conversion.InitMySQL(
		conversion.LoginOption(cfg.Addr, cfg.Username, cfg.Password))
	if err != nil {
		return nil, fmt.Errorf("init conversion:%w", err)
	}

	conversion.CachePath = filepath.Join(DefaultPath, cfg.Cache)
	var node conversion.Node
	switch cfg.Mode {
	case "cluster":
		node = conversion.NewClusterNode(cfg.ClusterAddr)
	default:
		node = conversion.NewSingleNode(cfg.SingleAddr)
	}

	conversion.RegisterNode(node)
	conversion.RegisterCache()
	conversion.RegisterDatabase(eng)
	//conversion.SetNodeAddress(cfg.SingleAddr)
	//if err := conversion.ConnectToNode(); err != nil {
	//	return nil, fmt.Errorf("can't connect to ipfs node:%w", err)
	//}
	t := conversion.NewTask()
	t.Limit = cfg.Limit
	t.SetAutoStop(false)
	return t, nil
}
