package main

import (
	_ "fmt"
	_ "strings"

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
	Rate(req *pb.Rating) error
	IsFollowing(req *pb.Follower) (bool, error)
	GetFollowing(follower_id string) ([]*pb.User, error)
	GetByEmail(email string) (*pb.User, error)
	GetByAppleUserId(apple_user_id string) (*pb.User, error)
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

func (repo *UserRepository) GetByAppleUserId(apple_user_id string) (*pb.User, error) {
	user := &pb.User{}
	if err := repo.db.Where("apple_user_id = ?", apple_user_id).
		First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) Create(user *pb.User) error {
	if err := repo.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) Edit(user *pb.User) error {
	if err := repo.db.Model(&user).Updates(map[string]interface{}{"username": user.Username, "description": user.Description, "avatar_url": user.AvatarUrl, "cover_url": user.CoverUrl, "age": user.Age, "push_token": user.PushToken}).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) ChangePassword(user *pb.User) error {
	if err := repo.db.Model(&user).Where("email = ?", user.Email).Update("password", user.Password).Error; err != nil {
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

func (repo *UserRepository) Rate(rating *pb.Rating) error {
	user, getErr := repo.Get(rating.UserId)
	if getErr != nil {
		return getErr
	}
	rating.Id = uuid.NewV4().String()
	if err := repo.db.Create(rating).Error; err != nil {
		return err
	}
	user.RatingCount = user.RatingCount + 1
	user.Rating = ((user.Rating * (float32(user.RatingCount) - float32(1))) + rating.Rate) / float32(user.RatingCount)
	if editErr := repo.db.Model(&user).Updates(map[string]interface{}{"rating_count": user.RatingCount, "rating": user.Rating}).Error; editErr != nil {
		return editErr
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

func (repo *UserRepository) GetFollowing(follower_id string) ([]*pb.User, error) {
	var users []*pb.User
	repo.db.Table("followers").Select("users.id").Joins("left join users on followers.user_id = users.id").Where("followers.follower_id = ?", follower_id).Scan(&users)
	return users, nil
}