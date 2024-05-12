package delivery

import (
	"fmt"
	"go-api-enigma/config"
	"go-api-enigma/delivery/controller"
	"go-api-enigma/repository"
	"go-api-enigma/usecase"

	"github.com/gin-gonic/gin"
)

type Server struct {
	uomUc usecase.UomUseCase
	// productUc usecase.ProductUseCase
	engine *gin.Engine
}

func (s *Server) Run() {
	err := s.engine.Run()
	if err != nil {
		panic(err)
	}
}

func (s *Server) initControllers() {
	// Inisialisasi controller
	controller.NewUomController(s.uomUc, s.engine)
}

func NewServer() *Server {
	cfg, error := config.NewConfig()
	if error != nil {
		fmt.Println(error)
	}

	conn, error := config.NewDbCOnnection(cfg)
	if error != nil {
		fmt.Println(error)
	}
	db := conn.Conn()

	// instace repo
	uomRepo := repository.NewUomRepository(db)
	// productRepo := repository.NewProductRepository(db)

	// instance Uc
	uomUC := usecase.NewUomUseCase(uomRepo)

	// controller
	engine := gin.Default()
	server := Server{
		uomUc:  uomUC,
		engine: engine,
	}
	server.initControllers()

	return &server
}
