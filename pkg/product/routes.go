package product

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/product/pb"
	"github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/product/routes"
)

type ProductService struct {
	client pb.ProductServiceClient
}

func InitRegisterRoutes(router *gin.RouterGroup) {
	ps := ProductService{client: InitServiceClient()}
	router.POST("/create", ps.CreateProduct)
	router.GET("/detail/:id", ps.DetailProduct)
	router.GET("/list", ps.ListProduct)
}

func (s *ProductService) CreateProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, s.client)
}

func (s *ProductService) DetailProduct(ctx *gin.Context) {
	routes.DetailProduct(ctx, s.client)
}

func (s *ProductService) ListProduct(ctx *gin.Context) {
	routes.ListProduct(ctx, s.client)
}
