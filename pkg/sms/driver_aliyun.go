package sms

import (
	"gohub/pkg/logger"
)

// Aliyun 实现 sms.Driver interface
type Aliyun struct{}

// Send 实现 sms.Driver interface 的 Send 方法
func (s *Aliyun) Send(phone string, message Message, config map[string]string) bool {
	logger.DebugJSON("短信[阿里云]", "配置信息", message.Data)
	return true
}
