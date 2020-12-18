package routers

import (
	"net/http"
)

func InitRoutes() *http.ServeMux {
	h := http.NewServeMux()
	// Routes for the User entity
	router := SetRouters(h)
	return router
}
