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
func (stTouban StTouban) CreateTouban() error {
	err := db.Table(toubanTbl).Create(&stTouban).Error
	if err != nil {
		print(err)
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

func DeleteOrder(refId int) {
	db.Table(toubanTbl).Where("toubanId = ?", refId).Delete(&StOrder{})
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
