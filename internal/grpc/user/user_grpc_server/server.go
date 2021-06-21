package usergrpcserver

import (
	"context"

	"github.com/LibenHailu/inventory_auth/internal/core/domain"
	"github.com/LibenHailu/inventory_auth/internal/grpc/user/userpb"
	"github.com/LibenHailu/inventory_auth/internal/repositories/user"
)

type GrpcServer struct {
	userRepo user.UserPgxRepo
}

func (gs *GrpcServer) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {

	newUser := req.GetUser()

	data := &domain.User{
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Email:     newUser.Email,
		Password:  newUser.Password,
		RoleID:    uint(newUser.RoleId),
	}

	user, err := gs.userRepo.StoreUser(ctx, data)

	if err != nil {
		return nil, err[0]
	}

	return &userpb.CreateUserResponse{
		User: &userpb.User{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Password:  user.Password,
			RoleId:    int64(user.RoleID),
		},
	}, nil

}

// func main() {
// 	lis, err := net.Listen("tcp", "0.0.0.0:50051")
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v ", err)
// 	}

// 	s := grpc.NewServer()
// 	userpb.RegisterUserServiceServer(s, &GrpcServer{
// 		userRepo: user.UserPgxRepo{},
// 	})
// 	// ItemUsecase: itemUsecase,
// 	// })

// 	if err := s.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// }
