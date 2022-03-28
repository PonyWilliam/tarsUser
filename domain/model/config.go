package model

import (
	"log"

	"github.com/Andrew-M-C/tarsgo-tools/config"
)
type Db struct{
	Host string
	Port int
	User string
	Pwd string
	Dbname string
}

func GetMysqlConfig()*Db{
	db := &Db{}
	//从配置中心获取所有需要的配置，下面是mysql配置。
	tarsconfg,err := config.NewConfig()
	if err != nil{
		log.Fatal(err)
	}
	db.Host,_ = tarsconfg.GetString("/tars/db","dbhost","")
	db.Port,_ = tarsconfg.GetInt("tars/db","dbport")
	db.User,_ = tarsconfg.GetString("/tars/db","dbuser","")
	db.Pwd,_ = tarsconfg.GetString("/tars/db","dbpass","")
	db.Dbname,_ = tarsconfg.GetString("/tars/db","dbname","")
	return db
}