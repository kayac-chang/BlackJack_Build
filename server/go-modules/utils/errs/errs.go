package errs

import "fmt"

var (
	errPostArgsNotFound    string = "Args [%s] not found."
	errPostArgsInvalied    string = "Args [%s] invalied. (%v)"
	errVendorTokenInvalied string = "Vendor [%s] is invalied."
)

func NewErrPostArgsNotFound(key string) error {
	return fmt.Errorf(errPostArgsNotFound, key)
}

func NewErrVendorTokenInvalied(key string) error {
	return fmt.Errorf(errVendorTokenInvalied, key)
}

func NewErrPostArgsInvalied(key string, err error) error {
	return fmt.Errorf(errPostArgsInvalied, key, err)
}

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400

	ERROR_EXIST_TAG         = 10001
	ERROR_NOT_EXIST_TAG     = 10002
	ERROR_NOT_EXIST_ARTICLE = 10003

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
)

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "請求參數錯誤",
	ERROR_EXIST_TAG:                "標籤已存在",
	ERROR_NOT_EXIST_TAG:            "標籤不存在",
	ERROR_NOT_EXIST_ARTICLE:        "文章不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token驗證失敗",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超時",
	ERROR_AUTH_TOKEN:               "Token生成失敗",
	ERROR_AUTH:                     "Token錯誤",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
