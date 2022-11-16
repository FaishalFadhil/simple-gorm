package models

import (
	"gorm.io/gorm"
	"time"
)

// type Car struct{
// 	Merk string `json:"merk"`
// 	Harga int `json:"harga"`
// 	Typecars string `json:"typecars"`
// 	Id uint `json:"id"`
// }

type Car struct {
	Id        uint   `gorm:"PRIMARY_KEY"`
	Brand     string `gorm:"column:brand;varchar(100)"`
	Price     int    `gorm:"column:price;integer(11)"`
	Type      string `gorm:"column:type;varchar(100)"`
	Owner     []User `gorm:"foreignKey:UserId;association_foreignkey:Id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	Id        uint   `gorm:"PRIMARY_KEY"`
	UserId    uint   `gorm:"column:user_id"`
	Name      string `gorm:"column:name;varchar(100)"`
	Address   string `gorm:"column:address;text"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// type User struct {
// 	UserId    int     `gorm:"PRIMARY_KEY"`
// 	Email     string  `gorm:"column:email"`
// 	FirstName string  `gorm:"column:firstname"`
// 	LastName  string  `gorm:"column:lastname"`
// 	Assets    []Asset `gorm:"foreignkey:user_id"`
// }

// type Asset struct {
// 	AssetId int    `gorm:"PRIMARY_KEY"`
// 	UserId  int    `gorm:"column:user_id"`
// 	Slug    string `gorm:"column:slug"`
// 	Address string `gorm:"address"`
// }
