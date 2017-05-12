package util

import (
	"context"
	"fmt"
	"reflect"

	"github.com/t3reezhou/figure/figure/middleware"
	"github.com/t3reezhou/figure/lib/inject"
)

const injectKey = "web"

func Inject(ctx context.Context, params interface{}) error {
	ctxValue := ctx.Value(middleware.PARSEMIDDLEKEY)
	if ctxValue == nil {
		return fmt.Errorf("ctx without %s", middleware.PARSEMIDDLEKEY)
	}
	vals, ok := ctxValue.(map[string]interface{})
	if !ok {
		return fmt.Errorf("the kind of ParseMiddlectx is %s,not map[string]interface{}", reflect.TypeOf(ctxValue).Kind())
	}

	if err := inject.Inject(vals, params, injectKey); err != nil {
		return err
	}
	return nil
}
