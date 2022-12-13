package impl

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"github.com/kalandramo/keyauth/apps/token"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *service) saveToken(tk *token.Token) error {
	if _, err := s.col.InsertOne(context.TODO(), tk); err != nil {
		return exception.NewInternalServerError("inserted token(%s) document error, %s", tk.AccessToken, err)
	}

	return nil
}

func (s *service) updateToken(tk *token.Token) error {
	if _, err := s.col.UpdateOne(context.TODO(), bson.M{"_id": tk.AccessToken}, bson.M{"$set": tk}); err != nil {
		return exception.NewInternalServerError("update token(%s) document error, %s", tk.AccessToken, err)
	}

	return nil
}
