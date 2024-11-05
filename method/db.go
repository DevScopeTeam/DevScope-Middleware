package method

import (
	"DevScope-Middleware/config"
	"DevScope-Middleware/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func getDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(config.DSN), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	return db, err
}

func init() {
	db, err := getDB()
	sqlDB, _ := db.DB() //结束后关闭 DB
	defer sqlDB.Close()
	if err != nil {
		fmt.Println(err)
	}

	// 迁移 schema
	db.AutoMigrate(&model.DeveloperRank{})

	// 执行原生 SQL 以创建外键
}
