package service

import (
	"fmt"

	"github.com/glvd/conversion"
)

func initConversion(cfg *Config) (*conversion.Task, error) {
	eng, err := conversion.InitMySQL(
		conversion.LoginOption(cfg.Database.Addr, cfg.Database.Username, cfg.Database.Password))
	if err != nil {
		return nil, fmt.Errorf("init conversion:%w", err)
	}
	conversion.RegisterDatabase(eng)

	t := conversion.NewTask()
	t.Limit = cfg.ConversionLimit
	t.SetAutoStop(false)
	return t, nil
}
