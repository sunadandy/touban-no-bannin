package main

import (
	"encoding/json"
	"strings"
	"time"
	"touban/controller"
	"touban/email"
	"touban/model"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	// データベースと接続
	model.DBConnect()

	router := gin.Default()
	router.Use(static.Serve("/touban-no-bannin", static.LocalFile("dist", false)))
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
		// 毎日23時にタスクが走るように設定(コンテナの手タイムゾーンがUTCなので9時間の誤差を考慮)
		now := time.Now()
		target := time.Date(now.Year(), now.Month(), now.Day(), 14, 50, 00, 0, time.Local)
		if target.Before(now) {
			target = target.AddDate(0, 0, 1)
		}
		d := target.Sub(now)
		timer := time.NewTimer(d)
		<-timer.C

		// 日付情報の更新
		IsDateUpdateNeeded()

		// メールの配信
		email.ActionRemind()
	}
}

// ------------------------------------------------------------------------------------
// メンバーテーブルの最終実施日更新
// ------------------------------------------------------------------------------------
func IsDateUpdateNeeded() {
	// 当番テーブル取得
	stToubans := model.NewToubans()
	jsonData, _ := model.ReadTouban()
	json.Unmarshal([]byte(jsonData), &stToubans)
	// 日付情報取得
	now := time.Now()
	today := now.Format("2006-01-02")

	for _, touban := range stToubans {
		// 当番IDと一致するメンバーテーブル取得
		stMembers := model.NewMembers()
		jsonData := model.ReadMemberByToubanId(touban.Id)
		json.Unmarshal([]byte(jsonData), &stMembers)

		// 次回実施日が今日の日付だったメンバーを探して最終実施日を更新
		for i, stMember := range stMembers {
			if stMember.Next == today {
				// 最終実施日を今日に変更
				stMember.Last = today
				// 次回実施日を設定
				date := CalcuNextDate(stMembers, touban.Scheduling)
				stMember.Next = date
				DateUpdateMember(stMember)
				// 次回実施日を当番の開始日に設定
				touban.Start = stMembers[i+1].Next
				DateUpdateTouban(touban)
				break
			}
		}
	}
}

func CalcuNextDate(members model.Members, shedule string) string {
	var date string
	var maxDate time.Time

	// MAX次回実施日を取得
	for _, member := range members {
		next, _ := time.Parse("2006-01-02", member.Next)
		if maxDate.IsZero() || next.After(maxDate) {
			maxDate = next
		}
	}
	// 次回実施日計算
	interval := strings.Split(shedule, "-")[0]
	if interval == "1" {
		date = maxDate.AddDate(0, 0, 7).Format("2006-01-02")
	} else if interval == "2" {
		date = maxDate.AddDate(0, 0, 14).Format("2006-01-02")
	} else if interval == "3" {
		date = maxDate.AddDate(0, 0, 21).Format("2006-01-02")
	} else if interval == "4" {
		date = maxDate.AddDate(0, 0, 28).Format("2006-01-02")
	} else if interval == "0" {
		// 金曜日かどうか判定
		if maxDate.Weekday() == time.Friday {
			date = maxDate.AddDate(0, 0, 3).Format("2006-01-02")
		} else {
			date = maxDate.AddDate(0, 0, 1).Format("2006-01-02")
		}
	} else {
		date = maxDate.Format("2006-01-02")
	}

	return date
}

func DateUpdateMember(accessM controller.AccessMember) int64 {
	return accessM.UpdateMember()
}
func DateUpdateTouban(accessM controller.AccessTouban) int64 {
	return accessM.UpdateTouban()
}
