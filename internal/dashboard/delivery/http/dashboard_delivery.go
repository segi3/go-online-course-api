package dashboard

import (
	"net/http"

	useCase "online-course/internal/dashboard/usecase"
	"online-course/internal/middleware"
	"online-course/pkg/utils"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	useCase useCase.DashboardUseCase
}

func NewDashboardHandler(useCase useCase.DashboardUseCase) *DashboardHandler {
	return &DashboardHandler{useCase}
}

func (handler *DashboardHandler) Route(r *gin.RouterGroup) {
	dashboardHandler := r.Group("/api/v1")

	dashboardHandler.Use(middleware.AuthJwt, middleware.AuthAdmin)
	{
		dashboardHandler.GET("/dashboards", handler.GetDataDashboard)
	}
}

func (handler *DashboardHandler) GetDataDashboard(ctx *gin.Context) {
	data := handler.useCase.GetDataDashboard()

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", data))
}
