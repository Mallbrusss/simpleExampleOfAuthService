package changepassword

import (
	pb "auth/internal/authentication"
	"context"
)

type RpcChangePwd struct {
}

func (*RpcChangePwd) ChangePassword(context.Context, *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
	panic("unimplemented")
}
