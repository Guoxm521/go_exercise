package common

import "github.com/gin-gonic/gin"

type Resp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//响应
func Response(ctx *gin.Context, data interface{}) {
	_rs := parse(data)
	ctx.JSON(200, _rs)
}

//响应解析，返回统一数据格式
func parse(data interface{}) Resp {
	_code := 200
	_msg := "SUCCESS"
	var _rsData interface{}
	switch _t := data.(type) {
	case error:
		_code = 400
		_msg = _t.Error()
		_rsData = ""
	default:
		_rsData = data
	}
	return Resp{
		Code:    _code,
		Message: _msg,
		Data:    _rsData,
	}
}
