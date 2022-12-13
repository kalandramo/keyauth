package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/kalandramo/keyauth/apps/user"
)

var h = &handler{}

type handler struct {
	service user.ServiceServer
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(user.AppName)
	h.service = app.GetGrpcApp(user.AppName).(user.ServiceServer)
	return nil
}

func (h *handler) Name() string {
	return user.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{h.Name()}

	ws.Route(ws.POST("").To(h.CreateUser).
		Doc("create a user").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(user.CreateAccountRequest{}).
		Writes(response.NewData(user.User{})))

	ws.Route(ws.GET("/").To(h.QueryUser).
		Doc("get all user").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata("action", "list").
		Reads(user.QueryAccountRequest{}).
		Writes(response.NewData(user.Set{})).
		Returns(200, "OK", user.Set{}))
}

func init() {
	app.RegistryRESTfulApp(h)
}
