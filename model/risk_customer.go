package model

import "gorm.io/gorm"

type RiskCustomer struct {
	gorm.Model
	Name   string `gorm:"unique;size:30;not null"`
	Status int    `gorm:"type:tinyint;default:0"` // 0正常 1违约
}

// InsertNew
// 添加新违约客户
func (rc *RiskCustomer) InsertNew() error {
	err := db.Create(rc).Error
	return err
}

// Find
// 查找违约客户
func (rc *RiskCustomer) Find() error {
	err := db.Where("name = ?", rc.Name).First(rc).Error
	return err
}

// IsRisk
// 客户是否违约
func (rc *RiskCustomer) IsRisk() bool {
	if rc.Find() == nil {
		return rc.Status == 1
	}
	return false
}

func (rc *RiskCustomer) Save() error {
	err := db.Save(rc).Error
	return err
}

func GetCustomerList() (list []*RiskCustomer, err error) {
	err = db.Model(RiskCustomer{}).Where("status = 1").Find(&list).Error
	return
}
