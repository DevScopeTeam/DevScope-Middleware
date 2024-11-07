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
	db.AutoMigrate(&model.Tag{})
	db.AutoMigrate(&model.Domain{})

	// 执行原生 SQL 以创建外键
	// db.Exec("ALTER TABLE domains ADD CONSTRAINT fk_domain_tag FOREIGN KEY (tag_uuid) REFERENCES tags(uuid) ON DELETE CASCADE ON UPDATE CASCADE;")
	// db.Exec("ALTER TABLE domains ADD CONSTRAINT fk_domain_developer FOREIGN KEY (username) REFERENCES developer_ranks(username) ON DELETE CASCADE ON UPDATE CASCADE;")
}
