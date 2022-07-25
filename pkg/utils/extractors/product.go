package extractors

import (
	"github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/product/pb"
)

func GetProductEx(product *pb.Product) map[string]interface{} {
	result := map[string]interface{}{
		"id":        product.GetId(),
		"slug":      product.GetSlug(),
		"title":     product.GetTitle(),
		"body":      product.GetBody(),
		"price":     product.GetPrice(),
		"create_at": product.GetCreateAt(),
		"update_at": product.GetUpdateAt(),
	}
	return result
}

func GetListProductEx(products []*pb.Product) interface{} {
	result := make([]interface{}, len(products))
	for i := 0; i < len(products); i++ {
		result[i] = GetProductEx(products[i])
	}
	return result
}
