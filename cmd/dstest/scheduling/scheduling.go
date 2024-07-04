package scheduling

import (
	"github.com/egeberkaygulcan/dstest/cmd/dstest/config"
	"github.com/egeberkaygulcan/dstest/cmd/dstest/network"
)

type Scheduler interface {
	Init(config *config.Config)
	Reset()
	Shutdown()
	NextIteration()
	Next([]*network.Message, int) int
	GetClientRequest() int
}

type SchedulerType string

const (
	Random SchedulerType = "random"
	Pctcp  SchedulerType = "pctcp"
)

func NewScheduler(schedulerType SchedulerType) Scheduler {
	switch schedulerType {
	case Random:
		return new(RandomScheduler)
	default:
		return new(RandomScheduler)
	}
}
