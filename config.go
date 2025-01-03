package gd_sdk

import "encoding/json"

type Config struct {
	ApiURL    string `json:"api_url"`
	AppID     string `json:"app_id"`
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	Sign      string `json:"sign"`
}

func NewConfig(apiUrl, appID, appKey, appSecret string) *Config {
	sign := GenerateSignature(appKey, appSecret)
	return &Config{
		ApiURL:    apiUrl,
		AppID:     appID,
		AppKey:    appKey,
		AppSecret: appSecret,
		Sign:      sign,
	}
}

func (c *Config) String() string {
	data, _ := json.MarshalIndent(c, "", "  ")
	return string(data)
}
