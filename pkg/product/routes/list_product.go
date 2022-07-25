package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/product/pb"
	"github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/utils"
	"github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/utils/extractors"
	"io"
	"net/http"
)

func ListProduct(ctx *gin.Context, client pb.ProductServiceClient) {
	req := pb.ListProductRequest{Page: 1, PageSize: 20}
	stream, err := client.ListProduct(context.Background(), &req)
	if err != nil {
		utils.APIErrorResponse(ctx, http.StatusBadGateway, "GET", err.Error())
		return
	}
	var products []*pb.Product
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			utils.APIErrorResponse(ctx, http.StatusBadGateway, "GET", err.Error())
			return
		}
		products = append(products, msg.GetProduct())
	}
	utils.APIResponse(ctx, "Successfuly list all product.", http.StatusOK, "GET", extractors.GetListProductEx(products))
}
