package domain

import (
	"encoding/json"
	"fmt"

	"github.com/infraboard/mcube/http/request"
	"github.com/kalandramo/keyauth/common/types"
)

// NewCreateDomainRequest todo
func NewCreateDomainRequest() *CreateDomainRequest {
	return &CreateDomainRequest{
		Profile: &DomainProfile{},
	}
}

// Validate 校验请求是否合法
func (req *CreateDomainRequest) Validate() error {
	return validate.Struct(req)
}

// NewQueryDomainRequest 查询domian列表
func NewQueryDomainRequest(page *request.PageRequest) *QueryDomainRequest {
	return &QueryDomainRequest{
		Page: page,
	}
}

// Validate 校验请求合法
func (req *QueryDomainRequest) Validate() error {
	return nil
}

// NewDescribeDomainRequest 查询详情请求
func NewDescribeDomainRequest() *DescribeDomainRequest {
	return &DescribeDomainRequest{}
}

// NewDescribeDomainRequestWithName 查询详情请求
func NewDescribeDomainRequestWithName(name string) *DescribeDomainRequest {
	return &DescribeDomainRequest{
		Name: name,
	}
}

// Validate todo
func (req *DescribeDomainRequest) Validate() error {
	if req.Name == "" {
		return fmt.Errorf("name required")
	}

	return nil
}

// NewPutDomainRequest todo
func NewPutDomainRequest() *UpdateDomainInfoRequest {
	return &UpdateDomainInfoRequest{
		UpdateMode: types.UpdateMode_PUT,
	}
}

// Validate 更新校验
func (req *UpdateDomainInfoRequest) Validate() error {
	return validate.Struct(req)
}

// NewPatchDomainRequest todo
func NewPatchDomainRequest() *UpdateDomainInfoRequest {
	return &UpdateDomainInfoRequest{
		UpdateMode: types.UpdateMode_PATCH,
	}
}

// NewPutDomainSecurityRequest todo
func NewPutDomainSecurityRequest() *UpdateDomainSecurityRequest {
	return &UpdateDomainSecurityRequest{
		UpdateMode:      types.UpdateMode_PUT,
		SecuritySetting: NewDefaultSecuritySetting(),
	}
}

// Validate todo
func (req *UpdateDomainSecurityRequest) Validate() error {
	return validate.Struct(req)
}

// NewDeleteDomainRequestByName todo
func NewDeleteDomainRequestByName(name string) *DeleteDomainRequest {
	return &DeleteDomainRequest{
		Name: name,
	}
}

// Patch todo
func (req *DomainProfile) Patch(data *DomainProfile) {
	patchData, _ := json.Marshal(data)
	json.Unmarshal(patchData, req)
}
