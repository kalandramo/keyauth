package user

import (
	"encoding/json"
	"fmt"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/types/ftime"
	"golang.org/x/crypto/bcrypt"
)

const (
	// DefaultExpiresDays 默认多少天无登录系统就冻结该用户
	DefaultExpiresDays = 90
)

// NewDefaultUser 实例
func NewDefaultUser() *User {
	return &User{
		Profile: NewProfile(),
		Status: &Status{
			Locked: false,
		},
	}
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

// Block 锁用户
func (u *User) Block(reason string) {
	u.Status.Locked = true
	u.Status.LockedReason = reason
	u.Status.LockedTime = ftime.Now().Timestamp()
}

func (u *User) UnBlock() error {
	if !u.Status.Locked {
		return exception.NewBadRequest("user %s not block, don't unblock", u.Account)
	}

	u.Status.Locked = false
	u.Status.UnlockTime = ftime.Now().Timestamp()
	return nil
}

// ChangePassword 修改用户密码
func (u *User) ChangePassword(old, new string, maxHistory uint, needReset bool) error {
	// 确认旧密码
	if err := u.HashedPassword.CheckPassword(old); err != nil {
		return err
	}

	// 修改新密码
	newPass, err := NewHashedPassword(new)
	if err != nil {
		return exception.NewBadRequest(err.Error())
	}
	u.HashedPassword.Update(newPass, maxHistory, needReset)

	return nil
}

// HasDepartment todo
func (u *User) HasDepartment() bool {
	return u.DepartmentId != ""
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

// SetExpired 密码过期
func (p *Password) SetExpired() {
	p.IsExpired = true
}

// SetNeedReset 需要被重置
func (p *Password) SetNeedReset(format string, a ...interface{}) {
	p.NeedReset = true
	p.ResetReason = fmt.Sprintf(format, a...)
}

// CheckPassword 判断password 是否正确
func (p *Password) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(p.Password), []byte(password))
	if err != nil {
		return exception.NewUnauthorized("user or password not connect")
	}

	return nil
}

// IsHistory 检测是否是历史密码
func (p *Password) IsHistory(password string) bool {
	for _, pass := range p.History {
		err := bcrypt.CompareHashAndPassword([]byte(pass), []byte(password))
		if err == nil {
			return true
		}
	}

	return false
}

// HistoryCount 保存了几个历史密码
func (p *Password) HistoryCount() int {
	return len(p.History)
}

func (p *Password) rotaryHistory(maxHistory uint) {
	if uint(p.HistoryCount()) < maxHistory {
		p.History = append(p.History, p.Password)
	} else {
		remainHistory := p.History[:maxHistory]
		p.History = []string{p.Password}
		p.History = append(p.History, remainHistory...)
	}
}

// Update 更新密码
func (p *Password) Update(new *Password, maxHistory uint, needReset bool) {
	p.rotaryHistory(maxHistory)
	p.Password = new.Password
	p.NeedReset = needReset
	p.UpdateAt = ftime.Now().Timestamp()
	if !needReset {
		p.ResetReason = ""
	}
}

// NewProfile todo
func NewProfile() *Profile {
	return &Profile{}
}

// ValidateInitialized 判断初始化数据是否准备好了
func (pf *Profile) ValidateInitialized() error {
	if pf.Email != "" && pf.Phone != "" {
		return nil
	}

	return fmt.Errorf("email and phone required when initial")
}

// Patch todo
func (pf *Profile) Patch(data *Profile) {
	md := NewProfile()
	patchData, _ := json.Marshal(data)
	oldData, _ := json.Marshal(pf)
	json.Unmarshal(oldData, md)
	json.Unmarshal(patchData, md)
	*pf = *md
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
