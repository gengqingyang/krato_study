/*
 * @Author: gengqingyang gengqingyang@onething.net
 * @Date: 2022-12-14 16:50:34
 * @LastEditTime: 2022-12-14 18:09:21
 * @Description:
 */
package service

import (
	"context"
	v1 "krato_study/api/realworld/v1"
)

func (c *RealWorldService) GetCurrentUser(ctx context.Context, req *v1.GetCurrentUserRequest) (*v1.UserReply, error) {
	out := new(v1.UserReply)
	return out, nil
}

func (c *RealWorldService) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UserReply, error) {
	out := new(v1.UserReply)

	return out, nil
}

func (c *RealWorldService) FollowUser(ctx context.Context, req *v1.FollowUserRequest) (*v1.ProfileReply, error) {
	out := new(v1.ProfileReply)

	return out, nil
}

func (c *RealWorldService) UnfollowUser(ctx context.Context, req *v1.UnfollowUserRequest) (*v1.ProfileReply, error) {
	out := new(v1.ProfileReply)

	return out, nil
}

func (c *RealWorldService) GetProfile(ctx context.Context, req *v1.GetProfileRequest) (*v1.ProfileReply, error) {
	out := new(v1.ProfileReply)

	return out, nil
}
