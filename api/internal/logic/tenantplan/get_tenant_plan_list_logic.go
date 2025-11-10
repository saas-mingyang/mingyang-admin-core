package tenantplan

import (
	"context"

	"github.com/saas-mingyang/mingyang-admin-core/api/internal/svc"
	"github.com/saas-mingyang/mingyang-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTenantPlanListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTenantPlanListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTenantPlanListLogic {
	return &GetTenantPlanListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTenantPlanListLogic) GetTenantPlanList(req *types.TenantPlanListReq) (resp *types.TenantPlanListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
