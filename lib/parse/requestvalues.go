package parse

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func ParseRequest(r *http.Request, setFunc func(map[string]interface{})) (values map[string]interface{}) {
	values = ParseForm(r)
	if setFunc != nil {
		setFunc(values)
	}
	defer r.Body.Close()
	data, _ := ioutil.ReadAll(r.Body)
	for k, v := range ParseJson(data) {
		values[k] = v
	}
	return
}

func ParseForm(r *http.Request) map[string]interface{} {
	r.ParseForm()
	values := make(map[string]interface{}, 0)

	for k, v := range r.Form {
		if value, err := strconv.ParseInt(strings.Join(v, ""), 10, 64); err == nil {
			values[k] = value
			continue
		}
		values[k] = strings.Join(v, "")
	}
	return values
}

func ParseJson(data []byte) map[string]interface{} {
	values := make(map[string]interface{}, 0)
	var jsonValues = make(map[string]json.RawMessage, 0)
	json.Unmarshal(data, &jsonValues)
	for k, raw := range jsonValues {
		distributionJson(raw, k, values)
	}
	return values
}

func distributionJson(raw json.RawMessage, key string, valueContainer map[string]interface{}) {
	if len(raw) > 0 {
		switch raw[0] {
		case '"':
			parseJsonString(raw, key, valueContainer)
		case '[':
			parseJsonSlice(raw, key, valueContainer)
		case '{':
			parseJsonObect(raw, key, valueContainer)
		case 't', 'T', 'f', 'F':
			parseJsonBool(raw, key, valueContainer)
		default:
			parseJsonNormal(raw, key, valueContainer)
		}
	}
}

func parseJsonObect(object json.RawMessage, key string, valueContainer map[string]interface{}) {
	var jsonValues = make(map[string]json.RawMessage, 0)
	json.Unmarshal(object, &jsonValues)
	for k, raw := range jsonValues {
		distributionJson(raw, key+"@"+k, valueContainer)
	}
}

func parseJsonSlice(slice json.RawMessage, key string, valueContainer map[string]interface{}) {
	var jsonValues []json.RawMessage
	var values []interface{}
	err := json.Unmarshal(slice, &jsonValues)
	if err != nil {
		fmt.Println("parseJsonSlice wrong")
		return
	}
	for _, v := range jsonValues {
		if len(v) > 0 {
			switch v[0] {
			case '"':
				values = parseSliceString(v, values)
			case '[', '{':
				values = parseSliceObejectOrSlice(v, values)
			case 't', 'T', 'f', 'F':
				values = parseSliceBool(v, values)
			default:
				values = parseSliceNormal(v, values)
			}
		}
	}
	valueContainer[key] = values
}

func parseJsonString(dataString json.RawMessage, key string, valueContainer map[string]interface{}) {
	valueContainer[key] = string(dataString[1 : len(dataString)-1])
}

func parseJsonNormal(dataNormal json.RawMessage, key string, valueContainer map[string]interface{}) {
	if value, err := strconv.ParseInt(string(dataNormal), 10, 64); err == nil {
		valueContainer[key] = value
		return
	}
	if value, err := strconv.ParseFloat(string(dataNormal), 64); err == nil {
		valueContainer[key] = value
		return
	}
}

func parseJsonBool(dataBool json.RawMessage, key string, valueContainer map[string]interface{}) {
	if value, err := strconv.ParseBool(string(dataBool)); err == nil {
		valueContainer[key] = value
	}
}

func parseSliceString(dataString []byte, valueContainer []interface{}) []interface{} {
	return append(valueContainer, string(dataString[1:len(dataString)-1]))
}
func parseSliceObejectOrSlice(dataObject []byte, valueContainer []interface{}) []interface{} {
	return append(valueContainer, string(dataObject))
}

func parseSliceBool(dataBool []byte, valueContain []interface{}) []interface{} {
	if value, err := strconv.ParseBool(string(dataBool)); err == nil {
		return append(valueContain, value)
	}
	return valueContain
}

func parseSliceNormal(dataNormal []byte, valueContain []interface{}) []interface{} {
	if value, err := strconv.ParseInt(string(dataNormal), 10, 64); err == nil {
		return append(valueContain, value)

	}
	if value, err := strconv.ParseFloat(string(dataNormal), 64); err == nil {
		return append(valueContain, value)
	}
	return valueContain
}
