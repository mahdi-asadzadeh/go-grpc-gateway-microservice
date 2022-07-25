package product

import (
	"fmt"
	"github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/product/pb"
	"google.golang.org/grpc"
	"os"
)

func InitServiceClient() pb.ProductServiceClient {
	fmt.Println("Init product grpc service ...")
	client, err := grpc.Dial(os.Getenv("GRPC_PRODUCT_URL"), grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	return pb.NewProductServiceClient(client)
}
