package model

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	Username       string `gorm:"unique"`
	Email          string `gorm:"unique"`
	PasswordHashed string
}

func (u User) TableName() string {
	return "user"
}

func GetByEmail(db *gorm.DB, ctx context.Context, email string) (user *User, err error) {
	err = db.WithContext(ctx).Model(&User{}).Where(&User{Email: email}).First(&user).Error
	return
}

func Create(db *gorm.DB, ctx context.Context, user *User) error {
	return db.WithContext(ctx).Create(user).Error
}

// GetByID 根据用户ID获取用户信息
func GetByID(db *gorm.DB, ctx context.Context, userID int32) (*User, error) {
	var user User
	err := db.WithContext(ctx).Where("id = ?", userID).First(&user).Error
	return &user, err
}

// UpdatePwd 根据用户ID更新密码
func UpdatePwd(db *gorm.DB, ctx context.Context, userID int32, newHashedPwd string) (*User, error) {
	// 使用 Updates 进行部分字段更新
	var user User
	err := db.WithContext(ctx).Model(&User{}).Where("id = ?", userID).Update("password_hashed", newHashedPwd).First(&user).Error

	return &user, err
}

func DeleteByID(db *gorm.DB, ctx context.Context, userID int32) error {
	var user User
	return db.WithContext(ctx).Where("id = ?", userID).Delete(&user).Error
}

func GetPwdByID(db *gorm.DB, ctx context.Context, userID int32) (*User, error) {
	var user User
	err := db.WithContext(ctx).Where("id = ?", userID).First(&user).Error
	return &user, err
}
