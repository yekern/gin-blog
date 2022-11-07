package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	ctx *gin.Context
	msg *Message
}

type Message struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// NewResponse 构造函数
func NewResponse(ctx *gin.Context) *Response {
	return &Response{ctx: ctx, msg: &Message{
		Code:    http.StatusOK,
		Message: "ok",
		Data:    nil,
	}}
}

// Message 返回消息加数据
func (r *Response) Message(message string, data interface{}) {
	r.msg.Message = message
	r.msg.Data = data
	r.ctx.JSON(http.StatusOK, r.msg)
	r.ctx.Abort()
	return
}

// SuccessOk 返回成功信息
func (r *Response) SuccessOk() {
	r.Message("ok", nil)
}

// Data 返回数据信息
func (r *Response) Data(data interface{}) {
	r.Message("ok", data)
}

// Error 返回错误信息
func (r *Response) Error(message string) {
	r.SetCode(http.StatusBadRequest)
	r.Message(message, nil)
}

// SetCode 设置状态码
func (r *Response) SetCode(code int64) *Response {
	r.msg.Code = code
	return r
}
