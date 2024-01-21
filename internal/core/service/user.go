package service

import (
	"errors"
	utils "github.com/skantay/web-1-clean/internal/core/common/utls"
	"github.com/skantay/web-1-clean/internal/core/dto"
	"github.com/skantay/web-1-clean/internal/core/entity/error_code"
	"github.com/skantay/web-1-clean/internal/core/model/request"
	"github.com/skantay/web-1-clean/internal/core/model/response"
	"github.com/skantay/web-1-clean/internal/core/port/repository"
)

const (
	invalidUserNameErrMsg = "invalid username"
	invalidPasswordErrMsg = "invalid password"
)

type UserService interface {
	SignUp(request *request.SignUpRequest) *response.Response
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u userService) SignUp(request *request.SignUpRequest) *response.Response {
	if len(request.Username) == 0 {
		return u.createFailedResponse(error_code.InvalidRequest, invalidUserNameErrMsg)
	}

	if len(request.Password) == 0 {
		return u.createFailedResponse(error_code.InvalidRequest, invalidPasswordErrMsg)
	}

	currentTime := utils.GetUTCCurrentMillis()
	userDTO := dto.UserDTO{
		UserName:    request.Username,
		Password:    request.Password,
		DisplayName: u.getRandomDisplayName(request.Username),
		CreatedAt:   currentTime,
		UpdatedAt:   currentTime,
	}

	if err := u.userRepo.Insert(userDTO); err != nil {
		if errors.Is(err, repository.DuplicateUser) {
			return u.createFailedResponse(error_code.DuplicateUser, err.Error())
		}
		return u.createFailedResponse(error_code.InternalError, error_code.InternalErrMsg)
	}

	signUpData := response.SignUpDataResponse{
		DisplayName: userDTO.DisplayName,
	}

	return u.createSuccessResponse(signUpData)
}

func (u userService) getRandomDisplayName(username string) string {
	return username + "Random"
}

func (u userService) createFailedResponse(code error_code.ErrorCode, message string) *response.Response {
	return &response.Response{
		Status:       false,
		ErrorCode:    code,
		ErrorMessage: message,
	}
}

func (u userService) createSuccessResponse(data response.SignUpDataResponse) *response.Response {
	return &response.Response{
		Data:         data,
		Status:       true,
		ErrorCode:    error_code.Success,
		ErrorMessage: error_code.SuccessErrMsg,
	}
}
