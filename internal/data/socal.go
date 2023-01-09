/*
 * @Author: gengqingyang gengqingyang@onething.net
 * @Date: 2022-12-04 11:56:46
 * @LastEditTime: 2022-12-14 18:13:42
 * @Description:
 */
package data

import (
	"context"
	"krato_study/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type socalRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func NewCommentRepo(data *Data, logger log.Logger) biz.CommentRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func NewTagRepo(data *Data, logger log.Logger) biz.TagRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) CreateArticle(ctx context.Context, g *biz.Article) (*biz.Article, error) {
	return g, nil
}

func (r *userRepo) CreateComment(ctx context.Context, g *biz.Comment) (*biz.Comment, error) {
	return g, nil
}

func (r *userRepo) CreateTag(ctx context.Context, g *biz.Tag) (*biz.Tag, error) {
	return g, nil
}
