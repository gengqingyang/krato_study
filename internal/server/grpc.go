/*
 * @Author: gengqingyang gengqingyang@onething.net
 * @Date: 2022-11-26 23:15:23
 * @LastEditTime: 2023-01-09 16:55:17
 * @Description:
 */
package server

import (
	v1 "krato_study/api/realworld/v1"
	"krato_study/internal/conf"
	"krato_study/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, userService *service.RealWorldService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterRealWorldServer(srv, userService)
	return srv
}
