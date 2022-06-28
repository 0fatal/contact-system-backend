package api

import (
	"backend/model"
	"backend/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func GetReasonList(c *gin.Context) {
	var reasonList []*model.ApplyReason
	var err error
	if c.Query("type") == "refresh" {
		reasonList, err = model.GetApplyReasonList(1)
	} else {
		reasonList, err = model.GetApplyReasonList(0)
	}
	if err != nil {
		response.Fail().Msg(err.Error()).Send(c)
		return
	}
	response.Ok().Data(reasonList).Send(c)
}

func GetReasonDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		response.Fail().Msg("id错误").Send(c)
		return
	}

	reason := model.ApplyReason{
		Model: gorm.Model{
			ID: uint(id),
		},
	}

	err = reason.GetApplyReasonDetail()

	if err != nil {
		response.Fail().Msg(err.Error()).Send(c)
		return
	}
	response.Ok().Data(reason).Send(c)
}

func EnableReason(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		response.Fail().Msg("id错误").Send(c)
		return
	}

	var open bool
	if c.Query("open") == "1" {
		open = true
	} else {
		open = false
	}

	reason := model.ApplyReason{
		Model: gorm.Model{
			ID: uint(id),
		},
	}

	err = reason.EnableReason(open)

	if err != nil {
		response.Fail().Msg(err.Error()).Send(c)
		return
	}
	if c.Query("open") == "1" {
		response.Ok().Msg("启用成功").Send(c)
	} else {
		response.Ok().Msg("禁用成功").Send(c)
	}
}
