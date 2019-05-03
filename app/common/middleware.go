
package common

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func AuthHeader() gin.HandlerFunc {
  return func (c *gin.Context) {
    c.Next()
    requestHeader := c.GetHeader("Content-Type")
    if requestHeader != "application/json" {
      c.JSON(http.StatusUnsupportedMediaType,
        gin.H{
          "code": "Unsupported Media Type",
          "message": "application/jsonでリクエストして下さい",
        })
        c.Abort()
      }
    }
  }
