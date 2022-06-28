package model

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IdentifyApply struct {
	gorm.Model
	CheckerId     sql.NullInt64 `json:"-" gorm:"index"`
	TargetId      uint          `json:"-" gorm:"not null;index"`
	ExternalLevel string        `gorm:"type:varchar(20);not null"`
	Target        *RiskCustomer `gorm:"foreignKey:TargetId;references:ID"`  // 客户
	ReasonId      uint          `json:"-" gorm:"not null;index"`            // 认定原因
	RiskReason    *ApplyReason  `gorm:"foreignKey:ReasonId;references:ID"`  // 违约原因
	RiskLevel     int           `gorm:"default:0"`                          // 违约严重性 0: 低 1：中 2：高
	Remark        string        `gorm:"size:1024"`                          // 备注
	Appendix      []string      `gorm:"type:varchar(1024);default:''"`      // 附件
	ApplyStatus   int           `gorm:"type:tinyint;default:0"`             // 审核状态 0:待审核 1:同意 2:拒绝
	Checker       *User         `gorm:"foreignKey:CheckerId;references:ID"` // 审核人
}

func (ia IdentifyApply) Value() (driver.Value, error) {
	// 将appendix数组转为，分割
	return json.Marshal(ia)
}

func (ia *IdentifyApply) Scan(data interface{}) error {
	return json.Unmarshal(data.([]byte), ia)
}

func (ia *IdentifyApply) InsertNew() error {
	err := db.Create(ia).Error
	return err
}

func GetIdentifyApplyList() (list []*IdentifyApply, err error) {
	err = db.Preload(clause.Associations).Find(&list).Error
	return
}

func (ia *IdentifyApply) GetIdentifyDetail() (err error) {
	err = db.Preload(clause.Associations).Where("id = ?", ia.ID).First(ia).Error
	return
}

// Check
// 违约认定审核
func (ia *IdentifyApply) Check(approve bool, checker *User) (err error) {
	if approve {
		ia.ApplyStatus = 1
	} else {
		ia.ApplyStatus = 2
	}
	ia.CheckerId = sql.NullInt64{
		Int64: int64(checker.ID),
	}
	err = db.Model(RefreshApply{}).Where("id = ?", ia.ID).Update("apply_status", ia.ApplyStatus).Update("checker_id", checker.ID).Error
	return
}

func FindLatestIdentifyApply(targetId uint) (apply *IdentifyApply, err error) {
	apply = &IdentifyApply{}
	err = db.Preload(clause.Associations).Where("target_id = ?", targetId).Order("created_at desc").First(apply).Error
	return
}
