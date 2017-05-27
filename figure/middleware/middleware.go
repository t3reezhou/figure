package middleware

import (
	"log"
	"net/http"
)

const DEATHLINEKEY = "DEATHLINEKEY"

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
			deathLine := r.Context().Value(DEATHLINEKEY)
			if deathLine == nil {
				m.ServeHTTP(w, r)
			} else {
				log.Println(ERRORCOL+"[ERROR]"+ORIGIN, r.URL, deathLine)
			}
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
