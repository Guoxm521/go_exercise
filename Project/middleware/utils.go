package middleware

import "github.com/gin-gonic/gin"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func newResponse(code int, msg string, data interface{}) *Response {
	return &Response{code, msg, data}
}
func response(ctx *gin.Context, rs *Response) {
	ctx.JSON(200, map[string]interface{}{
		"code":    rs.Code,
		"message": rs.Message,
		"data":    rs.Data,
	})
	ctx.Abort()
}
