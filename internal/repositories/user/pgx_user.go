package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/LibenHailu/inventory_auth/internal/core/domain"
	"github.com/LibenHailu/inventory_auth/internal/core/ports"
	"github.com/jackc/pgx"
)

const (
	// insert a new user to an item table
	QueryInsertUser = "INSERT INTO users ( first_name,last_name,email,password,role_id) VALUES ($1,$2,$3,$4,$5)"
)

//implements UserReposotory interface
type UserPgxRepo struct {
	conn *pgx.Conn
}

// NewUserPgxRepo creates a new object of UserPgxRepo
func NewUserPgxRepo(db *pgx.Conn) ports.UserRepository {
	return &UserPgxRepo{conn: db}
}

// // returns all user from the database
// func (userRepo *UserPgxRepo)Users() ([]domain.User, []error){

// }
// func (userRepo *UserPgxRepo)User(id uint) (*domain.User, []error){

// }

//StoreUser stores user to the database
func (userRepo *UserPgxRepo) StoreUser(ctx context.Context, user *domain.User) (*domain.User, error) {

	fmt.Println("this is in store user")

	newUsr, err := userRepo.conn.Exec(ctx, QueryInsertUser, user.FirstName, user.LastName, user.Email, user.Password, user.RoleID)

	if err != nil {
		return nil, err
	}

	if newUsr.RowsAffected() > 0 {

		return user, nil
	}
	return nil, errors.New("Some thing went wrong")

}

// func (userRepo *UserPgxRepo)UserByEmail(email string) (*domain.User, []error){

// }
// func (userRepo *UserPgxRepo)(user *domain.User) (*domain.User, []error) {}
// func (userRepo *UserPgxRepo)DeleteUser(id uint) (*domain.User, []error){}
