package web

import "io"

// 服务端状态码
type StatusCode int

const (
	//Http状态码
	StatusCodeOK               StatusCode = 200 //http正常返回结果
	StatusCodeMovedPermanently StatusCode = 301 //http永久转移
	StatusCodeMovedTemporarily StatusCode = 302 //http临时转移
	StatusCodeNotFound         StatusCode = 404 //http页面未找到
	//框架内部状态码
	StatusCodeParamNotEnough  StatusCode = iota + 10000 //http参数不足
	StatusCodeParamNotCorrect                           //http参数不正确
	StatusCodePageNotFound                              //页面未找到
	//用户自定义状态码
	StatusCodeUserDefined StatusCode = 1000000 //用户自定义状态码起始码
)

// 可用于默认http方法(默认为Post)的返回结果
type Result interface {
	// WriteTo 将Result的内容写入writer
	WriteTo(writer io.Writer) error
}

// 可用于Get方法的返回结果
type GetResult Result

// 可用于Post方法的返回结果
type PostResult Result

// 可用于Put方法的返回结果
type PutResult Result

// 可用于Delete方法的返回结果
type DeleteResult Result

// 可用于Options方法的返回结果
type OptionsResult Result

// 可用于Head方法的返回结果
type HeadResult Result

// 可用于Trace方法的返回结果
type TraceResult Result

// 可用于Connect方法的返回结果
type ConnectResult Result

// 可用于Get和Post方法的返回结果
type GetPostResult Result

// 用于http的结果
type HttpResult interface {
	// Code 返回状态码
	Code() StatusCode
	// Message 返回状态信息
	Message() string
	// 实现Result接口
	Result
}
