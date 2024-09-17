package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"prova-esoft8s/data/request"
	"prova-esoft8s/data/response"
	"prova-esoft8s/helper"
	"prova-esoft8s/service"
	"strconv"
)

type ViagemController struct {
	viagemService service.ViagensService
}

func NewViagemController(service service.ViagensService) *ViagemController {
	return &ViagemController{viagemService: service}
}

func (controller *ViagemController) Create(ctx *gin.Context) {
	createViagemRequest := request.CreateViagensRequest{}
	err := ctx.ShouldBindJSON(&createViagemRequest)
	helper.ErrorPanic(err)

	controller.viagemService.Create(createViagemRequest)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ViagemController) Update(ctx *gin.Context) {
	updateViagemRequest := request.UpdateViagensRequest{}
	err := ctx.ShouldBindJSON(&updateViagemRequest)
	helper.ErrorPanic(err)

	viagemId := ctx.Param("viagemId")
	id, err := strconv.Atoi(viagemId)
	helper.ErrorPanic(err)

	updateViagemRequest.Id = id

	controller.viagemService.Update(updateViagemRequest)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ViagemController) Delete(ctx *gin.Context) {
	viagemId := ctx.Param("viagemId")
	id, err := strconv.Atoi(viagemId)
	helper.ErrorPanic(err)
	controller.viagemService.Delete(id)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *ViagemController) FindById(ctx *gin.Context) {
	viagemId := ctx.Param("viagemId")
	id, err := strconv.Atoi(viagemId)
	helper.ErrorPanic(err)

	viagemResponse := controller.viagemService.FindById(id)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   viagemResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ViagemController) FindAll(ctx *gin.Context) {
	viagemResponse := controller.viagemService.FindAll()

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   viagemResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
