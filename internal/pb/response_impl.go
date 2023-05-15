package pb

import (
	context "context"
	"rest-vs-grpc/internal/constants"
	"rest-vs-grpc/internal/entities"
)

type Server struct {
	UnimplementedServerGrpcServer
}

var ResponsePb = entitiesToResponseGRPC(constants.Response)

func (s *Server) GetUsers(context.Context, *Null) (*Response, error) {
	return ResponsePb, nil
}

func entitiesToResponseGRPC(resp entities.Response) *Response {
	users := make([]*Users, 0, len(resp.Users))

	for _, v := range resp.Users {
		users = append(users, &Users{
			FirstName: v.FirstName,
			LastName:  v.LastName,
			Address:   v.Address,
		})
	}

	return &Response{
		Users: users,
	}
}
