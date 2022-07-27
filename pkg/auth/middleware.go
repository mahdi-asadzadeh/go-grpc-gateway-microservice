package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/auth/pb"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	Client pb.AuthServiceClient
}

func (auth *AuthMiddleware) LoginRequired(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("authorization")
	if authorization == "" {
		ctx.JSON(http.StatusUnauthorized, "You must be authenticated")
		ctx.Abort()
		return
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token) < 2 {
		ctx.JSON(http.StatusUnauthorized, "You must be authenticated")
		ctx.Abort()
		return
	}
	res, err := auth.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "You must be authenticated")
		ctx.Abort()
		return
	}

	ctx.Set("userId", res.GetId())
	ctx.Next()
}
