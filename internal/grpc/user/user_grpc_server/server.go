// package usergrpcserver
package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/LibenHailu/inventory_auth/internal/core/domain"
	"github.com/LibenHailu/inventory_auth/internal/core/ports"
	userservice "github.com/LibenHailu/inventory_auth/internal/core/service/user_service"
	"github.com/LibenHailu/inventory_auth/internal/grpc/user/userpb"
	"github.com/LibenHailu/inventory_auth/internal/repositories/user"
	"github.com/LibenHailu/inventory_auth/pkg/postgres"
	"github.com/jackc/pgx"
	"google.golang.org/grpc"
)

var (
	conn *pgx.Conn = postgres.DbConn()
)

type GrpcServer struct {
	userService ports.UserService
	userpb.UnimplementedUserServiceServer
}

func (gs *GrpcServer) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {

	fmt.Println("this is from the server")
	data := &domain.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
		RoleID:    uint(req.RoleId),
	}

	// user, err := gs.userRepo.StoreUser(ctx, data)
	user, err := gs.userService.StoreUser(ctx, data)

	if err != nil {
		return nil, err
	}

	return &userpb.CreateUserResponse{

		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		RoleId:    int64(user.RoleID),
	}, nil

}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v ", err)
	}

	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, &GrpcServer{
		// userRepo: user.NewUserPgxRepo(conn),
		userService: userservice.NewUserService(user.NewUserPgxRepo(conn)),
	})
	// ItemUsecase: itemUsecase,
	// })

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	defer conn.Close(context.Background())
}
