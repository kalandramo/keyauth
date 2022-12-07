package user

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/http/request"
)

var validate = validator.New()

// NewCreateAccountRequest 构造创建用户请求参数
func NewCreateAccountRequest() *CreateAccountRequest {
	return &CreateAccountRequest{
		Profile:     NewProfile(),
		ExpiresDays: DefaultExpiresDays,
	}
}

// Validate 校验创建用户请求参数
func (c *CreateAccountRequest) Validate() error {
	return validate.Struct(c)
}

// NewCreateAccountRequest 构造查询用户请求参数,grpc使用
func NewQueryAccountRequest() *QueryAccountRequest {
	return &QueryAccountRequest{
		Page:           request.NewPageRequest(10, 1),
		WithDepartment: false,
		SkipItems:      false,
	}
}

// NewQueryAccountRequestFromHTTP 构造查询用户请求参数,http使用
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
