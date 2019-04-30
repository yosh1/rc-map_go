
package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "strconv"
  "app/controllers"
  "app/models"
)

// common

func badRequest(c *gin.Context) {
  c.JSON(http.StatusBadRequest, gin.H{
    "code": "BadRequest",
    "message": "テナントIDを指定して下さい。",
  })
}

func notFound(c *gin.Context) {
  c.JSON(http.StatusNotFound, gin.H {
    "code": "Not Found",
    "message": "対象のテナントのレコードが見つかりません",
  })
}

// setting

func getSetting(c *gin.Context) {
  tenantId, _ := strconv.Atoi(c.Param("tenantId"))
  ctrl := controllers.NewSetting()
  setting := ctrl.GetByTenantId(tenantId)
  if setting == false {
    notFound(c)
  }
  c.JSON(http.StatusOK, setting)
}

func postSetting(c *gin.Context) {
  tenantId, _ := strconv.Atoi(c.Param("tenantId"))
  var setting models.Setting
  c.BindJSON(&setting)
  ctrl := controllers.NewSetting()
  ctrl.PostSetting(tenantId, setting)
  c.JSON(http.StatusOK, gin.H{
    "message": "record created",
    "tenantId": tenantId,
  })
}

func putSetting(c *gin.Context) {
  tenantId, _ := strconv.Atoi(c.Param("tenantId"))
  var setting models.Setting
  c.BindJSON(&setting)
  ctrl := controllers.NewSetting()
  result := ctrl.PutSetting(tenantId, setting)
  if result == false {
    notFound(c)
  } else {
    c.JSON(http.StatusOK, gin.H{
      "message": "record updated",
      "tenantId": tenantId,
    })
  }
}

func deleteSetting(c *gin.Context) {
  tenantId, _ := strconv.Atoi(c.Param("tenantId"))
  ctrl := controllers.NewSetting()
  result := ctrl.DeleteSetting(tenantId)
  if result == false {
    notFound(c)
  } else {
    c.JSON(http.StatusOK, gin.H{
      "message": "record deleted",
      "tenantId": tenantId,
    })
  }
}
