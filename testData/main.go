package main

import (
	"encoding/json"
	"fmt"
	"github.com/dmlvhh/gd_sdk"
)

func main() {
	gd := gd_sdk.NewConfig("https://122.192.50.42:11060", "1847483747142742018", "QRAXmCI0DuJOvucX2Kh9FiyGZA38Eoba", "5xYx0jWdp9KE8UVsWK4y2UDBUlc8w14s")
	res, err := gd.BatchQuerySimFlow(&gd_sdk.BatchQuerySimFlowReq{
		Iccids:    []string{"89861590082420569601"},
		QueryDate: "202501",
	})
	if err != nil {
		return
	}
	marshal, _ := json.Marshal(res)
	fmt.Println(string(marshal))
}
