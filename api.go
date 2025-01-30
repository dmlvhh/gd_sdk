package gd_sdk

import (
	"encoding/json"
	"log"
)

// ChangeSimcardStatus 物联网卡停复机 此接口不可频繁调用，同号码，相同操作，5分钟内最多调用3次。
func (c *Config) ChangeSimcardStatus(req *ChangeSimcardStatusReq) (res *Response, err error) {
	request, err := c.ApiRequest("/ipa/api/v1/wlqd/changeSimcardStatus", req)
	if err != nil {
		log.Printf("changeSimcardStatus: %s", err)
		return
	}
	err = json.Unmarshal([]byte(request), &res)
	return
}

// BatchChangeSimcardStatus 物联卡批量停复机 此接口不可频繁调用，同号码，相同操作，5分钟内最多调用3次
func (c *Config) BatchChangeSimcardStatus(req *BatchChangeSimcardStatusReq) (res *BatchChangeSimcardStatusRes, err error) {
	request, err := c.ApiRequest("/ipa/api/v1/wlqd/batchChangeSimcardStatus", req)
	if err != nil {
		log.Printf("batchChangeSimcardStatus: %s", err)
		return
	}
	err = json.Unmarshal([]byte(request), &res)
	return
}

// ChangeSimcardLimit 物联卡单卡限速
func (c *Config) ChangeSimcardLimit(req *ChangeSimcardLimitReq) (res *Response, err error) {
	request, err := c.ApiRequest("/ipa/api/v1/wlqd/changeSimcardLimit", req)
	if err != nil {
		log.Printf("changeSimcardLimit: %s", err)
		return
	}
	err = json.Unmarshal([]byte(request), &res)
	return
}

// ManualChangeSimImei 机卡绑定手动更换IMEI号 手动更换IMEI号。相同号码，相同操作，5分钟内最多调用3次。
func (c *Config) ManualChangeSimImei(req *ManualChangeSimImeiReq) (res *Response, err error) {
	request, err := c.ApiRequest("/ipa/api/v1/wlqd/manualChangeSimImei", req)
	if err != nil {
		log.Printf("manualChangeSimImei: %s", err)
		return
	}
	err = json.Unmarshal([]byte(request), &res)
	return
}

// BatchQuerySimFlow 物联网卡批量流量使用情况查询
func (c *Config) BatchQuerySimFlow(req *BatchQuerySimFlowReq) (res *BatchQuerySimFlowRes, err error) {
	request, err := c.ApiRequest("/ipa/api/v1/wlqd/batchQuerySimFlow", req)
	if err != nil {
		log.Printf("batchQuerySimFlow: %s", err)
		return
	}
	err = json.Unmarshal([]byte(request), &res)
	return
}

// QueryResultsAsynchronously 异步执行结果查询
func (c *Config) QueryResultsAsynchronously(req *QueryResultsAsynchronouslyReq) (res *QueryResultsAsynchronouslyRes, err error) {
	request, err := c.ApiRequest("/ipa/api/v1/wlqd/queryResultsAsynchronously", req)
	if err != nil {
		log.Printf("queryResultsAsynchronously: %s", err)
		return
	}
	err = json.Unmarshal([]byte(request), &res)
	return
}

// QueryBasicInfoBatch 物联卡批量查询
func (c *Config) QueryBasicInfoBatch(req *QueryBasicInfoBatchReq) (res *QueryBasicInfoBatchRes, err error) {
	request, err := c.ApiRequest("/ipa/api/v1/wlqd/queryBasicInfoBatch", req)
	if err != nil {
		log.Printf("queryBasicInfoBatch: %s", err)
		return
	}
	err = json.Unmarshal([]byte(request), &res)
	return
}
