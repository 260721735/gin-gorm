package controller

import (
	"gin-gorm/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	ID     uint
	Name   string
	Age    int
	Gender string
	// 假设后面还有几百个字段...
}

type APIUser struct {
	ID   uint
	Name string
}

func GetTest(c *gin.Context) {
	db.Db.Model(&User{}).Limit(10).Find(&APIUser{})
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "GetTest",
		"data": nil,
	})
}
