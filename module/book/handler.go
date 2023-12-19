package book

import (
	"fmt"
	"idempotency/pkg/mlogger"
	"idempotency/pkg/web"
	"net/http"
)

type Handler struct {
	service BookUsecaseAssumer
	log     mlogger.Logger
}

func NewHandler(service BookUsecaseAssumer, log mlogger.Logger) *Handler {
	return &Handler{
		service: service,
		log:     log,
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	isbn, err := web.ReadStrIDParam(r)
	if err != nil {
		h.log.WarnT(ctx, "invalid id params", err)
		web.ErrorResponse(w, 400, err.Error())
	}

	fmt.Println(isbn, "======")

	result, err := h.service.Get(ctx, isbn)
	if err != nil {
		h.log.ErrorT(ctx, "error get", err)
		web.ErrorResponse(w, 404, err.Error()) // ignore mapping error code
		return
	}

	env := web.Envelope{
		"data": result,
	}

	err = web.WriteJSON(w, http.StatusOK, env, nil)
	if err != nil {
		web.ServerErrorResponse(w, r, err)
		return
	}
}

func (h *Handler) Insert(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req BookEntity
	err := web.ReadJSON(w, r, &req)
	if err != nil {
		h.log.WarnT(ctx, "bad json", err)
		web.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.service.Insert(ctx, req)
	if err != nil {
		h.log.ErrorT(ctx, "error insert", err)
		web.ErrorResponse(w, 400, err.Error()) // ignore mapping error code
		return
	}

	env := web.Envelope{
		"data": result,
	}

	err = web.WriteJSON(w, http.StatusCreated, env, nil)
	if err != nil {
		web.ServerErrorResponse(w, r, err)
		return
	}
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	isbn, err := web.ReadStrIDParam(r)
	if err != nil {
		h.log.WarnT(ctx, "invalid id params", err)
		web.ErrorResponse(w, 400, err.Error())
	}

	var req BookEntity
	err = web.ReadJSON(w, r, &req)
	if err != nil {
		h.log.WarnT(ctx, "bad json", err)
		web.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	req.ISBN = isbn

	result, err := h.service.Update(ctx, req)
	if err != nil {
		h.log.ErrorT(ctx, "error update", err)
		web.ErrorResponse(w, 400, err.Error()) // ignore mapping error code
		return
	}

	env := web.Envelope{
		"data": result,
	}

	err = web.WriteJSON(w, http.StatusOK, env, nil)
	if err != nil {
		web.ServerErrorResponse(w, r, err)
		return
	}
}
