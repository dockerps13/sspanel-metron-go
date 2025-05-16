package models

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Username     string    `gorm:"uniqueIndex;size:50" json:"username"`
	Email        string    `gorm:"uniqueIndex;size:100" json:"email"`
	PasswordHash string    `gorm:"size:255" json:"-"`
	TransferEnable uint64   `json:"transfer_enable"` // 可用流量，单位字节
	Upload       uint64    `json:"upload"`         // 已用上传流量
	Download     uint64    `json:"download"`       // 已用下载流量
	ExpiredAt    time.Time `json:"expired_at"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	result := db.Where("username = ?", username).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("用户不存在")
	}
	return &user, result.Error
}

func GetUserByID(id uint) (*User, error) {
	var user User
	result := db.First(&user, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("用户不存在")
	}
	return &user, result.Error
}

func CreateUser(user *User) error {
	return db.Create(user).Error
}

func UpdateUserTraffic(userID uint, upload, download uint64) error {
	return db.Model(&User{}).Where("id = ?", userID).
		Updates(map[string]interface{}{
			"upload":   gorm.Expr("upload + ?", upload),
			"download": gorm.Expr("download + ?", download),
		}).Error
}

func ResetUserTraffic(userID uint) error {
	return db.Model(&User{}).Where("id = ?", userID).
		Updates(map[string]interface{}{
			"upload":   0,
			"download": 0,
		}).Error
}
