package util

import (
	"context"
	"net/http"

	"github.com/t3reezhou/figure/figure/middleware"
)

func Write(r *http.Request, result interface{}) {
	ctx := context.WithValue(r.Context(), middleware.WRITEMIDDLEKEY, result)
	*r = *r.WithContext(ctx)
}
