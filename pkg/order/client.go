package order

import (
	"fmt"
	"github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/order/pb"
	"google.golang.org/grpc"
	"os"
)

func InitServiceClient() pb.OrderServiceClient {
	fmt.Println("Init order grpc service ...")
	client, err := grpc.Dial(os.Getenv("GRPC_ORDER_URL"), grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	return pb.NewOrderServiceClient(client)
}
