package dto

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (e *Result) SetCode(code int) {
	e.Code = code
}
func (e *Result) SetMsg(msg string) {
	e.Msg = msg
}
func (e *Result) SetData(data any) {
	e.Data = data
}

func OK(c *gin.Context, data any, msg string) {
	result := &Result{
		Code: http.StatusOK,
		Msg:  msg,
		Data: data,
	}
	c.AbortWithStatusJSON(http.StatusOK, result)
}

func Fail(c *gin.Context, data any, msg string) {
	result := &Result{
		Code: http.StatusBadRequest,
		Msg:  msg,
		Data: data,
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, result)
}
