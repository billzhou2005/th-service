package main

import (
	"net/http"
	"strconv"
	"th-service/controllers"

	"th-service/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := setupRouter()
	_ = r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(CORSMiddleware())

	r.GET("/cardgen/:id", cardgen)
	tableRepo := controllers.Newtable()
	r.POST("/tables", tableRepo.CreateTable)
	r.GET("/tables", tableRepo.GetTables)
	r.GET("/tables/:id", tableRepo.GetTable)
	r.PUT("/tables/:id", tableRepo.UpdateTable)
	r.DELETE("/tables/:id", tableRepo.DeleteTable)

	userRepo := controllers.New()
	r.POST("/users", userRepo.CreateUser)
	r.GET("/users", userRepo.GetUsers)
	r.GET("/users/:id", userRepo.GetUser)
	r.PUT("/users/:id", userRepo.UpdateUser)
	r.DELETE("/users/:id", userRepo.DeleteUser)

	return r
}
func cardgen(c *gin.Context) {
	id, _ := c.Params.Get("id")
	//numofplayers := 9
	numofplayers, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Input id not number": err})
		return
	}
	if numofplayers < 1 || numofplayers > 9 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Number range not 1-9": err})
		return
	}

	players := models.Cardgen(numofplayers)
	c.JSON(http.StatusOK, players)

	//Save the cargen to database:gin_gorm/tables
	controllers.SaveCardgenToTable(numofplayers, players)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
