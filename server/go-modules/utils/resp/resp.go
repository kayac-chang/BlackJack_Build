package resp

type RespData struct {
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
	// Message string      `json:"msg"`
}

func Resp(statusCode int, data interface{}, err error) RespData {
	res := RespData{
		Code: statusCode,
		Data: data,
	}
	if err != nil {
		res.Data = err.Error()
	}
	return res
}
