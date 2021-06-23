package postgres

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx"
)

// Creating database connection
func DbConn() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:admin@localhost:5432/inventory_auth")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	log.Println("database connected successfully")
	// defer conn.Close(context.Background())

	return conn

}
