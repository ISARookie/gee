package gee

import "net/http"

// H /** 自定义类型H 为string-interface map
type H map[string]interface{}

type Context struct {
	Write http.ResponseWriter
}
