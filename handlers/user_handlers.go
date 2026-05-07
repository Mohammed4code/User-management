package handlers

import (
	"fmt"
	"net/http"
	"user-management/db"
	"user-management/models"

	"github.com/gin-gonic/gin"
)

// 1. عرض قائمة المستخدمين (GET /users)
func ListUsers(c *gin.Context) {
	var users []models.User
	db.DB.Find(&users)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"users": users,
	})
}

// 2. عرض صفحة النموذج للإضافة (GET /users/create)  
func ShowCreateForm(c *gin.Context) {
	c.HTML(http.StatusOK, "create.html", nil)
}

// 3. معالجة بيانات الإضافة (POST /users/create)
func CreateUser(c *gin.Context) {
	var input models.AddUser

	if err := c.ShouldBind(&input); err != nil {
		c.HTML(http.StatusBadRequest, "create.html", gin.H{"Error": "تأكد من صحة البيانات: " + err.Error()})
		return
	}

	fmt.Println("Received data:")
	fmt.Println("Fullname:", input.Fullname)
	fmt.Println("Email:", input.Email)
	fmt.Println("Phonenumber:", input.Phonenumber)
	fmt.Println("Position:", input.Position)
	fmt.Println("Age:", input.Age)

	// تحقق من عدم وجود بيانات فارغة
	if input.Fullname == "" || input.Email == "" || input.Age == 0 {
		c.HTML(http.StatusBadRequest, "create.html", gin.H{"Error": "الاسم والبريد الإلكتروني والعمر مطلوبة"})
		return
	}

	user := models.User{
		Fullname:    input.Fullname,
		Email:       input.Email,
		Phonenumber: input.Phonenumber,
		Position:    input.Position,
		Age:         input.Age,
	}

	if err := db.DB.Create(&user).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "create.html", gin.H{"Error": "تعذر حفظ البيانات: " + err.Error()})
		return
	}

	c.Redirect(http.StatusSeeOther, "/users")
}

// 4. عرض صفحة التعديل (GET /users/edit/:id)
func ShowEditForm(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := db.DB.First(&user, id).Error; err != nil {
		c.HTML(http.StatusNotFound, "index.html", gin.H{"error": "المستخدم غير موجود"})
		return
	}
	c.HTML(http.StatusOK, "edit.html", gin.H{
		"user": user,
	})
}

// 5. معالجة التعديل (POST /users/update/:id)
func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	// جلب المستخدم الحالي من قاعدة البيانات
	var user models.User
	if err := db.DB.First(&user, id).Error; err != nil {
		c.HTML(http.StatusNotFound, "edit.html", gin.H{"Error": "المستخدم غير موجود"})
		return
	}

	// ربط البيانات المدخلة
	var input models.UpdateUser
	if err := c.ShouldBind(&input); err != nil {
		c.HTML(http.StatusBadRequest, "edit.html", gin.H{"Error": "بيانات غير صالحة"})
		return
	}

	// تحديث الحقول إذا كانت موجودة
	if input.Fullname != nil && *input.Fullname != "" {
		user.Fullname = *input.Fullname
	}
	if input.Email != nil && *input.Email != "" {
		user.Email = *input.Email
	}
	if input.Phonenumber != nil {
		user.Phonenumber = *input.Phonenumber
	}
	if input.Position != nil {
		user.Position = *input.Position
	}
	if input.Age != nil && *input.Age > 0 {
		user.Age = *input.Age
	}

	// حفظ التحديثات في قاعدة البيانات
	if err := db.DB.Save(&user).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "edit.html", gin.H{"Error": "فشل تحديث البيانات: " + err.Error()})
		return
	}

	c.Redirect(http.StatusSeeOther, "/users")
}

// 6. حذف المستخدم (GET /users/delete/:id)
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	db.DB.Unscoped().Delete(&models.User{}, id)
	c.Redirect(http.StatusSeeOther, "/users")
}