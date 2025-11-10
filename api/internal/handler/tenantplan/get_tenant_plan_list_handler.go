package tenantplan

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/saas-mingyang/mingyang-admin-core/api/internal/logic/tenantplan"
	"github.com/saas-mingyang/mingyang-admin-core/api/internal/svc"
	"github.com/saas-mingyang/mingyang-admin-core/api/internal/types"
)

// swagger:route post /tenant_plan/list tenantplan GetTenantPlanList
//
// Get tenant plan list | 获取租户套餐列表
//
// Get tenant plan list | 获取租户套餐列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: TenantPlanListReq
//
// Responses:
//  200: TenantPlanListResp

func GetTenantPlanListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TenantPlanListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := tenantplan.NewGetTenantPlanListLogic(r.Context(), svcCtx)
		resp, err := l.GetTenantPlanList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
