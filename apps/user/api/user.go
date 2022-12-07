package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/kalandramo/keyauth/apps/user"
)

func (h *handler) CreateUser(r *restful.Request, w *restful.Response) {
	req := user.NewCreateAccountRequest()

	if err := request.GetDataFromRequest(r.Request, req); err != nil {
		response.Failed(w, err)
	}

	set, err := h.service.CreateAccount(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	response.Success(w.ResponseWriter, set)
}

func (h *handler) QueryUser(r *restful.Request, w *restful.Response) {
	req := user.NewQueryAccountRequestFromHTTP(r.Request)
	h.log.Debugf("query user: %s", req)
	set, err := h.service.QueryAccount(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}
	response.Success(w.ResponseWriter, set)
}
