/*
 * @Author: gengqingyang gengqingyang@onething.net
 * @Date: 2022-12-18 14:00:52
 * @LastEditTime: 2022-12-18 17:39:28
 * @Description:
 */
package server

import (
	"fmt"
	"krato_study/internal/customer_error"
	"net/http"

	"github.com/go-kratos/kratos/v2/errors"

	krato_http "github.com/go-kratos/kratos/v2/transport/http"
)

func FromError(err error) *customer_error.HttpError {
	if err == nil {
		return nil
	}

	fmt.Println("err is ",err.Error())

	if se := new(customer_error.HttpError); errors.As(err, &se) {
		return se
	}

	if se := new(errors.Error); errors.As(err, &se) {
		return customer_error.NewHttpError(int(se.Code), map[string]interface{}{
			se.Reason: se.Message,
		})
	}

	return &customer_error.HttpError{Errors: map[string]interface{}{
		"error": "internal server error",
	}, Code: 500}
}

func errorEncoder(w http.ResponseWriter, r *http.Request, err error) {

	se := FromError(err)
	codec, _ := krato_http.CodecForRequest(r, "Accept")
	body, err := codec.Marshal(se)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/"+codec.Name())
	if se.Code < 600 && se.Code > 99 {
		w.WriteHeader(se.Code)
	} else {
		w.WriteHeader(500)
	}
	w.Write(body)

}
