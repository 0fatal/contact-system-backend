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

// GetRefreshApplyList
// 获取重生申请列表
func GetRefreshApplyList(c *gin.Context) {
	applyList, err := model.GetRefreshApplyList()
	if err != nil {
		response.Fail().Msg(err.Error()).Send(c)
		return
	}
	response.Ok().Data(applyList).Send(c)
}

// GetRefreshDetail
// 获取重生申请详情
func GetRefreshDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		response.Fail().Msg("id错误").Send(c)
		return
	}

	apply := model.RefreshApply{
		Model: gorm.Model{
			ID: uint(id),
		},
	}

	err = apply.GetRefreshDetail()

	if err != nil {
		response.Fail().Msg(err.Error()).Send(c)
		return
	}
	response.Ok().Data(apply).Send(c)
}

// CreateNewRefresh
// 创建新重生申请
func CreateNewRefresh(c *gin.Context) {
	var data dto.CreateRefreshApply
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

	if !customer.IsRisk() {
		response.Fail().Msg("客户未违规，无法发起重生").Send(c)
		return
	}

	identityReply, _ := model.FindLatestIdentifyApply(customer.ID)

	apply := model.RefreshApply{
		ReasonId: data.RefreshReason,
		ApplyId:  identityReply.ID,
	}

	err = apply.InsertNew()

	if err != nil {
		response.Fail().Msg(err.Error()).Send(c)
		return
	}
	response.Ok().Msg("创建申请成功").Send(c)
}

// CheckRefreshApply
// 审核重生申请
func CheckRefreshApply(c *gin.Context) {
	var data dto.CheckRefreshApply

	err := c.ShouldBindJSON(&data)
	if HandleDTOVerifyError(err, c) {
		return
	}
	refreshApply := model.RefreshApply{
		Model: gorm.Model{
			ID: data.ApplyId,
		},
	}

	if err := refreshApply.GetRefreshDetail(); err != nil {
		response.Fail().Msg("该申请不存在").Send(c)
		return
	}

	if refreshApply.Status != 0 {
		response.Fail().Msg("请求已被处理").Send(c)
		return
	}

	if err := refreshApply.Check(data.Approve, util.GetCurrentUser(c)); err != nil {
		response.Fail().Msg(err.Error()).Send(c)
		return
	}

	if data.Approve {

		identifyApply := model.IdentifyApply{
			Model: gorm.Model{
				ID: refreshApply.ApplyId,
			},
		}
		_ = identifyApply.GetIdentifyDetail()
		customer := identifyApply.Target
		customer.Status = 0 // 客户状态改为正常
		_ = customer.Save()
		response.Ok().Msg("已重生").Send(c)
		return
	}

	response.Ok().Msg("拒绝成功").Send(c)
}
