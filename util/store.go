package util

import (
	"backend/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func GetCurrentUser(c *gin.Context) *model.User {
	_uid, _ := c.Cookie("user")
	uid, _ := strconv.Atoi(_uid)
	user := model.User{
		Model: gorm.Model{
			ID: uint(uid),
		},
	}

	_ = user.GetInfo()
	return &user
}
