package registeruser

import (
	pb "auth/internal/authentication"
	"auth/internal/entities"
	storages "auth/internal/storages/postgres"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	ctx.Deadline()
	userRepository := NewRepository(storages.DB)
	usernameExist, err := userRepository.isUserExists(req.Username)
	if err != nil {
		return nil, err
	}

	if usernameExist {
		return nil, errors.New("username already exist")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &entities.User{
		Username: string(req.Username),
		Password: string(hashedPassword),
		Email:    string(req.Email),
		FullName: string(req.FullName),
	}
	if err := userRepository.saveUser(*user); err != nil {
		return nil, err
	}

	return &pb.RegisterUserResponse{
		Success: true,
		Message: "Success registrate",
	}, nil
}
