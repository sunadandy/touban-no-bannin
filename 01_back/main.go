package main

import (
	"encoding/json"
	"strconv"
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
	router.Use(static.Serve("/", static.LocalFile("dist", false)))
	router.GET("/touban-no-bannin/touban", controller.GetTouban)
	router.POST("/touban-no-bannin/touban", controller.PostTouban)
	router.PUT("/touban-no-bannin/touban", controller.PutTouban)
	router.DELETE("/touban-no-bannin/touban", controller.DeleteTouban)
	router.GET("/touban-no-bannin/member", controller.GetMember)
	router.POST("/touban-no-bannin/member", controller.PostMember)
	router.PUT("/touban-no-bannin/member", controller.PutMember)
	router.DELETE("/touban-no-bannin/member", controller.DeleteMember)

	// リマインダーメール再送（未送信だった場合のみ）
	email.Reminder()

	go DailyTask()
	router.Run(":3000")
}

func DailyTask() {
	for {
		now := time.Now()
		// 毎日23時にタスクが走るように設定(コンテナのタイムゾーンがUTCなので9時間の誤差を考慮)
		target := time.Date(now.Year(), now.Month(), now.Day(), 14, 45, 00, 0, time.Local)
		// 目標時刻が現在よりも前（＝過去）の場合、1日加算して目標を翌日に設定
		if target.Before(now) {
			target = target.AddDate(0, 0, 1)
		}
		// 現在時刻と目標との差(duration)を取得
		d := target.Sub(now)
		// タイマーを設定して待機
		timer := time.NewTimer(d)
		<-timer.C

		// 日付情報の更新
		IsDateUpdateNeeded()

		// 1分待機してからメールの配信
		oneMinutes := time.Minute * 1
		timer = time.NewTimer(oneMinutes)
		<-timer.C

		// メールの配信
		email.Reminder()
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

		// 次回実施日が今日の日付だったメンバーを探して最終実施日と次回実施日を更新
		var latestNextDate time.Time // 最遅次回実施日
		var updatedMemberIdx int
		for i, stMember := range stMembers {
			// 次回実施日を格納して現在の最遅実施日と比較
			next, _ := time.Parse("2006-01-02", stMember.Next)
			if latestNextDate.IsZero() || next.After(latestNextDate) {
				latestNextDate = next
			}
			// 最終実施日を今日に変更
			if stMember.Next == today {
				stMember.Last = today
				updatedMemberIdx = i
			}
		}
		// 次回実施日を設定
		stMembers[updatedMemberIdx].Next = CalcuNextDate(latestNextDate, touban.Scheduling)
		// データベース更新
		DateUpdateMember(stMembers[updatedMemberIdx])
		DateUpdateTouban(touban)
	}
}

func CalcuNextDate(latestNextDate time.Time, shedule string) string {
	var nextDate string
	now := time.Now()

	// 次回実施日計算
	interval := strings.Split(shedule, "-")[0]
	week := strings.Split(shedule, "-")[1]
	day := strings.Split(shedule, "-")[2]
	date := strings.Split(shedule, "-")[3]
	// 毎週
	if interval == "1" {
		nextDate = latestNextDate.AddDate(0, 0, 7).Format("2006-01-02")
		// 隔週
	} else if interval == "2" {
		nextDate = latestNextDate.AddDate(0, 0, 14).Format("2006-01-02")
		// ３週間毎
	} else if interval == "3" {
		nextDate = latestNextDate.AddDate(0, 0, 21).Format("2006-01-02")
		// １か月ごと
	} else if interval == "4" {
		// 毎月指定日
		if date != "0" {
			// フロントエンドでの日付計算が正しい前提とすると、バックエンドでの毎月指定は1か月加算するだけでいい。
			nextDate = latestNextDate.AddDate(0, 1, 0).Format("2006-01-02")
		} else {
			dayInt, _ := strconv.Atoi(day)
			weekInt, _ := strconv.Atoi(week)
			// 来月の初日を取得
			firstDayOfNextMonth := time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, now.Location())
			// 指定集の指定曜日の日付を算出
			diff := dayInt - int(firstDayOfNextMonth.Weekday())
			nextDate = firstDayOfNextMonth.AddDate(0, 0, 7*(weekInt-1)+diff).Format("2006-01-02")
		}
		// 毎日
	} else if interval == "0" {
		// 金曜日かどうか判定
		if latestNextDate.Weekday() == time.Friday {
			nextDate = latestNextDate.AddDate(0, 0, 3).Format("2006-01-02")
		} else {
			nextDate = latestNextDate.AddDate(0, 0, 1).Format("2006-01-02")
		}
	} else {
		nextDate = latestNextDate.Format("2006-01-02")
	}

	return nextDate
}

func DateUpdateMember(accessM controller.AccessMember) int64 {
	return accessM.UpdateMember()
}
func DateUpdateTouban(accessM controller.AccessTouban) int64 {
	return accessM.UpdateTouban()
}
