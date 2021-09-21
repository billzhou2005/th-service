package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

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

func (repository *TableRepo) SaveCardgenToTable(tableid int, numberofplayers int, players [9]models.Player) {
	var table models.Table
	tablelevelmap := make(map[string]int)
	//var repository *TableRepo

	table.TableID = tableid
	//table.Tablelevel = "primary"
	table.Numofplayers = numberofplayers

	jsonfile, err := os.Open("./models/tablelevel.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonfile.Close()
	bytevalue, _ := ioutil.ReadAll(jsonfile)
	json.Unmarshal(bytevalue, &tablelevelmap)

	//calculate the tablelevel, primary/intermediate/advanced/vip/royalvip
	basenum := 10000000
	divres := tableid / basenum
	for k, v := range tablelevelmap {
		if divres == v {
			table.Tablelevel = k
		}
	}

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

	errdb := models.CreateTable(repository.Db, &table)
	if errdb != nil {
		fmt.Println("models.CreateTable", errdb)
		return
	}
}

//create table from cardgen
func (repository *TableRepo) CreateTableFromCardgen(c *gin.Context) {
	//	tableid, _ := c.Params.Get("tableid")
	//	numofp, _ := c.Params.Get("numofp")
	tableidstr := c.Query("tableid")
	numofpstr := c.Query("numofp")

	//numofplayers := 9
	numofplayers, err := strconv.Atoi(numofpstr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"numofplayers not number": err})
		return
	}
	if numofplayers < 1 || numofplayers > 9 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"numofplayers range not 1-9": err})
		return
	}

	tableid, errtableid := strconv.Atoi(tableidstr)
	if errtableid != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"tableid not number": err})
		return
	}
	if tableid < 10000000 || tableid >= 60000000 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"tableid range not 10000000~59999999": err})
		return
	}

	players := models.Cardgen(numofplayers)
	c.JSON(http.StatusOK, players)

	//sava players to db
	repository.SaveCardgenToTable(tableid, numofplayers, players)
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
