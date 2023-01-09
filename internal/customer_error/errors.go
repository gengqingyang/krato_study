/*
 * @Author: gengqingyang gengqingyang@onething.net
 * @Date: 2022-12-18 14:48:57
 * @LastEditTime: 2022-12-18 17:32:49
 * @Description:
 */
package customer_error

import "fmt"

type HttpError struct {
	Errors map[string]interface{} `json:"errors"`
	Code   int                    `json:"-"`
}

func (e *HttpError) Error() string {
	return fmt.Sprintf("http error: %d", e.Code)
}

func NewHttpError(code int, errs map[string]interface{}) *HttpError {
	return &HttpError{
		Code:   code,
		Errors: errs,
	}
}
