package middleware

import (
	"net/http"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParseMiddle(t *testing.T) {
	Convey("test ParseMiddle", t, func() {
		parseMiddle := ParseMiddle{}
		r, _ := http.NewRequest("GET", "http://localhost/test?test=test", strings.NewReader(`{"id":2}`))
		Convey("test ServeHTTP", func() {
			parseMiddle.ServeHTTP(nil, r)
			So(r.Context().Value(PARSEMIDDLEKEY), ShouldNotBeNil)
			ctx := r.Context().Value(PARSEMIDDLEKEY)
			values, ok := ctx.(map[string]interface{})
			So(ok, ShouldBeTrue)
			So(values, ShouldContainKey, "test")
			So(values["test"], ShouldEqual, "test")
			So(values, ShouldContainKey, "id")
			So(values["id"], ShouldEqual, 2)
		})
	})
}
