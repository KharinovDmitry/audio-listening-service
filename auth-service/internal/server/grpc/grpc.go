package grpc

import (
	"context"
	_ "protos/gen/go/auth"
)

type AuthService interface {
	Login(ctx context.Context, login string, password string) (token string, err error)
	RegisterNewListener(ctx context.Context, login string, password string) (userID int64, err error)
	RegisterNewArtist(ctx context.Context, login string, password string, name string) (userID int64, err error)
}

type serverAPI struct {
	ssov1.UnimplementedAuthServer // Хитрая штука, о ней ниже
	auth                          Auth
}

func Register(gRPCServer *grpc.Server, authService AuthService) {
	ssov1.RegisterAuthServer(gRPCServer, &serverAPI{auth: auth})
}

func (s *serverAPI) Login(
	ctx context.Context,
	in *ssov1.LoginRequest,
) (*ssov1.LoginResponse, error) {
	// TODO
}

func (s *serverAPI) Register(
	ctx context.Context,
	in *ssov1.RegisterRequest,
) (*ssov1.RegisterResponse, error) {
	// TODO
}
