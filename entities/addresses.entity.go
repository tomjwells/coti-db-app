package entities

import (
	"time"
)

type Addresses struct {
	ID          int32     `json:"id" gorm:"column:id;type:int(11) NOT NULL AUTO_INCREMENT"`
	AddressHash string    `json:"addressHash" gorm:"column:addressHash;type:varchar(200) COLLATE utf8_unicode_ci NOT NULL UNIQUE;index:addressHash_INDEX"`
	CreateTime  time.Time `json:"createTime" gorm:"column:createTime;type:timestamp NOT NULL;default:CURRENT_TIMESTAMP;"`
	UpdateTime  time.Time `json:"updateTime" gorm:"column:updateTime;type:timestamp NOT NULL;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;"`
}

func NewAddresses(addressHash string) *Addresses {
	instance := new(Addresses)
	instance.AddressHash = addressHash

	return instance
}

func (Addresses) TableName() string {
	return "addresses"
}
