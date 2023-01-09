/*
 * @Author: gengqingyang gengqingyang@onething.net
 * @Date: 2022-12-04 12:00:19
 * @LastEditTime: 2022-12-14 18:19:45
 * @Description:
 */
package service

import (
	v1 "krato_study/api/realworld/v1"
	"krato_study/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRealWorldService)

type RealWorldService struct {
	v1.UnimplementedRealWorldServer

	sc  *biz.SocalUsecase
	uc  *biz.UserUsecase
	log *log.Helper
}

// NewGreeterService new a greeter service.
func NewRealWorldService(uc *biz.UserUsecase, sc *biz.SocalUsecase, logger log.Logger) *RealWorldService {
	return &RealWorldService{
		sc:  sc,
		uc:  uc,
		log: log.NewHelper(logger),
	}
}
