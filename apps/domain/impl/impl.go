package impl

import (
	"context"

	"github.com/infraboard/keyauth/apps/domain"
	"github.com/kalandramo/keyauth/conf"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

var srv = &service{}

type service struct {
	col           *mongo.Collection
	enableCache   bool
	notifyCachPre string
	domain.UnimplementedServiceServer
}

func (s *service) Config() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}
	dc := db.Collection("domain")

	indexs := []mongo.IndexModel{
		{
			Keys:    bsonx.Doc{{Key: "name", Value: bsonx.Int32(-1)}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bsonx.Doc{{Key: "ldap_config.base_dn", Value: bsonx.Int32(-1)}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bsonx.Doc{{Key: "create_at", Value: bsonx.Int32(-1)}},
		},
	}

	_, err = dc.Indexes().CreateMany(context.Background(), indexs)
	if err != nil {
		return err
	}

	s.col = dc

	return nil
}
