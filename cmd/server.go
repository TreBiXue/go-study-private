package main

import "github.com/gin-gonic/gin"

type Server struct {
	engine *gin.Engine
}

func (s *Server) Start() {
	s.engine.Run()
}

func NewServer(engine *gin.Engine) *Server {
	return &Server{
		engine: engine,
	}
}
