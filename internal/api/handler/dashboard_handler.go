package handler

import (
	"dashboard-service/internal/domain/dashboard"
	"github.com/labstack/echo/v4"
	"net/http"
)

// DashboardHandler handles HTTP requests for todos
type DashboardHandler struct {
	dashboardService dashboard.DashboardService
}

func NewDashboardHandler(dashboardService dashboard.DashboardService) *DashboardHandler {
	return &DashboardHandler{
		dashboardService: dashboardService,
	}
}

// RegisterRoutes registers
func (h *DashboardHandler) RegisterRoutes(e *echo.Echo) {
	dashboardGroup := e.Group("/api/dashboard")
	dashboardGroup.POST("", h.fetchDashboardMetrics)

}

func (h *DashboardHandler) fetchDashboardMetrics(c echo.Context) error {

	var filters []string
	results, err := h.dashboardService.FetchMetrics(c.Request(), filters)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
