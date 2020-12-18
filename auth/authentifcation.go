package auth

import (
	"net/http"

	"github.com/golang/glog"
)

func ValidateToken(request *http.Request) bool {
	authToken := request.Header.Get("X-Auth-Token")
	if authToken == "" || authToken != "a2abe96296ce" {
		glog.Error("invalid auth token")
		return false
	}
	return true
}
