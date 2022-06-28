package dto

type CreateIdentifyApply struct {
	RiskReason    uint     `json:"risk_reason" binding:"required"`
	RiskLevel     int      `json:"risk_level" binding:"required"`
	ExternalLevel string   `json:"external_level" binding:"required"`
	Remark        string   `json:"remark"`
	Appendix      []string `json:"appendix"`
	TargetName    string   `json:"target_name" binding:"required"`
}

type CheckIdentifyApply struct {
	Approve bool `json:"approve"`  // 是否同意
	ApplyId uint `json:"apply_id"` // 申请号
}
