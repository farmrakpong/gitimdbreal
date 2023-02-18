package userRepository

import (
	_ "database/sql"
	"fmt"

	db "github.com/farmrakpong/goimdbreal/db"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

//	type Register struct {
//		ID           string `json:"id"`
//		Name         string `json:"name"`
//		TotalProduct string `json:"totalProduct"`
//	}
type User struct {
	gorm.Model
	ID           uint `gorm:"primaryKey;autoIncrement"`
	Name         string
	TotalProduct string
}

func UserRepository() string {
	c := db.Connect()
	err := c.AutoMigrate(&User{})
	if err != nil {
		fmt.Println("<----err UserRepository----->")
		return "err"
	}
	v := c.Create(&User{Name: "D42", TotalProduct: "100"})
	// fmt.Println(v.Error)
	if v.Error != nil {
		fmt.Println("<----err CreateUser----->")
		return "err"
	}
	return "ok"
}
