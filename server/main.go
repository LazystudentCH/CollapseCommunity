package main

import (
	"server/services"
	"server/utils"
)

func init() {
	//读取配置文件
	utils.LoadConfig()
	//链接数据库
	utils.ConnDb()
	//解析参数是否需要更新数据库表
	utils.MigrateTable(utils.GetMysqlMasterConn())
	return
}
type A struct {
	Name string
	Age int
	Gender string
}

func main() {
	//启动Http服务
	go services.RunHttpService()

	services.RunHotSubjectService()

}
