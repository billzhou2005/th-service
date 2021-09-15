package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	CardID   int    `json:"cardid"`
}

/*{
	"username": "player1",
	"cards": {
		"cardone": {
			"points": 6,
			"suits": 4
		},
		"cardtwo": {
			"points": 9,
			"suits": 1
		},
		"cardthree": {
			"points": 9,
			"suits": 3
		}
	},
	"cardstype": "highcard",
	"cifirst": 9,
	"cisecond": 9,
	"cithird": 6,
	"Cardsscore": 0
}*/
//create a user
func CreateUser(db *gorm.DB, User *User) (err error) {
	err = db.Create(User).Error
	if err != nil {
		return err
	}
	return nil
}

//get users
func GetUsers(db *gorm.DB, User *[]User) (err error) {
	err = db.Find(User).Error
	if err != nil {
		return err
	}
	return nil
}

//get user by id
func GetUser(db *gorm.DB, User *User, id string) (err error) {
	err = db.Where("id = ?", id).First(User).Error
	if err != nil {
		return err
	}
	return nil
}

//update user
func UpdateUser(db *gorm.DB, User *User) (err error) {
	db.Save(User)
	return nil
}

//delete user
func DeleteUser(db *gorm.DB, User *User, id string) (err error) {
	db.Where("id = ?", id).Delete(User)
	return nil
}
