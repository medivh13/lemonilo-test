package controllers

import (
	"net/http"

	"github.com/medivh13/lemonilo-test/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
