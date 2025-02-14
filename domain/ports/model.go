package ports

// Port represents the struct format of a port from the json file
// my assumption was that those fields are present in any port item in the json file
// but I didn't check all of them to actually get the unified format
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

type NewPort struct {
	ID string `json:"id"`
	Port
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
