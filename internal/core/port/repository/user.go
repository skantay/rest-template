package repository

import (
	"database/sql"
	"errors"
	"github.com/skantay/web-1-clean/internal/core/dto"
	"io"
)

var (
	DuplicateUser = errors.New("duplicate user")
)

type UserRepository interface {
	Insert(user dto.UserDTO) error
}

type Database interface {
	io.Closer
	GetDB() *sql.DB
}
