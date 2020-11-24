package main

import (
	"strings"
	"strconv"
	"math/rand"

	pb "github.com/charles-hashdak/cleartoo-services/user-service/proto/user"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Repository interface {
	GetAll() ([]*pb.User, error)
	Get(id string) (*pb.User, error)
	Create(user *pb.User) error
	GetByEmail(email string) (*pb.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func (repo *UserRepository) GetAll() ([]*pb.User, error) {
	var users []*pb.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) Get(id string) (*pb.User, error) {
	user := &pb.User{}
	user.Id = id
	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetByEmail(email string) (*pb.User, error) {
	user := &pb.User{}
	if err := repo.db.Where("email = ?", email).
		First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) Create(user *pb.User) error {
	user.Id = uuid.NewV4().String()
	var usernameId = 1000 + rand.Intn(9999-1000)
	user.Username = strings.Split(user.Email, string('@'))[0] + strconv.Itoa(usernameId)
	user.Rating = 0;
	user.FollowersCount = 0;
	user.FollowingCount = 0;
	if err := repo.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}