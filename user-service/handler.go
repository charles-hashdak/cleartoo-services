package main

import (
	"fmt"
	"errors"
	"log"
	"context"
	"strings"
	"strconv"
	"math/rand"

	pb "github.com/charles-hashdak/cleartoo-services/user-service/proto/user"
	micro "github.com/micro/go-micro/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"golang.org/x/crypto/bcrypt"
)

type Authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *pb.User) (string, string, error)
	EncodeRefreshToken(user *pb.User) (string, error)
}

type service struct {
	repo   			Repository
	tokenService 	Authable
	Publisher    	micro.Publisher
}

func (srv *service) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	user, err := srv.repo.Get(req.Id)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}

func (srv *service) Edit(ctx context.Context, req *pb.User, res *pb.Response) error {
	err := srv.repo.Edit(req)
	if err != nil {
		return err
	}
	res.User = req
	return nil
}

func (srv *service) Follow(ctx context.Context, follower *pb.Follower, res *pb.FollowResponse) error {
	err := srv.repo.Follow(follower)
	if err != nil {
		return err
	}
	res.Followed = true
	return nil
}

func (srv *service) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	users, err := srv.repo.GetAll()
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}

func (srv *service) Auth(ctx context.Context, req *pb.User, res *pb.Response) error {
    _, ok := metadata.FromIncomingContext(ctx)
    if !ok {
        return status.Errorf(codes.DataLoss, "Failed to get metadata")
    }
	user, err := srv.repo.GetByEmail(req.Email)
	if err != nil {
		return err
	}

	// Compares our given password against the hashed password
	// stored in the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return status.Errorf(codes.Internal, "wrong_password")
	}

	token, refreshToken, err := srv.tokenService.Encode(user)
	if err != nil {
		return err
	}
	header := metadata.Pairs("Set-Cookie", "refresh-jwt="+refreshToken+"; Max-Age=5256000; SameSite=none", "Access-Control-Allow-Methods", "POST")
	if err := grpc.SetHeader(ctx, header); err != nil {
		log.Println(err)
		return status.Errorf(codes.Internal, "unable to send 'Set-Cookie' header")
	}
	if err := grpc.SendHeader(ctx, header); err != nil {
		log.Println(err)
		return status.Errorf(codes.Internal, "unable to send 'Set-Cookie' header")
	}
	user.Password = "";
	res.Token = new(pb.Token)
	res.Token.Token = token
	res.User = user
	return nil
}

func (srv *service) Create(ctx context.Context, req *pb.User, res *pb.Response) error {

	log.Println("Creating user: ", req)

	var usernameId = 1000 + rand.Intn(9999-1000)
	req.Username = strings.Split(req.Email, string('@'))[0] + strconv.Itoa(usernameId)
	req.Rating = 0;
	req.FollowersCount = 0;
	req.FollowingCount = 0;

	// Generates a hashed version of our password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New(fmt.Sprintf("error hashing password: %v", err))
	}

	req.Password = string(hashedPass)
	if err := srv.repo.Create(req); err != nil {
		return errors.New(fmt.Sprintf("error creating user: %v", err))
	}

	token, refreshToken, err := srv.tokenService.Encode(req)
	if err != nil {
		return err
	}
	header := metadata.New(map[string]string{"Set-Cookie": "refresh-jwt="+refreshToken+"; Max-Age=5256000; SameSite=none"})
	if err := grpc.SendHeader(ctx, header); err != nil {
		return status.Errorf(codes.Internal, "unable to send 'Set-Cookie' header")
	}

	req.Password = "";
	res.User = req
	res.Token = &pb.Token{Token: token}

	/*
		if err := srv.Publisher.Publish(ctx, req); err != nil {
			return errors.New(fmt.Sprintf("error publishing event: %v", err))
		}*/

	return nil
}

func (srv *service) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {

	// Decode token
	claims, err := srv.tokenService.Decode(req.Token)

	if err != nil {
		return err
	}

	if claims.User.Id == "" {
		return errors.New("invalid user")
	}

	res.Valid = true

	return nil
}