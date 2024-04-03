package enter

import (
	"code-storm/rpc/model/usermodel"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitGorm(MysqlDatasource string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(MysqlDatasource), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connect to database success")

	// Migrate the schema
	e := db.AutoMigrate(&usermodel.User{})
	if e != nil {
		return nil
	}
	return db
}
