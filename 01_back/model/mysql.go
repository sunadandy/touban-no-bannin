package model

import (
	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// 当番テーブル(JSON)をマッピングする構造体
type StTouban struct {
	Id            int    `json:"id" gorm:"AUTO_INCREMENT"`
	Title         string `json:"title"`
	Owner         string `json:"owner"`
	Start         string `json:"start"`
	Interval_type int    `json:"interval_type"`
	Mailing       bool   `json:"mailing"`
	Timing        int    `json:"timing"`
	Message       string `json:"message"`
	Password      string `json:"password"`
	Cc            string `json:"cc"`
}

// 順番テーブル(JSON)をマッピングする構造体
type StMember struct {
	Touban_id       int    `json:"touban_id"`
	Name            string `json:"name"`
	Employee_number string `json:"emplyoee_number"`
	Email           string `json:"email"`
	Order_number    int    `json:"order_number"`
	Last            string `json:"last"`
	Next            string `json:"next"`
}
type Members []StMember

var db *gorm.DB
var toubanTbl string = "toubanTbl"
var memberTbl string = "memberTbl"

func DBConnect() {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	// PROTOCOL := "tcp(touban_cont:3306)"
	PROTOCOL := "tcp(172.19.0.2:3306)" //MySQLコンテナのIP
	DB := "toubanDB"

	connection := USER + ":" + PASS + "@" + PROTOCOL + "/" + DB

	var err error
	db, err = gorm.Open(DBMS, connection)
	if err != nil {
		panic(err.Error())
	}
}

func NewTouban() StTouban {
	stTouban := StTouban{}
	return stTouban
}

func NewMember() StMember {
	stMember := StMember{}
	return stMember
}

func NewMembers() Members {
	members := Members{}
	return members
}

// ---------------------------------------------------------------
// 当番テーブルCRUD
// ---------------------------------------------------------------
func (stTouban StTouban) CreateTouban() error {
	err := db.Table(toubanTbl).Create(&stTouban).Error
	if err != nil {
		return err
	}
	return nil
}

func ReadTouban() (string, error) {
	stToubans := []StTouban{}
	db.Table(toubanTbl).Find(&stToubans)
	jsonData, err := json.Marshal(&stToubans)
	if err != nil {
		print(err)
		return "", err
	}
	return string(jsonData), nil
}

func (stTouban StTouban) UpdateTouban() int64 {
	id := stTouban.Id
	result := db.Table(toubanTbl).Where("id = ?", id).Update(&stTouban)
	// Updateはエラーを返さないので更新された行数で成否を判断
	return result.RowsAffected
}

func DeleteTouban(id int) {
	db.Table(toubanTbl).Where("id = ?", id).Delete(&StTouban{})
}

// ---------------------------------------------------------------
// 順番テーブルCRUD
// ---------------------------------------------------------------
func (stMember StMember) CreateMember() error {
	err := db.Table(memberTbl).Create(&stMember).Error
	if err != nil {
		return err
	}
	return nil
}

func ReadMember() string {
	stMember := []StMember{}
	db.Table(memberTbl).Find(&stMember)
	jsonData, err := json.Marshal(&stMember)
	if err != nil {
		print(err)
	}
	return string(jsonData)
}

func (stMember StMember) UpdateMember() error {
	return nil
}

func DeleteMember(toubanId int) {
	db.Table(memberTbl).Where("touban_id = ?", toubanId).Delete(&StMember{})
}
