package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// 1. الموديل الأساسي (قاعدة البيانات)
type User struct {
	gorm.Model                
	Fullname     string  `gorm:"size:255"`
	Phonenumber  string  `gorm:"size:20"`
	Email        string  `gorm:"unique"`
	Position     string 
	Password     string 
}

// 2. موديل الإضافة (Request Body)
type AddUser struct {
    Fullname     string `form:"Fullname" binding:"required"`   
    Phonenumber  string `form:"Phonenumber"`
    Email        string `form:"Email" binding:"required"`         
    Position     string `form:"Position"`
    Password     string `form:"Password" binding:"required"`     
}


type Displayuser struct {
    ID         uint   `json:"id"`
    Fullname   string `json:"fullname"`
	Position   string `json:"position"`
}

// 4. موديل التعديل (Update Request)
type UpdateUser struct {
    Fullname    *string `form:"Fullname"`     
    Phonenumber *string `form:"Phonenumber"`  
    Position    *string `form:"Position"`     
    Email       *string `form:"Email"`        
    Password    *string `form:"Password"`     
}



// دالة لتشفير كلمة المرور
func (u *User) HashPassword() error {
    hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    u.Password = string(hashed)
    return nil
}


func (u *User) BeforeCreate(tx *gorm.DB) error {
    if u.Password != "" {
        return u.HashPassword()
    }
    return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
    // تحقق إذا تم تغيير كلمة المرور
    if tx.Statement.Changed("Password") && u.Password != "" {
        return u.HashPassword()
    }
    return nil
}