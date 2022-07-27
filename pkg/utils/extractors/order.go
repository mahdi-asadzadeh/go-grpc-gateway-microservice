package extractors

import "github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/order/pb"

func GetOrderEx(order *pb.Order) map[string]interface{} {
	result := map[string]interface{}{
		"id":         order.GetId(),
		"price":      order.GetPrice(),
		"product_id": order.GetProductId(),
		"user_id":    order.GetUserId(),
		"quantity":   order.GetQuantity(),
	}
	return result
}

func GetListOrderEx(orders []*pb.ListOrderResponse) []interface{} {
	result := make([]interface{}, len(orders))
	for i := 0; i < len(orders); i++ {
		order := map[string]interface{}{
			"id":         orders[i].GetId(),
			"price":      orders[i].GetPrice(),
			"product_id": orders[i].GetProductId(),
			"user_id":    orders[i].GetUserId(),
			"quantity":   orders[i].GetQuantity(),
			"create_at":  orders[i].GetCreateAt(),
			"update_at":  orders[i].GetUpdateAt(),
		}
		result[i] = order
	}
	return result
}
