package dashboard

import "net/http"

type DashboardServiceImpl struct {
	repo Repository
}

func NewDashboardService(repo Repository) *DashboardServiceImpl {
	return &DashboardServiceImpl{
		repo: repo,
	}
}

func (s *DashboardServiceImpl) FetchMetrics(request *http.Request, filters []string) (*Metrics, error) {

	data := &Metrics{
		MonthlyRevenue: map[string]RevenueOffice{
			"2025-01": {MonthlyRevenue: 15000.75, UnreservedOffice: 5},
			"2025-02": {MonthlyRevenue: 17200.00, UnreservedOffice: 3},
			"2025-03": {MonthlyRevenue: 19850.25, UnreservedOffice: 2},
		},
	}

	return data, nil
}
