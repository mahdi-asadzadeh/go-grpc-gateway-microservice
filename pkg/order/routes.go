package order

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/auth"
	"github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/order/pb"
	"github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/order/routes"
)

type OrderService struct {
	client pb.OrderServiceClient
}

func InitRegisterRoutes(router *gin.RouterGroup) {
	ps := OrderService{client: InitServiceClient()}
	auth := auth.AuthMiddleware{Client: auth.InitServiceClient()}
	router.Use(auth.LoginRequired)
	{
		router.POST("/create", ps.CreateOrder)
		router.DELETE("/delete/:id", ps.DeleteOrder)
		router.GET("/list", ps.ListOrder)
	}

}

func (s *OrderService) CreateOrder(ctx *gin.Context) {
	routes.CreateOrder(ctx, s.client)
}

func (s *OrderService) DeleteOrder(ctx *gin.Context) {
	routes.DeleteOrder(ctx, s.client)
}

func (s *OrderService) ListOrder(ctx *gin.Context) {
	routes.ListOrder(ctx, s.client)
}
