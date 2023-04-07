package app

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-studying/api/router"
	_ "go-studying/docs"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine        *gin.Engine
	accountRouter *router.AccountRouter
	centerRouter  *router.CenterRouter
	venderRouter  *router.VenderRouter
	nyukaRouter   *router.NyukaRouter
}

func (s *Server) Start() {
	//url := ginSwagger.URL("http://127.0.0.1:8000/swagger/doc.json")
	s.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	s.accountRouter.Login(s.engine)
	s.centerRouter.Center(s.engine)
	s.venderRouter.Vender(s.engine)
	s.nyukaRouter.Nyuka(s.engine)
	s.engine.Run()
}

func NewServer(engine *gin.Engine, accountRouter *router.AccountRouter, centerRouter *router.CenterRouter,
	venderRouter *router.VenderRouter, nyukaRouter *router.NyukaRouter) *Server {
	return &Server{
		engine:        engine,
		accountRouter: accountRouter,
		centerRouter:  centerRouter,
		venderRouter:  venderRouter,
		nyukaRouter:   nyukaRouter,
	}
}
