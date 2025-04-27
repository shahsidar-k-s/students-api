package students

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/shahsidar-k-s/students-api/internal/responser"
	"github.com/shahsidar-k-s/students-api/internal/types"
)

func GetAllStudents() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responser.ResponseWriter(w, http.StatusBadGateway, "Hello Student's")
	}
}

func AddNewStudent() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var userData types.User
		error := json.NewDecoder(r.Body).Decode(&userData)
		if errors.Is(error, io.EOF) {
			responser.ResponseWriter(w, http.StatusBadRequest, responser.GeneralError(error, http.StatusBadRequest))
			return
		}
		// this can also be done
		// w.Header().Set("Content-type", "application/json")
		// w.WriteHeader(http.StatusOK)
		// json.NewEncoder(w).Encode(userData)
		responser.ResponseWriter(w, http.StatusOK, userData)
	}
}
func UpdateStudent() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

	}
}
