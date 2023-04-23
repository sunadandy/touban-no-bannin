package main

import (
	"touban/controller"
	"touban/model"

	"github.com/gin-gonic/gin"
)

func main() {
	// データベースと接続
	model.DBConnect()

	router := gin.Default()
	router.GET("/touban", controller.GetTouban)
	router.POST("/touban", controller.PostTouban)
	router.GET("/order", controller.GetOrder)
	router.POST("/order", controller.PostOrder)
	router.Run(":3000")
}
