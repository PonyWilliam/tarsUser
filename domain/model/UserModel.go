package model

import "gorm.io/plugin/soft_delete"
type User struct{
	ID int64 `gorm:"primary_key;auto_increment;not_null" json:"id"`
	Name string `json:"name"`//真实姓名
	IDcard string `json:"idcard" gorm:"unique"`//身份证号
	Level int64 `json:"level"`//等级
	Score int64 `json:"score"`//信誉分
	Phone string `json:"phone"`//电话
	State int64 `json:"state"` // 1->在业 2->学生 3->失业
	Promise bool `json:"promise"` //信用与否
	//逾期记录关联表进行设计
	Username string `json:"user_name" gorm:"not_null;unique"`
	Password string `json:"password"`
	Money int64 `json:"money"`
	Created  int64 `gorm:"autoCreateTime"`//创建时间
	Updated  int64 `gorm:"autoUpdateTime"`//修改时间
	DeletedAt soft_delete.DeletedAt//删除时间，数据库存档一段时间，可以做个定时脚本定时删除
	
}