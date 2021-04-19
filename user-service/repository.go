package main

import (
	"fmt"

	pb "github.com/charles-hashdak/cleartoo-services/user-service/proto/user"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Repository interface {
	GetAll() ([]*pb.User, error)
	Get(id string) (*pb.User, error)
	Create(user *pb.User) error
	Edit(user *pb.User) error
	ChangePassword(user *pb.User) error
	Follow(req *pb.Follower) error
	IsFollowing(req *pb.Follower) (bool, error)
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
	fmt.Println(email)
	user := &pb.User{}
	if err := repo.db.Where("email = ?", email).
		First(&user).Error; err != nil {
		return nil, err
	}
	fmt.Println(user)
	return user, nil
}

func (repo *UserRepository) Create(user *pb.User) error {
	fmt.Println(user)
	user.Id = uuid.NewV4().String()
	// var usernameId = 1000 + rand.Intn(9999-1000)
	// user.Username = strings.Split(user.Email, string('@'))[0] + strconv.Itoa(usernameId)
	user.Rating = 0;
	user.FollowersCount = 0;
	user.FollowingCount = 0;
	if err := repo.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) Edit(user *pb.User) error {
	fmt.Println(user)
	if err := repo.db.Model(&user).Updates(map[string]interface{}{"username": user.Username, "description": user.Description, "avatar_url": user.AvatarUrl, "cover_url": user.CoverUrl, "age": user.Age}).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) ChangePassword(user *pb.User) error {
	if err := repo.db.Model(&user).Updates(map[string]interface{}{"password": user.Password}).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) Follow(follower *pb.Follower) error {
	isFollowing, isFollowingErr := repo.IsFollowing(follower)
	if isFollowingErr != nil {
		return isFollowingErr
	}
	user, getErr := repo.Get(follower.UserId)
	if getErr != nil {
		return getErr
	}
	followerUser, getFErr := repo.Get(follower.FollowerId)
	if getFErr != nil {
		return getFErr
	}
	if(isFollowing){
		if err := repo.db.Delete(follower).Error; err != nil {
			return err
		}
		user.FollowersCount = user.FollowersCount - 1
		followerUser.FollowingCount = followerUser.FollowingCount - 1
	}else{
		follower.Id = uuid.NewV4().String()
		if err := repo.db.Create(follower).Error; err != nil {
			return err
		}
		user.FollowersCount = user.FollowersCount + 1
		followerUser.FollowingCount = followerUser.FollowingCount + 1
	}
	if editErr := repo.db.Model(&user).Updates(map[string]interface{}{"followers_count": user.FollowersCount}).Error; editErr != nil {
		return editErr
	}
	if editFErr := repo.db.Model(&followerUser).Updates(map[string]interface{}{"following_count": followerUser.FollowingCount}).Error; editFErr != nil {
		return editFErr
	}
	return nil
}

func (repo *UserRepository) IsFollowing(follower *pb.Follower) (bool, error) {
	var count int64
	repo.db.Table("followers").Where("follower_id = ? AND user_id = ?", follower.FollowerId, follower.UserId).Count(&count)
	if(count > 0){
		return true, nil
	}else{
		return false, nil
	}
}