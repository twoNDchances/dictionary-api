package models

import (
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type Database struct {
	Username string
	Password string
	Hostname string
	Port     int
	Database string
}

func (d *Database) validateArgs() []error {
	errs := make([]error, 0)
	mapStringArgs := map[string]string{
		"username": d.Username,
		"password": d.Password,
		"hostname": d.Hostname,
		"database": d.Database,
	}
	for key, value := range mapStringArgs {
		if value == "" {
			errs = append(errs, fmt.Errorf("%s can't leave blank", key))
		}
	}
	mapIntegerArgs := map[string]int{
		"port": d.Port,
	}
	for key, value := range mapIntegerArgs {
		if value <= 0 {
			errs = append(errs, fmt.Errorf("%s can't less than or equal 0", key))
		}
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func (d *Database) getDsn() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		d.Username,
		d.Password,
		d.Hostname,
		d.Port,
		d.Database,
	)
}

func (d *Database) connect() (*gorm.DB, error) {
	if errs := d.validateArgs(); errs != nil {
		return nil, fmt.Errorf("error when validate database args: %w", errors.Join(errs...))
	}
	DB, err := gorm.Open(mysql.Open(d.getDsn()), &gorm.Config{
		Logger: logger.Discard,
	})
	if err != nil {
		return nil, fmt.Errorf("error when connect to database: %s", err.Error())
	}
	return DB, nil
}

func (d *Database) Init() (err error) {
	database := Database{
		Username: d.Username,
		Password: d.Password,
		Hostname: d.Hostname,
		Port:     d.Port,
		Database: d.Database,
	}
	DB, err = database.connect()
	if err != nil {
		return fmt.Errorf("error when init database: %s", err.Error())
	}
	DB.AutoMigrate(&Dictionary{})
	return nil
}
