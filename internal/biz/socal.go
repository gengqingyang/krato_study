/*
 * @Author: gengqingyang gengqingyang@onething.net
 * @Date: 2022-11-26 23:15:23
 * @LastEditTime: 2022-12-14 17:59:49
 * @Description:
 */
package biz

import (
	"context"

	v1 "krato_study/api/realworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrSocalNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "socal not found")
)

// GreeterRepo is a Greater repo.
type ArticleRepo interface {
	CreateArticle(ctx context.Context, article *Article) (*Article, error)
}

type CommentRepo interface {
	CreateComment(ctx context.Context, article *Comment) (*Comment, error)
}

type TagRepo interface {
	CreateTag(ctx context.Context, article *Tag) (*Tag, error)
}

type Article struct {
	name string
}

type Comment struct {
}

type Tag struct {
}

// GreeterUsecase is a Greeter usecase.
type SocalUsecase struct {
	ar  ArticleRepo
	cr  CommentRepo
	tr  TagRepo
	log *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewSocalUsecase(ar ArticleRepo, cr CommentRepo, tr TagRepo, logger log.Logger) *SocalUsecase {
	return &SocalUsecase{
		ar:  ar,
		cr:  cr,
		tr:  tr,
		log: log.NewHelper(logger),
	}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *SocalUsecase) CreateArticle(ctx context.Context) (*Article, error) {
	uc.log.WithContext(ctx).Infof("Create article")
	return nil, nil
	// return uc.ar.Save(ctx, g)
}
