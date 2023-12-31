// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package sqlc

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type Rank string

const (
	RankBronze   Rank = "bronze"
	RankSilver   Rank = "silver"
	RankGold     Rank = "gold"
	RankPlatinum Rank = "platinum"
	RankGarnet   Rank = "garnet"
	RankRuby     Rank = "ruby"
	RankDiamond  Rank = "diamond"
)

func (e *Rank) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Rank(s)
	case string:
		*e = Rank(s)
	default:
		return fmt.Errorf("unsupported scan type for Rank: %T", src)
	}
	return nil
}

type NullRank struct {
	Rank  Rank `json:"rank"`
	Valid bool `json:"valid"` // Valid is true if Rank is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullRank) Scan(value interface{}) error {
	if value == nil {
		ns.Rank, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Rank.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullRank) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Rank), nil
}

type Role string

const (
	RoleAdmin      Role = "admin"
	RoleManager    Role = "manager"
	RoleBankTeller Role = "bankTeller"
	RoleCustomer   Role = "customer"
)

func (e *Role) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Role(s)
	case string:
		*e = Role(s)
	default:
		return fmt.Errorf("unsupported scan type for Role: %T", src)
	}
	return nil
}

type NullRole struct {
	Role  Role `json:"role"`
	Valid bool `json:"valid"` // Valid is true if Role is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullRole) Scan(value interface{}) error {
	if value == nil {
		ns.Role, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Role.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullRole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Role), nil
}

type Status string

const (
	StatusDeleted  Status = "deleted"
	StatusInactive Status = "inactive"
	StatusHolded   Status = "holded"
	StatusActive   Status = "active"
)

func (e *Status) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Status(s)
	case string:
		*e = Status(s)
	default:
		return fmt.Errorf("unsupported scan type for Status: %T", src)
	}
	return nil
}

type NullStatus struct {
	Status Status `json:"status"`
	Valid  bool   `json:"valid"` // Valid is true if Status is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullStatus) Scan(value interface{}) error {
	if value == nil {
		ns.Status, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Status.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Status), nil
}

type Account struct {
	ID        int64              `db:"id" json:"id"`
	Type      string             `db:"type" json:"type"`
	Balance   int64              `db:"balance" json:"balance"`
	AccStatus Status             `db:"acc_status" json:"acc_status"`
	CreatedAt pgtype.Timestamptz `db:"created_at" json:"created_at"`
}

type AccountHolder struct {
	AccID     int64              `db:"acc_id" json:"acc_id"`
	UserID    int64              `db:"user_id" json:"user_id"`
	CreatedAt pgtype.Timestamptz `db:"created_at" json:"created_at"`
}

type Entry struct {
	ID        int64 `db:"id" json:"id"`
	AccountID int64 `db:"account_id" json:"account_id"`
	// can be negative or positive
	Amount    int64              `db:"amount" json:"amount"`
	CreatedAt pgtype.Timestamptz `db:"created_at" json:"created_at"`
}

type Session struct {
	ID           pgtype.UUID        `db:"id" json:"id"`
	Username     string             `db:"username" json:"username"`
	RefreshToken string             `db:"refresh_token" json:"refresh_token"`
	UserAgent    string             `db:"user_agent" json:"user_agent"`
	ClientIp     string             `db:"client_ip" json:"client_ip"`
	IsBlocked    bool               `db:"is_blocked" json:"is_blocked"`
	ExpiresAt    pgtype.Timestamptz `db:"expires_at" json:"expires_at"`
	CreatedAt    pgtype.Timestamptz `db:"created_at" json:"created_at"`
}

type Transfer struct {
	ID             int64 `db:"id" json:"id"`
	FromAccountID  int64 `db:"from_account_id" json:"from_account_id"`
	ToAccountID    int64 `db:"to_account_id" json:"to_account_id"`
	TransferedByID int64 `db:"transfered_by_id" json:"transfered_by_id"`
	// must be positive
	Amount    int64              `db:"amount" json:"amount"`
	CreatedAt pgtype.Timestamptz `db:"created_at" json:"created_at"`
}

type User struct {
	ID                int64              `db:"id" json:"id"`
	FirstName         string             `db:"first_name" json:"first_name"`
	LastName          string             `db:"last_name" json:"last_name"`
	Username          string             `db:"username" json:"username"`
	Nic               string             `db:"nic" json:"nic"`
	HashedPassword    string             `db:"hashed_password" json:"hashed_password"`
	PasswordChangedAt pgtype.Timestamptz `db:"password_changed_at" json:"password_changed_at"`
	Email             string             `db:"email" json:"email"`
	IsEmailVerified   bool               `db:"is_email_verified" json:"is_email_verified"`
	EmailChangedAt    pgtype.Timestamptz `db:"email_changed_at" json:"email_changed_at"`
	Phone             string             `db:"phone" json:"phone"`
	IsPhoneVerified   bool               `db:"is_phone_verified" json:"is_phone_verified"`
	PhoneChangedAt    pgtype.Timestamptz `db:"phone_changed_at" json:"phone_changed_at"`
	AccStatus         Status             `db:"acc_status" json:"acc_status"`
	CustomerRank      Rank               `db:"customer_rank" json:"customer_rank"`
	IsAnEmployee      bool               `db:"is_an_employee" json:"is_an_employee"`
	IsACustomer       bool               `db:"is_a_customer" json:"is_a_customer"`
	Role              pgtype.Text        `db:"role" json:"role"`
	Department        pgtype.Text        `db:"department" json:"department"`
	CreatedAt         pgtype.Timestamptz `db:"created_at" json:"created_at"`
}

type VerifyEmail struct {
	ID         int64              `db:"id" json:"id"`
	Username   string             `db:"username" json:"username"`
	Email      string             `db:"email" json:"email"`
	SecretCode string             `db:"secret_code" json:"secret_code"`
	IsUsed     bool               `db:"is_used" json:"is_used"`
	CreatedAt  pgtype.Timestamptz `db:"created_at" json:"created_at"`
	ExpiredAt  pgtype.Timestamptz `db:"expired_at" json:"expired_at"`
}

type VerifyPhone struct {
	ID         int64              `db:"id" json:"id"`
	Username   string             `db:"username" json:"username"`
	Phone      string             `db:"phone" json:"phone"`
	SecretCode string             `db:"secret_code" json:"secret_code"`
	IsUsed     bool               `db:"is_used" json:"is_used"`
	CreatedAt  pgtype.Timestamptz `db:"created_at" json:"created_at"`
	ExpiredAt  pgtype.Timestamptz `db:"expired_at" json:"expired_at"`
}
