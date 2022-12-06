package impl

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/kalandramo/keyauth/apps/user"
	"github.com/kalandramo/keyauth/conf"
)

// Service 服务实例
var svr = &service{}

type service struct {
	col *mongo.Collection
	log logger.Logger
	user.UnimplementedServiceServer
}

func (s *service) Config() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}
	mc := db.Collection(s.Name())

	indexList := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{Key: "create_at", Value: bsonx.Int32(-1)}},
		},
		{
			Keys: bsonx.Doc{{Key: "department_id", Value: bsonx.Int32(-1)}},
		},
	}
	_, err = mc.Indexes().CreateMany(context.Background(), indexList)
	if err != nil {
		return err
	}

	s.col = mc
	s.log = zap.L().Named(s.Name())
	return nil
}

func (s *service) Name() string {
	return user.AppName
}

func (s *service) Registry(server *grpc.Server) {
	user.RegisterServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}
