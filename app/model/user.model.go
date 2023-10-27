package model

import (
	"database/sql"
	"time"

	"github.com/Ahmad940/assetly_server/pkg/util/nullable"
	"gopkg.in/guregu/null.v4"
)

type UserRole string

const (
	AdminRoleAdmin      = string("admin")
	ConsultantRoleAdmin = string("consultant")
	UserRoleAdmin       = string("user")
)

type User struct {
	ID string `json:"id" gorm:"primaryKey; type:varchar; not null; unique"`

	FirstName   string      `json:"first_name" gorm:"type:varchar"`
	LastName    string      `json:"last_name" gorm:"type:varchar"`
	CountryCode string      `json:"country_code" gorm:"type:varchar; not null"`
	PhoneNumber string      `json:"phone_number" gorm:"type:varchar; not null"`
	Email       string      `json:"email" gorm:"type:varchar;not null;unique"`
	UserImage   string      `json:"user_image" gorm:"type:varchar"`
	Session     null.String `json:"-" gorm:"type:varchar"`

	Role string `json:"role" gorm:"type:varchar; check:role IN ('admin', 'mod', 'user'); not null; default:user"`

	UserDetail *UserDetail `json:"user_detail" gorm:"foreignKey:UserID;references:ID"`
	Wallet     *Wallet     `json:"wallet" gorm:"foreignKey:UserID;references:ID"`

	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"-" gorm:"index"`
}

type UserDetail struct {
	ID string `json:"id" gorm:"primaryKey; type:varchar; not null; unique"`

	UserID  string    `json:"user_id" gorm:"type:varchar;index;unique"`
	Country string    `json:"country" gorm:"type:varchar"`
	State   string    `json:"state" gorm:"type:varchar"`
	LGA     string    `json:"lga" gorm:"type:varchar"`
	Area    string    `json:"area" gorm:"type:varchar"`
	DOB     time.Time `json:"dob" gorm:"type:timestamp"`
	Gender  string    `json:"gender" gorm:"type:varchar"`
}

type CreateUser struct {
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	CountryCode string `json:"country_code" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Email       string `json:"email" validate:"required"`
	UserImage   string `json:"user_image"`

	Country string    `json:"country" validate:"required"`
	State   string    `json:"state" validate:"required"`
	LGA     string    `json:"lga" validate:"required"`
	Area    string    `json:"area" validate:"required"`
	DOB     time.Time `json:"dob" validate:"required"`
	Gender  string    `json:"gender" validate:"required"`
}

type UpdateUser struct {
	ID string `json:"id" validate:"required"`

	FullName    string                  `json:"full_name"`
	CountryCode string                  `json:"country_code"`
	PhoneNumber string                  `json:"phone_number"`
	DOB         nullable.CustomNullTime `json:"dob"`
	Gender      string                  `json:"gender"`
}

type UpdateUserAdmin struct {
	ID          string `json:"id" validate:"required"`
	FullName    string `json:"full_name"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role" gorm:"type:varchar; check:role IN ('admin', 'user'); not null; default:user"`
}
