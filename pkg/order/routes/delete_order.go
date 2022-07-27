package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/order/pb"
	"github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/utils"
	"net/http"
	"strconv"
)

func DeleteOrder(ctx *gin.Context, client pb.OrderServiceClient) {
	orderIdStr := ctx.Param("id")
	orderId, _ := strconv.Atoi(orderIdStr)

	userId, _ := ctx.Get("userId")
	req := pb.DeleteOrderRequest{OrderId: int64(orderId), UserId: userId.(int64)}
	_, err := client.DeleteOrder(context.Background(), &req)
	if err != nil {
		utils.APIErrorResponse(ctx, http.StatusBadRequest, "DELETE", err.Error())
		return
	}
	utils.APIResponse(ctx, "Successfuly delete order.", http.StatusOK, "DELETE", nil)
}
