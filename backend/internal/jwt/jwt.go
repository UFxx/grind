package jwt

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWT struct {
	secret string
	exp    int64
}

func New(secret string, exp int64) *JWT {

	return &JWT{
		secret: secret,
		exp:    exp,
	}
}

func (j JWT) CreateToken(userID uuid.UUID) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &CustomTokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(j.exp) * time.Second)),
		},
		userID,
	})

	return token.SignedString([]byte(j.secret))
}

func (j JWT) GinJWTAuthMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, ErrorResp{
				Errors: Errors{"authorization header is required"},
			})
			ctx.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.ParseWithClaims(
			tokenString,
			&CustomTokenClaims{},
			func(token *jwt.Token) (interface{}, error) {

				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}

				return []byte(j.secret), nil
			},
		)

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, ErrorResp{
				Errors: Errors{"invalid or expired token"},
			})
			ctx.Abort()
			return
		}

		ctx.Set("user_id", token.Claims.(*CustomTokenClaims).UserID)
		ctx.Next()
	}
}
