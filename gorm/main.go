package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type User struct {
	ID       *string `gorm:"default:'uuid'"`
	Name     string  `gorm:"default:'galeone'"`
	Age      int64
	Birthday time.Time
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@/hlj?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		logrus.Info(err)
		os.Exit(1)
	}
	user := User{Name: "jinzhu", Age: 18, Birthday: time.Now()}
	db.NewRecord(user)
	db.Create(&user)
	db.NewRecord(user)
	//db.Set("gorm:insert_option", "ON CONFLICT").Create(&user)
	//通过主键查询第一条记录
	user1 := User{}
	db.First(&user1)
	logrus.Info(user1)
	//// SELECT * FROM users ORDER BY id LIMIT 1;

	// Get one record, no specfied order
	user2 := User{}
	db.Take(&user2)
	logrus.Info(user2)
	//// SELECT * FROM users LIMIT 1;

	// Get last record, order by primary key
	user3 := User{}
	db.Last(&user3)
	logrus.Info(user3)
	//// SELECT * FROM users ORDER BY id DESC LIMIT 1;

	var users []User
	// Get all records
	db.Find(&users)
	logrus.Info(users)
	//// SELECT * FROM users;

	// Get record with primary key (only works for integer primary key)
	user4 := User{}
	db.First(&user4, 10)
	logrus.Info(user4)
	//// SELECT * FROM users WHERE id = 10;

	db.Create(User{Name: "z", Age: 23, Birthday: time.Now()})
	defer db.Close()
}
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	uuidID := uuid.NewV4().String()
	//_ = scope.SetColumn("id", uuidID)
	user.ID = &uuidID
	return nil
}
