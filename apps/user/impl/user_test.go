package impl_test

import (
	"context"
	"testing"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/kalandramo/keyauth/apps/user"
	"github.com/kalandramo/keyauth/conf"
	"github.com/stretchr/testify/assert"
)

var ins user.ServiceServer

func TestCreateAccount(t *testing.T) {
	should := assert.New(t)

	req := user.NewCreateAccountRequest()
	req.Domain = "ck"
	req.Account = "kalandramo"
	req.Password = "12345678"

	user, err := ins.CreateAccount(context.Background(), req)
	if should.NoError(err) {
		t.Log(user)
	}
}

func TestQueryAccount(t *testing.T) {
	should := assert.New(t)

	req := user.NewQueryAccountRequest()
	req.Domain = "ck"

	userSet, err := ins.QueryAccount(context.Background(), req)
	if should.NoError(err) {
		t.Log(userSet)
	}
}

func init() {
	// 通过环境变量加载测试配置
	if err := conf.LoadConfigFromEnv(); err != nil {
		panic(err)
	}

	// 全局日志对象初始化
	zap.DevelopmentSetup()

	// 初始化所有实例
	if err := app.InitAllApp(); err != nil {
		panic(err)
	}

	ins = app.GetGrpcApp(user.AppName).(user.ServiceServer)
}
