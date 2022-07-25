package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/product/pb"
	"github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/utils"
	"github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/utils/extractors"
	"net/http"
)

type CreateProductInput struct {
	Slug  string  `json:"slug"`
	Title string  `json:"title"`
	Body  string  `json:"body"`
	Price float32 `json:"price"`
}

func CreateProduct(ctx *gin.Context, client pb.ProductServiceClient) {
	input := CreateProductInput{}
	if err := ctx.BindJSON(&input); err != nil {
		utils.APIErrorResponse(ctx, http.StatusBadRequest, "POST", err.Error())
		return
	}
	req := pb.CreateProductRequest{
		Slug:  input.Slug,
		Title: input.Title,
		Body:  input.Body,
		Price: input.Price,
	}
	res, err := client.CreateProduct(context.Background(), &req)
	if err != nil {
		utils.APIErrorResponse(ctx, http.StatusBadRequest, "POST", err.Error())
		return
	}
	utils.APIResponse(ctx, "Successfuly create product.", http.StatusCreated, "POST", extractors.GetProductEx(res.GetProduct()))
}
