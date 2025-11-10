package tenantplan

import (
	"context"

	"github.com/saas-mingyang/mingyang-admin-core/api/internal/svc"
	"github.com/saas-mingyang/mingyang-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTenantPlanByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTenantPlanByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTenantPlanByIdLogic {
	return &GetTenantPlanByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTenantPlanByIdLogic) GetTenantPlanById(req *types.IDReq) (resp *types.TenantPlanInfo, err error) {
	// todo: add your logic here and delete this line

	return
}
