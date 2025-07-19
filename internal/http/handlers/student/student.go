package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/lokendraJadon041422/studentsApi/internal/response"
	"github.com/lokendraJadon041422/studentsApi/internal/types"
)

func CreateStudent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("error should not be empty")))
			return
		}
		//  err := response.WriteJson(w, http.StatusBadRequest, student);
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		if err := validator.New().Struct(student); err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(err.(validator.ValidationErrors)))
			return
		}
		slog.Info("Student created successfully", "student", student)
		response.WriteJson(w, http.StatusCreated, student)
	}
}
