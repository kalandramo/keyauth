package domain

import (
	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/types/ftime"
)

var validate = validator.New()

// New 新建一个domain
func New(req *CreateDomainRequest) (*Domain, error) {
	d := &Domain{
		CreateAt:        ftime.Now().Timestamp(),
		UpdateAt:        ftime.Now().Timestamp(),
		Owner:           req.Owner,
		Name:            req.Name,
		Profile:         req.Profile,
		Enabled:         true,
		SecuritySetting: NewDefaultSecuritySetting(),
	}

	return d, nil
}

func NewDefaultDomain() *Domain {
	return &Domain{
		Profile:         &DomainProfile{},
		SecuritySetting: NewDefaultSecuritySetting(),
	}
}

// NewDomainSet Set实例
func NewDomainSet() *Set {
	return &Set{
		Items: []*Domain{},
	}
}

// Length 总个数
func (ds *Set) Length() int {
	return len(ds.Items)
}

// Add 添加Item
func (ds *Set) Add(d *Domain) {
	ds.Items = append(ds.Items, d)
}
