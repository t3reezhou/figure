package middleware

import "net/http"

type AfterMiddle interface {
	After(http.ResponseWriter, *http.Request)
}

func Middleware(handler func(w http.ResponseWriter, r *http.Request), middlewares []http.Handler) http.Handler {
	h := http.HandlerFunc(handler)
	var newHandler http.Handler
	if middlewares != nil && len(middlewares) != 0 {
		m := Middleware(h, middlewares[1:])
		newHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			middlewares[0].ServeHTTP(w, r)
			m.ServeHTTP(w, r)
			if after, ok := middlewares[0].(AfterMiddle); ok {
				after.After(w, r)
			}
		})
	} else {
		newHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		})
	}
	return newHandler
}
