package utils

import (
	"github.com/jinzhu/gorm"
	"os"
	"server/models"
	"flag"
)

func MigrateTable(db *gorm.DB)  {
	flag.String("migrate","","更新数据库表")
	flag.Parse()
	if flag.Arg(0) != "migrate"{
		return
	}

	//更新数据表
	db.Set("gorm:table_options", "ENGINE=InnoDB").Set("gorm:table_options", "charset=utf8").AutoMigrate(
		&models.User{},
		&models.Subject{},
		&models.SubjectComment{},
	)
	os.Exit(0)
	return
}
