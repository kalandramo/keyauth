package impl

import (
	"github.com/kalandramo/keyauth/apps/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type queryDomain struct {
	*domain.QueryDomainRequest
}

func newQueryDomainRequest(req *domain.QueryDomainRequest) *queryDomain {
	return &queryDomain{
		QueryDomainRequest: req,
	}
}

func (r *queryDomain) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryDomain) FindFilter() bson.M {
	filter := bson.M{}

	filter["owner"] = r.Owner
	return filter
}

type descDomain struct {
	*domain.DescribeDomainRequest
}

func newDescDomainRequest(req *domain.DescribeDomainRequest) (*descDomain, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return &descDomain{req}, nil
}

func (r *descDomain) FindFilter() bson.M {
	filter := bson.M{}

	if r.Name != "" {
		filter["_id"] = r.Name
	}

	return filter
}
