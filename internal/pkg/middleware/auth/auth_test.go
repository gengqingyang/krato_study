/*
 * @Author: gengqingyang gengqingyang@onething.net
 * @Date: 2022-12-16 16:33:57
 * @LastEditTime: 2022-12-16 16:58:14
 * @Description:
 */
package auth

import (
	"fmt"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken("gengqingyang","hello")
	if err != nil {
		t.Log("generate err:", err.Error())
		return
	}
	fmt.Println("token is:", token)
}
