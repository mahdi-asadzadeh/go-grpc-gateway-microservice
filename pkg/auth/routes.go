package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/auth/pb"
	"github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/auth/routes"
)

type AuthService struct {
	Client pb.AuthServiceClient
}

func InitRegisterRoutes(router *gin.RouterGroup) {
	ps := AuthService{Client: InitServiceClient()}
	router.POST("/register", ps.Register)
	router.POST("/login", ps.Login)
}

func (s *AuthService) Register(ctx *gin.Context) {
	routes.Register(ctx, s.Client)
}

func (s *AuthService) Login(ctx *gin.Context) {
	routes.Login(ctx, s.Client)
}
