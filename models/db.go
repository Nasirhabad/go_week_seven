package models

import (
	"week7/config"
)

func MigrateDB() {
	config.DB.AutoMigrate(&User{}, &Package{})
}
