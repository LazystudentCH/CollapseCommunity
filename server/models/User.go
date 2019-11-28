package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name string `json:"name"`
	Gender string `json:"gender"`
	PassWord string `gorm:"not null",json:"password"`
	Phone string `json:"phone",gorm:"unique;not null;unique_index"`
	Avatar string `gorm:"default:'http://q0u8r9imq.bkt.clouddn.com/bbb.jpeg'",json:"avatar"`
	VerifyCode string `gorm:"-",json:"verifyCode"`
	IsAdmin bool `json:"is_admin",gorm:"default:'0'"`

}
