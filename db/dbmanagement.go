package db

import (
    "fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
    "user-management/models"
)

//الوصول الي قاعد البيانات عن تريط المتغير
var DB *gorm.DB

func ConnectDatabase() {
	username := "root"
	password := "mysql"
	host := "127.0.0.1"
	port := "3306"
	dbname := "user_management_db"

	// 1. الاتصال بدون تحديد قاعدة البيانات (لإنشائها أولاً)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to MySQL server: " + err.Error())
	}

	// 2. إنشاء قاعدة البيانات إذا لم تكن موجودة
	createDBQuery := "CREATE DATABASE IF NOT EXISTS " + dbname
	if err := database.Exec(createDBQuery).Error; err != nil {
		panic("Failed to create database: " + err.Error())
	}
	fmt.Printf("Database '%s' is ready\n", dbname)

	// 3. الاتصال بقاعدة البيانات المحددة
	dsnWithDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbname)

	DB, err = gorm.Open(mysql.Open(dsnWithDB), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	// 4. AutoMigrate لإنشاء الجداول
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println("Migration Error:", err)
	}

	fmt.Println("Database connection successfully connected!")
}