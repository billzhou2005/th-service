package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `json:"username"`
	Userid    string `json:"userid"`
	Passwd    string `json:"passwd"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
	Balance   int    `json:"balance"`
	CardID    int    `json:"cardid"`
}

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

//get user by username
func GetUserByUsername(db *gorm.DB, User *[]User, username string) (err error) {
	err = db.Where("username = ?", username).Find(User).Error
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
