package token

import (
	"fmt"
	"time"
)

// NewDefaultToken todo
func NewDefaultToken() *Token {
	return &Token{}
}

// HasNamespace
func (t *Token) HasNamespace(ns string) bool {
	for _, v := range t.AvailableNamespace {
		if v == ns {
			return true
		}
	}

	return false
}

// IsRefresh todo
func (t *Token) IsRefresh() bool {
	return t.GrantType == GrantType_REFRESH
}

// IsOwner todo
func (t *Token) IsOwner(account string) bool {
	return t.Account == account
}

// BlockMessage todo
func (t *Token) BlockMessage() string {
	if !t.IsBlock {
		return ""
	}

	return fmt.Sprintf("token blocked at %d, reason: %s", t.BlockAt, t.BlockReason)
}

// IsAvailable 判断一个token的可用性
func (t *Token) IsAvailable() error {
	if t.IsBlock {
		return fmt.Errorf("token is blocked")
	}

	if t.CheckAccessIsExpired() {
		return fmt.Errorf("token is expired")
	}

	return nil
}

// CheckAccessIsExpired 检测token是否过期
func (t *Token) CheckAccessIsExpired() bool {
	if t.AccessExpiredAt == 0 {
		return false
	}

	return time.Unix(t.AccessExpiredAt/1000, 0).Before(time.Now())
}

// WithRemoteIP todo
func (t *Token) WithRemoteIP(ip string) {
	t.RemoteIp = ip
}

// GetRemoteIP todo
func (t *Token) GetRemoteIP() string {
	return t.RemoteIp
}

// WithUerAgent todo
func (t *Token) WithUerAgent(ua string) {
	t.UserAgent = ua
}

// CheckRefreshIsExpired 检测刷新token是否过期
func (t *Token) CheckRefreshIsExpired() bool {
	if t.RefreshExpiredAt == 0 {
		return false
	}

	return time.Unix(t.RefreshExpiredAt/1000, 0).Before(time.Now())
}

// CheckTokenApplication 判断token是否属于该应用
func (t *Token) CheckTokenApplication(applicationId string) error {
	if t.ApplicationId != applicationId {
		return fmt.Errorf("the token is not issue by this application %s", applicationId)
	}

	return nil
}

// Desensitize 数据脱敏
func (t *Token) Desensitize() {
	t.RefreshToken = ""
}

// EndAt token结束时间
func (t *Token) EndAt() int64 {
	if t.IsBlock {
		return t.BlockAt
	}

	if t.CheckAccessIsExpired() {
		return t.AccessExpiredAt
	}

	if t.CheckRefreshIsExpired() {
		return t.RefreshExpiredAt
	}

	return 0
}

// NewTokenSet 实例化
func NewTokenSet() *Set {
	return &Set{
		Items: []*Token{},
	}
}

func (s *Set) Add(item *Token) {
	s.Items = append(s.Items, item)
}
