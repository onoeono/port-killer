package strategy

import (
	"port-killer/backend/dto"
	"runtime"
	"sync"
)

var (
	instance *PortUsageContext
	once     sync.Once
)

type PortUsageStrategy interface {
	GetPortUsage() ([]dto.PortUsageDTO, error)
}

type PortUsageContext struct {
	Strategy PortUsageStrategy
}

func newUsageContext() *PortUsageContext {
	var strategy PortUsageStrategy
	switch runtime.GOOS {
	case "windows":
		strategy = &WindowsPortUsageStrategy{}
	case "linux", "darwin":
		strategy = &LinuxPortUsageStrategy{}
	default:
		return nil
	}

	return &PortUsageContext{Strategy: strategy}
}

func GetUsageContext() *PortUsageContext {
	once.Do(func() {
		instance = newUsageContext()
	})
	return instance
}
