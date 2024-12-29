package gd_sdk_test

import (
	"fmt"
	"github.com/dmlvhh/gd_sdk"
	"testing"
)

func TestHttpPostRequest(t *testing.T) {
	request, err := gd_sdk.ApiRequest("/batchQuerySimFlow", gd_sdk.ChangeSimcardStatusReq{
		Iccid:    "89861590082420569601",
		OperType: "0",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(request)
}
