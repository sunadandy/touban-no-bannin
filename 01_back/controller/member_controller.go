package controller

import (
	"net/http"
	"strconv"
	"touban/model"

	"github.com/gin-gonic/gin"
)

type AccessMember interface {
	CreateMember() error
	UpdateMember() error
}

func GetMember(c *gin.Context) {
	jsonData := model.ReadMember()
	c.String(http.StatusOK, jsonData)
}

func PostMember(c *gin.Context) {
	stMembers := model.NewMembers() //オーダー構造体を配列で取得
	var err error

	// POSTで受け取ったJSONを構造体にマッピング
	err = c.ShouldBindJSON(&stMembers)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		// データベースに追加
		for _, jsonData := range stMembers {
			err = AddMember(jsonData)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": `${jsonData}が正しく書き込めなかったレコードがあります。`})
			} else {
				c.JSON(http.StatusOK, gin.H{"data": "Success!"})
			}
		}
	}

}
func AddMember(accessO AccessMember) error {
	return accessO.CreateMember()
}

func PutMember(c *gin.Context) {
	print("TBD")
}

func DeleteMember(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		print("Invalid ID")
	} else {
		model.DeleteMember(id)
	}
}
