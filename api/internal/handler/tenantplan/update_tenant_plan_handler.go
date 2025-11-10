package tenantplan

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/saas-mingyang/mingyang-admin-core/api/internal/logic/tenantplan"
	"github.com/saas-mingyang/mingyang-admin-core/api/internal/svc"
	"github.com/saas-mingyang/mingyang-admin-core/api/internal/types"
)

// swagger:route post /tenant_plan/update tenantplan UpdateTenantPlan
//
// Update tenant plan | 更新租户套餐
//
// Update tenant plan | 更新租户套餐
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: TenantPlanInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateTenantPlanHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TenantPlanInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := tenantplan.NewUpdateTenantPlanLogic(r.Context(), svcCtx)
		resp, err := l.UpdateTenantPlan(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
