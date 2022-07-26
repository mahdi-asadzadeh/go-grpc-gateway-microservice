package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/auth"
	"github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/config"
	"github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/order"
	"github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/product"
)

func main() {
	fmt.Println("Starting api gateway server ...")
	fmt.Println("Loading config envs ...")
	config.LoadSettings(true)
	router := gin.Default()

	apiRouteGroup := router.Group("/api")
	product.InitRegisterRoutes(apiRouteGroup.Group("/product"))
	auth.InitRegisterRoutes(apiRouteGroup.Group("/auth"))
	order.InitRegisterRoutes(apiRouteGroup.Group("/order"))

	router.Run(":8080")
}
