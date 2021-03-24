package helpers

import (
	"encoding/json"
	"net/http"
)

func BodyJsonDecode(w http.ResponseWriter, r *http.Request, data interface{}) {
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {

	}
}
