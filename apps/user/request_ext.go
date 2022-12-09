package user

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/http/request"
)

var validate = validator.New()

// NewCreateAccountRequest 创建用户请求参数
func NewCreateAccountRequest() *CreateAccountRequest {
	return &CreateAccountRequest{
		Profile:     NewProfile(),
		ExpiresDays: DefaultExpiresDays,
	}
}

// NewCreateUserRequestWithLDAPSync todo
func NewCreateUserRequestWithLDAPSync(username, password string) *CreateAccountRequest {
	req := NewCreateAccountRequest()
	req.CreateType = CreateType_LADP_SYNC
	req.Account = username
	req.Password = password

	return req
}

// NewCreateUserRequestWithWXWORKSync todo
func NewCreateUserRequestWithWXWORKSync(username, password string) *CreateAccountRequest {
	req := NewCreateAccountRequest()
	req.CreateType = CreateType_WXWORK_SYNC
	req.Account = username
	req.Password = password

	return req
}

// Validate 校验创建用户请求参数
func (c *CreateAccountRequest) Validate() error {
	return validate.Struct(c)
}

// NewCreateAccountRequest 查询用户请求参数,grpc使用
func NewQueryAccountRequest() *QueryAccountRequest {
	return &QueryAccountRequest{
		Page:           request.NewPageRequest(10, 1),
		WithDepartment: false,
		SkipItems:      false,
	}
}

// NewQueryAccountRequestFromHTTP 查询用户请求参数,http使用
func NewQueryAccountRequestFromHTTP(r *http.Request) *QueryAccountRequest {
	qs := r.URL.Query()

	return &QueryAccountRequest{
		Page:     request.NewPageRequestFromHTTP(r),
		Keywords: qs.Get("keywords"),
		Domain:   qs.Get("domain"),
	}
}

// Validate 校验查询用户请求参数
func (q *QueryAccountRequest) Validate() error {
	return nil
}

// NewDescriptAccountRequest 查询详情请求
func NewDescriptAccountRequest() *DescribeAccountRequest {
	return &DescribeAccountRequest{}
}

// NewDescriptAccountRequestWithAccount 查询详情请求
func NewDescriptAccountRequestWithAccount(account string) *DescribeAccountRequest {
	return &DescribeAccountRequest{Account: account}
}

// Validate 校验详情查询
func (req *DescribeAccountRequest) Validate() error {
	if req.Account == "" {
		return fmt.Errorf("account is required")
	}
	return nil
}
