package handler

import (
	"net/http"

	"github.com/593349399/zeroshop/apps/app/api/internal/logic"
	"github.com/593349399/zeroshop/apps/app/api/internal/svc"
	"github.com/593349399/zeroshop/apps/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func OrderListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewOrderListLogic(r.Context(), svcCtx)
		resp, err := l.OrderList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
