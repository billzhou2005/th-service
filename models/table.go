package models

import (
	"gorm.io/gorm"
)

type Table struct {
	gorm.Model
	TableID      int    `json:"tableid"`
	Numofplayers int    `json:"numofplayers"`
	Tablelevel   string `json:"tablelevel"`
	Player1      string `json:"player1"`
	Player2      string `json:"player2"`
	Player3      string `json:"player3"`
	Player4      string `json:"player4"`
	Player5      string `json:"player5"`
	Player6      string `json:"player6"`
	Player7      string `json:"player7"`
	Player8      string `json:"player8"`
	Player9      string `json:"player9"`
}

// create Table
func CreateTable(db *gorm.DB, Table *Table) (err error) {
	err = db.Create(Table).Error
	if err != nil {
		return err
	}
	return nil
}

//get tables
func GetTables(db *gorm.DB, Table *[]Table) (err error) {
	err = db.Find(Table).Error
	if err != nil {
		return err
	}
	return nil
}

//get Table by id
func GetTable(db *gorm.DB, Table *Table, id string) (err error) {
	err = db.Where("id = ?", id).First(Table).Error
	if err != nil {
		return err
	}
	return nil
}

//update Table
func UpdateTable(db *gorm.DB, Table *Table) (err error) {
	db.Save(Table)
	return nil
}

//delete Table
func DeleteTable(db *gorm.DB, Table *Table, id string) (err error) {
	db.Where("id = ?", id).Delete(Table)
	return nil
}
