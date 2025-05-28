package dashboard

import "net/http"

type Metrics struct {
	MonthlyRevenue map[string]RevenueOffice
}

type RevenueOffice struct {
	MonthlyRevenue   float64
	UnreservedOffice int32
}
type DashboardService interface {
	FetchMetrics(request *http.Request, filters []string) (*Metrics, error)
}
