package auth

import (
	"fmt"
	"github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/auth/pb"
	"google.golang.org/grpc"
	"os"
)

func InitServiceClient() pb.AuthServiceClient {
	fmt.Println("Init auth grpc service ...")
	client, err := grpc.Dial(os.Getenv("GRPC_AUTH_URL"), grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	return pb.NewAuthServiceClient(client)
}
