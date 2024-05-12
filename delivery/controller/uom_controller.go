package controller

import (
	"go-api-enigma/model"
	"go-api-enigma/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UomController struct {
	uomUC  usecase.UomUseCase
	router *gin.Engine
}

func (u *UomController) createUomHandler(c *gin.Context) {
	var uom model.Uom
	if err := c.ShouldBindJSON(&uom); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	uom.Id = uuid.NewString()
	if err := u.uomUC.CreateNew(uom); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, uom)
}

func (u *UomController) listUomHandler(c *gin.Context) {
	uoms, err := u.uomUC.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, uoms)
}

func (u *UomController) getById(c *gin.Context) {
	id := c.Param("id")
	uom, err := u.uomUC.FindById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message" : err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, uom)
}

func NewUomController(uomUc usecase.UomUseCase, engine *gin.Engine) {
	controller := UomController{
		uomUC:  uomUc,
		router: engine,
	}

	rg := engine.Group("/api/v1")
	rg.POST("/uoms", controller.createUomHandler)
	rg.GET("/uoms", controller.listUomHandler)
	rg.GET("/uoms/:id", controller.getById)
}
