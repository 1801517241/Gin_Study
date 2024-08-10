package res

import "github.com/gin-gonic/gin"

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

var codeMap = map[int]string{
	1001: "权限错误",
	1002: "角色错误",
}

func response(context *gin.Context, code int, data any, msg string) {
	context.JSON(200, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func OK(context *gin.Context, data any, msg string) {
	response(context, 0, data, msg)
}

func OkWithData(context *gin.Context, data any) {
	OK(context, data, "成功")
}

func OkWithMsg(context *gin.Context, msg string) {
	OK(context, gin.H{}, msg)
}

func Fail(context *gin.Context, code int, data any, msg string) {
	response(context, code, data, msg)
}

func FailWithMsg(context *gin.Context, msg string) {
	response(context, 1001, nil, msg)
}

func FailWithCode(context *gin.Context, code int) {
	msg, ok := codeMap[code]
	if !ok {
		msg = "服务错误"
	}
	response(context, code, nil, msg)
}
