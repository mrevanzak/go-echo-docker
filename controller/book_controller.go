package controller

import (
	"Praktikum/config"
	"Praktikum/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var books []model.Book

func GetBooksAllController(c echo.Context) error {
	if err := config.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes get all books",
		"books":   books,
	})
}

func GetBookController(c echo.Context) error {
	var book model.Book
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = config.DB.Where("id = ?", idInt).First(&book).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes get book",
		"book":    book,
	})
}

func CreateBookController(c echo.Context) error {
	book := model.Book{}
	if err := c.Bind(&book); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes create book",
		"book":    book,
	})
}

func DeleteBookController(c echo.Context) error {
	var book model.Book
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = config.DB.Where("id = ?", idInt).First(&book).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Delete(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes delete book",
	})
}

func UpdateBookController(c echo.Context) error {
	var book model.Book
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = c.Bind(&book)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = config.DB.Model(books).Where("id = ?", idInt).Updates(&model.Book{
		Id:     idInt,
		Title:  book.Title,
		Author: book.Author,
		Year:   book.Year,
	}).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes update book",
		"book":    book,
	})
}
