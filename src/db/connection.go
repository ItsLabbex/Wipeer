package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type DBInit struct {
	Host     string
	User     string
	Password string
	DBname   string
}

func (i *DBInit) ConnectionDB() {
	// String de la base de datos
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", i.User, i.Password, i.Host, i.DBname)

	// Realizamos la conexion con la base de datos
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	// Comprobamos que la conexion con la base de datos es correcta
	if err != nil {
		panic(err)
	}
}
