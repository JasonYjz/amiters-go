package resp

type JsonResult struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func JsonData(code int, data interface{}) *JsonResult {
	return &JsonResult{
		Code: code,
		Data: data,
	}
}

func SucData() *JsonResult {
	return &JsonResult{
		Code: 200,
		Data: nil,
	}
}

func FailData() *JsonResult {
	return &JsonResult{
		Code: 400,
		Data: "Internal Error",
	}
}
