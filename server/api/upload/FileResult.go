package upload

import (
	"encoding/json"
	"net/http"
)

func Ok() map[string]interface{} {
	return map[string]interface{}{"code": 0, "msg": "sucess"}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	if json.NewEncoder(w).Encode(data) != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
