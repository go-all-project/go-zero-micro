package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go_zero/apps/app/web/internal/logic"
	"go_zero/apps/app/web/internal/svc"
)

func homeBannerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewHomeBannerLogic(r.Context(), svcCtx)
		resp, err := l.HomeBanner()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
