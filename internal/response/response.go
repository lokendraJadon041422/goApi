package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

const (
	StatusOK    = "OK"
	StatusError = "ERROR"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) Response {

	return Response{
		Status: http.StatusInternalServerError,
		Error:  err.Error(),
	}
}

func ValidationError(err validator.ValidationErrors) Response {
	var errMsg []string
	for _, err := range err {
		switch err.ActualTag() {
		case "required":
			errMsg = append(errMsg, fmt.Sprintf("field %s is required", err.Field()))
		default:
			errMsg = append(errMsg, fmt.Sprintf("field %s is invalid", err.Field()))
		}

	}
	return Response{
		Status: http.StatusBadRequest,
		Error:  strings.Join(errMsg, ", "),
	}
}
