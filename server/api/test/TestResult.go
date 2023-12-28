package test

import (
	"encoding/json"
	"net/http"
)

func Allow(super bool) map[string]interface{} {
	return map[string]interface{}{"is_superuser": super, "result": "allow"}
}

func Deny(super bool) map[string]interface{} {
	return map[string]interface{}{"is_superuser": super, "result": "deny"}
}

func Ignore(super bool) map[string]interface{} {
	return map[string]interface{}{"is_superuser": super, "result": "ignore"}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	if json.NewEncoder(w).Encode(data) != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
