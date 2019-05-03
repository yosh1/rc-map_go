
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
	db.AutoMigrate(models.Setting{}, models.RailroadCrossing{})

	// DEBUG:
	db.Create(models.Setting{
		TenantId: 1,
		PaymentType: 1,
		CardType: 1,
		DestBankAccount: "三井住友銀行吹田支店",
		Tos: "利用規約",
	})
	db.Create(models.Setting{
		TenantId: 2,
		PaymentType: 2,
		CardType: 2,
		DestBankAccount: "UFJ銀行",
		Tos: "利用規約",
	})

	db.Create(models.RailroadCrossing{
		Name: "志垣 大智",
		TelNumber: 123456789,
		Mail: "hoge-hoge@gmail.com",
		SaleName: "Go言語入門",
		Type: 1,
		IsSuccess: 1,
		Date: "2018/12/12",
	})

	db.Create(models.RailroadCrossing{
		Name: "関口 颯",
		TelNumber: 123456789,
		Mail: "hoge-hoge@gmail.com",
		SaleName: "Docker言語入門",
		Type: 1,
		IsSuccess: 1,
		Date: "2018/01/12",
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
		setting.GET("/:tenantId", getSetting)
		setting.POST("/:tenantId", postSetting)
		setting.PUT("/:tenantId", putSetting)
		setting.DELETE("/:tenantId", deleteSetting)
	}

	sales := r.Group("/sales")
	{
		sales.GET("", badRequest)
		sales.POST("", badRequest)
		sales.GET("/:id", badRequest)
		sales.PUT("/:id", badRequest)
		sales.DELETE("/:id", badRequest)
	}

	RailroadCrossing := r.Group("RailroadCrossing")
	{
		RailroadCrossing.GET("", badRequest)
		RailroadCrossing.GET("/:userId", badRequest)
	}

	cancel := r.Group("/cancel")
	{
		cancel.POST("/:userId/:salesId", badRequest)
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}
