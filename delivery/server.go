package delivery

import (
	"fmt"
	"go-api-enigma/config"
	"go-api-enigma/delivery/controller"
	"go-api-enigma/delivery/middleware"
	"go-api-enigma/manager"

	"github.com/gin-gonic/gin"
	"github.com/siruspen/logrus"
)

type Server struct {
	ucManager manager.UseCaseManager
	// productUc usecase.ProductUseCase
	engine *gin.Engine
	host   string
	log    *logrus.Logger
}

func (s *Server) Run() {
	s.initMiddleware()
	// s.initControllers()
	err := s.engine.Run(s.host)
	if err != nil {
		panic(err)
	}
}

func (s *Server) initMiddleware() {
	s.engine.Use(middleware.LogRequestMiddleware(s.log))
}

func (s *Server) initControllers() {
	s.engine.Use(middleware.LogRequestMiddleware(s.log))
	// Inisialisasi controller
	controller.NewUomController(s.ucManager.UomUseCase()).Route(s.engine)
}

func NewServer() *Server {
	cfg, error := config.NewConfig()
	if error != nil {
		fmt.Println(error)
	}

	infraManager, error := manager.NewInfraManager(cfg)
	if error != nil {
		fmt.Println(error)
	}
	rm := manager.NewRepoManager(infraManager)
	ucm := manager.NewUseCaseManager(rm)
	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	log := logrus.New()
	engine := gin.Default()
	server := Server{
		ucManager: ucm,
		engine:    engine,
		host:      host,
		log:       log,
	}
	server.initControllers()
	return &server
}
