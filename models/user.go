package models

import (
	"gorm.io/gorm"
)

// الموديل الأساسي (قاعدة البيانات)
type User struct {
	gorm.Model                
	Fullname     string  `gorm:"size:255"`
	Phonenumber  string  `gorm:"size:20"`
	Email        string  `gorm:"unique"`
	Position     string 
	Age          int     `gorm:"default:0"`  
}

// موديل الإضافة (Request Body)
type AddUser struct {
	Fullname     string `form:"Fullname" binding:"required"`   
	Phonenumber  string `form:"Phonenumber"`
	Email        string `form:"Email" binding:"required"`         
	Position     string `form:"Position"`
	Age          int    `form:"Age" binding:"required"`    
}

type Displayuser struct {
	ID         uint   `json:"id"`
	Fullname   string `json:"fullname"`
	Position   string `json:"position"`
	Age        int    `json:"age"`
}

// موديل التعديل (Update Request)
type UpdateUser struct {
	Fullname    *string `form:"Fullname"`     
	Phonenumber *string `form:"Phonenumber"`  
	Position    *string `form:"Position"`     
	Email       *string `form:"Email"`        
	Age         *int    `form:"Age"`        
}