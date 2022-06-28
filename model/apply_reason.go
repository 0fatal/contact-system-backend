package model

import "gorm.io/gorm"

type ApplyReason struct {
	gorm.Model
	Reason string `gorm:"size:255;not null"`     // 违约原因
	Status bool   `gorm:"not null"`              // 是否启用
	Type   int    `gorm:"type:tinyint;not null"` // 0违约认定 1重生认定
}

func GetApplyReasonList(typ int) (list []*ApplyReason, err error) {
	err = db.Model(ApplyReason{}).Where("type = ?", typ).Find(&list).Error
	return
}

// EnableReason
// 启用违约原因
func (ar *ApplyReason) EnableReason(open bool) (err error) {
	ar.Status = open
	if open {
		err = db.Model(ApplyReason{}).Where("id = ?", ar.ID).Update("status", 1).Error
	} else {
		err = db.Model(ApplyReason{}).Where("id = ?", ar.ID).Update("status", 0).Error
	}
	return
}

func (ar *ApplyReason) GetApplyReasonDetail() (err error) {
	err = db.Where("id = ?", ar.ID).First(ar).Error
	return
}

// PreInsertApplyReason
// 预插入原因
func PreInsertApplyReason() {
	// 违约认定原因
	reasons := []ApplyReason{
		ApplyReason{
			Reason: "6 个月内，交易对手技术性或资金等原因，给当天结算带来头寸缺口 2 次以上",
			Status: true,
			Type:   0,
		},
		ApplyReason{
			Reason: "6 个月内因各种原因导致成交后撤单 2 次以上",
			Status: true,
			Type:   0,
		},
		ApplyReason{
			Reason: "未能按照合约规定支付或延期支付利息，本金或其他交付义务（不包括在宽限期内延 期支付）",
			Status: true,
			Type:   0,
		},
		ApplyReason{
			Reason: "关联违约：如果集团（内部联系较紧密的集团）或集团内任一公司（较重要的子公司， 一旦发生违约会对整个集团造成较大影响的）发生违约，可视情况作为集团内所有成 员违约的触发条件",
			Status: true,
			Type:   0,
		},
		ApplyReason{
			Reason: "发生消极债务置换：债务人提供给债权人新的或重组的债务，或新的证券组合、现金 或资产低于原有金融义务；或为了债务人未来避免发生破产或拖欠还款而进行的展期 或重组",
			Status: true,
			Type:   0,
		},
		ApplyReason{
			Reason: "申请破产保护，发生法律接管，或者处于类似的破产保护状态",
			Status: true,
			Type:   0,
		},
		ApplyReason{
			Reason: "在其他金融机构违约（包括不限于：人行征信记录中显示贷款分类状态不良类情况， 逾期超过 90 天等），或外部评级显示为违约级别",
			Status: true,
			Type:   0,
		},
	}

	// 重生认定原因
	reasons = append(reasons, ApplyReason{
		Reason: "正常结算后解除",
		Status: true,
		Type:   0,
	}, ApplyReason{
		Reason: "在其他金融机构违约解除，或外部评级显示为非违约级别",
		Status: true,
		Type:   0,
	}, ApplyReason{
		Reason: "计提比例小于设置界限",
		Status: true,
		Type:   0,
	}, ApplyReason{
		Reason: "连续 12 个月内按时支付本金和利息",
		Status: true,
		Type:   0,
	}, ApplyReason{
		Reason: "客户的还款意愿和还款能力明显好转，已偿付各项逾期本金、逾期利息和其他费用（包 括罚息等），且连续 12 个月内按时支付本金、利息",
		Status: true,
		Type:   0,
	}, ApplyReason{
		Reason: "导致违约的关联集团内其他发生违约的客户已经违约重生，解除关联成员的违约设定",
		Status: true,
		Type:   0,
	})
	db.CreateInBatches(&reasons, len(reasons))
}
