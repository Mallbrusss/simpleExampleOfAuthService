package loginuser

import (
	pb "auth/internal/authentication"
	"context"
)

type RpcLogin struct{}

func (*RpcLogin) LoginUser(context.Context, *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	panic("unimplemented")
}
