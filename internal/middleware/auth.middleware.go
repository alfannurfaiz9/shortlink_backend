package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/alfannurfaiz9/shortlink_backend/internal/dto"
	"github.com/alfannurfaiz9/shortlink_backend/pkg"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

func AuthMiddleware(rdb *redis.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearer := ctx.GetHeader("Authorization")
		if bearer == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Success: false,
				Message: "please login first",
			})
			return
		}

		result := strings.Split(bearer, " ")
		if len(result) != 2 || result[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Success: false,
				Message: "invalid bearer token",
			})
			return
		}

		token := &pkg.JWTClaims{}
		err := token.DecodeToken(result[1])

		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenInvalidIssuer) {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
					Success: false,
					Message: "invalid token",
				})
				return
			}
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
				Success: false,
				Message: "internal server error",
			})
			return
		}

		key := fmt.Sprintf("shortlink:token:%s", result[1])
		_, err = rdb.Get(ctx, key).Result()
		if err == nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Success: false,
				Message: "please login first",
			})
			return
		}

		ctx.Set("id", token.Id)
		ctx.Next()
	}
}
