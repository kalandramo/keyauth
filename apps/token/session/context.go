package session

import "github.com/kalandramo/keyauth/apps/token"

// Session 请求上下文信息
type Session struct {
	tk *token.Token
}

// NewSession todo
func NewSession() *Session {
	return &Session{}
}

// WithToken 携带token
func (s *Session) WithToken(tk *token.Token) {
	s.tk = tk
}

// WithTokenGetter geter
func (s *Session) WithTokenGetter(gt Getter) {
	s.tk = gt.Getter()
}

// GetToken 获取token
func (s *Session) GetToken() *token.Token {
	return s.tk
}

// GetAccount todo
func (s *Session) GetAccount() string {
	if s.tk == nil {
		return "Nil"
	}

	return s.tk.Account
}

type Getter interface {
	Getter() *token.Token
}
