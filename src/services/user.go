package services

import (
	"context"
	"google.golang.org/grpc"
	"spider/proto/user"
)

type userServce struct {


}

func (pd *userServce) GetUserList(ctx context.Context, in *user.GetUserListRequest) (out *user.GetUserCourseResponse, err error) {
	response := &user.GetUserCourseResponse{}
	response.IsSuccess = true
	return response,nil
}

func RegisterUserServiceServer(s *grpc.Server) {
	user.RegisterUserServiceServer(s, &userServce{})
}