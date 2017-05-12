package parse

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParseJson(t *testing.T) {
	Convey("testParseJson", t, func() {
		data := `{
    "id": 22868,
    "fileid": 22868,
    "groupid": 6079,
    "parentid": 0,
    "fname": "test.doc",
    "ftype": "sharefile",
    "fver": 2,
    "fsize": 3038,
    "fsha": "8e3378ea78ca402c859fc29cd9ed2d0e377a435e",
    "storeid": "afq0ndfsLEyA(l9zAAvejjN46njKQCyFn8Kc2e0tDjd6Q14",
    "store": 1,
    "secure_guid": "test",
    "deleted": false,
    "ctime": 1481601634,
    "mtime": 1481601789,
    "a": {
        "b": {
            "c": 1
        }
    },
    "test": [1,2,3,4,5],
    "impush": true
}`
		result := ParseJson([]byte(data))
		So(result, ShouldHaveLength, 18)
	})
	Convey("test int64", t, func() {
		data := `{
		"data_int64" : 3
		}`
		result := ParseJson([]byte(data))
		So(result, ShouldContainKey, "data_int64")
		So(result["data_int64"], ShouldEqual, 3)
		So(result["data_int64"], ShouldHaveSameTypeAs, int64(3))
	})
	Convey("test string", t, func() {
		data := `{
		"data_string": "test"
		}`
		result := ParseJson([]byte(data))
		So(result, ShouldContainKey, "data_string")
		So(result["data_string"], ShouldEqual, "test")
		So(result["data_string"], ShouldHaveSameTypeAs, "test")
	})
	Convey("test bool", t, func() {
		data := `{
		"data_bool": true
		}`
		result := ParseJson([]byte(data))
		So(result, ShouldContainKey, "data_bool")
		So(result["data_bool"], ShouldEqual, true)
		So(result["data_bool"], ShouldHaveSameTypeAs, true)
	})
	Convey("test float", t, func() {
		data := `{
		"data_float": 3.1415
		}`
		result := ParseJson([]byte(data))
		So(result, ShouldContainKey, "data_float")
		So(result["data_float"], ShouldEqual, 3.1415)
		So(result["data_float"], ShouldHaveSameTypeAs, 3.1415)
	})
	Convey("test slice int64", t, func() {
		data := `{
		"data_slice_int64":[3,27,28]
		}`
		result := ParseJson([]byte(data))
		So(result, ShouldContainKey, "data_slice_int64")
		So(result["data_slice_int64"], ShouldHaveSameTypeAs, []interface{}{})
		dataSliceInt64 := result["data_slice_int64"].([]interface{})
		So(dataSliceInt64, ShouldHaveLength, 3)
		So(dataSliceInt64[0], ShouldEqual, int64(3))
		So(dataSliceInt64[1], ShouldEqual, int64(27))
		So(dataSliceInt64[2], ShouldEqual, int64(28))
	})
	Convey("test slice float64", t, func() {
		data := `{
		"data_slice_float":[3.1415,27.1314,28.7466]
		}`
		result := ParseJson([]byte(data))
		So(result, ShouldContainKey, "data_slice_float")
		So(result["data_slice_float"], ShouldHaveSameTypeAs, []interface{}{})
		dataSliceFloat := result["data_slice_float"].([]interface{})
		So(dataSliceFloat, ShouldHaveLength, 3)
		So(dataSliceFloat[0], ShouldEqual, float64(3.1415))
		So(dataSliceFloat[1], ShouldEqual, float64(27.1314))
		So(dataSliceFloat[2], ShouldEqual, float64(28.7466))
	})
	Convey("test slice bool", t, func() {
		data := `{
		"data_slice_bool":[true,false,true]
		}`
		result := ParseJson([]byte(data))
		So(result, ShouldContainKey, "data_slice_bool")
		So(result["data_slice_bool"], ShouldHaveSameTypeAs, []interface{}{})
		dataSliceBool := result["data_slice_bool"].([]interface{})
		So(dataSliceBool, ShouldHaveLength, 3)
		So(dataSliceBool[0], ShouldEqual, true)
		So(dataSliceBool[1], ShouldEqual, false)
		So(dataSliceBool[2], ShouldEqual, true)
	})
	Convey("test slice string", t, func() {
		data := `{
		"data_slice_string":["3","27","28"]
		}`
		result := ParseJson([]byte(data))
		So(result, ShouldContainKey, "data_slice_string")
		So(result["data_slice_string"], ShouldHaveSameTypeAs, []interface{}{})
		dataSliceString := result["data_slice_string"].([]interface{})
		So(dataSliceString, ShouldHaveLength, 3)
		So(dataSliceString[0], ShouldEqual, "3")
		So(dataSliceString[1], ShouldEqual, "27")
		So(dataSliceString[2], ShouldEqual, "28")
	})
	Convey("test object", t, func() {
		data := `{
		"a":{
			"b":{
				"c":4,
				"d":27.1314,
				"e":true,
				"f":"test",
				"g": [3,27,28]
		}
	}
}`
		result := ParseJson([]byte(data))
		So(result, ShouldHaveLength, 5)
		So(result, ShouldContainKey, "a@b@c")
		So(result["a@b@c"], ShouldEqual, int64(4))
		So(result, ShouldContainKey, "a@b@d")
		So(result["a@b@d"], ShouldEqual, float64(27.1314))
		So(result, ShouldContainKey, "a@b@e")
		So(result["a@b@e"], ShouldEqual, true)
		So(result, ShouldContainKey, "a@b@f")
		So(result["a@b@f"], ShouldEqual, "test")
		So(result, ShouldContainKey, "a@b@g")
		So(result["a@b@g"], ShouldHaveSameTypeAs, []interface{}{})
		sliceInt64 := result["a@b@g"].([]interface{})
		So(sliceInt64, ShouldHaveLength, 3)
		So(sliceInt64[0], ShouldEqual, 3)
		So(sliceInt64[1], ShouldEqual, 27)
		So(sliceInt64[2], ShouldEqual, 28)
	})
}
