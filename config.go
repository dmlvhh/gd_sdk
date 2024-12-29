package gd_sdk

import "encoding/json"

type Config struct {
	ApiURL    string `json:"api_url"`
	AppID     string `json:"app_id"`
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	Sign      string `json:"sign"`
}

// NewConfig 创建一个新的配置实例
func NewConfig(conf *Config) *Config {
	sign := GenerateSignature(conf.AppKey, conf.AppSecret)
	return &Config{
		ApiURL:    conf.ApiURL,
		AppID:     conf.AppID,
		AppKey:    conf.AppKey,
		AppSecret: conf.AppSecret,
		Sign:      sign,
	}
}
func (c *Config) String() string {
	data, _ := json.MarshalIndent(c, "", "  ")
	return string(data)
}
