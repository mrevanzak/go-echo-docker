package controller

import (
	"Praktikum/config"
	"Praktikum/middleware"
	"Praktikum/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var users []model.User

func GetUsersAllController(c echo.Context) error {
	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes get all users",
		"users":   users,
	})
}

func GetUserController(c echo.Context) error {
	var user model.User
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = config.DB.Where("id = ?", idInt).First(&user).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes get user",
		"user":    user,
	})
}

func CreateUserController(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes create user",
		"user":    user,
	})
}

func DeleteUserController(c echo.Context) error {
	var user model.User
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = config.DB.Where("id = ?", idInt).First(&user).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes delete user",
	})
}

func UpdateUserController(c echo.Context) error {
	var user model.User
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = c.Bind(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = config.DB.Model(users).Where("id = ?", idInt).Updates(&model.User{
		Id:       idInt,
		Email:    user.Email,
		Name:     user.Name,
		Password: user.Password,
	}).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes update user",
		"user":    user,
	})
}

func LoginUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)
	err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "failed login",
			"error":   err.Error(),
		})
	}

	var token string
	token, err = middleware.CreateToken(user.Id, user.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "failed create token",
			"error":   err.Error(),
		})
	}

	userResponse := model.UserResponse{user.Id, user.Email, user.Name, token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"data":    userResponse,
	})
}
