package repository

import (
	"crypto/md5"
	"fmt"
	"io"

	"github.com/PonyWilliam/tarsUser/domain/model"
	"gorm.io/gorm"
)

// 对外暴露接口

var db *gorm.DB

func init(){

	
}

type IUser interface{
	Init() error
	Login(username string,password string)(bool) //登陆接口
	Registry(name string,idcard string,phone string,username string,password string)(int64,error) // 注册接口
	FindByID(id int64)*model.User
	FindByIDCard(idcard string)*model.User
	Find()[]model.User
	FindByName(name string)[]model.User
	FindByUserName(username string)*model.User
	DelByID(id int64)(error)
	UpdateUser(user model.User)(error)
}
type UserRepo struct{
	Db *gorm.DB
}
func InitUser(db *gorm.DB)IUser{
	//判定是否有表，如果没有就新建，否则就直接使用
	u := &UserRepo{Db: db}
	return u
}

func(u UserRepo)Init() error{
	if(!u.Db.Migrator().HasTable(&model.User{})){
		//创建表
		fmt.Println("创建表")
		u.Db.Migrator().CreateTable(&model.User{})
	}else{
		fmt.Println("无需创建")
	}
	return nil
}

func (u UserRepo)Login(username string,password string)(bool){
	//token，err
	//登陆结果返回
	user := &model.User{}
	u.Db.Where("username = ?",username).First(&user)
	if user.Password == EncodeMD5(password){
		return true
	}
	return false
}
func (u UserRepo)Registry(name string,idcard string,phone string,username string,password string)(int64,error){
	md5_pwd := EncodeMD5(password)//明文密码转成md5
	user := &model.User{}
	user.Level = 5 //默认等级
	user.Name = name
	user.IDcard = idcard
	user.Username = username
	user.Password = md5_pwd
	user.Promise = true//默认信用等级
	user.Phone = phone
	user.Money = 100000 //注册免费送10w块钱进行秒杀
	result := u.Db.Create(&user)
	if result.Error != nil{
		return -1,result.Error
	}
	return user.ID,nil
}

func(u UserRepo)FindByID(id int64)*model.User{
	user := &model.User{}
	u.Db.Where("id = ?",id).First(&user)
	return user
}
func(u UserRepo)FindByIDCard(idcard string)*model.User{
	user := &model.User{}
	u.Db.Where("idcard = ?",idcard).First(&user)
	return user
}
func(u UserRepo)Find()[]model.User{
	var user []model.User
	u.Db.Find(&user)
	return user
}
func(u UserRepo)FindByName(name string)[]model.User{
	var user []model.User
	u.Db.Where("name = ?",&user)
	return user
}
func(u UserRepo)FindByUserName(username string)*model.User{
	user := &model.User{}
	u.Db.Where("username = ?",&user)
	return user
}
func(u UserRepo)DelByID(id int64)(error){
	u.Db.Delete(&model.User{},id)//根据主键删除
	return nil
}
func(u UserRepo)UpdateUser(user model.User)(error){
	// Name string `json:"name"`//真实姓名
	// IDcard string `json:"idcard" gorm:"unique"`//身份证号
	// Level int64 `json:"level"`//等级
	// Score int64 `json:"score"`//信誉分
	// Phone string `json:"phone"`//电话
	// State int64 `json:"state"` // 1->在业 2->学生 3->失业
	// Promise bool `json:"promise"` //信用与否
	// //逾期记录关联表进行设计
	// Username string `json:"user_name" gorm:"not_null;unique"`
	// Password string `json:"password"`
	// Money int64 `json:"money"`
	db.Model(&user).Updates(user)
	return nil
}

func EncodeMD5(pwd string)string{
	h := md5.New()
	_, _ = io.WriteString(h,pwd)
	sum := fmt.Sprintf("%x",h.Sum([]byte("sanxiangbank")))//独立签名
	return sum
}

