package domain

type BrewdudeApiResponse struct {
	Success bool        `json:"success"`
	Message *string     `json:"message,omitempty"`
	Errors  []string    `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
