package handler

import (
	"net/http"

	"github.com/ThoPham02/research_management/service/topic/api/internal/logic"
	"github.com/ThoPham02/research_management/service/topic/api/internal/svc"
	"github.com/ThoPham02/research_management/service/topic/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateTopicSubcommitteeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateTopicSubcommitteeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUpdateTopicSubcommitteeLogic(r.Context(), svcCtx)
		resp, err := l.UpdateTopicSubcommittee(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
