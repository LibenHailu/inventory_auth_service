package main

import (
	userservice "github.com/LibenHailu/inventory_auth/internal/core/service/user_service"
	userhdl "github.com/LibenHailu/inventory_auth/internal/handlers/user_hdl"
	"github.com/LibenHailu/inventory_auth/internal/repositories/user"
	"github.com/LibenHailu/inventory_auth/pkg/postgres"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
)

var (
	conn *pgx.Conn = postgres.DbConn()
)

func main() {
	userSrv := userservice.NewUserService(user.NewUserPgxRepo(conn))
	userHandler := userhdl.NewHTTPHandler(userSrv)

	router := gin.New()
	router.POST("/user", userHandler.StoreUser)

	router.Run(":8080")
}
