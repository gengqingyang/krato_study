/*
 * @Author: gengqingyang gengqingyang@onething.net
 * @Date: 2022-11-26 23:15:23
 * @LastEditTime: 2022-12-14 18:04:53
 * @Description:
 */
package server

import (
	"github.com/google/wire"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer,NewHTTPServer)
