package response

import (
	"net/http"
)

// 定义标准响应状态码
const (
	CodeSuccess          = 200   // 成功
	CodeBadRequest       = 400   // 错误的请求
	CodeUnauthorized     = 401   // 未授权
	CodeForbidden        = 403   // 禁止访问
	CodeNotFound         = 404   // 资源不存在
	CodeInternalError    = 500   // 服务器内部错误
	CodeValidationError  = 422   // 参数验证错误
	CodeServiceError     = 503   // 服务不可用
)

// Response 标准响应结构
type Response struct {
	Code       int         `json:"code"`                // 状态码
	Message    string      `json:"message"`             // 消息
	Data       interface{} `json:"data,omitempty"`      // 数据
	RequestId  string      `json:"request_id,omitempty"`// 请求ID，用于追踪
	Errors     []string    `json:"errors,omitempty"`    // 错误详情
	Pagination *Pagination `json:"pagination,omitempty"` // 分页信息
}

// Pagination 分页信息
type Pagination struct {
	Current  int   `json:"current"`  // 当前页
	PageSize int   `json:"pageSize"` // 每页大小
	Total    int64 `json:"total"`    // 总记录数
}

// Success 成功响应
func Success(data interface{}) *Response {
	return &Response{
		Code:    CodeSuccess,
		Message: "success",
		Data:    data,
	}
}

// SuccessWithPagination 带分页的成功响应
func SuccessWithPagination(data interface{}, current, pageSize int, total int64) *Response {
	return &Response{
		Code:    CodeSuccess,
		Message: "success",
		Data:    data,
		Pagination: &Pagination{
			Current:  current,
			PageSize: pageSize,
			Total:    total,
		},
	}
}

// Error 错误响应
func Error(code int, message string) *Response {
	return &Response{
		Code:    code,
		Message: message,
	}
}

// ErrorWithDetails 带详细错误信息的错误响应
func ErrorWithDetails(code int, message string, errors []string) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Errors:  errors,
	}
}

// BadRequest 400错误
func BadRequest(message string) *Response {
	return Error(CodeBadRequest, message)
}

// Unauthorized 401错误
func Unauthorized(message string) *Response {
	if message == "" {
		message = "unauthorized access"
	}
	return Error(CodeUnauthorized, message)
}

// Forbidden 403错误
func Forbidden(message string) *Response {
	if message == "" {
		message = "forbidden access"
	}
	return Error(CodeForbidden, message)
}

// NotFound 404错误
func NotFound(message string) *Response {
	if message == "" {
		message = "resource not found"
	}
	return Error(CodeNotFound, message)
}

// ValidationError 422错误
func ValidationError(message string, errors []string) *Response {
	return ErrorWithDetails(CodeValidationError, message, errors)
}

// InternalError 500错误
func InternalError(message string) *Response {
	if message == "" {
		message = "internal server error"
	}
	return Error(CodeInternalError, message)
}

// ServiceUnavailable 503错误
func ServiceUnavailable(message string) *Response {
	if message == "" {
		message = "service temporarily unavailable"
	}
	return Error(CodeServiceError, message)
}

// WithRequestID 添加请求ID
func (r *Response) WithRequestID(requestID string) *Response {
	r.RequestId = requestID
	return r
}

// IsSuccess 判断是否成功响应
func (r *Response) IsSuccess() bool {
	return r.Code == CodeSuccess
}

// GetHTTPStatusCode 获取对应的HTTP状态码
func (r *Response) GetHTTPStatusCode() int {
	switch r.Code {
	case CodeSuccess:
		return http.StatusOK
	case CodeBadRequest:
		return http.StatusBadRequest
	case CodeUnauthorized:
		return http.StatusUnauthorized
	case CodeForbidden:
		return http.StatusForbidden
	case CodeNotFound:
		return http.StatusNotFound
	case CodeValidationError:
		return http.StatusUnprocessableEntity
	case CodeServiceError:
		return http.StatusServiceUnavailable
	default:
		return http.StatusInternalServerError
	}
} 