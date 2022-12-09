package password

import "regexp"

// Validator 密码校验器
type Validater interface {
	Reset(pass string)
	IncludeNumbers() bool
	IncludeLowercaseLetters() bool
	IncludeUppercaseLetters() bool
	IncludeSymbols() bool
	LengthOK(limit int) bool
}

// Validator 校验密码强度
type Validator struct {
	pass      string
	numReg    string
	lowerReg  string
	upperReg  string
	symbolReg string
}

// NewValidator 生成校验器
func NewValidator(pass string) *Validator {
	return &Validator{
		pass:      pass,
		numReg:    `[0-9]{1}`,
		lowerReg:  `[a-z]{1}`,
		upperReg:  `[A-Z]{1}`,
		symbolReg: `[!@#~$%^&*()+|_]{1}`,
	}
}

func (v *Validator) Reset(pass string) {
	v.pass = pass
}

// IncludeNumbers 是否包含数字
func (v *Validator) IncludeNumbers() bool {
	return v.match(v.numReg)
}

// IncludeLowercaseLetters 是否包含小写字母
func (v *Validator) IncludeLowercaseLetters() bool {
	return v.match(v.lowerReg)
}

// IncludeUppercaseLetters 是否包含大写字母
func (v *Validator) IncludeUppercaseLetters() bool {
	return v.match(v.upperReg)
}

// IncludeSymbols 是否包含特殊字符
func (v *Validator) IncludeSymbols() bool {
	return v.match(v.symbolReg)
}

// LengthOK 长度是否合法
func (v *Validator) LengthOK(limit int) bool {
	return len(v.pass) >= limit
}

func (v *Validator) match(reg string) bool {
	if matched, err := regexp.MatchString(reg, v.pass); !matched || err != nil {
		return false
	}

	return true
}
