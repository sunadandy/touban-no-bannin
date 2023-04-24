package controller

import (
	"net/http"
	"strconv"
	"touban/model"

	"github.com/gin-gonic/gin"
)

type AccessTouban interface {
	CreateTouban() error
	UpdateTouban() int64
}

func GetTouban(c *gin.Context) {
	jsonData, _ := model.ReadTouban()
	c.String(http.StatusOK, jsonData)
}

func PostTouban(c *gin.Context) {
	stTouban := model.NewTouban()
	var err error

	// POSTで受け取ったJSONを構造体にマッピング
	err = c.ShouldBindJSON(&stTouban)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	// データベースに追加
	err = AddTouban(stTouban)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": "Success!"})
}
func AddTouban(accessT AccessTouban) error {
	err := accessT.CreateTouban()
	return err
}

func PutTouban(c *gin.Context) {
	stTouban := model.NewTouban()

	// POSTで受け取ったJSONを構造体にマッピング
	err := c.ShouldBindJSON(&stTouban)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// レコード更新
	changedColumns := UpdateTouban(stTouban)
	if changedColumns != 0 {
		c.JSON(http.StatusOK, gin.H{"data": "Success!"})
	}
	// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}
func UpdateTouban(accessT AccessTouban) int64 {
	ret := accessT.UpdateTouban()
	return ret
}

func DeleteTouban(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		print("Invalid ID")
	}
	model.DeleteOrder(id)
}
