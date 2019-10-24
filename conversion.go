package service

import (
	"fmt"

	"github.com/glvd/conversion"
)

func initConversion(cfg *ConversionConfig) (*conversion.Task, error) {
	eng, err := conversion.InitMySQL(
		conversion.LoginOption(cfg.Addr, cfg.Username, cfg.Password))
	if err != nil {
		return nil, fmt.Errorf("init conversion:%w", err)
	}
	conversion.RegisterDatabase(eng)

	t := conversion.NewTask()
	t.Limit = cfg.Limit
	t.SetAutoStop(false)
	return t, nil
}
