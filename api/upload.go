package api

import (
	"backend/response"
	"backend/util"
	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.Fail().Msg(err.Error()).Send(c)
		return
	}
	_url, err := util.Upload(file)
	if err != nil {
		response.Fail().Msg(err.Error()).Send(c)
		return
	}
	response.Ok().Data(Data{
		"url": _url,
	}).Send(c)
}
