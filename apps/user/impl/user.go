package impl

import (
	"context"

	"github.com/kalandramo/keyauth/apps/user"
)

// CreateAccount 创建用户
func (s *service) CreateAccount(ctx context.Context, req *user.CreateAccountRequest) (*user.User, error) {
	u, err := user.New(req)
	if err != nil {
		return nil, err
	}

	// 如果是管理员创建的账号需要用户自己重置密码
	if u.CreateType.IsIn(user.CreateType_DOMAIN_CREATED) {
		u.HashedPassword.SetNeedReset("admin created user need reset when first login")
	}

	if err := s.saveAccount(u); err != nil {
		return nil, err
	}

	u.HashedPassword = nil

	return u, nil
}

// QueryAccount 查询用户
func (s *service) QueryAccount(ctx context.Context, req *user.QueryAccountRequest) (*user.Set, error) {
	// 校验请求
	r, err := NewQueryAccountRequest(req)
	if err != nil {
		return nil, err
	}

	return s.queryAccount(ctx, r)
}
