package main

import (
	"github.com/gin-gonic/gin"
	"port-killer/backend/service"
)

func main() {
	router := gin.Default()

	router.GET("/list", service.ListUsage)

	router.Run(":1123")
	//context := strategy.GetUsageContext()
	//if context == nil {
	//	panic("Error")
	//}

	//usage, err := context.Strategy.GetPortUsage()
	//if err != nil {
	//	panic(err)
	//}
	//marshal, err := json.Marshal(usage)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(marshal))
}
