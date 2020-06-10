package http

import "github.com/tpphu/gobox/service"
import "github.com/gin-gonic/gin"

type Http struct {
	service.Runable
	*gin.Engine
	addr string
}

type Option func(a *Http)

func Address(addr string) Option{
	return func(a *Http) {
		a.addr = addr
	}
}


func Default(opts ...Option) *Http {
	engine := gin.New()
	http := &Http{}
	http.Engine = engine
	http.addr = ":3000"
	return http
}

func (s *Http) Init() {
	s.Use(gin.Logger(), gin.Recovery())
	s.GET("/ping", func(c *gin.Context){
		c.String(200, "pong")
	})
}

func (s *Http) Run() {
	s.Engine.Run(s.addr)
}

func (s *Http) Shutdown() {
	
}

// add [[route, handler],[route, handler]] = group
// add [middleware,middleware,middleware...]
