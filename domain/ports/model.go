package ports

type Port struct {
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Coordinates []float64 `json:"coordinates"`
	Province    string    `json:"province"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}

type UpdatePort struct {
	Name        *string    `json:"name"`
	City        *string    `json:"city"`
	Country     *string    `json:"country"`
	Alias       *[]string  `json:"alias"`
	Regions     *[]string  `json:"regions"`
	Coordinates *[]float64 `json:"coordinates"`
	Province    *string    `json:"province"`
	Timezone    *string    `json:"timezone"`
	Unlocs      *[]string  `json:"unlocs"`
	Code        *string    `json:"code"`
}

func StringToPointerString(s string) *string {
	return &s
}
