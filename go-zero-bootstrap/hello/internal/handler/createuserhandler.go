package handler

import (
	"net/http"

	"hello/internal/logic"
	"hello/internal/svc"
	"hello/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func CreateUserHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCreateUserLogic(r.Context(), ctx)
		err := l.CreateUser(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
