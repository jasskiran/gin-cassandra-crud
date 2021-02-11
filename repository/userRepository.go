package repository

import (
	"fmt"
	"gin-cassandra-crud/models"
	"github.com/gocql/gocql"
	"log"
)

type UserRepository interface {
	Create(service *gocql.Session, out *models.User) error
	GetById(service *gocql.Session, id int) (*models.User, error)
	EditUser(service *gocql.Session, id int, user *models.User) error
	DeleteUser(service *gocql.Session, id int) error
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (u userRepository) Create(service *gocql.Session, user *models.User) error {

	fmt.Println("user.Id ", user.Id, &user.Id)
	if err := service.Query(`INSERT INTO test.user(id, name, email, phone) VALUES (?, ?, ?, ?)`,
		user.Id, user.Name, user.Email, user.Phone).Exec(); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (u userRepository) GetById(service *gocql.Session, id int) (*models.User, error) {
	var user models.User

	if err := service.Query(`SELECT name, email, phone FROM user WHERE id = ?`,
		id).Scan(&user.Name, &user.Email, &user.Phone); err != nil {
		fmt.Errorf("internal server err %s", err)
		return nil, err
	}

	return &user, nil
}

func (u userRepository)EditUser(service *gocql.Session, id int, user *models.User) error{

	if err := service.Query("UPDATE user SET name = ?, email = ?, phone = ? WHERE id = ?",
		&user.Name, &user.Email, &user.Phone, id).Exec(); err != nil {
		fmt.Println("Error while updating")
		fmt.Println(err)
		return err
	}
	return nil
}

func (u userRepository)DeleteUser(service *gocql.Session, id int) error {

	if err := service.Query("DELETE FROM test.user WHERE id = ?", id).Exec(); err != nil {
		fmt.Println("Error while deleting")
		fmt.Println(err)
		return err
	}

	return nil
}
