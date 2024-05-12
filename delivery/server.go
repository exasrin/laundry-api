package delivery

import (
	"fmt"
	"go-api-enigma/config"
	"go-api-enigma/delivery/controller"
	"go-api-enigma/manager"

	"github.com/gin-gonic/gin"
)

type Server struct {
	ucManager manager.UseCaseManager
	// productUc usecase.ProductUseCase
	engine *gin.Engine
	host   string
}

func (s *Server) Run() {
	err := s.engine.Run(s.host)
	if err != nil {
		panic(err)
	}
}

func (s *Server) initControllers() {
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
	engine := gin.Default()
	server := Server{
		ucManager: ucm,
		engine:    engine,
		host:      host,
	}
	server.initControllers()
	return &server
}
