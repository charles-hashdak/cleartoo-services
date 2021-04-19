package main

import (
	"fmt"
  	"net/smtp"
	"errors"
	"log"
	"context"
	"strings"
	"strconv"
	"math/rand"

	pb "github.com/charles-hashdak/cleartoo-services/user-service/proto/user"
	micro "github.com/micro/go-micro/v2"
	"github.com/oliveroneill/exponent-server-sdk-golang/sdk"
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

func (srv *service) SendNotification(ctx context.Context, req *pb.Notification, res *pb.Response) error {
	user, err := repo.Get(req.UserId)

	if err != nil{
		return err
	}

    pushToken, err := expo.NewExponentPushToken(user.PushToken)
    if err != nil {
        panic(err)
    }

    client := expo.NewPushClient(nil)

    response, err := client.Publish(
        &expo.PushMessage{
            To: pushToken,
            Body: "Check your sales!",
            Data: map[string]string{"withSome": "data"},
            Sound: "default",
            Title: "New order!",
            Priority: expo.DefaultPriority,
        },
    )

    if err != nil {
        return err
    }

    if response.ValidateResponse() != nil {
        fmt.Println(response.PushMessage.To, "failed")
    }
    
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

func (srv *service) IsFollowing(ctx context.Context, follower *pb.Follower, res *pb.IsFollowingResponse) error {
	isFollowing, err := srv.repo.IsFollowing(follower)
	if err != nil {
		return err
	}
	res.IsFollowing = isFollowing
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
	req.Rating = 0;
  	req.AvatarUrl= ""
  	req.CoverUrl= ""
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

func (srv *service) FacebookLogin(ctx context.Context, req *pb.User, res *pb.Response) error {
	user, err := srv.repo.GetByEmail(req.Email)
	if err != nil {
		return err
	}
	fmt.Println(user)
	fmt.Println(req)
	if(user.Email != req.Email){
		var usernameId = 1000 + rand.Intn(9999-1000)
		req.Username = strings.Split(user.Email, string('@'))[0] + strconv.Itoa(usernameId)
      	req.Name= ""
      	req.Company= ""
      	req.Description= ""
      	req.Rating= 0
      	req.AvatarUrl= ""
      	req.CoverUrl= ""
      	req.FollowersCount= 0
      	req.FollowingCount= 0
		if err = srv.repo.Create(req); err != nil {
			return errors.New(fmt.Sprintf("error creating user: %v", err))
		}

		token, refreshToken, err := srv.tokenService.Encode(req)
		if err != nil {
			return err
		}
		header := metadata.New(map[string]string{"Set-Cookie": "refresh-jwt="+refreshToken+"; Max-Age=5256000; SameSite=none"})
		if err = grpc.SendHeader(ctx, header); err != nil {
			return status.Errorf(codes.Internal, "unable to send 'Set-Cookie' header")
		}

		req.Password = "";
		res.User = req
		res.Token = &pb.Token{Token: token}
	}else{
		token, refreshToken, err := srv.tokenService.Encode(user)
		if err != nil {
			return err
		}
		header := metadata.New(map[string]string{"Set-Cookie": "refresh-jwt="+refreshToken+"; Max-Age=5256000; SameSite=none"})
		if err = grpc.SendHeader(ctx, header); err != nil {
			return status.Errorf(codes.Internal, "unable to send 'Set-Cookie' header")
		}

		req.Password = "";
		res.User = user
		res.Token = &pb.Token{Token: token}
	}
	return nil
}

func (srv *service) ChangePassword(ctx context.Context, req *pb.User, res *pb.Response) error {
	// Generates a hashed version of our password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New(fmt.Sprintf("error hashing password: %v", err))
	}

	req.Password = string(hashedPass)
	if err := srv.repo.ChangePassword(req); err != nil {
		return errors.New(fmt.Sprintf("error creating user: %v", err))
	}

	req.Password = "";
	res.User = req

	return nil
}

func (srv *service) ResetPassword(ctx context.Context, req *pb.User, res *pb.Response) error {
	user, err1 := srv.repo.GetByEmail(req.Email)
	if err1 != nil {
		return err1
	}
	if(user.Email != req.Email){
		return errors.New(fmt.Sprintf("account not found"))
	}
	// Generates a hashed version of our password
	plainPassword := generatePassword()
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.New(fmt.Sprintf("error hashing password: %v", err))
	}

	req.Password = string(hashedPass)
	if err := srv.repo.ChangePassword(req); err != nil {
		return errors.New(fmt.Sprintf("error creating user: %v", err))
	}

	from := "no_reply@cleartoo.co.th"
	password := "QZyPXnA9"

	// Receiver email address.
	to := []string{
		req.Email,
	}

	// smtp server configuration.
	smtpHost := "smtp.cleartoo.co.th"
	smtpPort := "25"

	// Message.
	message := []byte("To: "+req.Email+"\r\n" +
		"Subject: Password reset!\r\n" +
		"\r\n" +
		"New password = "+plainPassword)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	mailErr := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if mailErr != nil {
		return mailErr
	}

	req.Password = "";
	res.User = req

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

var (
    lowerCharSet   = "abcdedfghijklmnopqrst"
    upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    specialCharSet = "!@#$%&*"
    numberSet      = "0123456789"
    allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

func generatePassword() string {
    var password strings.Builder

    //Set special character
    for i := 0; i < 1; i++ {
        random := rand.Intn(len(specialCharSet))
        password.WriteString(string(specialCharSet[random]))
    }

    //Set numeric
    for i := 0; i < 1; i++ {
        random := rand.Intn(len(numberSet))
        password.WriteString(string(numberSet[random]))
    }

    //Set uppercase
    for i := 0; i < 1; i++ {
        random := rand.Intn(len(upperCharSet))
        password.WriteString(string(upperCharSet[random]))
    }

    remainingLength := 5
    for i := 0; i < remainingLength; i++ {
        random := rand.Intn(len(allCharSet))
        password.WriteString(string(allCharSet[random]))
    }
    inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}