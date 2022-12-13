package security

import (
	"context"

	"github.com/kalandramo/keyauth/apps/token"
)

// Checker 安全检测
type Checker interface {
	MaxTryChecker
	// ExceptionLockChecker
	// IPProtectChecker
}

// MaxTryChecker todo 失败重试限制
type MaxTryChecker interface {
	MaxFailedRetryCheck(context.Context, *token.IssueTokenRequest) error
	UpdateFailedRetry(context.Context, *token.IssueTokenRequest) error
}

// ExceptionLockChecker 异地登录限制
type ExceptionLockChecker interface {
	OtherPlaceLoggedInCheck(context.Context, *token.Token) error
	NotLoginDaysCheck(context.Context, *token.Token) error
}

// IPProtectChecker todo
type IPProtectChecker interface {
	IPProtectCheck(context.Context, *token.IssueTokenRequest) error
}
