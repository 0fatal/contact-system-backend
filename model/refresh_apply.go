package model

import (
	"database/sql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RefreshApply struct {
	gorm.Model
	ReasonId      uint          `json:"-" gorm:"not null;index"`
	ApplyId       uint          `gorm:"not null;index"`
	CheckerId     sql.NullInt64 `json:"-" gorm:"index"`
	Checker       *User         `gorm:"foreignKey:CheckerId;references:ID"` // 审核人
	RefreshReason *ApplyReason  `gorm:"foreignKey:ReasonId;references:ID"`
	Status        int           `gorm:"type:tinyint;default:0"` // 0待审核 1审核通过 2审核不通过
}

func (ra *RefreshApply) InsertNew() error {
	return db.Create(ra).Error
}

func GetRefreshApplyList() (list []*RefreshApply, err error) {
	err = db.Preload(clause.Associations).Find(&list).Error
	return
}

func (ra *RefreshApply) GetRefreshDetail() (err error) {
	err = db.Preload(clause.Associations).First(ra).Error
	return
}

func (ra *RefreshApply) Check(approve bool, checker *User) (err error) {
	if approve {
		ra.Status = 1
	} else {
		ra.Status = 2
	}
	ra.CheckerId = sql.NullInt64{
		Int64: int64(checker.ID),
	}
	err = db.Model(RefreshApply{}).Where("id = ?", ra.ID).Update("status", ra.Status).Update("checker_id", checker.ID).Error
	return
}
