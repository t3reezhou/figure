package middleware

import (
	"context"
	"log"
	"net/http"
	"time"
)

type LogMiddle struct{}

const (
	LOGMIDDLEKEY = "LogMiddleKey"
	INFOCOL      = "\033[32;1m"
	WARNOCOL     = "\033[33;1m"
	OUTTIME      = 2 * time.Second // 2s
)

func (m *LogMiddle) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	start := time.Now()
	log.Printf(INFOCOL+"[INFO]"+ORIGIN+" start %s\n", r.URL)
	ctx := context.WithValue(r.Context(), LOGMIDDLEKEY, start)
	*r = *r.WithContext(ctx)
}

func (m *LogMiddle) After(rw http.ResponseWriter, r *http.Request) {
	start := r.Context().Value(LOGMIDDLEKEY).(time.Time)
	if cost := time.Since(start); cost > OUTTIME {
		log.Printf(WARNOCOL+"[WARN]"+ORIGIN+" %s timeout: %ds", r.URL, cost)
	}
	log.Printf(INFOCOL+"[INFO]"+ORIGIN+" end %s\n", r.URL)
}
