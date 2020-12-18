package routers

import (
	"net/http"

	"github.com/gkganesh126/multi-email-sender/controllers"
)

func SetRouters(router *http.ServeMux) *http.ServeMux {
	router.HandleFunc("/send-email", controllers.SendEmail)
	return router
}
