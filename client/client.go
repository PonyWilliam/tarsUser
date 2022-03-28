package main

import (
	"fmt"

	"github.com/TarsCloud/TarsGo/tars"

	"github.com/PonyWilliam/tarsUser/tars-protocol/Users"
)
var app1 *Users.Login
var app2 *Users.ManageUser
var loginobj string
var ManageUserobj string
func RegistryTest(){
	comm := tars.NewCommunicator()
	comm.StringToProxy(loginobj, app1)
	ret,err := app1.Registry("谭俊伟","431103200009260311","17794516068","tanjunwei","xiaowei123")
	fmt.Println(ret)
	if err != nil{
		fmt.Println(err)
	}
}

func LoginTest(){
	comm := tars.NewCommunicator()
	comm.StringToProxy(loginobj, app1)
	ret,err := app1.Login("tanjunwei","xiaowei123")
	fmt.Println(ret)
	if err != nil{
		fmt.Println(err)
	}
}
func FindByIDTest(id int32){
	comm := tars.NewCommunicator()
	comm.StringToProxy(ManageUserobj, app2)
	user := &Users.User{}
	_,err := app2.FindByID(id,user)
	fmt.Println(user)
	if err != nil{
		fmt.Println(err)
	}
}
func FindByUserName(username string){
	comm := tars.NewCommunicator()
	comm.StringToProxy(ManageUserobj, app2)
	user := &Users.User{}
	_,err := app2.FindByUserName(username,user)
	fmt.Println(user)
	if err != nil{
		fmt.Println(err)
	}
}
func FindByIDCardTest(idcard string){
	comm := tars.NewCommunicator()
	comm.StringToProxy(ManageUserobj, app2)
	user := &Users.User{}
	_,err := app2.FindByIDCard(idcard,user)
	fmt.Println(user)
	if err != nil{
		fmt.Println(err)
	}
}
func FindTest(){
	comm := tars.NewCommunicator()
	comm.StringToProxy(ManageUserobj, app2)
	var users *[]Users.User
	users = new([]Users.User)
	_,err := app2.Find(users)
	fmt.Println(*users)
	if err != nil{
		fmt.Println(err)
	}
}
func FindNameTest(name string){
	comm := tars.NewCommunicator()
	comm.StringToProxy(ManageUserobj, app2)
	var users *[]Users.User
	users = new([]Users.User)
	_,err := app2.FindByName(name,users)
	fmt.Println(*users)
	if err != nil{
		fmt.Println(err)
	}
}
func UpdateTest(){
	comm := tars.NewCommunicator()
	comm.StringToProxy(ManageUserobj, app2)
	user := &Users.User{}
	user.Id = 1
	user.Name = "tanjunwei"
	app2.UpdateUser(user)
}
func init(){
	app1 = new(Users.Login)
	app2 = new(Users.ManageUser)
	loginobj = fmt.Sprintf("Users.Users.LoginObj@tcp -h 127.0.0.1 -p 10015 -t 60000")
	ManageUserobj = fmt.Sprintf("Users.Users.ManagerUserObj@tcp -h 127.0.0.1 -p 10017 -t 60000")
}
func main() {
	// RegistryTest()	功能正常
	// LoginTest()	功能正常
	// FindByIDTest(1)	功能正常
	// FindTest()	fix -> 功能正常（指针逻辑问题）
	// FindNameTest("谭俊伟")	fix -> 功能正常（rp中数据库语法问题）
	// FindByUserName("tanjunwei")	fix -> 功能正常（rp中数据库语法问题）
	// FindByIDCardTest("431103200009260311")	功能正常
	UpdateTest()
}
