/*
 * @Author: gengqingyang gengqingyang@onething.net
 * @Date: 2022-12-04 11:55:23
 * @LastEditTime: 2022-12-14 17:38:40
 * @Description:

 */
package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewSocalUsecase, NewUserUsecase)
