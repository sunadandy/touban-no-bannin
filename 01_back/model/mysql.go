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
	Mail          bool   `json:"mail"`
	Remind        int    `json:"remind_type"`
	Message       string `json:"message"`
	Password      string `json:"password"`
}

// 順番テーブル(JSON)をマッピングする構造体
type StOrder struct {
	Touban_id       int    `json:"toubanId"`
	Name            string `json:"name"`
	Employee_number string `json:"employeeNo"`
	Order_number    int    `json:"order"`
	Last            string `json:"last"`
	Next            string `json:"next"`
}
type Orders []StOrder

var db *gorm.DB
var toubanTbl string = "toubanTbl"
var orderTbl string = "orderTbl"

func DBConnect() {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	// PROTOCOL := "tcp(touban_cont:3306)"
	PROTOCOL := "tcp(172.19.0.2:3306)" //コンテナのIP
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

func NewOrder() StOrder {
	stOrder := StOrder{}
	return stOrder
}

func NewOrders() Orders {
	orders := Orders{}
	return orders
}

// ---------------------------------------------------------------
// 当番テーブルCRUD
// ---------------------------------------------------------------
func (stTouban StTouban) CreateTouban() (int, error) {
	err := db.Table(toubanTbl).Create(&stTouban).Error
	if err != nil {
		return -1, err
	}

	// 追加されたレコードのIDを取得
	var maxID int
	if err = db.Table(toubanTbl).Select("max(id)").Scan(&maxID).Error; err != nil {
		panic(err)
	}
	return maxID, nil
}

func ReadTouban() string {
	stToubans := []StTouban{}
	db.Table(toubanTbl).Find(&stToubans)
	jsonData, err := json.Marshal(&stToubans)
	if err != nil {
		print(err)
	}
	return string(jsonData)
}

// ---------------------------------------------------------------
// 順番テーブルCRUD
// ---------------------------------------------------------------
func (stOrder StOrder) CreateOrder() error {
	err := db.Table(orderTbl).Create(&stOrder).Error
	if err != nil {
		return err
	}
	return nil
}

func ReadOrder() string {
	stOrder := []StOrder{}
	db.Table(orderTbl).Find(&stOrder)
	jsonData, err := json.Marshal(&stOrder)
	if err != nil {
		print(err)
	}
	return string(jsonData)
}

func (stOrder StOrder) UpdateOrder() error {
	return nil
}

// func ReadOrderByToubanID(id string) string {
// 	stOrder := []StOrder{}
// 	// テーブルIDでフィルタリングして取得
// 	db.Table(orderTbl).Where("touban_id = ?", id).Find(&stOrder)
// 	jsonData, err := json.Marshal(&stOrder)
// 	if err != nil {
// 		print(err)
// 	}
// 	return string(jsonData)
// }
