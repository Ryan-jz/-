// Package response 提供统一的响应处理工具
// 用于标准化 API 响应格式
package response

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// JsonRes 标准 JSON 响应结构
type JsonRes struct {
	Code    int         `json:"code"`    // 状态码
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据
}

// Json 返回标准 JSON 数据
// 用于成功响应
func Json(r *ghttp.Request, code int, message string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}

	r.Response.WriteJson(JsonRes{
		Code:    code,
		Message: message,
		Data:    responseData,
	})
}

// JsonExit 返回 JSON 数据并退出当前请求
// 用于错误响应或需要立即终止请求的场景
func JsonExit(r *ghttp.Request, code int, message string, data ...interface{}) {
	Json(r, code, message, data...)
	r.Exit()
}

// Success 返回成功响应
func Success(r *ghttp.Request, data ...interface{}) {
	Json(r, 200, "操作成功", data...)
}

// SuccessWithMessage 返回带自定义消息的成功响应
func SuccessWithMessage(r *ghttp.Request, message string, data ...interface{}) {
	Json(r, 200, message, data...)
}

// Fail 返回失败响应
func Fail(r *ghttp.Request, message string) {
	JsonExit(r, 500, message)
}

// FailWithCode 返回带状态码的失败响应
func FailWithCode(r *ghttp.Request, code int, message string) {
	JsonExit(r, code, message)
}

// PageRes 分页响应结构
type PageRes struct {
	List     interface{} `json:"list"`     // 数据列表
	Total    int         `json:"total"`    // 总记录数
	Page     int         `json:"page"`     // 当前页码
	PageSize int         `json:"pageSize"` // 每页大小
}

// Page 返回分页数据
func Page(r *ghttp.Request, list interface{}, total, page, pageSize int) {
	Json(r, 200, "查询成功", g.Map{
		"list":     list,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}
