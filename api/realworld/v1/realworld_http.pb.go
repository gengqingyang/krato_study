// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.3
// - protoc             v3.19.4
// source: realworld/v1/realworld.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationRealWorldAddCommentsToAnArticle = "/realworld.v1.RealWorld/AddCommentsToAnArticle"
const OperationRealWorldCreateArticle = "/realworld.v1.RealWorld/CreateArticle"
const OperationRealWorldDeleteArticle = "/realworld.v1.RealWorld/DeleteArticle"
const OperationRealWorldDeleteComment = "/realworld.v1.RealWorld/DeleteComment"
const OperationRealWorldFavoriteArticle = "/realworld.v1.RealWorld/FavoriteArticle"
const OperationRealWorldFeedArticles = "/realworld.v1.RealWorld/FeedArticles"
const OperationRealWorldFollowUser = "/realworld.v1.RealWorld/FollowUser"
const OperationRealWorldGetArticle = "/realworld.v1.RealWorld/GetArticle"
const OperationRealWorldGetCommentsFromAnArticle = "/realworld.v1.RealWorld/GetCommentsFromAnArticle"
const OperationRealWorldGetCurrentUser = "/realworld.v1.RealWorld/GetCurrentUser"
const OperationRealWorldGetProfile = "/realworld.v1.RealWorld/GetProfile"
const OperationRealWorldGetTags = "/realworld.v1.RealWorld/GetTags"
const OperationRealWorldListArticles = "/realworld.v1.RealWorld/ListArticles"
const OperationRealWorldLogin = "/realworld.v1.RealWorld/Login"
const OperationRealWorldRegister = "/realworld.v1.RealWorld/Register"
const OperationRealWorldUnfavoriteArticle = "/realworld.v1.RealWorld/UnfavoriteArticle"
const OperationRealWorldUnfollowUser = "/realworld.v1.RealWorld/UnfollowUser"
const OperationRealWorldUpdateArticle = "/realworld.v1.RealWorld/UpdateArticle"
const OperationRealWorldUpdateUser = "/realworld.v1.RealWorld/UpdateUser"

type RealWorldHTTPServer interface {
	AddCommentsToAnArticle(context.Context, *AddCommentsToAnArticleRequest) (*SingleCommentReply, error)
	CreateArticle(context.Context, *CreateArticleRequest) (*ArticleReplay, error)
	DeleteArticle(context.Context, *DeleteArticleRequest) (*NullReplay, error)
	DeleteComment(context.Context, *DeleteCommentRequest) (*NullReplay, error)
	FavoriteArticle(context.Context, *FavoriteArticleRequest) (*SingleArticleReply, error)
	FeedArticles(context.Context, *FeedArticlesRequest) (*MultipleArticlesReply, error)
	FollowUser(context.Context, *FollowUserRequest) (*ProfileReply, error)
	GetArticle(context.Context, *GetArticleRequest) (*SingleArticleReply, error)
	GetCommentsFromAnArticle(context.Context, *GetCommentsFromAnArticleRequest) (*MultipleCommentsReply, error)
	GetCurrentUser(context.Context, *GetCurrentUserRequest) (*UserReply, error)
	GetProfile(context.Context, *GetProfileRequest) (*ProfileReply, error)
	GetTags(context.Context, *GetTagsRequest) (*TagsReply, error)
	ListArticles(context.Context, *ListArticlesRequest) (*MultipleArticlesReply, error)
	Login(context.Context, *LoginRequest) (*UserReply, error)
	Register(context.Context, *RegisterRequest) (*UserReply, error)
	UnfavoriteArticle(context.Context, *UnfavoriteArticleRequest) (*SingleArticleReply, error)
	UnfollowUser(context.Context, *UnfollowUserRequest) (*ProfileReply, error)
	UpdateArticle(context.Context, *UpdateArticleRequest) (*ArticleReplay, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UserReply, error)
}

func RegisterRealWorldHTTPServer(s *http.Server, srv RealWorldHTTPServer) {
	r := s.Route("/")
	r.POST("/api/users/login", _RealWorld_Login0_HTTP_Handler(srv))
	r.POST("/api/users", _RealWorld_Register0_HTTP_Handler(srv))
	r.GET("/api/user", _RealWorld_GetCurrentUser0_HTTP_Handler(srv))
	r.PUT("/api/user", _RealWorld_UpdateUser0_HTTP_Handler(srv))
	r.GET("/api/profiles/:username", _RealWorld_GetProfile0_HTTP_Handler(srv))
	r.POST("/api/profiles/:username/follow", _RealWorld_FollowUser0_HTTP_Handler(srv))
	r.DELETE("/api/profiles/:username/follow", _RealWorld_UnfollowUser0_HTTP_Handler(srv))
	r.GET("/api/articles", _RealWorld_ListArticles0_HTTP_Handler(srv))
	r.GET("/api/articles/feed", _RealWorld_FeedArticles0_HTTP_Handler(srv))
	r.GET("/api/articles/:slug", _RealWorld_GetArticle0_HTTP_Handler(srv))
	r.POST("/api/articles", _RealWorld_CreateArticle0_HTTP_Handler(srv))
	r.PUT("/api/articles/:slug", _RealWorld_UpdateArticle0_HTTP_Handler(srv))
	r.DELETE("/api/articles/:slug", _RealWorld_DeleteArticle0_HTTP_Handler(srv))
	r.POST("/api/articles/:slug/comments", _RealWorld_AddCommentsToAnArticle0_HTTP_Handler(srv))
	r.GET("/api/articles/:slug/comments", _RealWorld_GetCommentsFromAnArticle0_HTTP_Handler(srv))
	r.DELETE("/api/articles/:slug/comments/:id", _RealWorld_DeleteComment0_HTTP_Handler(srv))
	r.POST("/api/articles/:slug/favorite", _RealWorld_FavoriteArticle0_HTTP_Handler(srv))
	r.DELETE("/api/articles/:slug/favorite", _RealWorld_UnfavoriteArticle0_HTTP_Handler(srv))
	r.GET("/api/tags", _RealWorld_GetTags0_HTTP_Handler(srv))
}

func _RealWorld_Login0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRealWorldLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_Register0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRealWorldRegister)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_GetCurrentUser0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCurrentUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRealWorldGetCurrentUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCurrentUser(ctx, req.(*GetCurrentUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_UpdateUser0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRealWorldUpdateUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUser(ctx, req.(*UpdateUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_GetProfile0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetProfileRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRealWorldGetProfile)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetProfile(ctx, req.(*GetProfileRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ProfileReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_FollowUser0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FollowUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRealWorldFollowUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FollowUser(ctx, req.(*FollowUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ProfileReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_UnfollowUser0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UnfollowUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRealWorldUnfollowUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UnfollowUser(ctx, req.(*UnfollowUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ProfileReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_ListArticles0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListArticlesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRealWorldListArticles)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListArticles(ctx, req.(*ListArticlesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MultipleArticlesReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_FeedArticles0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FeedArticlesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRealWorldFeedArticles)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FeedArticles(ctx, req.(*FeedArticlesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MultipleArticlesReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_GetArticle0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetArticleRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRealWorldGetArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetArticle(ctx, req.(*GetArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SingleArticleReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_CreateArticle0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateArticleRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRealWorldCreateArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateArticle(ctx, req.(*CreateArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ArticleReplay)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_UpdateArticle0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateArticleRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRealWorldUpdateArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateArticle(ctx, req.(*UpdateArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ArticleReplay)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_DeleteArticle0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteArticleRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRealWorldDeleteArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteArticle(ctx, req.(*DeleteArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NullReplay)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_AddCommentsToAnArticle0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddCommentsToAnArticleRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRealWorldAddCommentsToAnArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddCommentsToAnArticle(ctx, req.(*AddCommentsToAnArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SingleCommentReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_GetCommentsFromAnArticle0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCommentsFromAnArticleRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRealWorldGetCommentsFromAnArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCommentsFromAnArticle(ctx, req.(*GetCommentsFromAnArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MultipleCommentsReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_DeleteComment0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteCommentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRealWorldDeleteComment)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteComment(ctx, req.(*DeleteCommentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NullReplay)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_FavoriteArticle0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FavoriteArticleRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRealWorldFavoriteArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FavoriteArticle(ctx, req.(*FavoriteArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SingleArticleReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_UnfavoriteArticle0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UnfavoriteArticleRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRealWorldUnfavoriteArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UnfavoriteArticle(ctx, req.(*UnfavoriteArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SingleArticleReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_GetTags0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetTagsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRealWorldGetTags)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetTags(ctx, req.(*GetTagsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TagsReply)
		return ctx.Result(200, reply)
	}
}

type RealWorldHTTPClient interface {
	AddCommentsToAnArticle(ctx context.Context, req *AddCommentsToAnArticleRequest, opts ...http.CallOption) (rsp *SingleCommentReply, err error)
	CreateArticle(ctx context.Context, req *CreateArticleRequest, opts ...http.CallOption) (rsp *ArticleReplay, err error)
	DeleteArticle(ctx context.Context, req *DeleteArticleRequest, opts ...http.CallOption) (rsp *NullReplay, err error)
	DeleteComment(ctx context.Context, req *DeleteCommentRequest, opts ...http.CallOption) (rsp *NullReplay, err error)
	FavoriteArticle(ctx context.Context, req *FavoriteArticleRequest, opts ...http.CallOption) (rsp *SingleArticleReply, err error)
	FeedArticles(ctx context.Context, req *FeedArticlesRequest, opts ...http.CallOption) (rsp *MultipleArticlesReply, err error)
	FollowUser(ctx context.Context, req *FollowUserRequest, opts ...http.CallOption) (rsp *ProfileReply, err error)
	GetArticle(ctx context.Context, req *GetArticleRequest, opts ...http.CallOption) (rsp *SingleArticleReply, err error)
	GetCommentsFromAnArticle(ctx context.Context, req *GetCommentsFromAnArticleRequest, opts ...http.CallOption) (rsp *MultipleCommentsReply, err error)
	GetCurrentUser(ctx context.Context, req *GetCurrentUserRequest, opts ...http.CallOption) (rsp *UserReply, err error)
	GetProfile(ctx context.Context, req *GetProfileRequest, opts ...http.CallOption) (rsp *ProfileReply, err error)
	GetTags(ctx context.Context, req *GetTagsRequest, opts ...http.CallOption) (rsp *TagsReply, err error)
	ListArticles(ctx context.Context, req *ListArticlesRequest, opts ...http.CallOption) (rsp *MultipleArticlesReply, err error)
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *UserReply, err error)
	Register(ctx context.Context, req *RegisterRequest, opts ...http.CallOption) (rsp *UserReply, err error)
	UnfavoriteArticle(ctx context.Context, req *UnfavoriteArticleRequest, opts ...http.CallOption) (rsp *SingleArticleReply, err error)
	UnfollowUser(ctx context.Context, req *UnfollowUserRequest, opts ...http.CallOption) (rsp *ProfileReply, err error)
	UpdateArticle(ctx context.Context, req *UpdateArticleRequest, opts ...http.CallOption) (rsp *ArticleReplay, err error)
	UpdateUser(ctx context.Context, req *UpdateUserRequest, opts ...http.CallOption) (rsp *UserReply, err error)
}

type RealWorldHTTPClientImpl struct {
	cc *http.Client
}

func NewRealWorldHTTPClient(client *http.Client) RealWorldHTTPClient {
	return &RealWorldHTTPClientImpl{client}
}

func (c *RealWorldHTTPClientImpl) AddCommentsToAnArticle(ctx context.Context, in *AddCommentsToAnArticleRequest, opts ...http.CallOption) (*SingleCommentReply, error) {
	var out SingleCommentReply
	pattern := "/api/articles/:slug/comments"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRealWorldAddCommentsToAnArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...http.CallOption) (*ArticleReplay, error) {
	var out ArticleReplay
	pattern := "/api/articles"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRealWorldCreateArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...http.CallOption) (*NullReplay, error) {
	var out NullReplay
	pattern := "/api/articles/:slug"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRealWorldDeleteArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...http.CallOption) (*NullReplay, error) {
	var out NullReplay
	pattern := "/api/articles/:slug/comments/:id"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRealWorldDeleteComment))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) FavoriteArticle(ctx context.Context, in *FavoriteArticleRequest, opts ...http.CallOption) (*SingleArticleReply, error) {
	var out SingleArticleReply
	pattern := "/api/articles/:slug/favorite"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRealWorldFavoriteArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) FeedArticles(ctx context.Context, in *FeedArticlesRequest, opts ...http.CallOption) (*MultipleArticlesReply, error) {
	var out MultipleArticlesReply
	pattern := "/api/articles/feed"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRealWorldFeedArticles))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) FollowUser(ctx context.Context, in *FollowUserRequest, opts ...http.CallOption) (*ProfileReply, error) {
	var out ProfileReply
	pattern := "/api/profiles/:username/follow"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRealWorldFollowUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) GetArticle(ctx context.Context, in *GetArticleRequest, opts ...http.CallOption) (*SingleArticleReply, error) {
	var out SingleArticleReply
	pattern := "/api/articles/:slug"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRealWorldGetArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) GetCommentsFromAnArticle(ctx context.Context, in *GetCommentsFromAnArticleRequest, opts ...http.CallOption) (*MultipleCommentsReply, error) {
	var out MultipleCommentsReply
	pattern := "/api/articles/:slug/comments"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRealWorldGetCommentsFromAnArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) GetCurrentUser(ctx context.Context, in *GetCurrentUserRequest, opts ...http.CallOption) (*UserReply, error) {
	var out UserReply
	pattern := "/api/user"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRealWorldGetCurrentUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) GetProfile(ctx context.Context, in *GetProfileRequest, opts ...http.CallOption) (*ProfileReply, error) {
	var out ProfileReply
	pattern := "/api/profiles/:username"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRealWorldGetProfile))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) GetTags(ctx context.Context, in *GetTagsRequest, opts ...http.CallOption) (*TagsReply, error) {
	var out TagsReply
	pattern := "/api/tags"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRealWorldGetTags))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) ListArticles(ctx context.Context, in *ListArticlesRequest, opts ...http.CallOption) (*MultipleArticlesReply, error) {
	var out MultipleArticlesReply
	pattern := "/api/articles"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRealWorldListArticles))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*UserReply, error) {
	var out UserReply
	pattern := "/api/users/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRealWorldLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) Register(ctx context.Context, in *RegisterRequest, opts ...http.CallOption) (*UserReply, error) {
	var out UserReply
	pattern := "/api/users"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRealWorldRegister))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) UnfavoriteArticle(ctx context.Context, in *UnfavoriteArticleRequest, opts ...http.CallOption) (*SingleArticleReply, error) {
	var out SingleArticleReply
	pattern := "/api/articles/:slug/favorite"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRealWorldUnfavoriteArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) UnfollowUser(ctx context.Context, in *UnfollowUserRequest, opts ...http.CallOption) (*ProfileReply, error) {
	var out ProfileReply
	pattern := "/api/profiles/:username/follow"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRealWorldUnfollowUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) UpdateArticle(ctx context.Context, in *UpdateArticleRequest, opts ...http.CallOption) (*ArticleReplay, error) {
	var out ArticleReplay
	pattern := "/api/articles/:slug"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRealWorldUpdateArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...http.CallOption) (*UserReply, error) {
	var out UserReply
	pattern := "/api/user"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRealWorldUpdateUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
