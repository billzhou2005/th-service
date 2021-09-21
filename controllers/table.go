package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"th-service/database"
	"th-service/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TableRepo struct {
	Db *gorm.DB
}

func Newtable() *TableRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Table{})
	return &TableRepo{Db: db}
}

func SaveCardgenToTable(numberofplayers int, players [9]models.Player) {
	var table models.Table
	//var repository *TableRepo

	table.TableID = 10020
	table.Tablelevel = "primary"
	table.Numofplayers = numberofplayers
	b, err := json.Marshal(players[0])
	if err != nil {
		panic(err)
	}
	table.Player1 = string(b)
	b, err = json.Marshal(players[1])
	if err != nil {
		panic(err)
	}
	table.Player2 = string(b)
	b, err = json.Marshal(players[2])
	if err != nil {
		panic(err)
	}
	table.Player3 = string(b)
	b, err = json.Marshal(players[3])
	if err != nil {
		panic(err)
	}
	table.Player4 = string(b)
	b, err = json.Marshal(players[4])
	if err != nil {
		panic(err)
	}
	table.Player5 = string(b)
	b, err = json.Marshal(players[5])
	if err != nil {
		panic(err)
	}
	table.Player6 = string(b)
	b, err = json.Marshal(players[6])
	if err != nil {
		panic(err)
	}
	table.Player7 = string(b)
	b, err = json.Marshal(players[7])
	if err != nil {
		panic(err)
	}
	table.Player8 = string(b)
	b, err = json.Marshal(players[8])
	if err != nil {
		panic(err)
	}
	table.Player9 = string(b)

	db := database.InitDb()
	db.AutoMigrate(&models.Table{})
	db.Create(&table)
	fmt.Println("create table.tableID 10020,convert json array to string")
	//return
}

//create Table
func (repository *TableRepo) CreateTable(c *gin.Context) {
	var table models.Table
	c.BindJSON(&table)
	err := models.CreateTable(repository.Db, &table)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, table)
}

//get Tables
func (repository *TableRepo) GetTables(c *gin.Context) {
	var table []models.Table
	err := models.GetTables(repository.Db, &table)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, table)
}

//get Table by id
func (repository *TableRepo) GetTable(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var table models.Table
	err := models.GetTable(repository.Db, &table, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, table)
}

// update Table
func (repository *TableRepo) UpdateTable(c *gin.Context) {
	var table models.Table
	id, _ := c.Params.Get("id")
	err := models.GetTable(repository.Db, &table, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&table)
	err = models.UpdateTable(repository.Db, &table)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, table)
}

// delete Table
func (repository *TableRepo) DeleteTable(c *gin.Context) {
	var table models.Table
	id, _ := c.Params.Get("id")
	err := models.DeleteTable(repository.Db, &table, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Table deleted successfully"})
}
