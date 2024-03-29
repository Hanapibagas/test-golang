package handler

import (
	"Test-Golang/app/middlewares"
	"Test-Golang/features/user"
	"Test-Golang/utils/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService user.UserServiceInterface
}

func NewUser(service user.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (handler *UserHandler) RegisterUser(c echo.Context) error {
	newUser := UserRequestRegister{}

	errBind := c.Bind(&newUser)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid."+errBind.Error(), nil))
	}

	user := RequestUserRegisterToCore(newUser)

	_, token, errRegister := handler.userService.Register(user)
	if errRegister != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error insert data. insert failed"+errRegister.Error(), nil))
	}

	responseData := UserResponRegister{
		Name:  newUser.Name,
		Email: newUser.Email,
		Role:  user.Role,
		Token: token,
	}

	return c.JSON(http.StatusCreated, responses.WebResponse("insert success", responseData))
}

func (handler *UserHandler) LoginUser(c echo.Context) error {
	var reqData = UserRequestLogin{}

	errBind := c.Bind(&reqData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid."+errBind.Error(), nil))
	}

	result, token, err := handler.userService.Login(reqData.Email, reqData.Password)
	if err != nil {
		return c.JSON(http.StatusForbidden, responses.WebResponse("Email atau password tidak boleh kosong "+err.Error(), nil))
	}

	responseData := map[string]any{
		"id":    result.ID,
		"name":  result.Name,
		"email": result.Email,
		"role":  result.Role,
		"toke":  token,
	}

	return c.JSON(http.StatusOK, responses.WebResponse("insert success", responseData))
}

func (handler *UserHandler) UpdatePassword(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	updatePassword := UpdatePasswordRequest{}
	errBind := c.Bind(&updatePassword)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid"+errBind.Error(), nil))
	}

	userCore := RequestToUpdatePassword(updatePassword)

	errUpdate := handler.userService.UpdatePassword(uint(userId), userCore)
	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error editing data. "+errUpdate.Error(), nil))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "Successful Operation",
	})
}

func (handler *UserHandler) GetById(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	result, errGetByID := handler.userService.GetById(uint(userId))
	if errGetByID != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid"+errGetByID.Error(), nil))
	}

	userResult := CoreResponUserById(*result)

	return c.JSON(http.StatusOK, responses.WebResponse("insert success", userResult))
}
