package main

import (
	"fmt"
	"log"

	"github.com/PonyWilliam/tarsUser/domain/model"
	"github.com/PonyWilliam/tarsUser/domain/repository"
	"github.com/PonyWilliam/tarsUser/tars-protocol/Users"
	"github.com/TarsCloud/TarsGo/tars"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var rp repository.IUser

func init(){
	db_data := model.GetMysqlConfig()
	//初始化数据库连接，并对数据库表进行检查
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",db_data.User,db_data.Pwd,db_data.Host,db_data.Port,db_data.Dbname)
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatal(err)
	}
	//传递db给User Repository进行维护
	rp = repository.InitUser(db)
	_ = rp.Init()
}

func main() {


	mux := &tars.TarsHttpMux{}
	r := gin.Default()
	r.GET("/",func(ctx *gin.Context){
		ctx.JSON(200,gin.H{
			"code":200,
			"msg":"ok",
		})
	})
	mux.Handle("/",r)
    // Get server config
    cfg := tars.GetServerConfig()
	tars.AddHttpServant(mux,cfg.App + "." + cfg.Server + ".HttpObj")
    // New servant imp
    imp := new(LoginImp)
    // New servant
    app := new(Users.Login)
    // Register Servant
    app.AddServantWithContext(imp, cfg.App+"."+cfg.Server+".LoginObj")
    // Run application
    tars.Run()
}