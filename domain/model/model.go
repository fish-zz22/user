package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	//主键
	ID int64 `gorm:"primary_key;not_null;auto_increment"`
	//用户名称
	UserName string `gorm:"unique_index;not_null"`
	//添加需要的字段
	FirstName string
	//...
	//密码
	HashPassword string
}

type IUserRepository interface {
	//初始化数据表
	InitTable() error
	//根据用户名称查找用户信息
	FindUserByName(string) (*User,error)
	//根据用户ID查找用户信息
	FindUserByID(int64) (*User,error)
	//创建用户
	CreateUser(*User) (int64,error)
	//根据用户ID删除用户
	DeleteUserByID(int64) error
	//更新用户信息
	UpdateUser(*User) error
	//查找所有用
	FindAll() ([]*User,error)
}

//创建UserRepository
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{mysqlDb: db}
}

type UserRepository struct {
	mysqlDb *gorm.DB
}

//初始化表
func (u *UserRepository) InitTable() error  {
	return u.mysqlDb.CreateTable(&User{}).Error
}
//根据用户名称查找用户信息
func (u *UserRepository)FindUserByName(name string) (user *User,err error){
	user = &User{}
	return user,u.mysqlDb.Where("user_name = ?",name).Find(user).Error
}

//根据用户ID查找用户信息
func (u *UserRepository)FindUserByID(userID int64) (user *User,err error){
	user = &User{}
	return user,u.mysqlDb.First(user,userID).Error
}

//创建用户
func (u *UserRepository)CreateUser(user *User) (userID int64,err error) {
	return user.ID,u.mysqlDb.Create(user).Error
}

//根据用户ID删除用户
func (u *UserRepository)DeleteUserByID(userID int64) error {
	return u.mysqlDb.Where("id = ?",userID).Delete(&User{}).Error
}
//更新用户信息
func (u *UserRepository)UpdateUser(user *User) error {
	return u.mysqlDb.Model(user).Update(&user).Error
}

//查找所有用
func (u *UserRepository)FindAll() (userAll []*User,err error){
	return userAll,u.mysqlDb.Find(&userAll).Error
}