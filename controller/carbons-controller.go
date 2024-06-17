package controller

import (
	"net/http"
	"sampah/dto"
	"sampah/helper"
	"sampah/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CarbonsController interface {
	GetDetailCarbons(ctx *gin.Context)
	GetFootPrint(ctx *gin.Context)
	InsertCarbons(ctx *gin.Context)
}

type carbonsController struct {
	carbonsService service.CarbonsService
	jwtService     service.JWTService
}

func NewCarbonsController(carb service.CarbonsService , jwt service.JWTService) CarbonsController {
	return &carbonsController{
		carbonsService: carb,
		jwtService:     jwt,
	}
}

func (c *carbonsController) GetDetailCarbons(ctx *gin.Context) {
	user_id := ctx.Param("user_id")
	id, err := strconv.Atoi(user_id)
	if err != nil {
        panic(err)
    }
	carbons := c.carbonsService.GetDetailCarbons(id)
	res := helper.BuildResponse(true, "Success", carbons)
	ctx.JSON(http.StatusOK, res)
}
func (c *carbonsController) GetFootPrint(ctx *gin.Context) {
	user_id := ctx.Param("user_id")
	id, err := strconv.Atoi(user_id)
	if err != nil {
        panic(err)
    }
	carbons := c.carbonsService.GetFootPrint(id)
	res := helper.BuildResponse(true, "Success", carbons)
	ctx.JSON(http.StatusOK, res)
}

func (c *carbonsController) InsertCarbons(ctx *gin.Context){
	var carbDTO dto.CarbonsCreateDTO
	errDTO := ctx.ShouldBind(&carbDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	art := c.carbonsService.InsertCarbons(carbDTO)
	res := helper.BuildResponse(true, "Success", art)
	ctx.JSON(http.StatusOK, res)
}