/*
 * @Author: gengqingyang gengqingyang@onething.net
 * @Date: 2022-11-26 23:15:23
 * @LastEditTime: 2022-12-18 17:40:59
 * @Description:
 */
package biz

import (
	"context"
	"fmt"
	"time"

	v1 "krato_study/api/realworld/v1"
	"krato_study/internal/conf"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

type User struct {
	Email        string
	Username     string
	Bio          string
	Image        interface{}
	PasswordHash string
}

type UserLogin struct {
	Email    string
	Token    string
	Username string
	Bio      string
	Image    interface{}
}

func hashPassword(pwd string) string {

	b, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("hash password err:", err.Error())
		return ""
	}
	fmt.Println("hash password is:", string(b))
	return string(b)
}

func verifyPassword(hashed, input string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(input)); err != nil {
		return false
	}
	return true
}

// GreeterRepo is a Greater repo.
type UserRepo interface {
	CreateUser(context.Context, *User) (uint, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserById(ctx context.Context, id int) (*User, error)
}

type ProfileRepo interface {
}

// GreeterUsecase is a Greeter usecase.
type UserUsecase struct {
	ur  UserRepo
	pr  ProfileRepo
	jwt *conf.Jwt
	log *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewUserUsecase(ur UserRepo, pr ProfileRepo, jwt *conf.Jwt, logger log.Logger) *UserUsecase {
	return &UserUsecase{
		ur:  ur,
		pr:  pr,
		jwt: jwt,
		log: log.NewHelper(logger),
	}
}

func (uc *UserUsecase) GenerateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"nbf":      time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	tokenString, err := token.SignedString([]byte(uc.jwt.Secret))
	if err != nil {
		fmt.Println("GenerateToken err is:", err)
		return "", err
	}

	return tokenString, nil
}

// Register creates a Greeter, and returns the new Greeter.
func (uc *UserUsecase) Register(ctx context.Context, username, password, email string) (*UserLogin, error) {
	uc.log.WithContext(ctx).Infof("Register user")
	user := &User{
		Username:     username,
		PasswordHash: hashPassword(password),
		Email:        email,
	}
	uid, err := uc.ur.CreateUser(ctx, user)
	if err != nil {
		fmt.Println("err is: ", err.Error())
		return nil, err
	}

	u, err := uc.ur.GetUserById(ctx, int(uid))
	if err != nil {
		fmt.Println("err is: ", err.Error())
		return nil, err
	}

	token, err := uc.GenerateToken(username)
	if err != nil {
		return nil, err
	}

	userLogin := &UserLogin{
		Username: u.Username,
		Email:    u.Email,
		Token:    token,
		Bio:      u.Bio,
		Image:    u.Image,
	}

	return userLogin, nil
}

// Login creates a Greeter, and returns the new Greeter.
func (uc *UserUsecase) Login(ctx context.Context, email, password string) (*UserLogin, error) {
	uc.log.WithContext(ctx).Infof("GetUserByEmail")
	user, err := uc.ur.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if !verifyPassword(user.PasswordHash, password) {
		return nil, errors.New(201,"user","密码错误")
	}

	token, err := uc.GenerateToken(user.Username)
	if err != nil {
		return nil, err
	}

	return &UserLogin{
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
		Bio:      user.Bio,
		Image:    user.Image,
	}, nil
}
