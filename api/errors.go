package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Error struct{
	Code int `json:"code"`
	Err string `json:"err"`
}

func NewError(code int , err string) Error {
	return Error{
		Code: code,
		Err: err,
	}
}

// Error Implements the error interface
func(e Error) Error()string  {
	return e.Err
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	apiError, ok := err.(Error)
	if  ok{
		return c.Status(apiError.Code).JSON(apiError)
	}
	apiErr := NewError(http.StatusInternalServerError,err.Error())
	return c.Status(apiErr.Code).JSON(apiErr)
}

func ErrInvalidID() Error {
	return Error{
		Code: http.StatusBadRequest,
		Err: "inavlid ID",
	}
}

func ErrUnAuthorised() Error {
	return Error{
		Code: http.StatusUnauthorized,
		Err: "unauthorised",
	}
}

func ErrBadRequest() Error {
	return Error{
		Code: http.StatusBadRequest,
		Err: "invalid JSON request",
	}
}

func ErrResourceNotFound(res string) Error {
	return Error{
		Code: http.StatusNotFound,
		Err: res + " resource not found",
	}
}