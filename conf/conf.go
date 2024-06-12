package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"toDoListDemo/model"
)

// 用这些变量接收配置文件里面的值
var (
	AppMode  string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	RedisAddr   string
	RedisPw     string
	RedisDbName string
)

func Init() {
	file, err := ini.Load("D:\\godemo\\toDoListDemo\\conf\\config.ini") //读取配置文件,这样就可以通过file来访问这个文件中的各个参数了
	if err != nil {
		log.Println("配置文件读写错误，请检查文件路径:", err)
		panic(err)
	}

	//mysql路径
	path := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "123456", "127.0.0.1:3306", "todolist_db")
	model.DatabaseConnect(path) //连接数据库

	//读取file中的服务器，Mysql，Redis的配置信息
	LoadServer(file)
	LoadMysqlData(file)
	LoadRedis(file)
}

// 这些变量名在本文件的最上方已经定义了
func LoadServer(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}

func LoadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}

func LoadRedis(file *ini.File) {
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
}
