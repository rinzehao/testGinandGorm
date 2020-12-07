package common

type HttpResp struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	ErrCode string      `json:"err_code"`
	ErrMsg  string      `json:"err_msg"`
}
