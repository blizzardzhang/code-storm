package enter

import (
	"code-storm/rpc/model/usermodel"
	"code-storm/rpc/user/ent"
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
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

func InitEnt(Datasource string) {
	client, err := ent.Open("mysql", Datasource)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
