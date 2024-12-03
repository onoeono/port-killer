package service

import (
	"github.com/gin-gonic/gin"
	"port-killer/backend/dto"
	"port-killer/backend/strategy"
)

func ListUsage(ctx *gin.Context) {
	context := strategy.GetUsageContext()
	if context == nil || context.Strategy == nil {
		dto.OK(ctx, nil, "Error")
	}
	usage, err := context.Strategy.GetPortUsage()
	if err != nil {
		dto.OK(ctx, nil, "Error")
	}
	dto.OK(ctx, usage, "")
}
