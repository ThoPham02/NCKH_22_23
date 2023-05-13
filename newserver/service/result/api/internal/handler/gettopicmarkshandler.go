package handler

import (
	"net/http"

	"github.com/ThoPham02/research_management/service/result/api/internal/logic"
	"github.com/ThoPham02/research_management/service/result/api/internal/svc"
	"github.com/ThoPham02/research_management/service/result/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetTopicMarksHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetTopicMarksReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetTopicMarksLogic(r.Context(), svcCtx)
		resp, err := l.GetTopicMarks(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
