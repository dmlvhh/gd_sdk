package gd_sdk

type Request struct {
	Sign      string      `json:"sign"`
	Body      interface{} `json:"body"`
	AppId     string      `json:"appId"`
	Timestamp int64       `json:"timestamp"`
}

type ChangeSimcardStatusReq struct {
	Iccid    string `json:"iccid"`
	OperType string `json:"operType"`
}

type BatchChangeSimcardStatusReq struct {
	Iccids   []string `json:"iccids"`
	OperType string   `json:"operType"`
}

type ChangeSimcardLimitReq struct {
	Iccid string `json:"iccid"`
	Code  string `json:"code"`
}

type ManualChangeSimImeiReq struct {
	Iccid   string `json:"iccid"`
	OldImei string `json:"oldImei"`
	NewImei string `json:"newImei"`
}

type BatchQuerySimFlowReq struct {
	Iccids    []string `json:"iccids"`
	QueryDate string   `json:"queryDate"`
}

type QueryResultsAsynchronouslyReq struct {
	SerialNumber string `json:"serialNumber"`
}

type QueryBasicInfoBatchReq struct {
	Iccids []string `json:"iccids"`
}
