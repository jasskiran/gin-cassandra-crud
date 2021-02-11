package controllers

import (
	"fmt"
	"gin-cassandra-crud/models"
	"gin-cassandra-crud/repository"
	"github.com/gocql/gocql"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

// user wrapper
type UserController struct {
	userRepo repository.UserRepository
	Logger   *logrus.Logger
	Service  *gocql.Session
}

func NewUserController(userRepo repository.UserRepository, Logger *logrus.Logger, Service *gocql.Session) *UserController {
	return &UserController{
		userRepo: userRepo,
		Logger:   Logger,
		Service: Service,
	}
}

func (controller *UserController) CreateUser(c *gin.Context) {

	var user models.User

	// App level validation
	err := c.BindJSON(&user)
	//err := c.BindJSON(&user)
	if err != nil {
		controller.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, fmt.Sprint("bindErr"))
		return
	}

	fmt.Println("request user", user)

	// validating details for creating new user
	use, err := models.NewUser(controller.Logger, user.Name, user.Email, user.Phone)
	if err != nil {
		controller.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, fmt.Sprintf("not validated"))
		return
	}
	//var session struct {
	//	Session *gocql.Session
	//}
	//if err := session.Session.Query(`INSERT INTO test.user(id, name, email, phone) VALUES (?, ?, ?, ?)`,
	//	use.Id, use.Name, use.Email, use.Phone).Exec(); err != nil {
	//	controller.Logger.Error(err.Error())
	//	c.JSON(http.StatusInternalServerError, fmt.Sprintf("Something wrong on our server"))
	//	return
	//}
	// Inserting data
	err = controller.userRepo.Create(controller.Service, use)
	if err != nil {
		controller.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Something wrong on our server"))
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (controller *UserController) GetUser(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")

	fmt.Println("req id: ", id)
	Id, err := strconv.Atoi(id)
	if err != nil {
		controller.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Something wrong on our server"))
		return
	}
	fmt.Println("iddd", Id)
	// validate the id
	//if Id == 0 {
	//	c.JSON(http.StatusNotFound, "Not found")
	//}
	user, err := repository.UserRepository.GetById(controller.userRepo, controller.Service, Id)
	if err != nil {
		controller.Logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Something wrong on our server"))
		return
	}
	c.JSON(http.StatusOK, user)
}

func (controller *UserController) EditUser(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")

	Id, err := strconv.Atoi(id)
	if err != nil {
		controller.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Something wrong on our server"))
		return
	}
	var user models.User

	// App level validation
	bindErr := c.BindJSON(&user)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(bindErr))
	}

	// Check if resource exist
	u, _ := controller.userRepo.GetById(controller.Service, Id)
	fmt.Println("u.Id", u.Id)
	//if u.Id == 0 {
	//	c.JSON(http.StatusNotFound, "Not found")
	//	return
	//}

	// Updating data
	updatedUser := controller.userRepo.EditUser(controller.Service, Id, &user)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, updatedUser)
	}
}

func (controller *UserController) DeleteUser(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")

	Id, err := strconv.Atoi(id)
	if err != nil {
		controller.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Something wrong on our server"))
		return
	}

	// Check if resource exist
	user, _ := controller.userRepo.GetById(controller.Service, Id)
	fmt.Println("user.Id", user.Id)
	//if user.Id == 0 {
	//	c.JSON(http.StatusNotFound, "Not found")
	//	return
	//}

	// Deleting data
	err = controller.userRepo.DeleteUser(controller.Service, Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusNoContent, "Successful Deletion")

}
