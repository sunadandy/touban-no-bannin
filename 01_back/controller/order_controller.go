package controller

import (
	"net/http"
	"strconv"
	"touban/model"

	"github.com/gin-gonic/gin"
)

type AccessOrder interface {
	CreateOrder() error
	UpdateOrder() error
}

func GetOrder(c *gin.Context) {
	// id := c.Query("id")
	// jsonData := model.ReadOrder(id)
	jsonData := model.ReadOrder()
	c.String(http.StatusOK, jsonData)
}

func PostOrder(c *gin.Context) {
	stOrders := model.NewOrders() //オーダー構造体を配列で取得
	var err error

	// POSTで受け取ったJSONを構造体にマッピング
	err = c.ShouldBindJSON(&stOrders)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// データベースに追加
	for _, jsonData := range stOrders {
		err = AddOrder(jsonData)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}
}
func AddOrder(accessO AccessOrder) error {
	err := accessO.CreateOrder()
	return err
}

func DeleteOrder(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		print("Invalid ID")
	}
	model.DeleteTouban(id)
}
