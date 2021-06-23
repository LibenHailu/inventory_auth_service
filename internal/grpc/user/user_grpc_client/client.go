// package usergrpcclient
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/LibenHailu/inventory_auth/internal/grpc/user/userpb"
	"google.golang.org/grpc"
)

var cc *grpc.ClientConn

func connect() *grpc.ClientConn {
	if cc == nil {

		cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

		if err != nil {
			log.Fatalf("could not found connect %v ", err)
		}

		return cc
	}

	return cc

}

func CreateUser(firstName string, lastName string, email string, password string, role_id int64) {
	conn := connect()

	c := userpb.NewUserServiceClient(conn)

	req := &userpb.CreateUserRequest{

		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
		RoleId:    role_id,
	}

	fmt.Println(req)
	res, err := c.CreateUser(context.Background(), req)

	if err != nil {
		log.Printf("error while calling User RPC: %v", err)
	}

	log.Printf("Response form User: %v", res)

	// defer conn.Close()
}

func main() {
	CreateUser("li", "ds", "sds", "sd", 1)
}
