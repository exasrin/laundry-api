package controller

import (
	"fmt"
	"go-api-enigma/model"
	"go-api-enigma/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UomController struct {
	uomUC usecase.UomUseCase
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
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, uom)
}

func (u *UomController) updateHandler(c *gin.Context) {
	fmt.Println("==========================")
	var uom model.Uom
	if err := c.ShouldBindJSON(&uom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := u.uomUC.Update(uom); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully update uom",
	})
}

func (u *UomController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := u.uomUC.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	message := fmt.Sprintf("successfully delete uom with id %s", id)
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func (u *UomController) Route(engine *gin.Engine) {
	rg := engine.Group("/api/v1")
	rg.POST("/uoms", u.createUomHandler)
	rg.GET("/uoms", u.listUomHandler)
	rg.GET("/uoms/:id", u.getById)
	rg.PUT("/uoms", u.updateHandler)
	rg.DELETE("/uoms/:id", u.deleteHandler)
}

func NewUomController(uomUc usecase.UomUseCase) *UomController {
	return &UomController{
		uomUC: uomUc,
	}
}
