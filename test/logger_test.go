package test

import (
	"testing"

	"github.com/sirupsen/logrus"
	"webhook.com/app"
)

//日志测试
func TestLogger(t *testing.T) {
	// Info 级别日志
	app.Logger().WithFields(logrus.Fields{
		"name": "Info Test",
	}).Info("记录一下日志", "Info")
	// Error级别日志
	app.Logger().WithFields(logrus.Fields{
		"name": "Error Test",
	}).Error("记录错误日志", "Error")
	// Warn级别日志
	app.Logger().WithFields(logrus.Fields{
		"name": "警告",
	}).Warn("记录警告日志", "Warn")
	// Debug级别日志
	app.Logger().WithFields(logrus.Fields{
		"name": "Debug Test",
	}).Debug("记录Debug日志", "Debug")
}
