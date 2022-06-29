package model

import (
	"crypto/md5"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func init() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		"contact",    // 用户名
		"contact233", // 密码
		"", // host
		3306,          // 端口
		"contract_db", // 数据库名
	)

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 数据表名为单数
		},
		//DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		panic(err)
	}
	hasTable := db.Migrator().HasTable("apply_reason")
	// 数据库迁移
	_ = db.AutoMigrate(&User{}, &ApplyReason{}, &RiskCustomer{}, &IdentifyApply{}, &RefreshApply{})
	if !hasTable {
		PreInsertApplyReason()
		PreInsertUser()
	}
}

func Md5(str string) string {
	data := []byte(str)
	res := md5.Sum(data)
	return fmt.Sprintf("%x", res)
}
