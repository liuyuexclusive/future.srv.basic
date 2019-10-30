package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name            string
	Mobile          string
	Email           string
	Pwd             string
	Salt            string
	Access          string
	Avatar          string
	LastOn          *time.Time
	UserRoleMapList []UserRoleMap
}

type Role struct {
	gorm.Model
	Name            string
	UserRoleMapList []UserRoleMap
}

type UserRoleMap struct {
	ID     uint
	UserID uint
	RoleID uint
	User   User
	Role   Role
}

type Message struct {
	gorm.Model
	From          string
	Title         string
	Content       string
	MessageToList []MessageTo
}

type MessageTo struct {
	gorm.Model
	MessageID uint
	To        string
	Status    uint
	Message   Message
}
