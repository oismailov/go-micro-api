package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/oismailov/go-micro-api/api"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/account/:number", api.GetAccount)
		v1.POST("/order", api.Order)
	}

	return r
}

func GetPort() string {
	return fmt.Sprintf(":%d", 8080)
}
