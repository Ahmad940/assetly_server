package model

import (
	"time"
)

const (
	TransactionTypeCredit = "credit"
	TransactionTypeDebit  = "debit"

	TransactionStatusSuccess = "success"
	TransactionStatusFailed  = "failed"
	TransactionStatusPending = "pending"

	TransactionCategoryTopUp    = "top up"
	TransactionCategoryWithdraw = "withdraw"
	TransactionCategoryTransfer = "transfer"
	TransactionCategoryAirtime  = "airtime"
	TransactionCategoryData     = "data"
	TransactionCategoryReversal = "reversal"
)

type TransactionHistory struct {
	ID string `json:"id" gorm:"primaryKey; type:varchar; not null; unique"`

	WalletNo            string  `json:"wallet_no" gorm:"type:varchar;"`
	Amount              float64 `json:"amount" gorm:"type:numeric(10,2);default:0;index"`
	Narration           string  `json:"narration" gorm:"type:varchar"`
	TransactionType     string  `json:"transaction_type" gorm:"type:varchar;check:transaction_type IN ('credit', 'debit');index"`
	TransactionStatus   string  `json:"transaction_status" gorm:"type:varchar;check:transaction_status IN ('success', 'pending', 'failed', 'reversed');index"`
	TransactionCategory string  `json:"transaction_category" gorm:"type:varchar;check:transaction_category IN ('top up', 'withdraw', 'transfer', 'airtime', 'data', 'reversal');index"`

	RecipientName      string `json:"recipient_name" gorm:"type:varchar"`
	RecipientBank      string `json:"recipient_bank" gorm:"type:varchar"`
	RecipientAccountNo string `json:"recipient_account_no" gorm:"type:varchar"`

	CreatedAt time.Time `json:"created_at"`
}
