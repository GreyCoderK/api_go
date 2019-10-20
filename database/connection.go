package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect(dbUser, dbPwd, dbName string) (*gorm.DB, error) {
	ConnectionString := fmt.Sprintf(
		"%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,
		dbPwd,
		dbName,
	)

	return gorm.Open("mysql", ConnectionString)
}
