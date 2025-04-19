package services

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

// 用于动态管理后台 server
type RouteServer struct {
	echo *echo.Echo
}

func NewRouteServer() *RouteServer {
	return &RouteServer{
		echo: echo.New(),
	}
}

func (s *RouteServer) initRouteV1() {
	g := s.echo.Group("/v1")
	g.GET("/servers", nil)
}

func (s *RouteServer) Start(addr string) {
	if err := s.echo.Start(addr); err != nil {
		fmt.Println("start mcpone server")
	}
}
