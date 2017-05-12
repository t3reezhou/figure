package inject

import (
	"fmt"
	"reflect"
	"strings"
)

func Inject(vals map[string]interface{}, dst interface{}, structTag string) (err error) {
	defer func() {
		e := recover()
		if e != nil {
			if v, ok := e.(error); ok {
				err = fmt.Errorf("Panic: %v", v.Error())
			} else {
				err = fmt.Errorf("Panic: %v", e)
			}
		}
	}()
	dt := reflect.TypeOf(dst)
	dv := reflect.ValueOf(dst)
	if dv.Kind() != reflect.Ptr || dv.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("not a pointer of struct")
	}
	for i := 0; i < dt.Elem().NumField(); i++ {
		field := dt.Elem().Field(i)
		fv := dv.Elem().Field(i)
		ft := field.Type

		if field.Anonymous || !fv.CanSet() {
			continue
		}

		// vals, ok := ctx.Value(key).(map[string]interface{})
		// if !ok {
		// 	return fmt.Errorf("ctx without %s", key)
		// }
		val, err := GetValue(vals, field, structTag)
		if err != nil {
			return err
		}
		if val == nil {
			continue
		}

		err = distribution(ft.Kind(), reflect.ValueOf(val), fv, ft)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetValue(vals map[string]interface{}, f reflect.StructField, structTag string) (interface{}, error) {
	tag := f.Tag.Get(structTag)
	tags := parseTag(tag)

	var name, option string
	if len(tags) > 0 {
		name = tags[0]
	}
	if len(tags) > 1 {
		option = tags[1]
	}

	if name == "-" {
		return nil, nil
	}
	if name == "" {
		name = strings.ToLower(f.Name)
	}

	// check tag option
	val, ok := vals[name]
	if !ok && option == "required" { // value not found
		return nil, fmt.Errorf("'%v' not found", name)
	}
	return val, nil
}

func distribution(kind reflect.Kind, val, dst reflect.Value, ft reflect.Type) error {
	if val.Type().Kind() == reflect.Interface {
		val = reflect.ValueOf(val.Interface())
	}
	switch kind {
	case reflect.Slice:
		return convertSlice(val, dst, ft)
	case reflect.String:
		if val.Type().Kind() == reflect.String {
			convertString(val, dst)
			return nil
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if val.Type().Kind() <= reflect.Int64 && val.Type().Kind() >= reflect.Int {
			convertInt(val, dst)
			return nil
		}
	case reflect.Float32, reflect.Float64:
		if val.Type().Kind() <= reflect.Float64 && val.Type().Kind() >= reflect.Float32 {
			convertFloat(val, dst)
			return nil
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if val.Type().Kind() <= reflect.Uint64 && val.Type().Kind() >= reflect.Uint {
			convertUint(val, dst)
			return nil
		}
	case reflect.Bool:
		if val.Type().Kind() == reflect.Bool {
			converBool(val, dst)
			return nil
		}
	}
	return fmt.Errorf("inject type is %s,but get the value type is %s,values is %+v",
		kind, val.Type().Kind(), val)
}

func convertSlice(sliceValue, fv reflect.Value, ft reflect.Type) error {
	trueType := ft.Elem()

	dstslice := reflect.MakeSlice(ft, 0, 0)
	for i := 0; i < sliceValue.Len(); i++ {
		trueValue := reflect.New(trueType).Elem()

		err := distribution(trueType.Kind(), sliceValue.Index(i), trueValue, ft)
		if err != nil {
			return err
		}
		dstslice = reflect.Append(dstslice, trueValue)
	}
	fv.Set(dstslice)
	return nil
}

func convertInt(value, dst reflect.Value) {
	dst.SetInt(value.Int())
}

func convertString(value, dst reflect.Value) {
	dst.SetString(value.String())
}

func convertUint(value, dst reflect.Value) {
	dst.SetUint(value.Uint())
}
func convertFloat(value, dst reflect.Value) {
	dst.SetFloat(value.Float())
}

func converBool(value, dst reflect.Value) {
	dst.SetBool(value.Bool())
}

func parseTag(tag string) []string {
	return strings.Split(tag, ",")
}
