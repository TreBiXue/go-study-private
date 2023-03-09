package app

import (
	"go-studying/api/router"

	"github.com/gin-gonic/gin"
)

// type Server struct {
// 	engine *gin.Engine
// }

// func (s *Server) Start() {
// 	s.engine.Run()
// }

// func NewServer(engine *gin.Engine) *Server {
// 	return &Server{engine: engine}
// }

type Server struct {
	engine        *gin.Engine
	accountRouter *router.AccountRouter
	centerRouter  *router.CenterRouter
	venderRouter  *router.VenderRouter
}

func (s *Server) Start() {
	s.accountRouter.Login(s.engine)
	s.centerRouter.Center(s.engine)
	s.venderRouter.Vender(s.engine)
	s.engine.Run()
}

func NewServer(engine *gin.Engine, accountRouter *router.AccountRouter, centerRouter *router.CenterRouter,
	venderRouter *router.VenderRouter) *Server {
	return &Server{
		engine:        engine,
		accountRouter: accountRouter,
		centerRouter:  centerRouter,
		venderRouter:  venderRouter,
	}
}
