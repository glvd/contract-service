package service

import "github.com/glvd/conversion"

func initConversion(cfg *Config) *conversion.Task {
	t := conversion.NewTask()
	t.Limit = cfg.ConversionLimit
	t.SetAutoStop(false)
	return t
}
