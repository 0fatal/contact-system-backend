package api

import (
	"backend/dto"
	"backend/model"
	"backend/response"
	"backend/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

// GetIdentifyApplyList
// 获取违约认定申请列表
func GetIdentifyApplyList(c *gin.Context) {
	applyList, err := model.GetIdentifyApplyList()
	if err != nil {
		response.Fail().Msg(err.Error()).Send(c)
		return
	}
	response.Ok().Data(applyList).Send(c)
}

// GetIdentifyDetail
// 获取违约认定申请详情
func GetIdentifyDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		response.Fail().Msg("id错误").Send(c)
		return
	}

	apply := model.IdentifyApply{
		Model: gorm.Model{
			ID: uint(id),
		},
	}

	err = apply.GetIdentifyDetail()

	if err != nil {
		response.Fail().Msg(err.Error()).Send(c)
		return
	}
	response.Ok().Data(apply).Send(c)
}

// CreateNewIdentify
// 创建新违约认定申请
func CreateNewIdentify(c *gin.Context) {
	var data dto.CreateIdentifyApply
	err := c.ShouldBindJSON(&data)
	if HandleDTOVerifyError(err, c) {
		return
	}

	// 验证客户是否已经违规
	customer := model.RiskCustomer{
		Name: data.TargetName,
	}

	if customer.Find() != nil {
		_ = customer.InsertNew()
		_ = customer.Find()
	}

	if customer.IsRisk() {
		response.Fail().Msg("客户已经违规，无法再次发起违规认定").Send(c)
		return
	}

	if apply, err := model.FindLatestIdentifyApply(customer.ID); err != nil || apply.ApplyStatus == 0 {
		response.Fail().Msg("客户上一条认定申请未处理，无法发起违规认定").Send(c)
		return
	}

	apply := model.IdentifyApply{
		Appendix:      data.Appendix,
		Remark:        data.Remark,
		RiskLevel:     data.RiskLevel,
		ReasonId:      data.RiskReason,
		ExternalLevel: data.ExternalLevel,
		TargetId:      customer.ID,
	}

	err = apply.InsertNew()

	if err != nil {
		response.Fail().Msg(err.Error()).Send(c)
		return
	}
	response.Ok().Msg("创建申请成功").Send(c)
}

// CheckIdentityApply
// 违约认定申请
func CheckIdentityApply(c *gin.Context) {
	var data dto.CheckIdentifyApply

	err := c.ShouldBindJSON(&data)
	if HandleDTOVerifyError(err, c) {
		return
	}

	identityApply := model.IdentifyApply{
		Model: gorm.Model{
			ID: data.ApplyId,
		},
	}

	if err := identityApply.GetIdentifyDetail(); err != nil {
		response.Fail().Msg("该申请不存在").Send(c)
		return
	}

	if identityApply.ApplyStatus != 0 {
		response.Fail().Msg("请求已被处理").Send(c)
		return
	}

	if err := identityApply.Check(data.Approve, util.GetCurrentUser(c)); err != nil {
		response.Fail().Msg(err.Error()).Send(c)
		return
	}

	if data.Approve {
		customer := identityApply.Target
		customer.Status = 1 // 客户状态改为违约
		_ = customer.Save()
		response.Ok().Msg("已列入违约名单").Send(c)
		return
	}

	response.Ok().Msg("拒绝成功").Send(c)
}
