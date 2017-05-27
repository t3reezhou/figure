package middleware

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/t3reezhou/figure/lib/parse"
)

const PARSEMIDDLEKEY = "PARSEMIDDLEKEY"

type ParseMiddle struct{}

func (m *ParseMiddle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	Values := parse.ParseRequest(r, func(values map[string]interface{}) {
		for k, v := range mux.Vars(r) {
			if value, err := strconv.ParseInt(v, 10, 64); err == nil {
				values[k] = value
				continue
			}
			values[k] = v
		}
	})
	ctx := context.WithValue(r.Context(), PARSEMIDDLEKEY, Values)
	*r = *r.WithContext(ctx)
}
