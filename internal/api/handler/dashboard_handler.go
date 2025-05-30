package handler

import (
	"dashboard-service/internal/domain/dashboard"
	"github.com/labstack/echo/v4"
	"net/http"
)

// DashboardHandler handles HTTP requests for todos
type DashboardHandler struct {
	dashboardService dashboard.ReservationService
}

func NewDashboardHandler(dashboardService dashboard.ReservationService) *DashboardHandler {
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

	results := []string{"hello"}
	return c.JSON(http.StatusOK, results)
}
