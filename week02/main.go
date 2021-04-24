package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pkg/errors"
)

var db *gorm.DB

func init() {
	// 数据库连接
	var err error
	db, err = gorm.Open("mysql", "test:test@(localhost)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("mysql connect error: %s", err)
		panic(err)
	}
}

// User 用户
type User struct {
	ID   uint
	Name string
}

// GetByID 通过Id搜索记录
func (user *User) GetByID(id int) error {
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return errors.Wrap(err, fmt.Sprintf("filter by id=%d", id))
	}
	return nil
}

func main() {
	defer db.Close()

	user := new(User)
	err := user.GetByID(1)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 一些降级处理等
			user.Name = "anon"
		} else {
			fmt.Printf("Get user error, %s\n", err)
			return
		}
	}
	fmt.Printf("The user is: %s\n", user.Name)
}
