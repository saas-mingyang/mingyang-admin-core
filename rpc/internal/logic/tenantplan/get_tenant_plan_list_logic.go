package tenantplan

import (
	"context"

	"github.com/saas-mingyang/mingyang-admin-core/rpc/internal/svc"
	"github.com/saas-mingyang/mingyang-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTenantPlanListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTenantPlanListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTenantPlanListLogic {
	return &GetTenantPlanListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTenantPlanListLogic) GetTenantPlanList(in *core.IDReq) (*core.TenantPlanInfo, error) {
	// todo: add your logic here and delete this line

	return &core.TenantPlanInfo{}, nil
}
