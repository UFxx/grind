package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type (
	CustomTokenClaims struct {
		jwt.RegisteredClaims
		UserID uuid.UUID
	}

	ErrorResp struct {
		Errors Errors `json:"errors"`
	}

	Errors struct {
		Other string `json:"other"`
	}
)
