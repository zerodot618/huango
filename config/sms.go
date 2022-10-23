// Package config 站点配置信息
package config

import "zerodot618/huango/pkg/config"

func init() {
	config.Add("sms", func() map[string]interface{} {
		return map[string]interface{}{
			// 默认使用
			"default": config.Env("SMS_TYPE", "aliyun"),

			// 默认是阿里云的测试 sign_name 和 template_code
			"aliyun": map[string]interface{}{
				"access_key_id":     config.Env("SMS_ACCESS_ID"),
				"access_key_secret": config.Env("SMS_ACCESS_SECRET"),
				"sign_name":         config.Env("SMS_SIGN_NAME", "阿里云短信测试"),
				"template_code":     config.Env("SMS_TEMPLATE_CODE", "SMS_154950909"),
			},
			// 腾讯云
			"tencent": map[string]interface{}{
				"access_key_id":     config.Env("SMS_ACCESS_ID"),
				"access_key_secret": config.Env("SMS_ACCESS_SECRET"),
				"sign_name":         config.Env("SMS_SIGN_NAME"),
				"template_code":     config.Env("SMS_TEMPLATE_CODE"),
				"sdk_app_id":        config.Env("SMS_SDK_APP_ID"),
			},
		}
	})
}
