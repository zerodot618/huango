// Package sms 发送短信
package sms

import (
	"sync"
	"zerodot618/huango/pkg/config"
)

// Message 是短信的结构体
type Message struct {
	Template string
	Data     map[string]string

	Content string
}

// SMS 是我们发送短信的操作类
type SMS struct {
	Driver Driver
}

var driverTypes = map[string]Driver{
	"aliyun":  &Aliyun{},
	"tencent": &Tencent{},
}

// once 单例模式
var once sync.Once

// internalSMS 内部使用的 SMS 对象
var internalSMS *SMS

// NewSMS 单例模式获取
func NewSMS() *SMS {
	once.Do(func() {
		internalSMS = &SMS{
			Driver: driverTypes[config.GetString("sms.default")],
		}
	})

	return internalSMS
}

func (sms *SMS) Send(phone string, message Message) bool {
	return sms.Driver.Send(phone, message, config.GetStringMapString("sms."+config.GetString("sms.default")))
}
