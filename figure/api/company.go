package api

type CompanyHandler struct {
}

func NewComPanyHandler() *CompanyHandler {
	return &CompanyHandler{}
}

// func (c *CompanyHandler) Get(params martini.Params) interface{} {
// 	result := make(map[string]interface{}, 0)
// 	result["companyid"] = 10086
// 	result["name"] = "ALTER"
// 	return result
// }
