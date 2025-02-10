package web

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/kurdilesmana/backend-product-service/internal/core/models/helperModel"
	"github.com/kurdilesmana/backend-product-service/pkg/convert"
	"github.com/labstack/echo/v4"
)

// ResponseFormatter returning formatted JSON response
func ResponseFormatter(ctx echo.Context, code int, message string, body interface{}, err error) error {
	var response = helperModel.BaseResponseModel{
		Message: message,
		Data:    nil,
		Error:   "",
	}

	if err != nil {
		response.Error = err.Error()
	} else {
		response.Data = body
	}

	return ctx.JSON(code, response)
}

// ResponseErrValidation returning formatted JSON response
func ResponseErrValidation(ctx echo.Context, message string, errMap map[string]interface{}) error {

	var b strings.Builder
	for k, v := range errMap {
		b.WriteString(fmt.Sprintf("%s : %v, ", k, v))
	}
	errorString := strings.TrimRight(b.String(), ", ")

	var response = helperModel.ErrRespValidationModel{
		Message:         message,
		Data:            nil,
		ErrorValidation: errMap,
		Error:           errorString,
	}

	return ctx.JSON(http.StatusBadRequest, response)
}

// ResponseErrWithFormValidation returning formatted JSON response
func ResponseErrWithFormatValidation(ctx echo.Context, message string, validation map[string]interface{}) error {
	return ctx.JSON(
		http.StatusBadRequest,
		struct {
			Message        string                 `json:"message"`
			FormValidation map[string]interface{} `json:"form_validation"`
		}{Message: message, FormValidation: validation},
	)
}

func ResponseErrValidationWithCode(ctx echo.Context, message string, err error, code int) error {
	var msg string
	var errMap string
	var fieldRequired []string

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				field := fmt.Sprintf("%v,", convert.ToSnakeCase(err.Field()))
				fieldRequired = append(fieldRequired, field)

			}

			if len(fieldRequired) > 0 {
				errMap = fmt.Sprintf("REQUIRED: field %v harus diisi. ", fieldRequired)
			}
		}
	}

	response := map[string]interface{}{
		"message":          msg,
		"data":             nil,
		"error_validation": errMap,
	}

	return ctx.JSON(code, response)
}

// ResponseFormatter returning formatted JSON response with meta
func ResponseFormatterWithMeta(ctx echo.Context, code int, message string, body interface{}, meta interface{}, err error) error {
	var response map[string]interface{}

	if err != nil {
		response = map[string]interface{}{
			"message": message,
			"data":    body,
			"error":   err.Error(),
			"meta":    meta,
		}
	} else {
		response = map[string]interface{}{
			"message": message,
			"data":    body,
			"error":   nil,
			"meta":    meta,
		}
	}

	return ctx.JSON(code, response)
}

func ResponseErrValidationWithDefaultMessage(ctx echo.Context, message string, errMap map[string]interface{}, code int) error {
	var msg string

	if len(errMap) == 0 {
		msg = message
	} else {
		for _, value := range errMap {
			msg = value.(string)
			break
		}
	}

	response := map[string]interface{}{
		"message":          msg,
		"error_validation": errMap,
	}

	return ctx.JSON(code, response)
}
