package api

import (
	"github.com/JayantiTA/backend-weights/internal/entity"
	"github.com/JayantiTA/backend-weights/internal/usecase"
	"github.com/JayantiTA/backend-weights/utils"
	"github.com/gin-gonic/gin"
)

type WeightAPI struct {
	weightUsecase usecase.Weight
}

func NewWeightAPI(weightUsecase usecase.Weight) WeightAPI {
	return WeightAPI{
		weightUsecase: weightUsecase,
	}
}

func (api *WeightAPI) GetAll(c *gin.Context) {
	weights, err := api.weightUsecase.GetAll()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "error getting weights",
		})
		return
	}
	c.IndentedJSON(200, weights)
}

func (api *WeightAPI) Get(c *gin.Context) {
	dateUnix := c.Param("date")
	date, err := utils.UnixToTime(dateUnix)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "error parsing date",
		})
		return
	}
	weight, err := api.weightUsecase.Get(&date)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "error getting weight",
		})
		return
	}
	c.IndentedJSON(200, weight)
}

func (api *WeightAPI) CreateOrUpdate(c *gin.Context) {
	var input entity.WeightDto
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "error binding json",
		})
		return
	}
	err = api.weightUsecase.CreateOrUpdate(input.Date, input.Max, input.Min)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "error creating or updating weight",
		})
		return
	}
	c.IndentedJSON(200, gin.H{
		"message": "success",
	})
}

func (api *WeightAPI) Delete(c *gin.Context) {
	dateUnix := c.Param("date")
	date, err := utils.UnixToTime(dateUnix)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "error parsing date",
		})
		return
	}
	err = api.weightUsecase.Delete(&date)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "error deleting weight",
		})
		return
	}
	c.IndentedJSON(200, gin.H{
		"message": "success",
	})
}
