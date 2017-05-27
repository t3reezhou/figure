package api

import "net/http"

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}
func (h *UserHandler) Login(rw http.ResponseWriter, r *http.Request) {
	http.SetCookie(rw, &http.Cookie{
		Name:  "csrf",
		Value: "1234",
	})
}
