package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	// for debuging
	// "fmt"
	// "reflect"
	"app/common"
	"app/models"
)

func migrate(db *gorm.DB) {

	// create scheme
	db.AutoMigrate(models.Setting{}, models.Purchase{})

	// DEBUG:
	db.Create(models.Setting{
		Number:    1,
		Longitude: 1,
		Aatitude:  1,
	})

	db.Create(models.Purchase{
		Number:    1,
		Longitude: 1,
		Aatitude:  1,
	})

}

func main() {
	db := common.Init()
	migrate(db)
	defer db.Close()

	r := gin.Default()

	r.Static("/swagger", "./swagger")

	// use middleware
	//r.Use(common.AuthHeader())

	setting := r.Group("/setting")
	{
		setting.GET("", badRequest)
		setting.POST("", badRequest)
		setting.PUT("", badRequest)
		setting.DELETE("", badRequest)
		setting.GET("/", getSetting)
		setting.POST("/", postSetting)
		setting.PUT("/", putSetting)
		setting.DELETE("/", deleteSetting)
	}

	purchase := r.Group("purchase")
	{
		purchase.GET("", badRequest)
	}

	cancel := r.Group("/cancel")
	{
		cancel.POST("/", badRequest)
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}
