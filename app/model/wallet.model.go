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

type TransactionHistory struct {
	ID string `json:"id" gorm:"primaryKey; type:varchar; not null; unique"`

	Amount              float64 `json:"amount" gorm:"type:numeric(10,2);default:0;index"`
	Narration           string  `json:"narration" gorm:"type:varchar"`
	TransactionType     string  `json:"transaction_type" gorm:"type:varchar;check:role IN ('credit', 'debit');index"`
	TransactionStatus   string  `json:"transaction_status" gorm:"type:varchar;check:role IN ('success', 'pending', 'failed', 'reversed');index"`
	TransactionCategory string  `json:"transaction_category" gorm:"type:varchar;check:role IN ('top up', 'withdraw', 'transfer', 'airtime', 'data', 'reversal');index"`

	RecipientName      string `json:"recipient_name" gorm:"type:varchar"`
	RecipientBank      string `json:"recipient_bank" gorm:"type:varchar"`
	RecipientAccountNo string `json:"recipient_account_no" gorm:"type:varchar"`

	CreatedAt time.Time `json:"created_at"`
}
