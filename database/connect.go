package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	_, err := gorm.Open(mysql.Open("root:root@/golangreact"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database.")
	}

}
