package main

import (
	"context"
	"fmt"

	"github.com/PonyWilliam/tarsUser/domain/model"
	"github.com/PonyWilliam/tarsUser/tars-protocol/Users"
)
type ManagerUserImp struct {
}

func (imp *ManagerUserImp) Init() error {
	//initialize servant here:
	//...
	return nil
}

// Destroy servant destory
func (imp *ManagerUserImp) Destroy() {
	//destroy servant here:
	//...
}

func (imp *ManagerUserImp)FindByID(ctx context.Context,id int32, u *Users.User) (ret bool, err error){
	user := rp.FindByID(int64(id))
	if user != nil{
		ChangeToTarsUser(u,user)
	}else{
		u = nil
	}
	return true,nil
}
func (imp *ManagerUserImp)FindByIDCard(ctx context.Context,idcard string, u *Users.User) (ret bool, err error){
	user := rp.FindByIDCard(idcard)
	if user != nil{
		ChangeToTarsUser(u,user)
	}else{
		u = nil
	}
	return true,nil
}
func (imp *ManagerUserImp)Find(ctx context.Context,u *[]Users.User)(ret bool, err error){
	users := rp.Find()
	for _,v:=range users{
		var temp Users.User
		ChangeToTarsUser(&temp,&v)
		fmt.Println(temp)
		*u = append(*u, temp)
	}
	return true,nil
}
func (imp *ManagerUserImp)FindByName(ctx context.Context,name string,u *[]Users.User)(ret bool, err error){
	users := rp.FindByName(name)
	fmt.Println(users)
	for _,v:=range users{
		var temp Users.User
		ChangeToTarsUser(&temp,&v)
		*u = append(*u, temp)
	}
	return true,nil
}
func (imp *ManagerUserImp)FindByUserName(ctx context.Context,username string, u *Users.User) (ret bool, err error){
	user := rp.FindByUserName(username)
	if user != nil{
		ChangeToTarsUser(u,user)
	}else{
		u = nil
	}
	return true,nil
}
func (imp *ManagerUserImp)DelByID(ctx context.Context,id int32)(ret bool,err error){
	rp.DelByID(int64(id))
	return true,nil
}
func (imp *ManagerUserImp)UpdateUser(ctx context.Context,u *Users.User)(ret bool,err error){
	var temp model.User
	ChangeToModelUser(&temp,u)
	rp.UpdateUser(temp)
	return true,nil
}
func ChangeToModelUser(u1 *model.User,u2 *Users.User){
	u1.ID = int64(u2.Id)
	u1.Name = u2.Name
	u1.IDcard = u2.Idcard
	u1.Level = int64(u2.Level)
	u1.Score = int64(u2.Score)
	u1.Phone = u2.Phone
	u1.State = int64(u2.State)
	u1.Promise = u2.Promise
	u1.Username = u2.Username
	u1.Password = "NULL"
	u1.Money = int64(u2.Money)
}
func ChangeToTarsUser(u1 *Users.User,u2 *model.User){
	u1.Id = int32(u2.ID)
	u1.Name = u2.Name
	u1.Idcard = u2.IDcard
	u1.Level = int32(u2.Level)
	u1.Score = int32(u2.Score)
	u1.Phone = u2.Phone
	u1.State = int32(u2.State)
	u1.Promise = u2.Promise
	u1.Username = u2.Username
	u1.Money = int32(u2.Money)
}