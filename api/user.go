package api

import (
	"backend/dto"
	"backend/model"
	"backend/response"
	"backend/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Login(c *gin.Context) {
	var data dto.UserLogin
	err := c.ShouldBindJSON(&data)
	if HandleDTOVerifyError(err, c) {
		return
	}

	user := model.User{
		Username: data.Username,
		Password: data.Password,
	}

	if err := user.Login(); err != nil {
		response.Fail().Msg(err.Error()).Send(c)
		return
	}
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("user", strconv.Itoa(int(user.ID)), 3600, "/", "", true, false)

	response.Ok().Msg("登录成功").Send(c)
}

func GetInfo(c *gin.Context) {
	user := util.GetCurrentUser(c)
	response.Ok().Data(Data{
		"username": user.Username,
		"nickname": user.Nickname,
		"role":     user.Role,
	}).Send(c)
}
