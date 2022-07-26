package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/auth/pb"
	"github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/utils"
	"net/http"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context, client pb.AuthServiceClient) {
	input := LoginInput{}
	if err := ctx.BindJSON(&input); err != nil {
		utils.APIErrorResponse(ctx, http.StatusBadRequest, "POST", err.Error())
		return
	}
	req := pb.LoginRequest{User: &pb.User{Email: input.Email, Password: input.Password}}
	res, err := client.Login(context.Background(), &req)
	if err != nil {
		utils.APIErrorResponse(ctx, http.StatusBadGateway, "POST", err.Error())
		return
	}
	result := map[string]interface{}{"token": res.GetToken(), "email": input.Email}
	utils.APIResponse(ctx, "Successfuly login user.", http.StatusOK, "POST", result)
}
