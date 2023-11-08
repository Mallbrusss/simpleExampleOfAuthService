package resetpassword

import (
	pb "auth/internal/authentication"
	"context"
)

type RpcResetPwd struct{}

func (*RpcResetPwd) ResetPassword(context.Context, *pb.ResetPasswordRequest) (*pb.ResetPasswordResponse, error) {
	panic("unimplemented")
}
