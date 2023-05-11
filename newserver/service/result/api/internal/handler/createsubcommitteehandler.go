package handler

import (
	"net/http"

	"github.com/ThoPham02/research_management/service/result/api/internal/logic"
	"github.com/ThoPham02/research_management/service/result/api/internal/svc"
	"github.com/ThoPham02/research_management/service/result/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateSubcommitteeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateSubcommitteeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCreateSubcommitteeLogic(r.Context(), svcCtx)
		resp, err := l.CreateSubcommittee(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
