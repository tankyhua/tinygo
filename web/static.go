package web

import (
	"path/filepath"
	"strings"
)

// http Content-Type,建立文件扩展名和内容类型的关系
var contentType = map[string]string{
	"apk": "application/vnd.android.package-archive",
}

// SetContentType 设置指定扩展名对应的http内容类型
func SetContentType(ext, ct string) {
	contentType[ext] = ct
}

// ContentType 获取指定扩展名对应的http内容类型
func ContentType(ext string) string {
	var ct, ok = contentType[strings.ToLower(ext)]
	if !ok {
		return ""
	}
	return ct
}

// 静态文件执行器
type StaticExecutor struct {
	CommonExecutor
	path string
}

// NewStaticExecutor 创建静态文件执行器
func NewStaticExecutor(path string) *StaticExecutor {
	var se = new(StaticExecutor)
	se.path = path
	return se
}

// Excute 执行
func (this *StaticExecutor) Execute() interface{} {
	var context, ok = this.Context.(*Context)
	if ok {
		context.End = this.End
		if this.ExecutePreFilters() {
			var result Result = nil
			if context.HttpContext.Request.Method == "GET" {
				//返回文件
				var filePath = context.HttpContext.Request.URL.Path
				if !strings.Contains(filePath, "..") {
					var r = this.Router()
					var count = 0
					for r != nil {
						r = r.Parent()
						count++
					}
					var paths = strings.Split(filePath, "/")
					filePath = filepath.Join(this.path, strings.Join(paths[count:], "/"))
					result = context.File(filePath)
				}
			}
			if result == nil {
				result = context.NotFound()
			}
			//执行过滤器
			if this.ExecutePostFilters(result) {
				return result
			}
		}
	}
	return nil
}

// 文件执行器,用于返回特定文件
type FileExecutor struct {
	CommonExecutor
	path string
}

// NewFileExecutor 创建文件执行器
func NewFileExecutor(path string) *FileExecutor {
	var fe = new(FileExecutor)
	fe.path = path
	return fe
}

// Excute 执行
func (this *FileExecutor) Execute() interface{} {
	var context, ok = this.Context.(*Context)
	if ok {
		context.End = this.End
		if this.ExecutePreFilters() {
			//返回文件
			var file = context.File(this.path)
			//执行过滤器
			if this.ExecutePostFilters(file) {
				return file
			}
		}
	}
	return nil
}