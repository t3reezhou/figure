package middleware

import (
	"context"
	"net/http"

	"github.com/t3reezhou/figure/figure/status/errors"
)

type CSRFMiddle struct{}

const CSRFCOOKIEKEY = "CSRFCOOKIEKEY"

var SAFTMETHS = map[string]struct{}{
	"GET":  struct{}{},
	"HEAD": struct{}{},
}

func (m *CSRFMiddle) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if _, ok := SAFTMETHS[r.Method]; ok {
		return
	}
	err := m.CheckCSRF(r)
	if err != nil {
		context := context.WithValue(r.Context(), DEATHLINEKEY, err)
		*r = *r.WithContext(context)
		rw.Write([]byte(err.Error()))
	}
}

func (m *CSRFMiddle) CheckCSRF(r *http.Request) error {
	csrfCookie, err := r.Cookie(CSRFCOOKIEKEY)
	if err != nil {
		return errors.NewError(CSRFCOOKIEKEY, 403)
	}

	r.ParseForm()
	csrfRequest := r.Form
	if csrfCookie != nil && csrfRequest != nil {
		if csrfCookie.Value == csrfRequest.Get("csrf") {
			return nil
		}
	}
	return errors.NewError("csrfRequest", 403)
}
