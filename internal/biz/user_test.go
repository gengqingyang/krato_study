/*
 * @Author: gengqingyang gengqingyang@onething.net
 * @Date: 2022-12-16 14:24:39
 * @LastEditTime: 2022-12-16 14:33:49
 * @Description:
 */
package biz

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	hashPassword("abc")
}

func TestVerfyPassword(t *testing.T) {
	if verifyPassword("$2a$10$9KarGEFKOEWPIrYAmv0oU.yMpIEMfmu5ErmcTQeJ4MKKKpdxutkfW", "abc") {
		t.Log("verifyPassword pass")
	} else {
		t.Log("verifyPassword not ok")
	}
}
