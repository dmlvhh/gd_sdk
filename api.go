package gd_sdk

import (
	"encoding/json"
	"log"
)

//var gdSdk *Config
//
//func init() {
//	gdSdk = NewConfig(&Config{
//		ApiURL:    "https://122.192.50.42:11060/ipa/api/v1/wlqd",
//		AppID:     "1847483747142742018",
//		AppKey:    "QRAXmCI0DuJOvucX2Kh9FiyGZA38Eoba",
//		AppSecret: "5xYx0jWdp9KE8UVsWK4y2UDBUlc8w14s",
//	})
//}

// ChangeSimcardStatus 物联网卡停复机 此接口不可频繁调用，同号码，相同操作，5分钟内最多调用3次。
func ChangeSimcardStatus(req *ChangeSimcardStatusReq) (res Response, err error) {
	request, err := ApiRequest("/changeSimcardStatus", req)
	if err != nil {
		log.Fatalf("changeSimcardStatus: %s", err)
		return
	}
	err = json.Unmarshal([]byte(request), &res)
	return
}

// BatchChangeSimcardStatus 物联卡批量停复机 此接口不可频繁调用，同号码，相同操作，5分钟内最多调用3次
func BatchChangeSimcardStatus(req *BatchChangeSimcardStatusReq) (res BatchChangeSimcardStatusRes, err error) {
	request, err := ApiRequest("/batchChangeSimcardStatus", req)
	if err != nil {
		log.Fatalf("batchChangeSimcardStatus: %s", err)
		return
	}
	err = json.Unmarshal([]byte(request), &res)
	return
}

// ChangeSimcardLimit 物联卡单卡限速
func ChangeSimcardLimit(req *ChangeSimcardLimitReq) (res Response, err error) {
	request, err := ApiRequest("/changeSimcardLimit", req)
	if err != nil {
		log.Fatalf("changeSimcardLimit: %s", err)
		return
	}
	err = json.Unmarshal([]byte(request), &res)
	return
}

// ManualChangeSimImei 机卡绑定手动更换IMEI号 手动更换IMEI号。相同号码，相同操作，5分钟内最多调用3次。
func ManualChangeSimImei(req *ManualChangeSimImeiReq) (res Response, err error) {
	request, err := ApiRequest("/manualChangeSimImei", req)
	if err != nil {
		log.Fatalf("manualChangeSimImei: %s", err)
		return
	}
	err = json.Unmarshal([]byte(request), &res)
	return
}

// BatchQuerySimFlow 物联网卡批量流量使用情况查询
func BatchQuerySimFlow(req *BatchQuerySimFlowReq) (res BatchQuerySimFlowRes, err error) {
	request, err := ApiRequest("/batchQuerySimFlow", req)
	if err != nil {
		log.Fatalf("batchQuerySimFlow: %s", err)
		return
	}
	err = json.Unmarshal([]byte(request), &res)
	return
}

// QueryResultsAsynchronously 异步执行结果查询
func QueryResultsAsynchronously(req *QueryResultsAsynchronouslyReq) (res QueryResultsAsynchronouslyRes, err error) {
	request, err := ApiRequest("/queryResultsAsynchronously", req)
	if err != nil {
		log.Fatalf("queryResultsAsynchronously: %s", err)
		return
	}
	err = json.Unmarshal([]byte(request), &res)
	return
}

// QueryBasicInfoBatch 物联卡批量查询
func QueryBasicInfoBatch(req *QueryBasicInfoBatchReq) (res QueryBasicInfoBatchRes, err error) {
	request, err := ApiRequest("/queryBasicInfoBatch", req)
	if err != nil {
		log.Fatalf("queryBasicInfoBatch: %s", err)
		return
	}
	err = json.Unmarshal([]byte(request), &res)
	return
}
