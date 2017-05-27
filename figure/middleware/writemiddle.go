package middleware

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/t3reezhou/figure/figure/status/errors"
)

const (
	WRITEMIDDLEKEY = "WRITEMIDDLEKEY"
	ERRORCOL       = "\033[31;1m"
	ORIGIN         = "\033[0m"
)

type WriteMiddle struct{}

func (m *WriteMiddle) ServeHTTP(rw http.ResponseWriter, r *http.Request) {}

func (m *WriteMiddle) After(rw http.ResponseWriter, r *http.Request) {
	ctxValue := r.Context().Value(WRITEMIDDLEKEY)
	if ctxValue == nil {
		return
	}
	m.write(rw, r, ctxValue)
}

func (m *WriteMiddle) write(rw http.ResponseWriter, r *http.Request, result interface{}) {
	rw.Header().Set("Content-Type", "application/json;charset=utf-8")
	switch R := result.(type) {
	case *errors.Error:
		rw.WriteHeader(R.Code)
		rw.Write([]byte(R.OutPut))
		for _, err := range R.StackTrace() {
			log.Println(ERRORCOL+"[ERROR]"+ORIGIN, r.URL, err)
		}
	default:
		result, _ := json.Marshal(R)
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(result))
	}
}
