package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go_zero/mall/order/api/internal/logic"
	"go_zero/mall/order/api/internal/svc"
	"go_zero/mall/order/api/internal/types"
)

func getOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		// 创建 service
		l := logic.NewGetOrderLogic(r.Context(), svcCtx)
		// 使用 service
		resp, err := l.GetOrder(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
