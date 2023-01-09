/*
 * @Author: gengqingyang gengqingyang@onething.net
 * @Date: 2022-12-14 16:49:34
 * @LastEditTime: 2022-12-14 18:09:00
 * @Description:
 */
package service

import (
	"context"
	v1 "krato_study/api/realworld/v1"
)

func (c *RealWorldService) ListArticles(ctx context.Context, req *v1.ListArticlesRequest) (*v1.MultipleArticlesReply, error) {
	out := new(v1.MultipleArticlesReply)

	return out, nil
}

func (c *RealWorldService) FeedArticles(ctx context.Context, req *v1.FeedArticlesRequest) (*v1.MultipleArticlesReply, error) {
	out := new(v1.MultipleArticlesReply)

	return out, nil
}

func (c *RealWorldService) GetArticle(ctx context.Context, req *v1.GetArticleRequest) (*v1.SingleArticleReply, error) {
	out := new(v1.SingleArticleReply)

	return out, nil
}

func (c *RealWorldService) CreateArticle(ctx context.Context, req *v1.CreateArticleRequest) (*v1.ArticleReplay, error) {
	out := new(v1.ArticleReplay)

	return out, nil
}

func (c *RealWorldService) UpdateArticle(ctx context.Context, req *v1.UpdateArticleRequest) (*v1.ArticleReplay, error) {
	out := new(v1.ArticleReplay)

	return out, nil
}

func (c *RealWorldService) DeleteArticle(ctx context.Context, req *v1.DeleteArticleRequest) (*v1.NullReplay, error) {
	out := new(v1.NullReplay)

	return out, nil
}

func (c *RealWorldService) AddCommentsToAnArticle(ctx context.Context, req *v1.AddCommentsToAnArticleRequest) (*v1.SingleCommentReply, error) {
	out := new(v1.SingleCommentReply)
	return out, nil
}

func (c *RealWorldService) GetCommentsFromAnArticle(ctx context.Context, req *v1.GetCommentsFromAnArticleRequest) (*v1.MultipleCommentsReply, error) {
	out := new(v1.MultipleCommentsReply)

	return out, nil
}

func (c *RealWorldService) DeleteComment(ctx context.Context, req *v1.DeleteCommentRequest) (*v1.NullReplay, error) {
	out := new(v1.NullReplay)

	return out, nil
}

func (c *RealWorldService) FavoriteArticle(ctx context.Context, req *v1.FavoriteArticleRequest) (*v1.SingleArticleReply, error) {
	out := new(v1.SingleArticleReply)

	return out, nil
}

func (c *RealWorldService) UnfavoriteArticle(ctx context.Context, req *v1.UnfavoriteArticleRequest) (*v1.SingleArticleReply, error) {
	out := new(v1.SingleArticleReply)

	return out, nil
}

func (c *RealWorldService) GetTags(ctx context.Context, req *v1.GetTagsRequest) (*v1.TagsReply, error) {
	out := new(v1.TagsReply)

	return out, nil
}
