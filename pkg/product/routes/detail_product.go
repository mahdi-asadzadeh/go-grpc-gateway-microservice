package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/product/pb"
	"github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/utils"
	"github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/utils/extractors"
	"strconv"

	"net/http"
)

func DetailProduct(ctx *gin.Context, client pb.ProductServiceClient) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	req := pb.DetailProductRequest{Id: int64(id)}
	res, err := client.DetailProduct(context.Background(), &req)
	if err != nil {
		utils.APIErrorResponse(ctx, http.StatusNotFound, "GET", err.Error())
		return
	}
	utils.APIResponse(ctx, "Successfuly get product.", http.StatusOK, "GET", extractors.GetProductEx(res.GetProduct()))
}
