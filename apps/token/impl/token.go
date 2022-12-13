package impl

import (
	"context"
	"fmt"
	"time"

	"github.com/infraboard/mcube/exception"
	"github.com/kalandramo/keyauth/apps/token"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	// 登录错误统一返回信息
	AUTH_ERROR = "user or password not correct"
	// token有效时长
	DefaultTokenDuration = 10 * time.Minute
)

func (s *service) IssueToken(ctx context.Context, req *token.IssueTokenRequest) (*token.Token, error) {
	// 连续登录失败检测
	if err := s.loginBeforeCheck(ctx, req); err != nil {
		return nil, exception.NewBadRequest("安全检测失败, %s", err)
	}

	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate issue token error, %s", err)
	}

	// 根据不同授权模型来做不同的验证
	switch req.GrantType {
	case token.GrantType_PASSWORD:
	default:
		return nil, fmt.Errorf("grant type %s not implemented", req.GrantType)
	}

	return nil, status.Errorf(codes.Unimplemented, "method IssueToken not implemented")
}

func (s *service) ValidateToken(ctx context.Context, req *token.ValidateTokenRequest) (*token.Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateToken not implemented")
}

func (s *service) DescribeToken(ctx context.Context, req *token.DescribeTokenRequest) (*token.Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeToken not implemented")
}

func (s *service) RevolkToken(ctx context.Context, req *token.RevolkTokenRequest) (*token.Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevolkToken not implemented")
}

func (s *service) BlockToken(ctx context.Context, req *token.BlockTokenRequest) (*token.Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockToken not implemented")
}

func (s *service) ChangeNamespace(ctx context.Context, req *token.ChangeNamespaceRequest) (*token.Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeNamespace not implemented")
}

func (s *service) QueryToken(ctx context.Context, req *token.QueryTokenRequest) (*token.Set, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryToken not implemented")
}

func (s *service) DeleteToken(ctx context.Context, req *token.DeleteTokenRequest) (*token.DeleteTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteToken not implemented")
}

func (s *service) loginBeforeCheck(ctx context.Context, req *token.IssueTokenRequest) error {
	// 连续登录失败检测
	if err := s.checker.MaxFailedRetryCheck(ctx, req); err != nil {
		return exception.NewBadRequest("%s", err)
	}

	s.log.Debug("security check complete")
	return nil
}
