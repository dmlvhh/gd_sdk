package gd_sdk

import (
	"encoding/json"
	"fmt"
	"testing"
)

var gdSdk *Config

func init() {
	gdSdk = NewConfig("https://122.192.50.42:11060/ipa/api/v1/wlqd", "1847483747142742018", "QRAXmCI0DuJOvucX2Kh9FiyGZA38Eoba", "5xYx0jWdp9KE8UVsWK4y2UDBUlc8w14s")
}

func TestChangeSimcardStatus(t *testing.T) {
	res, err := gdSdk.ChangeSimcardStatus(&ChangeSimcardStatusReq{
		Iccid:    "89861590082420569601",
		OperType: "0",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	marshal, _ := json.Marshal(res)
	fmt.Println(string(marshal))
}

func TestBatchChangeSimcardStatus(t *testing.T) {
	res, err := gdSdk.BatchChangeSimcardStatus(&BatchChangeSimcardStatusReq{
		Iccids:   []string{"89861590082420569601"},
		OperType: "0",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	marshal, _ := json.Marshal(res)
	fmt.Println(string(marshal))
}

func TestChangeSimcardLimit(t *testing.T) {
	res, err := gdSdk.ChangeSimcardLimit(&ChangeSimcardLimitReq{
		Iccid: "89861590082420569601",
		Code:  "2-512kbps",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	marshal, _ := json.Marshal(res)
	fmt.Println(string(marshal))
}

func TestManualChangeSimImei(t *testing.T) {
	res, err := gdSdk.ManualChangeSimImei(&ManualChangeSimImeiReq{
		Iccid:   "89861590082420569601",
		OldImei: "1234567890111",
		NewImei: "1234567890112",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	marshal, _ := json.Marshal(res)
	fmt.Println(string(marshal))
}

func TestBatchQuerySimFlow(t *testing.T) {
	res, err := gdSdk.BatchQuerySimFlow(&BatchQuerySimFlowReq{
		Iccids:    []string{"89861590082420569601"},
		QueryDate: "202412",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	marshal, _ := json.Marshal(res)
	fmt.Println(string(marshal))
}

func TestQueryResultsAsynchronously(t *testing.T) {
	res, err := gdSdk.QueryResultsAsynchronously(&QueryResultsAsynchronouslyReq{
		SerialNumber: "17735778075672xxxx",
	})

	if err != nil {
		fmt.Println(err)
		return
	}
	marshal, _ := json.Marshal(res)
	fmt.Println(string(marshal))
}

func TestQueryBasicInfoBatch(t *testing.T) {
	res, err := gdSdk.QueryBasicInfoBatch(&QueryBasicInfoBatchReq{
		Iccids: []string{"89861590082420214994"},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	marshal, _ := json.Marshal(res)
	fmt.Println(string(marshal))
}
