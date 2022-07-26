package extractors

import "github.com/mahdi-asadzadeh/go-grpc-gateway-microservice/pkg/auth/pb"

func GetUserEx(user *pb.RegisterResponse) map[string]interface{} {
	result := map[string]interface{}{
		"id":        user.GetId(),
		"email":     user.GetEmail(),
		"create_at": user.GetCreateAt(),
		"update_at": user.GetUpdateAt(),
	}
	return result
}
