package responser

import (
	"encoding/json"
	"net/http"

	"github.com/shahsidar-k-s/students-api/internal/types"
)

func ResponseWriter(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)

}
func GeneralError(err error, StatusCode int) types.CustomErrors {

	return types.CustomErrors{
		StatusCode:   http.StatusBadRequest,
		ErrorMessage: err.Error() + " No Body found",
	}
}
