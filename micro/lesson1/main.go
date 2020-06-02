package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	reg := consul.NewRegistry(consul.Config(&api.Config{Address: "127.0.0.1:8500"}))
	server := web.NewService(
		web.Name("hello"),
		web.Address(":8080"),   // 设置地址
		web.Handler(routers()), // 将gin的router放入
		web.Registry(reg),
	)
	server.Run()
}

func routers() http.Handler {
	router := gin.Default()

	router.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"info": "api"})
	})
	router.GET("/news", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"info": "news"})
	})
	return router
}