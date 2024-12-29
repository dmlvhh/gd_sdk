package gd_sdk

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type BatchChangeSimcardStatusRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		SerialNumber string   `json:"serialNumber"`
		FailList     []string `json:"failList"`
	} `json:"data"`
}

type BatchQuerySimFlowRes struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Data struct {
		ResultList []struct {
			Iccid   string `json:"iccid"`
			ApnList []struct {
				ApnName string `json:"apnName"`
				Usage   string `json:"usage"`
			} `json:"apnList"`
		} `json:"resultList"`
		FailList []string `json:"failList"`
	} `json:"data"`
}

type QueryResultsAsynchronouslyRes struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Data struct {
		ResultList []struct {
			Iccid   string `json:"iccid"`
			Status  string `json:"status"`
			Message string `json:"message"`
		} `json:"resultList"`
	} `json:"data"`
}

type QueryBasicInfoBatchRes struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Data struct {
		InfoList []struct {
			Iccid      string `json:"iccid"`
			OpenDate   string `json:"openDate"`
			ActiveDate string `json:"activeDate"`
			IsRealName string `json:"isRealName"`
			Status     string `json:"status"`
			LifeCycle  string `json:"lifeCycle"`
		} `json:"infoList"`
		FailList []string `json:"failList"`
	} `json:"data"`
}
