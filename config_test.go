package gd_sdk_test

import (
	"fmt"
	"github.com/dmlvhh/gd_sdk"
	"testing"
)

var gdSdk *gd_sdk.Config

func TestNewConfig(t *testing.T) {
	gdSdk = gd_sdk.NewConfig(
		"dasdas",
		"dasdsa",
		"dsadad",
		"asdadada",
	)
	fmt.Println(gdSdk)
}
