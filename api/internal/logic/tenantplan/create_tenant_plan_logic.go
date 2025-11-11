package tenantplan

import (
	"context"
	"fmt"
	"github.com/saas-mingyang/mingyang-admin-core/api/internal/svc"
	"github.com/saas-mingyang/mingyang-admin-core/api/internal/types"
	"github.com/saas-mingyang/mingyang-admin-core/rpc/types/core"
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
	data, err := l.svcCtx.CoreRpc.CreateTenantPlan(l.ctx,
		&core.TenantPlanCreateReq{
			PackageName:       req.PackageName,
			MenuCheckStrictly: req.MenuCheckStrictly,
			Status:            req.Status,
			Remark:            req.Remark,
			MenuIds:           req.MenuIds,
			ApiIds:            req.ApiIds,
		})
	if err != nil {
		fmt.Printf("create tenant plan error: %v", err)
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
