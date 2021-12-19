package models

import (
	"gorm.io/gorm"
)

type BizInfo struct {
	gorm.Model
	Name        string `json:"name"`
	Tel         string `json:"tel"`
	Email       string `json:"email"`
	Description string `json:"description"`
	Locale      string `json:"locale"`
}

/*{
	"name": "name1",
	"tel": "highcard",
	"email": 9,
	"description": 9,
	"locale": 6,
}*/
//create a BizInfo
func CreateBizInfo(db *gorm.DB, BizInfo *BizInfo) (err error) {
	err = db.Create(BizInfo).Error
	if err != nil {
		return err
	}
	return nil
}

//get BizInfos
func GetBizInfos(db *gorm.DB, BizInfo *[]BizInfo) (err error) {
	err = db.Find(BizInfo).Error
	if err != nil {
		return err
	}
	return nil
}

//get BizInfo by id
func GetBizInfo(db *gorm.DB, BizInfo *BizInfo, id string) (err error) {
	err = db.Where("id = ?", id).First(BizInfo).Error
	if err != nil {
		return err
	}
	return nil
}

//update BizInfo
func UpdateBizInfo(db *gorm.DB, BizInfo *BizInfo) (err error) {
	db.Save(BizInfo)
	return nil
}

//delete BizInfo
func DeleteBizInfo(db *gorm.DB, BizInfo *BizInfo, id string) (err error) {
	db.Where("id = ?", id).Delete(BizInfo)
	return nil
}
