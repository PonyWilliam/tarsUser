package main

import (
	"context"
)

// LoginImp servant implementation
type LoginImp struct {
}

// Init servant init
func (imp *LoginImp) Init() error {
	//initialize servant here:
	//...
	return nil
}

// Destroy servant destory
func (imp *LoginImp) Destroy() {
	//destroy servant here:
	//...
}

func (imp *LoginImp) Login(ctx context.Context,username string, password string) (bool,error){
	ok := rp.Login(username,password)
	return ok,nil
}
func (imp *LoginImp) Registry (ctx context.Context,name string, idcard string, phone string, username string, password string) (bool,error) {
	_,err := rp.Registry(name,idcard,phone,username,password)
	if err != nil{
		return false,err
	}
	return true, nil
}
