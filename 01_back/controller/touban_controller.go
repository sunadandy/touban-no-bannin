package controller

import (
	"net/http"
	"touban/model"

	"github.com/gin-gonic/gin"
)

type AccessTouban interface {
	CreateTouban() (int, error)
}

func GetTouban(c *gin.Context) {
	jsonData := model.ReadTouban()
	c.String(http.StatusOK, jsonData)
}

func PostTouban(c *gin.Context) {
	stTouban := model.NewTouban() //当番構造体を配列で取得
	var id int
	var err error

	err = c.ShouldBindJSON(&stTouban)
	// POSTで受け取ったJSONを構造体にマッピング
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	// データベースに追加
	id, err = AddTouban(stTouban)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}
func AddTouban(accessT AccessTouban) (int, error) {
	id, err := accessT.CreateTouban()
	return id, err
}
