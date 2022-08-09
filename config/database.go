package config

import (
	"fmt"
	model "task-scheduler/datamodel"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DbConn *gorm.DB

func DbSetup() error {
	ds := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", Config.DbUser, Config.DbPassword, Config.DbHost, Config.DbName, "charset=utf8&loc=Asia%2FJakarta&parseTime=true")
	db, err := gorm.Open(mysql.Open(ds), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	sqlDB.SetConnMaxIdleTime(time.Minute * 5)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
	DbConn = db

	return err
}

func GetDatabaseConnection() (*gorm.DB, error) {
	sqlDB, err := DbConn.DB()
	if err != nil {
		return DbConn, err
	}
	if err := sqlDB.Ping(); err != nil {
		return DbConn, err
	}
	return DbConn, nil
}

func AutoMigrateDB() error {
	// Auto migrate database
	db, err := GetDatabaseConnection()
	if err != nil {
		return err
	}
	// Add required models here
	err = db.AutoMigrate(&model.Staff{}, &model.Role{}, &model.Task{}, &model.User{}, &model.CodeGenerator{})
	// Example for migrating multiple models
	// err:= db.AutoMigrate(&models.User{}, &models.Admin{}, &models.Guest{})
	return err
}
