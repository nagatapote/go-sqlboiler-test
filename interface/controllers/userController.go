package controllers

import (
	"go-gorm-test/domain/models"
	"go-gorm-test/usecase"
	"go-gorm-test/util"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type UserController interface {
	UserLogin(c echo.Context) (err error)
	UserTotp(c echo.Context) (err error)
	UserGetAll(c echo.Context) (err error)
	UserCreate(c echo.Context) (err error)
	UserUpdate(c echo.Context) (err error)
	UserDelete(c echo.Context) (err error)
}

type userController struct {
	Cuu usecase.UserUseCase
}

func NewUserController(cuu usecase.UserUseCase) UserController {
	return userController{cuu}
}

type (
	userlogin struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,gte=8,password"`
	}
	usertotp struct {
		Email string `json:"email" validate:"required,email"`
		Totp string `json:"totp" validate:"required"`
	}
	userpost struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,gte=8,password"`
	}
	userupdate struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,gte=8,password"`
	}
)

func (uc userController) UserLogin(c echo.Context) (err error) {
	ul := new(userlogin)
	err = util.BindValidate(c, ul)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, util.ErrorNoParameter)
	}
	post, statuscode, err := uc.Cuu.UserLoginUseCase(ul.Email, ul.Password)
	if err != nil {
		message := models.Message{
			Message: err.Error(),
		}
		return echo.NewHTTPError(statuscode, message)
	}
	return c.JSON(http.StatusOK, post)
}

func (uc userController) UserTotp(c echo.Context) (err error) {
	ut := new(usertotp)
	err = util.BindValidate(c, ut)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, util.ErrorNoParameter)
	}
	post, statuscode, err := uc.Cuu.UserTotpUseCase(ut.Email, ut.Totp)
	if err != nil {
		message := models.Message{
			Message: err.Error(),
		}
		return echo.NewHTTPError(statuscode, message)
	}
	return c.JSON(http.StatusOK, post)
}

func (uc userController) UserGetAll(c echo.Context) (err error) {
	post, statuscode, err := uc.Cuu.UserGetAllUseCase()
	if err != nil {
		message := models.Message{
			Message: err.Error(),
		}
		return echo.NewHTTPError(statuscode, message)
	}
	return c.JSON(http.StatusOK, post)
}

func (uc userController) UserCreate(c echo.Context) (err error) {
	// up := &userpost{
	// 	Email:          c.FormValue("email"),
	// 	Password: c.FormValue("password"),
	// }

	//NOTE: body content type application/jsonであれば、post dataをこれで受け取れる。
	up := new(userpost)
	err = util.BindValidate(c, up)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, util.ErrorNoParameter)
	}
	post, statuscode, err := uc.Cuu.UserCreateUseCase(up.Email, up.Password)
	if err != nil {
		message := models.Message{
			Message: err.Error(),
		}
		return echo.NewHTTPError(statuscode, message)
	}
	return c.JSON(http.StatusOK, post)
}

func (uc userController) UserUpdate(c echo.Context) (err error) {
	id := c.Param("id")
	uu := new(userupdate)
	err = util.BindValidate(c, uu)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, util.ErrorNoParameter)
	}
	f, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, util.ErrorServerError)
	}
	post, statuscode, err := uc.Cuu.UserUpdateUseCase(f, uu.Email, uu.Password)
	if err != nil {
		message := models.Message{
			Message: err.Error(),
		}
		return echo.NewHTTPError(statuscode, message)
	}
	return c.JSON(http.StatusOK, post)
}

func (uc userController) UserDelete(c echo.Context) (err error) {
	id := c.Param("id")
	f, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, util.ErrorServerError)
	}
	post, statuscode, err := uc.Cuu.UserDeleteUseCase(f)
	if err != nil {
		message := models.Message{
			Message: err.Error(),
		}
		return echo.NewHTTPError(statuscode, message)
	}
	return c.JSON(http.StatusOK, post)
}
