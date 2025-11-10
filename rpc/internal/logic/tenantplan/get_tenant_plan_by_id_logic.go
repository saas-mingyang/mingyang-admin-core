package tenantplan

import (
	"context"

	"github.com/saas-mingyang/mingyang-admin-core/rpc/internal/svc"
	"github.com/saas-mingyang/mingyang-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTenantPlanByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTenantPlanByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTenantPlanByIdLogic {
	return &GetTenantPlanByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTenantPlanByIdLogic) GetTenantPlanById(in *core.TenantPlanInfoReq) (*core.TenantPlanListResp, error) {
	// todo: add your logic here and delete this line

	return &core.TenantPlanListResp{}, nil
}
