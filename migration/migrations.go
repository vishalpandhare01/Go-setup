package migration

import (
	"github.com/vishalpandhare01/school_management_system/db"
	"github.com/vishalpandhare01/school_management_system/internals/model"
)

func Migration() error {
	DB := db.DatabaseConnection()
	DB.AutoMigrate(&model.Programing{})
	return nil
}
