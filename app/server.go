package app

import (
	"go-studying/api/router"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine       *gin.Engine
	apiRouter    *router.Router
	centerRouter *router.CenterRouter
}

func (s *Server) Start() {
	s.apiRouter.Login(s.engine)
	// s.centerRouter.Center(s.engine)
	s.engine.Run()
}

func NewServer(engine *gin.Engine, apiRouter *router.Router) *Server {
	return &Server{
		engine:    engine,
		apiRouter: apiRouter,
	}
}

//func NewServer(engine *gin.Engine, apiRouter *router.Router, centerRouter *router.CenterRouter) *Server {
//	return &Server{
//		engine:       engine,
//		apiRouter:    apiRouter,
//		centerRouter: centerRouter,
//	}
//}
