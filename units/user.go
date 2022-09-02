package units

import (
	"context"
)

type User struct {
	ID           uint
	TelegramID   uint   `db:"telegram_id"`
	IsBot        bool   `db:"is_bot"`
	FirstName    string `db:"first_name"`
	LastName     string `db:"last_name"`
	UserName     string `db:"user_name"`
	LanguageCode string `db:"language_code"`
	Notify       bool   `db:"notify"`
}

type UserPatch struct {
	FirstName *string
	LastName  *string
	UserName  *string
	Notify    *bool
}

type UserFilter struct {
	TelegramID *uint
	Notify     *bool

	Limit  int
	Offset int
}

type UserService interface {
	CreateUser(context.Context, *User) error

	UserByTelegramID(context.Context, uint) (*User, error)

	Users(context.Context, UserFilter) ([]*User, error)

	UpdateUser(context.Context, *User, UserPatch) error
}
