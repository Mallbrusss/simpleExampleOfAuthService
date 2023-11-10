package loginuser

import (
	pb "auth/internal/authentication"
	storages "auth/internal/storages/postgres"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type RpcLogin struct{}

func (*RpcLogin) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	ctx.Deadline()
	userRepository := NewRepository(storages.DB)
	isUsernameExist, err := userRepository.isUserNotExist(req.Username)
	if err != nil {
		return nil, err
	}
	if !isUsernameExist {
		return nil, errors.New("user with this username not exists")
	}

	hashPwd, err := userRepository.returnHashPassword(req.Username)
	if err != nil {
		return nil, err
	}

	checkPass := bcrypt.CompareHashAndPassword(hashPwd, []byte(req.Password))
	if checkPass != nil {
		switch {
		case errors.Is(checkPass, bcrypt.ErrMismatchedHashAndPassword):
			return nil, errors.New("incorrect password")
		case errors.Is(checkPass, bcrypt.ErrHashTooShort):
			return nil, errors.New("hashed password is too short")
		default:
			return nil, errors.New("error comparing passwords: " + checkPass.Error())
		}
	}
	return &pb.LoginUserResponse{
		Success: true,
		Message: "Success login",
	}, nil
}
