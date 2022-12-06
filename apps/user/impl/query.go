package impl

import (
	"github.com/kalandramo/keyauth/apps/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type queryAccountRequest struct {
	*user.QueryAccountRequest
}

func NewQueryAccountRequest(req *user.QueryAccountRequest) (*queryAccountRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return &queryAccountRequest{
		QueryAccountRequest: req,
	}, nil
}

func (req *queryAccountRequest) FindFilter() bson.M {
	// 默认查询条件
	// bson.M 无序map
	filter := bson.M{
		"type":   req.UserType,
		"domain": req.Domain,
	}

	// 有多个用户
	if len(req.Accounts) > 0 {
		// $in in
		filter["_id"] = bson.M{"$in": req.Accounts}
	}

	// 部门id
	if req.DepartmentId != "" {
		// 列出子账号
		if req.WithAllSub {
			filter["$or"] = bson.A{
				bson.M{
					"department_id": bson.M{"$regex": req.DepartmentId, "$options": "im"},
				},
			}
		} else {
			filter["department_id"] = req.DepartmentId
		}
	}

	// 搜索关键词列表
	if req.Keywords != "" {
		filter["$or"] = bson.A{
			bson.M{"_id": bson.M{"$regex": req.Keywords, "$options": "im"}},
			bson.M{"profile.mobile": bson.M{"$regex": req.Keywords, "$options": "im"}},
			bson.M{"profile.email": bson.M{"$regex": req.Keywords, "$options": "im"}},
		}
	}

	return filter
}

func (req *queryAccountRequest) FindOptions() *options.FindOptions {
	pageSize := int64(req.Page.PageSize)
	skip := int64(req.Page.PageSize) * int64(req.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}
