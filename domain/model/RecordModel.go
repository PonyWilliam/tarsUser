package model
type Record struct{
	//逾期记录表
	ID int64 `gorm:"primary_key;auto_increment;not_null" json:"id"`
	UID int64 `json:""`//关联到用户表，逾期记录
}