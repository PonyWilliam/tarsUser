package main

import (
	"fmt"
	"log"

	"github.com/TarsCloud/TarsGo/tars"

	"github.com/PonyWilliam/tarsUser/tars-protocol/Users"
)

func main() {
	comm := tars.NewCommunicator()
	obj := fmt.Sprintf("Users.Users.LoginObj@tcp -h 127.0.0.1 -p 10015 -t 60000")
	app := new(Users.Login)
	comm.StringToProxy(obj, app)
	ret,err := app.Registry("谭俊伟","431103200009260311","17794516068","tanjunwei","xiaowei123")
	fmt.Println(ret)
	if err != nil{
		log.Fatal(err)
	}
	
	// var out, i int32
	// i = 123
	// ret, err := app.Add(i, i*2, &out)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(ret, out)
}
