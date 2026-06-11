package dto

import "github.com/OnePulseOmni/pulse-wallet/internal/constant/model/db"

// CreateSampleRequest represents the request body for creating a sample
type CreateSampleRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Age      int32  `json:"age,omitempty"`
	IsActive bool   `json:"is_active,omitempty"`
}

// CreateSampleResponse represents the response for creating a sample
type CreateSampleResponse struct {
	Message string     `json:"message"`
	Sample  *db.Sample `json:"sample"`
}
