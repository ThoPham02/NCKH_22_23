package handler

import (
	"net/http"

	"github.com/ThoPham02/research_management/service/account/api/internal/logic"
	"github.com/ThoPham02/research_management/service/account/api/internal/svc"
	"github.com/ThoPham02/research_management/service/account/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetUsersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUsersReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetUsersLogic(r.Context(), svcCtx)
		resp, err := l.GetUsers(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
