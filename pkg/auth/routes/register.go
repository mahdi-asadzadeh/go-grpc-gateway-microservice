package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/auth/pb"
	"github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/utils"
	"github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/utils/extractors"
	"net/http"
)

type RegisterInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(ctx *gin.Context, client pb.AuthServiceClient) {
	input := RegisterInput{}
	if err := ctx.BindJSON(&input); err != nil {
		utils.APIErrorResponse(ctx, http.StatusBadRequest, "POST", err.Error())
		return
	}
	req := pb.RegisterRequest{User: &pb.User{Email: input.Email, Password: input.Password}}
	res, err := client.Register(context.Background(), &req)
	if err != nil {
		utils.APIErrorResponse(ctx, http.StatusBadGateway, "POST", err.Error())
		return
	}
	utils.APIResponse(ctx, "Successfuly register user.", http.StatusCreated, "POST", extractors.GetUserEx(res))
}
