package main

import (
	"th-service/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := setupRouter()
	_ = r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(CORSMiddleware())

	tableRepo := controllers.Newtable()
	r.GET("/cardgen", tableRepo.CreateTableFromCardgen)
	r.POST("/tables", tableRepo.CreateTable)
	r.GET("/tables", tableRepo.GetTables)
	r.GET("/tables/:id", tableRepo.GetTable)
	r.PUT("/tables/:id", tableRepo.UpdateTable)
	r.DELETE("/tables/:id", tableRepo.DeleteTable)

	bizinfoRepo := controllers.NewBizInfo()
	r.POST("/bizinfo", bizinfoRepo.CreateBizInfo)
	r.GET("/bizinfo", bizinfoRepo.GetBizInfos)
	r.GET("/bizinfo/:id", bizinfoRepo.GetBizInfo)
	r.PUT("/bizinfo/:id", bizinfoRepo.UpdateBizInfo)
	r.DELETE("/bizinfo/:id", bizinfoRepo.DeleteBizInfo)

	userRepo := controllers.New()
	r.POST("/users", userRepo.CreateUser)
	r.GET("/users", userRepo.GetUsers)
	r.GET("/users/:id", userRepo.GetUser)
	r.PUT("/users/:id", userRepo.UpdateUser)
	r.DELETE("/users/:id", userRepo.DeleteUser)

	return r
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
