package repository

import (
	"errors"
	"github.com/skantay/web-1-clean/internal/core/dto"
)

var (
	DuplicateUser = errors.New("duplicate user")
)

type UserRepository interface {
	Insert(user dto.UserDTO) error
}
