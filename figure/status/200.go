package status

// var (
// 	OK = "ok"
// )

var OK = &struct {
	Result  string `json:"result"`
	Message string `json:"msg"`
}{
	Result:  "ok",
	Message: "ok",
}
