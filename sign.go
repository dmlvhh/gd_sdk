package gd_sdk

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

//const APPID = "1847483747142742018"
//const APPKEY = "QRAXmCI0DuJOvucX2Kh9FiyGZA38Eoba"
//const APPSECRET = "5xYx0jWdp9KE8UVsWK4y2UDBUlc8w14s"

func GetNowUnixTime() int64 {
	return time.Now().UnixNano() / 1000000
}

func GenerateSignature(appKey, appSecret string) string {
	timestamp := GetNowUnixTime()
	signatureString := fmt.Sprintf("%s:%s:%d", appKey, appSecret, timestamp)
	hash := md5.New()
	hash.Write([]byte(signatureString))
	hashString := hex.EncodeToString(hash.Sum(nil))
	return hashString
}
