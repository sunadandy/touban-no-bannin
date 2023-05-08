package main

import (
	"encoding/json"
	"time"
	"touban/controller"
	"touban/email"
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

	go DailyTask()
	router.Run(":3000")
}

func DailyTask() {
	for {
		// 毎日23時にタスクが走るように設定
		now := time.Now()
		target := time.Date(now.Year(), now.Month(), now.Day(), 23, 50, 0, 0, time.Local)
		if target.Before(now) {
			target = target.AddDate(0, 0, 1)
		}
		d := target.Sub(now)
		timer := time.NewTimer(d)
		<-timer.C

		// 最終実施日の更新
		UpdateLastDate()
		// メールの配信
		email.ActionRemind()
	}
}

// ------------------------------------------------------------------------------------
// メンバーテーブルの最終実施日更新
// ------------------------------------------------------------------------------------
func UpdateLastDate() {
	stMembers := model.NewMembers()
	// メンバーテーブルを取得
	jsonData := model.ReadMember()
	err := json.Unmarshal([]byte(jsonData), &stMembers)
	if err == nil {
		now := time.Now()
		today := now.Format("2006-01-02")
		for _, stMember := range stMembers {
			if stMember.Next == today {
				stMember.Last = today
				// メンバー情報を更新(戻り値は使わない)
				LastDateUpdate(stMember)
			}
		}
	} else {
		print(err.Error())
	}

}
func LastDateUpdate(accessM controller.AccessMember) int64 {
	return accessM.UpdateMember()
}
