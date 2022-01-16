package usecase

import (
	"go-gorm-test/domain/repository"
	"go-gorm-test/util"
	"net/http"
)

type UserUseCase interface {
	UserGetUseCase() (resp interface{}, statuscode int, err error)
	// UserPostUseCase(ID int, Email string) (resp interface{}, statuscode int, err error)
	// UserPutUseCase(ID int) (resp interface{}, statuscode int, err error)
	// UserDeleteUseCase(ID int) (resp interface{}, statuscode int, err error)
}

type userUseCaceImpl struct {
	Ur repository.UserRepository
}

func NewUserUseCase(ur repository.UserRepository) UserUseCase {
	return userUseCaceImpl{ur}
}

func (uu userUseCaceImpl) UserGetUseCase() (resp interface{}, statuscode int, err error) {
	resp, err = uu.Ur.UserGet()
	if err != nil {
		return nil, http.StatusInternalServerError, util.ErrorServerError
	}
	return resp, http.StatusOK, nil
}

// func (uu userUseCaceImpl) UserDeleteUseCase(ID int) (resp interface{}, statuscode int, err error) {
// 	resp, err = uu.Ur.UserDelete(ID)
// 	if err != nil {
// 		return nil, http.StatusInternalServerError, util.ErrorServerError
// 	}
// 	return resp, http.StatusOK, nil
// }
