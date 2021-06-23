package userhdl

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/LibenHailu/inventory_auth/internal/core/domain"
	"github.com/LibenHailu/inventory_auth/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type UserHTTPHandler struct {
	userService ports.UserService
}

func NewHTTPHandler(userService ports.UserService) *UserHTTPHandler {
	return &UserHTTPHandler{
		userService: userService,
	}
}

//StoreUser accespts request as a primary adapter
func (us *UserHTTPHandler) StoreUser(c *gin.Context) {

	userData := domain.User{}

	err := c.ShouldBindJSON(&userData)

	// body, _ := ioutil.ReadAll(c.Request.Body)

	// fmt.Println("fas")
	// fmt.Println("asdf")
	// fmt.Println("asdf")
	// jsonData, _ := c.GetRawData()
	// fmt.Println(jsonData)

	// s := string(jsonData)

	// fmt.Println(s)
	// if err := json.Unmarshal(jsonData, &userData); err != nil {
	// 	panic(err)
	// }
	// fmt.Println(userData.RoleID)
	// fmt.Println(userData.FirstName)

	// c.JSON(http.StatusOK, gin.H{
	// 	"error": "no",
	// })

	if err != nil {
		log.Printf("err while processing the request %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("while processing req %v", err),
		})
		return
	}
	user, err := us.userService.StoreUser(context.Background(), &userData)

	if err != nil {
		log.Printf("err while processing the request %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("while processing req %v", err),
		})
		return
	}
	c.JSON(
		http.StatusOK, gin.H{
			"user": user,
		},
	)

}
