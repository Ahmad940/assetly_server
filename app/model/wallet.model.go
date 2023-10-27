package model

import (
	"database/sql"
	"time"
)

type Wallet struct {
	ID string `json:"id" gorm:"primaryKey; type:varchar; not null; unique"`

	WalletNo string  `json:"wallet_no" gorm:"type:varchar;check:(LENGTH(wallet_no) = 10);unique;index"`
	Balance  float64 `json:"balance" gorm:"type:numeric(10,2);default:0"`
	UserID   string  `json:"user_id" gorm:"type:varchar;index"`

	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"-" gorm:"index"`
}
