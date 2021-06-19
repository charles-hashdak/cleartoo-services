package main

import (
	"fmt"
  	"net/smtp"
	"errors"
	"log"
	"os"
	"bytes"
	"html/template"
	"context"
	"strings"
	"io/ioutil"
	"strconv"
	"math"
	"math/rand"
	"sync"
    "encoding/json"
	uuid "github.com/satori/go.uuid"

	pb "github.com/charles-hashdak/cleartoo-services/user-service/proto/user"
	orderPb "github.com/charles-hashdak/cleartoo-services/order-service/proto/order"
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
	globalMutex 	sync.Mutex
	orderClient 	orderPb.OrderService
}

func (srv *service) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	srv.globalMutex.Lock()
	defer srv.globalMutex.Unlock()
	user, err := srv.repo.Get(req.Id)
	if err != nil {
		return err
	}
	user.Rating = float32(math.Round(float64(user.Rating*10))/10)
	user.Password = "";
	res.User = user
	return nil
}

func (srv *service) SendNotification(ctx context.Context, req *pb.Notification, res *pb.Response) error {
	srv.globalMutex.Lock()
	defer srv.globalMutex.Unlock()
	user, err := srv.repo.Get(req.UserId)

	if err != nil{
		return err
	}

	if(user.PushToken != ""){
	    pushToken, err := expo.NewExponentPushToken(user.PushToken)
	    if err != nil {
	        fmt.Println(err)
	    }

	    client := expo.NewPushClient(nil)

		var dataInterface map[string]string
		err = json.Unmarshal([]byte(req.Data), &dataInterface)
	    if err != nil {
	        fmt.Println(err)
	    }

		fmt.Println(req.Data)
		fmt.Println(dataInterface)

	    response, err := client.Publish(
	        &expo.PushMessage{
	            To: []expo.ExponentPushToken{pushToken},
	            Body: req.Body,
	            Data: dataInterface,
	            Sound: "default",
	            Title: req.Title,
	            Priority: expo.DefaultPriority,
	        },
	    )

	    if err != nil {
	        return err
	    }

	    if response.ValidateResponse() != nil {
	        fmt.Println(response.PushMessage.To, "failed")
	    }
	}

	return nil
}

func (srv *service) Edit(ctx context.Context, req *pb.User, res *pb.Response) error {
	srv.globalMutex.Lock()
	defer srv.globalMutex.Unlock()
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

func (srv *service) Rate(ctx context.Context, rating *pb.Rating, res *pb.RateResponse) error {
	err := srv.repo.Rate(rating)
	if err != nil {
		return err
	}
	res.Rated = true
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

func (srv *service) GetFollowing(ctx context.Context, req *pb.User, res *pb.Response) error {
	users, err := srv.repo.GetFollowing(req.Id)
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}

func (srv *service) Auth(ctx context.Context, req *pb.User, res *pb.Response) error {
	srv.globalMutex.Lock()
	defer srv.globalMutex.Unlock()
	req.Email = strings.ToLower(req.Email)
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
	srv.globalMutex.Lock()
	defer srv.globalMutex.Unlock()
	req.Id = uuid.NewV4().String()
	req.Email = strings.ToLower(req.Email)
	req.Username = strings.ToLower(req.Username)
	// var usernameId = 1000 + rand.Intn(9999-1000)
	// user.Username = strings.Split(user.Email, string('@'))[0] + strconv.Itoa(usernameId)
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

	_, err = srv.orderClient.InitializeWallet(ctx, &orderPb.InitializeWalletRequest{
		UserId: req.Id,
	})

	if err != nil{
		return err
	}

	from := os.Getenv("EMAIL_USER")
	password := os.Getenv("EMAIL_PASSWORD")

	// Receiver email address.
	to := []string{
		req.Email,
	}

	// smtp server configuration.
	smtpHost := "us2.smtp.mailhostbox.com"
	smtpPort := "25"

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	files, err := ioutil.ReadDir("./")
    if err != nil {
        log.Fatal(err)
    }
 
    for _, f := range files {
            fmt.Println(f.Name())
    }

	t, err := template.ParseFiles("/var/templates/emails/welcome.html")
	if err != nil {
		return err
	}

	var body bytes.Buffer

	headers := "MIME-version: 1.0;\nContent-Type: text/html;"

	body.Write([]byte(fmt.Sprintf("From: Cleartoo <no_reply@cleartoo.co.th>\r\nTo: "+req.Email+"\r\nSubject: Welcome!\r\n%s\n\n", headers)))

	t.Execute(&body, struct {
		Username 	string
	}{
		Username: 	req.Username,
	})

	// Sending email.
	mailErr := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if mailErr != nil {
		return mailErr
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
	srv.globalMutex.Lock()
	defer srv.globalMutex.Unlock()
	user, err := srv.repo.GetByEmail(req.Email)
	if err != nil && err.Error() != "record not found" {
		fmt.Println(err)
		return err
	}
	if(err != nil && err.Error() == "record not found"){
		req.Id = uuid.NewV4().String()
		var usernameId = 1000 + rand.Intn(9999-1000)
		req.Username = strings.Split(req.Email, string('@'))[0] + strconv.Itoa(usernameId)
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

func (srv *service) AppleLogin(ctx context.Context, req *pb.User, res *pb.Response) error {
	srv.globalMutex.Lock()
	defer srv.globalMutex.Unlock()
	fmt.Println(req.AppleUserId)
	user, err := srv.repo.GetByAppleUserId(req.AppleUserId)
	if err != nil && err.Error() != "record not found" {
		fmt.Println(err)
		return err
	}
	if(err != nil && err.Error() == "record not found"){
		req.Id = uuid.NewV4().String()
		var usernameId = 1000 + rand.Intn(9999-1000)
		req.Username = strings.Split(req.Email, string('@'))[0] + strconv.Itoa(usernameId)
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
	srv.globalMutex.Lock()
	defer srv.globalMutex.Unlock()
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
	srv.globalMutex.Lock()
	defer srv.globalMutex.Unlock()
	req.Email = strings.ToLower(req.Email)
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

	from := os.Getenv("EMAIL_USER")
	password := os.Getenv("EMAIL_PASSWORD")

	// Receiver email address.
	to := []string{
		req.Email,
	}

	// smtp server configuration.
	smtpHost := "us2.smtp.mailhostbox.com"
	smtpPort := "25"

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	files, err := ioutil.ReadDir("./")
    if err != nil {
        log.Fatal(err)
    }
 
    for _, f := range files {
            fmt.Println(f.Name())
    }

	t, err := template.ParseFiles("/var/templates/emails/password_recovery.html")
	if err != nil {
		return err
	}

	var body bytes.Buffer

	headers := "MIME-version: 1.0;\nContent-Type: text/html;"

	body.Write([]byte(fmt.Sprintf("From: Cleartoo <no_reply@cleartoo.co.th>\r\nTo: "+req.Email+"\r\nSubject: Password reset!\r\n%s\n\n", headers)))

	t.Execute(&body, struct {
		Username 	string
		Password 	string
	}{
		Username: 	user.Username,
		Password:	plainPassword,
	})

	// Sending email.
	mailErr := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
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