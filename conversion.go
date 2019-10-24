package service

import "github.com/glvd/conversion"

func initConversion() *Task {
	t := conversion.NewTask()
	t.Limit = 1
	t.SetAutoStop(false)
	return t
}
