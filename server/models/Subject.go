package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Subject struct {
	gorm.Model
	User User `gorm:"foreignkey:UserId"`
	UserId uint `json:"user_id"`
	Title string `json:"title"` // 标题
	Content string `json:"content"`
	IsHot bool `json:"is_hot",gorm:"default:'0';index:'isHot'"`
}

type AddCommentRequest struct {
	Sid uint
	PubTime time.Time
}

type SubjectComment struct {
	gorm.Model
	User User `gorm:"foreignkey:UserId"`
	SubId uint `json:"sub_id"`
	UserId uint `json:"user_id"`
	Content string `json:"content"`
}

type PriseRecord struct {
	UserId uint `json:"user_id"`
	SubId uint `json:"sub_id"`
	HasPrise bool `json:"has_prise"`
}

func NewCommentRequest(sid uint,pubTime time.Time) *AddCommentRequest {
	return &AddCommentRequest{
		Sid:     sid,
		PubTime: pubTime,
	}
}