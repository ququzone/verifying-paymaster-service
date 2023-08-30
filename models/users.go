package models

import (
	"time"

	"gorm.io/gorm"

	"github.com/ququzone/verifying-paymaster-service/db"
)

type User struct {
	gorm.Model
	Address string `gorm:"type:varchar(42)"`
}

type ApiKeys struct {
	gorm.Model
	UserID      uint `json:"-"`
	User        User
	Key         string `gorm:"unique;type:varchar(32)"`
	Enable      bool
	Description string
}

func (a *ApiKeys) FindByKey(rep db.Repository, key string) (*ApiKeys, error) {
	var rec ApiKeys
	err := rep.Model(&ApiKeys{}).First(&rec, `"key" = ?`, key).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return &rec, err
}

type Account struct {
	gorm.Model
	Address     string `gorm:"unique;type:varchar(42)"`
	Enable      bool
	RemainGas   string `gorm:"type:varchar(30)"`
	UsedGas     string `gorm:"type:varchar(30)"`
	LastRequest time.Time
}
