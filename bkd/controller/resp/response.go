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
