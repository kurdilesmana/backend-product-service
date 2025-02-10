package healthCheckModel

type HealthCheck struct {
	AppDescription      AppDescription      `json:"app"`
	DatabaseDescription DatabaseDescription `json:"database"`
	Status              string              `json:"status"`
}

type AppDescription struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type DatabaseDescription struct {
	Engine string `json:"engine"`
	Status string `json:"status"`
}
