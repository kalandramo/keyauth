package user

import (
	"fmt"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/types/ftime"
	"golang.org/x/crypto/bcrypt"
)

const (
	// DefaultExpiresDays 默认多少天无登录系统就冻结该用户
	DefaultExpiresDays = 90
)

// NewProfile todo
func NewProfile() *Profile {
	return &Profile{}
}

// NewUserSet 实例
func NewUserSet() *Set {
	return &Set{
		Items: []*User{},
	}
}

func (s *Set) Add(user *User) {
	s.Items = append(s.Items, user)
}

// New 一个User实例
func New(req *CreateAccountRequest) (*User, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	password, err := NewHashedPassword(req.Password)
	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	u := &User{
		Domain:         req.Domain,
		CreateAt:       ftime.Now().Timestamp(),
		UpdateAt:       ftime.Now().Timestamp(),
		Profile:        req.Profile,
		DepartmentId:   req.DepartmentId,
		Account:        req.Account,
		CreateType:     req.CreateType,
		Type:           req.UserType,
		ExpiresDays:    req.ExpiresDays,
		Description:    req.Description,
		HashedPassword: password,
		Status: &Status{
			Locked: false,
		},
	}

	if req.DepartmentId != "" && req.Profile != nil {
		u.IsInitialized = true
	}

	return u, nil
}

// Desensitize 关键数据脱敏
func (u *User) Desensitize() {
	if u.HashedPassword != nil {
		u.HashedPassword.Password = ""
		u.HashedPassword.History = []string{}
	}
}

// NewHashedPassword 生产hash后的密码对象
func NewHashedPassword(password string) (*Password, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	}

	return &Password{
		Password: string(b),
		CreateAt: ftime.Now().Timestamp(),
		UpdateAt: ftime.Now().Timestamp(),
	}, nil
}

// SetNeedReset 需要被重置
func (p *Password) SetNeedReset(format string, a ...interface{}) {
	p.NeedReset = true
	p.ResetReason = fmt.Sprintf(format, a...)
}
