package models

import (
	"github.com/kalyanKumarPokkula/Go-jwt/helpers"
	"gorm.io/gorm"
)


type User struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey;autoIncrement"`
	Avatar string `grom:"type:varchar(100);default:https://via.placeholder.com/200x200.png" json:"avatar"` 
	UserName string `gorm:"type:varchar(100);not null" json:"username"`
	Email string `gorm:"unique;not null;type:varchar(100)" json:"email"`
	Password string `gorm:"type:varchar(100);not null" json:"password"`
	Role helpers.Role `gorm:"type:enum('admin', 'user');default:user" json:"role"`
	IsEmailVerified bool `gorm:"type:boolean;default:false" json:"isemailverified"`
	ForgotPasswordToken string `gorm:"type:varchar(100)" json:"forgotpasswordtoken" `
	ForgotPasswordExpiry string `gorm:"type:varchar(100)" json:"forgotpasswordexpiry"`
	EmailVerificationToken string `gorm:"2type:varchar(100)55" json:"emailverificationtoken"`
	EmailVerificationExpiry string `gorm:"type:varchar(100)" json:"emailverificationexpiry"`

}