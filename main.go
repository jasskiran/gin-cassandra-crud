package main

import (
	"gin-cassandra-crud/controllers"
	"gin-cassandra-crud/repository"
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"github.com/sirupsen/logrus"
	"os"
)

func main(){
	lvl, _ := logrus.ParseLevel("trace")
	logger := newLogger(lvl)

	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "test"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil{
		logger.Errorf("cannot create session %s", err)
	}
	defer session.Close()

	router := gin.Default()

	userRepository := repository.NewUserRepository()
	userController := controllers.NewUserController(userRepository, logger, session)

	router.POST("/create", userController.CreateUser)
	router.PUT("/update", userController.EditUser)
	router.GET("/user", userController.GetUser)
	router.DELETE("/delete", userController.DeleteUser)

	router.Run(":8000")
}

func newLogger(level logrus.Level) *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetReportCaller(true)
	logger.Level = level
	logger.Out = os.Stdout
	return logger
}
