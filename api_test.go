package gd_sdk

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestChangeSimcardStatus(t *testing.T) {
	res, err := ChangeSimcardStatus(&ChangeSimcardStatusReq{
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
	res, err := BatchChangeSimcardStatus(&BatchChangeSimcardStatusReq{
		//Iccids:   []string{"89861590082420569601"},
		Iccids:   []string{"89861590092420070096"},
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
	res, err := ChangeSimcardLimit(&ChangeSimcardLimitReq{
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
	res, err := ManualChangeSimImei(&ManualChangeSimImeiReq{
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
	res, err := BatchQuerySimFlow(&BatchQuerySimFlowReq{
		//Iccids:    []string{"89861590082420569601"},
		Iccids:    []string{"89861590092420070041"},
		QueryDate: "202501",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	marshal, _ := json.Marshal(res)
	fmt.Println(string(marshal))
}

func TestQueryResultsAsynchronously(t *testing.T) {
	res, err := QueryResultsAsynchronously(&QueryResultsAsynchronouslyReq{
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
		//Iccids: []string{"89861590082420214994"},
		Iccids: []string{"89861590092420070041"},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	marshal, _ := json.Marshal(res)
	fmt.Println(string(marshal))
}
