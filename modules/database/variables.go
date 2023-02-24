package database

import "gorm.io/gorm"

var (
	DatabaseConf TDatabase = TDatabase{}
	Db           *gorm.DB
)
