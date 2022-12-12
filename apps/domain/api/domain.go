package api

import (
	"net/http"

	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"
	"github.com/kalandramo/keyauth/apps/domain"
)

func (h *handler) CreateDomain(w http.ResponseWriter, r *http.Request) {
	req := domain.NewCreateDomainRequest()

	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	d, err := h.service.CreateDomain(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
}

func (h *handler) ListDomains(w http.ResponseWriter, r *http.Request) {
	// ctx := context.GetContext(r)
	// tk := ctx.AuthInfo(*token.token)

	page := request.NewPageRequestFromHTTP(r)
	req := domain.NewQueryDomainRequest(page)
	// req.Owner = tk.Account
	domains, err := h.service.QueryDomain(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, domains)
}

func (h *handler) GetDomain(w http.ResponseWriter, r *http.Request) {
	rCtx := context.GetContext(r)

	req := domain.NewDescribeDomainRequest()
	req.Name = rCtx.PS.ByName("name")

	d, err := h.service.DescribeDomain(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, d)
}

func (h *handler) PutDomain(w http.ResponseWriter, r *http.Request) {
	rCtx := context.GetContext(r)

	// 查找出原来的domain
	req := domain.NewPutDomainRequest()
	req.Name = rCtx.PS.ByName("name")

	// 解析需要更新的数据
	if err := request.GetDataFromRequest(r, req.Profile); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.UpdateDomain(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
	}

	response.Success(w, ins)
}

func (h *handler) PatchDomain(w http.ResponseWriter, r *http.Request) {
	rCtx := context.GetContext(r)

	// 查找出原来的domain
	req := domain.NewPatchDomainRequest()
	req.Name = rCtx.PS.ByName("name")

	// 解析需要更新的数据
	if err := request.GetDataFromRequest(r, req.Profile); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.UpdateDomain(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) DeleteDomain(w http.ResponseWriter, r *http.Request) {
	rCtx := context.GetContext(r)
	name := rCtx.PS.ByName("name")
	req := domain.NewDeleteDomainRequestByName(name)

	_, err := h.service.DeleteDomain(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, "delete ok")
}

func (h *handler) PutDomainSecurity(w http.ResponseWriter, r *http.Request) {
	rCtx := context.GetContext(r)

	// 查找出原来的domain
	req := domain.NewPutDomainSecurityRequest()
	req.Name = rCtx.PS.ByName("name")
	req.SecuritySetting = domain.NewDefaultSecuritySetting()

	// 解析需要更新的数据
	if err := request.GetDataFromRequest(r, req.SecuritySetting); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.UpdateDomainSecurity(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}
