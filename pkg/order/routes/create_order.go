package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/order/pb"
	"github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/utils"
	"github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/utils/extractors"
	"net/http"
)

type CreateOrderInput struct {
	Price     float64 `json:"price"`
	ProductId int64   `json:"productId"`
	Quantity  int64   `json:"quantity"`
}

func CreateOrder(ctx *gin.Context, client pb.OrderServiceClient) {
	input := CreateOrderInput{}
	if err := ctx.BindJSON(&input); err != nil {
		utils.APIErrorResponse(ctx, http.StatusBadRequest, "POST", err.Error())
		return
	}
	userId, _ := ctx.Get("userId")
	req := pb.CreateOrderRequest{
		Price:     float32(input.Price),
		ProductId: input.ProductId,
		UserId:    userId.(int64),
		Quantity:  int32(input.Quantity),
	}
	res, err := client.CreateOrder(context.Background(), &req)
	if err != nil {
		utils.APIErrorResponse(ctx, http.StatusBadRequest, "POST", err.Error())
		return
	}
	utils.APIResponse(ctx, "Successfuly create order.", http.StatusCreated, "POST", extractors.GetOrderEx(res.GetOrder()))
}
