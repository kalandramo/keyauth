package impl

import (
	"fmt"

	"github.com/kalandramo/keyauth/apps/token"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type describeTokenRequest struct {
	AccessToken  string
	RefreshToken string
}

func NewDescribeTokenRequestWithAccess(at string) *describeTokenRequest {
	return &describeTokenRequest{
		AccessToken: at,
	}
}

func NewDescribeTokenRequestWithRefresh(rt string) *describeTokenRequest {
	return &describeTokenRequest{
		RefreshToken: rt,
	}
}

func NewDescribeTokenRequest(req *token.DescribeTokenRequest) *describeTokenRequest {
	return &describeTokenRequest{
		AccessToken:  req.AccessToken,
		RefreshToken: req.RefreshToken,
	}
}

func (req *describeTokenRequest) String() string {
	return fmt.Sprintf("access_token: %s, refresh_token: %s", req.AccessToken, req.RefreshToken)
}

func (req *describeTokenRequest) FindFilter() bson.M {
	filter := bson.M{}

	if req.AccessToken != "" {
		filter["_id"] = req.AccessToken
	}

	if req.RefreshToken != "" {
		filter["refresh_token"] = req.RefreshToken
	}

	return filter
}

type queryTokenRequest struct {
	*token.QueryTokenRequest
}

func NewQueryTokenRequest(req *token.QueryTokenRequest) *queryTokenRequest {
	return &queryTokenRequest{req}
}

func (req *queryTokenRequest) FindOptions() *options.FindOptions {
	pageSize := int64(req.Page.PageSize)
	skip := int64(req.Page.PageSize) * int64(req.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (req *queryTokenRequest) FindFilter() bson.M {
	filter := bson.M{}

	if req.GrantType != token.GrantType_NULL {
		filter["grant_type"] = req.GrantType
	}

	if req.Account != "" {
		filter["account"] = req.Account
	}

	return filter
}

type deleteTokenRequest struct {
	*token.DeleteTokenRequest
}

func newDeleteTokenRequest(req *token.DeleteTokenRequest) *deleteTokenRequest {
	return &deleteTokenRequest{
		DeleteTokenRequest: req,
	}
}

func (req *deleteTokenRequest) String() string {
	return fmt.Sprintf("access_token: %s",
		req.AccessToken)
}

func (req *deleteTokenRequest) FindFilter() bson.M {
	filter := bson.M{}

	filter["domain"] = req.Domain
	filter["account"] = req.Account

	if len(req.AccessToken) > 0 {
		filter["_id"] = bson.M{"$in": req.AccessToken}
	}

	return filter
}
