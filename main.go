package main

import (
	"log"
	"user-management/db"
	"user-management/handlers"
	"user-management/models"

	"github.com/gin-gonic/gin"
)

func main() {
	//  الاتصال بقاعدة البيانات
	db.ConnectDatabase()

	
	err := db.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("فشل في تحديث قاعدة البيانات:", err)
	}

	//  إنشاء محرك Gin
	r := gin.Default()

	//  تحميل ملفات HTML والملفات الثابتة (CSS/JS)
	r.LoadHTMLGlob("templates/*.html")
	r.Static("/static", "./static")

	//  تعريف المسارات (Routes)
	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", handlers.ListUsers)            // عرض القائمة
		userRoutes.GET("/create", handlers.ShowCreateForm) // صفحة انشاء مستخدم
		userRoutes.POST("/create", handlers.CreateUser)    // انشاء مستخدم
		userRoutes.GET("/edit/:id", handlers.ShowEditForm) // صفحة التعديل
		userRoutes.POST("/update/:id", handlers.UpdateUser) // اجراء التعديل
		userRoutes.GET("/delete/:id", handlers.DeleteUser) // مسح مستخدم
		
	}

	//6. تشغيل السيرفر على المنفذ 8080
	log.Println("السيرفر يعمل على http://localhost:8080/users")
	r.Run(":8080")
}
