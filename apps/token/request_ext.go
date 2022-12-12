package token

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/request"
)

var validate = validator.New()

// NewIssueTokenRequest 默认请求
func NewIssueTokenRequest() *IssueTokenRequest {
	return &IssueTokenRequest{}
}

// NewIssueTokenByPassword
func NewIssueTokenByPassword(clientId, clientSecret, user, pass string) *IssueTokenRequest {
	req := NewIssueTokenRequest()
	req.ClientId = clientId
	req.ClientSecret = clientSecret
	req.Username = user
	req.Password = pass
	req.GrantType = GrantType_PASSWORD
	req.RemoteIp = "127.0.0.1"

	return req
}

func (req *IssueTokenRequest) Validate() error {
	switch req.GrantType {
	case GrantType_PASSWORD:
		if req.Username == "" || req.Password == "" {
			return fmt.Errorf("use %s grant type, username and password required", GrantType_PASSWORD)
		}
	case GrantType_REFRESH:
		if req.AccessToken == "" {
			return fmt.Errorf("use %s grant type, access_token required", GrantType_REFRESH)
		}
		if req.RefreshToken == "" {
			return fmt.Errorf("use %s grant type, refresh_token required", GrantType_REFRESH)
		}
	case GrantType_ACCESS:
		if req.AccessToken == "" {
			return fmt.Errorf("use %s grant type, access_token required", GrantType_ACCESS)
		}
	case GrantType_LDAP:
		if req.Username == "" || req.Password == "" {
			return fmt.Errorf("use %s grant type, username and password required", GrantType_LDAP)
		}
	case GrantType_CLIENT:
	case GrantType_AUTH_CODE:
		if req.AuthCode == "" {
			return fmt.Errorf("use %s grant type, code required", GrantType_AUTH_CODE)
		}
	case GrantType_WECHAT_WORK:
		if req.State == "" || req.Service == "" {
			return fmt.Errorf("use %s grant type, state required", GrantType_WECHAT_WORK)
		}
	default:
		return fmt.Errorf("grant type %s not implemented", req.GrantType)
	}

	return nil
}

// AbnormalUserCheckKey
func (req *IssueTokenRequest) AbnormalUserCheckKey() string {
	return "abnormal_" + req.Username
}

// WithUserAgent
func (req *IssueTokenRequest) WithUserAgent(userAgent string) {
	req.UserAgent = userAgent
}

// WithRemoteIPFromHTTP
func (req *IssueTokenRequest) WithRemoteIPFromHTTP(r *http.Request) {
	req.RemoteIp = request.GetRemoteIP(r)
}

// WithRemoteIP
func (req *IssueTokenRequest) WithRemoteIP(ip string) {
	req.RemoteIp = ip
}

func (req *IssueTokenRequest) IsLoginRequest() bool {
	if req.GrantType.Equal(GrantType_ACCESS) {
		return false
	}

	return true
}

// GetDomainNameFromAccount
func (req *IssueTokenRequest) GetDomainNameFromAccount() string {
	d := strings.Split(req.Username, "@")
	if len(d) == 2 {
		return d[1]
	}

	return ""
}

// NewBlockTokenRequest todo
func NewBlockTokenRequest(accessToken string, bt BlockType, reason string) *BlockTokenRequest {
	return &BlockTokenRequest{
		AccessToken: accessToken,
		BlockType:   bt,
		BlockReason: reason,
	}
}

// NewQueryDepartmentRequestFromHTTP 列表查询请求
func NewQueryDepartmentRequestFromHTTP(r *http.Request) (*QueryTokenRequest, error) {
	req := NewQueryTokenRequest(request.NewPageRequestFromHTTP(r))

	qs := r.URL.Query()
	gt := qs.Get("grant_type")
	grant_type, err := ParseGrantTypeFromString(gt)
	if err != nil {
		return nil, err
	}
	req.GrantType = grant_type

	return req, nil
}

// NewQueryTokenRequest
func NewQueryTokenRequest(page *request.PageRequest) *QueryTokenRequest {
	return &QueryTokenRequest{
		Page: page,
	}
}

// NewDescribeTokenRequestWithAccessToken
func NewDescribeTokenRequestWithAccessToken(at string) *DescribeTokenRequest {
	req := NewDescribeTokenRequest()
	req.AccessToken = at
	return req
}

// NewDescribeTokenRequest
func NewDescribeTokenRequest() *DescribeTokenRequest {
	return &DescribeTokenRequest{}
}

// Validate 校验
func (req *DescribeTokenRequest) Validate() error {
	if err := validate.Struct(req); err != nil {
		return err
	}

	if req.AccessToken == "" && req.RefreshToken == "" {
		return fmt.Errorf("describe token request validate error,access_token and refresh_token required one")
	}

	return nil
}

// NewRevolkTokenRequest 撤销Token请求
func NewRevolkTokenRequest(clientId, clientSecret string) *RevolkTokenRequest {
	return &RevolkTokenRequest{
		ClientId:      clientId,
		ClientSecret:  clientSecret,
		LogoutSession: true,
	}
}

// Validate
func (req *RevolkTokenRequest) Validate() error {
	if err := validate.Struct(req); err != nil {
		return err
	}

	return nil
}

// NewValidateTokenRequest
func NewValidateTokenRequest() *ValidateTokenRequest {
	return &ValidateTokenRequest{}
}

// Validate 校验参数
func (req *ValidateTokenRequest) Validate() error {
	if err := validate.Struct(req); err != nil {
		return err
	}

	if req.AccessToken == "" && req.RefreshToken == "" {
		return fmt.Errorf("access_token and refresh_token required one")
	}

	return nil
}

// MakeDescribeTokenRequest
func (req *ValidateTokenRequest) MakeDescribeTokenRequest() *DescribeTokenRequest {
	descReq := NewDescribeTokenRequest()
	descReq.AccessToken = req.AccessToken
	descReq.RefreshToken = req.RefreshToken

	return descReq
}

// MakeDescribeTokenRequest
func (req *RevolkTokenRequest) MakeDescribeTokenRequest() *DescribeTokenRequest {
	descReq := NewDescribeTokenRequest()
	descReq.AccessToken = req.AccessToken
	descReq.RefreshToken = req.RefreshToken

	return descReq
}

func NewDeleteTokenRequest() *DeleteTokenRequest {
	return &DeleteTokenRequest{}
}

func (req *DeleteTokenRequest) Validate() error {
	if len(req.AccessToken) == 0 {
		return exception.NewBadRequest("delete access token array need")
	}

	return nil
}

func NewDeleteTokenResponse() *DeleteTokenResponse {
	return &DeleteTokenResponse{}
}

func NewChangeNamespaceRequest() *ChangeNamespaceRequest {
	return &ChangeNamespaceRequest{}
}

func (req *ChangeNamespaceRequest) Validate() error {
	return validate.Struct(req)
}
