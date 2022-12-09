package domain

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/infraboard/mcube/exception"
	"github.com/kalandramo/keyauth/apps/user"
	"github.com/kalandramo/keyauth/common/password"
)

func NewDefaultSecuritySetting() *SecuritySetting {
	return &SecuritySetting{
		PasswordSecurity: NewDefaultPasswordSecurity(),
		LoginSecurity:    NewDefaultLoginSecurity(),
	}
}

func NewDefaultPasswordSecurity() *PasswordSecurity {
	return &PasswordSecurity{
		Length:                  8,
		IncludeNumber:           true,
		IncludeLowerLetter:      true,
		IncludeUpperLetter:      false,
		IncludeSymbols:          false,
		RepeatLimit:             1,
		PasswordExpiredDays:     90,
		BeforeExpiredRemindDays: 10,
	}
}

func NewDefaultLoginSecurity() *LoginSecurity {
	return &LoginSecurity{
		ExceptionLock: false,
		ExceptionLockConfig: &ExceptionLockConfig{
			OtherPlaceLogin: true,
			NotLoginDays:    30,
		},
		RetryLock: true,
		RetryLockConfig: &RetryLockConfig{
			RetryLimite:  5,
			LockedMinite: 30,
		},
		IpLimite: false,
		IpLimiteConfig: &IPLimiteConfig{
			Ip: []string{},
		},
	}
}

func (s *SecuritySetting) GetPasswordRepeateLimite() uint {
	if s.PasswordSecurity == nil {
		return 0
	}

	return uint(s.PasswordSecurity.RepeatLimit)
}

func (s *SecuritySetting) Patch(data *SecuritySetting) {
	b, err := json.Marshal(data)
	if err != nil {
		return
	}

	json.Unmarshal(b, s)
}

// Validate 校验对象合法性
func (p *PasswordSecurity) Validate() error {
	return validate.Struct(p)
}

// IsPasswordExpired 根据密码的更新时间判断密码是否过期
func (p *PasswordSecurity) IsPasswordExpired(pass *user.Password) error {
	if p.PasswordExpiredDays == 0 {
		return nil
	}

	detail := p.expiredDetail(time.Unix(pass.UpdateAt/1000, 0))
	if detail > 0 {
		return exception.NewPasswordExired("password expired %s days", detail)
	}

	return nil
}

// expiredDetail 计算密码距离过期的天数，正数表示已过期多少天，负数表示距离过期还有多少天。
func (p *PasswordSecurity) expiredDetail(updateAt time.Time) int {
	updateBefore := int(time.Now().Sub(updateAt).Hours() / 24)

	return int(updateBefore) - int(p.PasswordExpiredDays)
}

// SetPasswordNeedReset
func (p *PasswordSecurity) SetPasswordNeedReset(pass *user.Password) {
	// 密码永不过期, 不需要重置
	if p.PasswordExpiredDays == 0 {
		return
	}

	// 计算密码是否过期
	detail := p.expiredDetail(time.Unix(pass.UpdateAt/1000, 0))
	if detail > 0 {
		pass.SetExpired()
		return
	}

	// 计算是否即将过期, 提醒用户重置
	if -detail < int(p.BeforeExpiredRemindDays) {
		pass.SetNeedReset("密码%d天后过期, 请重置密码", -detail)
	}
}

// Check 校验密码是否符合规则
func (p *PasswordSecurity) Check(pass string) error {
	v := password.NewValidator(pass)

	if ok := v.LengthOK(int(p.Length)); !ok {
		return fmt.Errorf("password length less than %d", p.Length)
	}

	if p.IncludeNumber {
		if ok := v.IncludeNumbers(); !ok {
			return fmt.Errorf("must include number")
		}
	}

	if p.IncludeLowerLetter {
		if ok := v.IncludeLowercaseLetters(); !ok {
			return fmt.Errorf("must include lower letter")
		}
	}

	if p.IncludeUpperLetter {
		if ok := v.IncludeUppercaseLetters(); !ok {
			return fmt.Errorf("must include upper letter")
		}
	}

	if p.IncludeSymbols {
		if ok := v.IncludeSymbols(); !ok {
			return fmt.Errorf("must include symbol")
		}
	}

	return nil
}

// GenRandomPasswordConfig
// 生成密码生成器的规则
func (p *PasswordSecurity) GenRandomPasswordConfig() password.Config {
	return password.Config{
		Length:                  int(p.Length) + 4,
		IncludeSymbols:          true,
		IncludeNumbers:          true,
		IncludeLowercaseLetters: true,
		IncludeUppercaseLetters: true,
	}
}

// LockedMiniteDuration todo
func (r *RetryLockConfig) LockedMiniteDuration() time.Duration {
	return time.Duration(r.LockedMinite) * time.Minute
}
