package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"th-service/database"
	"th-service/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FtkRepo struct {
	Db *gorm.DB
}

func NewFtk() *FtkRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.FtkDb{})
	return &FtkRepo{Db: db}
}

func (repository *FtkRepo) SaveCardsToFtkDb(ftkCards models.FtkCards) {
	var ftkDb models.FtkDb

	ftkDb.RID = ftkCards.RID

	b, err := json.Marshal(ftkCards.PlayersCards[0])
	if err != nil {
		panic(err)
	}
	ftkDb.Player1 = string(b)
	b, err = json.Marshal(ftkCards.PlayersCards[1])
	if err != nil {
		panic(err)
	}
	ftkDb.Player2 = string(b)
	b, err = json.Marshal(ftkCards.PlayersCards[2])
	if err != nil {
		panic(err)
	}
	ftkDb.Player3 = string(b)
	b, err = json.Marshal(ftkCards.PlayersCards[3])
	if err != nil {
		panic(err)
	}
	ftkDb.Player4 = string(b)

	errdb := models.CreateFtkDb(repository.Db, &ftkDb)
	if errdb != nil {
		fmt.Println("models.CreateFtkDb", errdb)
	}
}

//create table from FtkGen
func (repository *FtkRepo) CreateFtkDbFromFtkGen(c *gin.Context) {
	rIDStr := c.Query("rid")

	rID, err := strconv.Atoi(rIDStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"rID not number": err})
		return
	}

	ftkCards := models.FtkGen(rID)
	c.JSON(http.StatusOK, ftkCards)

	repository.SaveCardsToFtkDb(ftkCards)
}

//create Table
func (repository *FtkRepo) CreateTable(c *gin.Context) {
	var ftkDb models.FtkDb
	c.BindJSON(&ftkDb)
	err := models.CreateFtkDb(repository.Db, &ftkDb)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, ftkDb)
}

//get Tables
func (repository *FtkRepo) GetTables(c *gin.Context) {
	var ftkDb []models.FtkDb
	err := models.GetFtkDbs(repository.Db, &ftkDb)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, ftkDb)
}

//get Table by id
func (repository *FtkRepo) GetTable(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var ftkDb models.FtkDb
	err := models.GetFtkDb(repository.Db, &ftkDb, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, ftkDb)
}

// update Table
func (repository *FtkRepo) UpdateTable(c *gin.Context) {
	var ftkDb models.FtkDb
	id, _ := c.Params.Get("id")
	err := models.GetFtkDb(repository.Db, &ftkDb, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&ftkDb)
	err = models.UpdateFtkDb(repository.Db, &ftkDb)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, ftkDb)
}

// delete Table
func (repository *FtkRepo) DeleteTable(c *gin.Context) {
	var ftkDb models.FtkDb
	id, _ := c.Params.Get("id")
	err := models.DeleteFtkDb(repository.Db, &ftkDb, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "FtkDb deleted successfully"})
}
