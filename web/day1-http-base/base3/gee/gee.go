package gee

import (
	"fmt"
	"net/http"
	"strconv"
)

// HandlerFunc 定义了handler函数的函数类型
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine 定义web请求路径到实际执行函数的映射
type Engine struct {
	router map[string]HandlerFunc
}

// New 初始化函数： 返回一个web service实体
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// 添加一个请求路径到请求处理函数的映射，并且还可以指定请求的方式（GET | POST）
func (engine *Engine) addRoute(method string, pattern string, handlerFunc HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handlerFunc
}

func (engine *Engine) GET(pattern string, handlerFunc HandlerFunc) {
	engine.addRoute("GET", pattern, handlerFunc)
}

func (engine *Engine) POST(pattern string, handlerFunc HandlerFunc) {
	engine.addRoute("POST", pattern, handlerFunc)
}

func (engine *Engine) Run(address int) (err error) {
	return http.ListenAndServe(":"+strconv.Itoa(address), engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	handlerFunc, ok := engine.router[key]
	if ok { //	自定了该请求的响应操作
		handlerFunc(w, req)
	} else { //	未找到该请求的响应操作
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
