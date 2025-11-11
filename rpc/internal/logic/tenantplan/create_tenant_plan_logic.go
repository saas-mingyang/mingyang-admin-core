package tenantplan

import (
	"context"
	"fmt"
	"github.com/saas-mingyang/mingyang-admin-common/i18n"
	"github.com/saas-mingyang/mingyang-admin-common/utils/pointy"
	"github.com/saas-mingyang/mingyang-admin-core/rpc/internal/svc"
	"github.com/saas-mingyang/mingyang-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/saas-mingyang/mingyang-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTenantPlanLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateTenantPlanLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTenantPlanLogic {
	return &CreateTenantPlanLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateTenantPlanLogic) CreateTenantPlan(in *core.TenantPlanCreateReq) (*core.BaseIDResp, error) {
	logx.Info("create tenant plan req:", in)
	result, err := l.svcCtx.DB.TenantPlan.Create().
		SetNotNilStatus(pointy.GetStatusPointer(in.Status)).
		SetNotNilAPIIds(in.ApiIds).
		SetPackageName(*in.PackageName).
		SetNotNilMenuIds(in.MenuIds).
		SetNotNilMenuCheckStrictly(in.MenuCheckStrictly).
		SetNotNilRemark(in.Remark).
		Save(l.ctx)
	if err != nil {
		fmt.Printf("create token error: %v", err)
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}
	return &core.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil
}
