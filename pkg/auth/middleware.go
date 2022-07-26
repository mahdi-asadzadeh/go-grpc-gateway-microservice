package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/auth/pb"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	client pb.AuthServiceClient
}

func (auth *AuthMiddleware) LoginRequired(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("authorization")

	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	res, err := auth.client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Set("userId", res.GetId())
	ctx.Next()
}
