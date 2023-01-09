/*
 * @Author: gengqingyang gengqingyang@onething.net
 * @Date: 2022-12-11 15:37:56
 * @LastEditTime: 2022-12-18 16:31:36
 * @Description:
 */
package service

import (
	"context"
	v1 "krato_study/api/realworld/v1"
)

func (r *RealWorldService) Login(ctx context.Context, req *v1.LoginRequest) (reply *v1.UserReply, err error) {

	// 自定义错误返回
	// return nil, customer_error.NewHttpErrot(402, map[string]interface{}{"body": []string{"can't find params"}})
	userLogin, err := r.uc.Login(ctx, req.User.Email, req.User.Password)
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Email:    userLogin.Email,
			Token:    userLogin.Token,
			Username: userLogin.Username,
			Bio:      userLogin.Bio,
		},
	}, nil
}

func (r *RealWorldService) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.UserReply, error) {
	userLogin, err := r.uc.Register(ctx, req.User.Username, req.User.Password, req.User.Email)
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Email:    userLogin.Email,
			Token:    userLogin.Token,
			Username: userLogin.Username,
			Bio:      userLogin.Bio,
		},
	}, nil
}
