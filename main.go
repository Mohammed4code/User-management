package main

import (
	"log"
	"user-management/db"
	"user-management/handlers"
	"user-management/models"
	"github.com/gin-gonic/gin"
)

func main() {
	// الاتصال بقاعدة البيانات
	db.ConnectDatabase()

	err := db.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("فشل في تحديث قاعدة البيانات:", err)
	}

	
	gin.SetMode(gin.ReleaseMode)

	
	r := gin.Default()

	
	r.SetTrustedProxies(nil)

	
	r.LoadHTMLFiles(
		"templates/index.html",
		"templates/create.html",
		"templates/edit.html",
	)

	r.Static("/static", "./static")

	// تعريف المسارات (Routes)
	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", handlers.ListUsers)            
		userRoutes.GET("/create", handlers.ShowCreateForm) 
		userRoutes.POST("/create", handlers.CreateUser)   
		userRoutes.GET("/edit/:id", handlers.ShowEditForm) 
		userRoutes.POST("/update/:id", handlers.UpdateUser) 
		userRoutes.GET("/delete/:id", handlers.DeleteUser) 
	}

	// تشغيل السيرفر على المنفذ 8080
	log.Println("السيرفر يعمل على http://localhost:8080/users")
	r.Run(":8080")
}