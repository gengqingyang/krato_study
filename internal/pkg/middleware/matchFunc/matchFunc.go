/*
 * @Author: gengqingyang gengqingyang@onething.net
 * @Date: 2022-12-16 17:43:43
 * @LastEditTime: 2022-12-16 18:02:11
 * @Description:
 */
package matchfunc

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware/selector"
)

func NewskipRouteListMatcher() selector.MatchFunc {

	skipRouteList := make(map[string]struct{})
	skipRouteList["/realworld.v1.RealWorld/Register"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := skipRouteList[operation]; ok {
			return false
		}
		return true
	}
}
