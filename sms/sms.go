package sms

import (
	"errors"
	"strings"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

// defaultSmsConfig 短信信息
var defaultSmsConfig = SmsConfig{}

// SmsConfig 短信配置信息
type SmsConfig struct {
	RegionID        string
	SignName        string
	AccessKeyID     string
	AccessKeySecret string
}

// Options 参数
type Options func(*SmsConfig)

// SetRegionID 短信发送地
func SetRegionID(regionID string) Options {
	return func(s *SmsConfig) {
		s.RegionID = regionID
	}
}

// SetSignName 模板签名
func SetSignName(signName string) Options {
	return func(s *SmsConfig) {
		s.SignName = signName
	}
}

// SetAccessKeyID 设置key
func SetAccessKeyID(accessKeyID string) Options {
	return func(s *SmsConfig) {
		s.AccessKeyID = accessKeyID
	}
}

// SetAccessKeySecret 秘钥
func SetAccessKeySecret(accessKeySecret string) Options {
	return func(s *SmsConfig) {
		s.AccessKeySecret = accessKeySecret
	}
}

// NewSms  实例化sms
func NewSms(ops ...Options) SmsConfig {
	for _, o := range ops {
		o(&defaultSmsConfig)
	}
	return defaultSmsConfig
}

// SendSms 短信发送
func (s SmsConfig) SendSms(phoneNumbers []string, templateCode, templateParam string) (respID string, err error) {
	conf := defaultSmsConfig
	client, err := dysmsapi.NewClientWithAccessKey(conf.RegionID, conf.AccessKeyID, conf.AccessKeySecret)
	if err != nil {
		return
	}
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = strings.Join(phoneNumbers, ",")
	request.SignName = conf.SignName
	request.TemplateCode = templateCode
	request.TemplateParam = templateParam

	response, err := client.SendSms(request)
	if err != nil {
		return
	}
	if strings.ToUpper(response.Code) != "OK" {
		err = errors.New(response.Message)
		return response.Code, err
	}
	return response.RequestId, nil
}
