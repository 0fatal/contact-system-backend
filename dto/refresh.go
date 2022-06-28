package dto

type CreateRefreshApply struct {
	RefreshReason uint   `json:"refresh_reason" binding:"required"` // 重生原因
	TargetName    string `json:"target_name" binding:"required"`    // 客户名
}

type CheckRefreshApply struct {
	Approve bool `json:"approve"  binding:"required"` // 是否同意
	ApplyId uint `json:"apply_id" binding:"required"` // 申请号
}
