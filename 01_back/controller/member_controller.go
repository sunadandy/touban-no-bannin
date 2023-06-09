package controller

import (
	"net/http"
	"strconv"
	"touban/model"

	"github.com/gin-gonic/gin"
)

type AccessMember interface {
	CreateMember() error
	UpdateMember() int64
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
func AddMember(accessM AccessMember) error {
	return accessM.CreateMember()
}

func PutMember(c *gin.Context) {
	stMembers := model.NewMembers() //オーダー構造体を配列で取得
	var err error

	// POSTで受け取ったJSONを構造体にマッピング
	err = c.ShouldBindJSON(&stMembers)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		// データベースに追加
		for _, jsonData := range stMembers {
			changedColumns := UpdateMember(jsonData)
			if changedColumns != 0 {
				c.JSON(http.StatusOK, gin.H{"data": "Success!"})
			}
			// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}
}
func UpdateMember(accessM AccessMember) int64 {
	return accessM.UpdateMember()
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
