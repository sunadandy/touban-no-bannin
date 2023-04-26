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
	router.PUT("/touban", controller.PutTouban)
	router.DELETE("/touban", controller.DeleteTouban)
	router.GET("/member", controller.GetMember)
	router.POST("/member", controller.PostMember)
	router.PUT("/member", controller.PutMember)
	router.DELETE("/member", controller.DeleteMember)

	router.Run(":3000")
}
