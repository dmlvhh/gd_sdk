package gd_sdk_test

import (
	"fmt"
	"github.com/dmlvhh/gd_sdk"
	"testing"
)

func TestGenerateSignature(t *testing.T) {
	appKey := "QRAXmCI0DuJOvucX2Kh9FiyGZA38Eoba"
	appSecret := "5xYx0jWdp9KE8UVsWK4y2UDBUlc8w14s"
	sign := gd_sdk.GenerateSignature(appKey, appSecret)
	fmt.Println(sign)
}

func TestGetNowUnixTime(t *testing.T) {
	time := gd_sdk.GetNowUnixTime()
	fmt.Println(time)
}
