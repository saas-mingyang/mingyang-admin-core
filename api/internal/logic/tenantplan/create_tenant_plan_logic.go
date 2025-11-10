package tenantplan

import (
	"context"

	"github.com/saas-mingyang/mingyang-admin-core/api/internal/svc"
	"github.com/saas-mingyang/mingyang-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTenantPlanLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTenantPlanLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTenantPlanLogic {
	return &CreateTenantPlanLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTenantPlanLogic) CreateTenantPlan(req *types.TenantPlanInfo) (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this line

	return
}
