package impl

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"github.com/kalandramo/keyauth/apps/user"
)

func (s *service) saveAccount(user *user.User) error {
	if _, err := s.col.InsertOne(context.TODO(), user); err != nil {
		return exception.NewInternalServerError("inserted user(%s) document error,%s", user.Account, err)
	}

	return nil
}

func (s *service) queryAccount(ctx context.Context, req *queryAccountRequest) (*user.Set, error) {
	userSet := user.NewUserSet()

	if !req.SkipItems {
		s.log.Debugf("find filter: %s", req.FindFilter())
		cursor, err := s.col.Find(ctx, req.FindFilter(), req.FindOptions())
		// // 查询出该空间下的用户列表
		// if req.NamespaceId != "" {
		// 	s.queryNamespacePolicy()
		// }
		if err != nil {
			return nil, exception.NewInternalServerError("find user error, error is %s", err)
		}

		// 循环
		for cursor.Next(context.TODO()) {
			u := new(user.User)
			if err := cursor.Decode(u); err != nil {
				return nil, exception.NewInternalServerError("decode user error, error is %s", err)
			}

			// 关键数据脱敏
			u.Desensitize()
			userSet.Add(u)
		}
	}

	count, err := s.col.CountDocuments(context.TODO(), req.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get device count error,error is %s", err)
	}
	userSet.Total = count

	return userSet, nil
}
