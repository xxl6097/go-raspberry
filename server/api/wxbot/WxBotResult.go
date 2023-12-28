package wxbot

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

func Ok() map[string]interface{} {
	return map[string]interface{}{"code": 0, "msg": "ok"}
}

func WXRespond(to, types, content string) map[string]interface{} {
	return map[string]interface{}{"to": to, "type": types, "content": content}
}

func DDRespond(msgtype, title, text string) map[string]interface{} {
	return map[string]interface{}{"msgtype": msgtype, msgtype: map[string]interface{}{"title": title, "text": text}}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	if json.NewEncoder(w).Encode(data) != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

type DDNSGOEntity struct {
	Ipv4Addr    string `json:"ipv4Addr"`
	Ipv4Domains string `json:"ipv4Domains"`
	Ipv4Result  string `json:"ipv4Result"`
	Ipv6Addr    string `json:"ipv6Addr"`
	Ipv6Domains string `json:"ipv6Domains"`
	Ipv6Result  string `json:"ipv6Result"`
}
