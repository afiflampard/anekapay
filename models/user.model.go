package models

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type User struct {
	gorm.Model
	Email    string `gorm:"column:email; type:varchar(255); not null; unique" json:"email"`
	Name     string `gorm:"column:name; type:varchar(255); not null" json:"name"`
	Mobile   string `gorm:"column:mobile; type:varchar(255) not null " json:"mobile"`
	Address  string `gorm:"column:alamat; type:varchar(255)" json:"alamat"`
	Password string `gorm:"column:password; type:varchar(255); not null " json:"password"`
	Photo    string `gorm:"column:photo; type: varchar(255)" json:"photo"`
	RoleID   uint   `gorm:"column:roleId" json:"roleId"`
	Role     Role   `gorm:"foreignKey:RoleID"`
}

type Role struct {
	gorm.Model
	Role string `gorm:"column:role"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID    uint64 `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}
