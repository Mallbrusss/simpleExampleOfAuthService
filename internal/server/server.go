package server

import (
	"auth/internal/authentication"
	changepassword "auth/internal/authentication/changePassword"
	loginuser "auth/internal/authentication/loginUser"
	registeruser "auth/internal/authentication/registerUser"
	resetpassword "auth/internal/authentication/resetPassword"
	"net"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	registeruser.RpcRegist
	loginuser.RpcLogin
	changepassword.RpcChangePwd
	resetpassword.RpcResetPwd
}

func StartGRPCServer() error {
	grpcServer := grpc.NewServer()
	authService := &GRPCServer{}
	authentication.RegisterAuthenticationServiceServer(grpcServer, authService)

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}

	if err := grpcServer.Serve(lis); err != nil {
		return err
	}

	return nil
}
