/*
 * @Author: gengqingyang gengqingyang@onething.net
 * @Date: 2022-12-04 11:56:46
 * @LastEditTime: 2022-12-18 17:45:36
 * @Description:
 */
package data

import (
	"context"

	"krato_study/internal/biz"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

type User struct {
	gorm.Model
	Email        string `gorm:"size:50"`
	Username     string `gorm:"size:50"`
	Bio          string `gorm:"size:50"`
	Image        string `gorm:"size:256"`
	PasswordHash string `gorm:"size:128"`
}

// NewGreeterRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func NewProfileRepo(data *Data, logger log.Logger) biz.ProfileRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) CreateUser(ctx context.Context, g *biz.User) (uint, error) {
	user := &User{
		Email:        g.Email,
		Username:     g.Username,
		Bio:          g.Bio,
		PasswordHash: g.PasswordHash,
	}
	rv := r.data.db.Create(user)
	if rv.Error != nil {
		return 0, rv.Error
	}
	return user.ID, nil
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (*biz.User, error) {
	user := &User{}
	rv := r.data.db.Model(&User{}).Where("email = ?", email).First(user)

	if errors.Is(rv.Error, gorm.ErrRecordNotFound) {
		return nil, errors.NotFound("user", "not find by email")
	}
	if rv.Error != nil {
		return nil, rv.Error
	}

	return &biz.User{
		Email:        user.Email,
		Username:     user.Username,
		Bio:          user.Bio,
		Image:        user.Image,
		PasswordHash: user.PasswordHash,
	}, nil
}

func (r *userRepo) GetUserById(ctx context.Context, id int) (*biz.User, error) {

	user := &User{}
	rv := r.data.db.Model(&User{}).Where("id = ?", id).Find(user)

	if rv.Error != nil {
		return nil, rv.Error
	}

	return &biz.User{
		Email:        user.Email,
		Username:     user.Username,
		Bio:          user.Bio,
		Image:        user.Image,
		PasswordHash: user.PasswordHash,
	}, nil
}
