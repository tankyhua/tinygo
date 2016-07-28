package tinygo

// Http方法常量
type HttpMethod string

const (
	HttpMethodGet    HttpMethod = "GET"    //GET方法
	HttpMethodPost   HttpMethod = "POST"   //POST方法
	HttpMethodPut    HttpMethod = "PUT"    //PUT方法
	HttpMethodDelete HttpMethod = "DELETE" //DELETE方法
)

// 默认配置文件名
const DefaultConfigPath = "web.cfg"

// 默认布局配置文件名
const DefaultLayoutConfigFileName = "layout.json"

// 默认视图文件扩展名
const DefaultTemplateExt = ".html"

// 默认模板文件内模板名,用于返回部分视图时使用
const DefaultTemplateName = "Content"

// 默认Session Cookie名
const DefaultSessionCookieName = "__SessionId"

// 默认CSRF Cookie 名
const DefaultCSRFCookieName = "__CSRFId"

// 默认CSRF 表单名
const DefaultCSRFTokenName = "csrf"

// 默认最大文件表单占用内存大小
const DefaultMaxMemory = 32 << 20 // 32 MB

// Api格式类型
type ApiType string

const (
	ApiTypeAuto ApiType = "auto" //根据请求自动判断
	ApiTypeJson ApiType = "json" //Json
	ApiTypeXml  ApiType = "xml"  //Xml
)