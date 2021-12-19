package controllers

import (
	"errors"
	"net/http"

	"th-service/database"
	"th-service/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BizInfoRepo struct {
	Db *gorm.DB
}

func NewBizInfo() *BizInfoRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.BizInfo{})
	return &BizInfoRepo{Db: db}
}

//create BizInfo
func (repository *BizInfoRepo) CreateBizInfo(c *gin.Context) {
	var bizinfo models.BizInfo
	c.BindJSON(&bizinfo)
	err := models.CreateBizInfo(repository.Db, &bizinfo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, bizinfo)
}

//get bizinfos
func (repository *BizInfoRepo) GetBizInfos(c *gin.Context) {
	var bizinfo []models.BizInfo
	err := models.GetBizInfos(repository.Db, &bizinfo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, bizinfo)
}

//get bizinfo by id
func (repository *BizInfoRepo) GetBizInfo(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var bizinfo models.BizInfo
	err := models.GetBizInfo(repository.Db, &bizinfo, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, bizinfo)
}

// update bizinfo
func (repository *BizInfoRepo) UpdateBizInfo(c *gin.Context) {
	var bizinfo models.BizInfo
	id, _ := c.Params.Get("id")
	err := models.GetBizInfo(repository.Db, &bizinfo, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&bizinfo)
	err = models.UpdateBizInfo(repository.Db, &bizinfo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, bizinfo)
}

// delete bizinfo
func (repository *BizInfoRepo) DeleteBizInfo(c *gin.Context) {
	var bizinfo models.BizInfo
	id, _ := c.Params.Get("id")
	err := models.DeleteBizInfo(repository.Db, &bizinfo, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "BizInfo deleted successfully"})
}
