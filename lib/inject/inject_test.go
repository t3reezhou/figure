package inject

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var TestStructTag = "web"

func TestInject(t *testing.T) {
	Convey("test a struct with int64 slice", t, func() {
		params := struct {
			SliceInt64 []int64 `web:"slice_int64"`
		}{}
		Convey("Inject a struct with int64 slice", func() {
			values := map[string]interface{}{
				"slice_int64": []int64{3, 27, 28},
			}
			err := Inject(values, &params, TestStructTag)
			So(err, ShouldBeNil)
			So(params.SliceInt64, ShouldResemble, values["slice_int64"])
		})
		Convey("struct is not a pointer", func() {
			err := Inject(nil, params, TestStructTag)
			So(err, ShouldResemble, fmt.Errorf(("not a pointer of struct")))
		})
		Convey("values's type is wrong", func() {
			values := map[string]interface{}{
				"slice_int64": []string{"3", "27", "28"},
			}
			So(params.SliceInt64, ShouldHaveLength, 0)
			err := Inject(values, &params, TestStructTag)
			So(err, ShouldNotBeNil)
		})
		Convey("values without the key", func() {
			err := Inject(map[string]interface{}{}, &params, TestStructTag)
			So(err, ShouldBeNil)
			So(params.SliceInt64, ShouldHaveLength, 0)
		})
	})
	Convey("struct with string data and it is required", t, func() {
		params := struct {
			DataString string `web:"data_string,required"`
		}{}
		Convey("values is empty", func() {
			err := Inject(map[string]interface{}{}, &params, TestStructTag)
			So(err, ShouldNotBeNil)
			So(params.DataString, ShouldBeZeroValue)
		})
		Convey("value is not empty", func() {
			values := map[string]interface{}{
				"data_string": "happy hacking",
			}
			err := Inject(values, &params, TestStructTag)
			So(err, ShouldBeNil)
			So(params.DataString, ShouldEqual, "happy hacking")
		})
		Convey("values's type is wrong", func() {
			values := map[string]interface{}{
				"data_string": 123,
			}
			err := Inject(values, &params, TestStructTag)
			So(err, ShouldNotBeNil)
			So(params.DataString, ShouldBeZeroValue)
		})
	})
	Convey("struct with int64 data and it is -", t, func() {
		params := struct {
			DataEmpty int64 `web:"-"`
		}{}
		Convey("values is not empty", func() {
			values := map[string]interface{}{
				"-": 123,
			}
			err := Inject(values, &params, TestStructTag)
			So(err, ShouldBeNil)
			So(params.DataEmpty, ShouldBeZeroValue)
		})
	})
	Convey("struct with bool data and it without structTag", t, func() {
		params := struct {
			DataWithoutStructTag bool
			int64
		}{int64: 2333}
		Convey("values is not empty", func() {
			values := map[string]interface{}{
				"datawithoutstructtag": true,
			}
			err := Inject(values, &params, TestStructTag)
			So(err, ShouldBeNil)
			So(params.DataWithoutStructTag, ShouldEqual, true)
			So(params.int64, ShouldEqual, 2333)
		})
	})
	Convey("struct with int64 data", t, func() {
		params := struct {
			DataInt64 int64 `web:"data_int64"`
		}{}
		values := map[string]interface{}{
			"data_int64": 2333,
		}
		err := Inject(values, &params, TestStructTag)
		So(err, ShouldBeNil)
		So(params.DataInt64, ShouldEqual, 2333)
	})

}
