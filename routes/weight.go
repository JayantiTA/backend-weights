package route

import (
	"github.com/JayantiTA/backend-weights/api"
	"github.com/JayantiTA/backend-weights/internal/repository"
	"github.com/JayantiTA/backend-weights/internal/service"
	"github.com/JayantiTA/backend-weights/internal/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitWeightRoutes(db *gorm.DB, route *gin.Engine) {
	weightRepository := repository.NewWeight(db)
	weightService := service.NewWeight(weightRepository)
	weightUsecase := usecase.NewWeight(weightService)
	weightAPI := api.NewWeightAPI(weightUsecase)

	groupRoute := route.Group("/api")
	groupRoute.GET("/weights", weightAPI.GetAll)
	groupRoute.GET("/weight/:date", weightAPI.Get)
	groupRoute.POST("/weight", weightAPI.CreateOrUpdate)
	groupRoute.DELETE("/weight", weightAPI.Delete)
}
