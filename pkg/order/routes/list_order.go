package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/order/pb"
	"github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/utils"
	"github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/utils/extractors"
	"io"
	"net/http"
)

func ListOrder(ctx *gin.Context, client pb.OrderServiceClient) {
	userId, _ := ctx.Get("userId")
	req := pb.ListOrderRequest{Page: 1, PageSize: 10, UserId: userId.(int64)}
	stream, err := client.ListOrder(context.Background(), &req)
	if err != nil {
		utils.APIErrorResponse(ctx, http.StatusBadGateway, "GET", err.Error())
		return
	}
	var orders []*pb.ListOrderResponse
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
		orders = append(orders, res)
	}
	utils.APIResponse(ctx, "Successfuly list orders.", 200, "GET", extractors.GetListOrderEx(orders))
}
