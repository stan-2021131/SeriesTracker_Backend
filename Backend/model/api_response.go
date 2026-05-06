package model

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Error estándar
type ErrorResponse struct {
	Message string `json:"message"`
}

// Rating request (entrada)
type RatingRequest struct {
	Puntaje int `json:"puntaje" example:"5"`
}
