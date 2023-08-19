package gee


// Engine implement the interface of ServeHTTP
type Engine struct {
	router *router
}

// New is the constructor of gee.Engine
fun New() *Engine {
	return &Engine{router: newRouter()}
}