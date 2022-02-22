package entities

import (
	"time"
)

type Currency struct {
	ID         int32     `json:"id" gorm:"column:id;type:int(11) NOT NULL AUTO_INCREMENT"`
	Hash       string    `json:"hash" gorm:"column:hash;type:varchar(200) COLLATE utf8_unicode_ci NOT NULL"`
	CreateTime time.Time `json:"createTime" gorm:"column:createTime;type:timestamp NOT NULL;default:CURRENT_TIMESTAMP;"`
	UpdateTime time.Time `json:"updateTime" gorm:"column:updateTime;type:timestamp NOT NULL;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;"`
}

func (Currency) TableName() string {
	return "currencies"
}

func NewCurrency(currencyHash string) *Currency {
	instance := new(Currency)
	instance.Hash = currencyHash
	return instance
}
