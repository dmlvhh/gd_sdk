package gd_sdk

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func ApiRequest(url string, data any) (res string, err error) {

	//appKey := "QRAXmCI0DuJOvucX2Kh9FiyGZA38Eoba"
	//appSecret := "5xYx0jWdp9KE8UVsWK4y2UDBUlc8w14s"
	//sign := GenerateSignature(gdSdk.AppKey, gdSdk.AppSecret)
	//url = "https://122.192.50.42:11060/ipa/api" + url
	//fmt.Println(gdSdk)
	reqData, err := json.Marshal(Request{
		Sign:      gdSdk.Sign,
		Body:      data,
		AppId:     gdSdk.AppID,
		Timestamp: GetNowUnixTime(),
	})
	if err != nil {
		log.Printf("Error Marshal: %s", err)
		return
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // 不安全，仅用于测试
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Post(gdSdk.ApiURL+url, "application/json", bytes.NewBuffer(reqData))
	if err != nil {
		log.Printf("Error making POST request: %s", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %s", err)
		return
	}
	return string(body), nil
}
