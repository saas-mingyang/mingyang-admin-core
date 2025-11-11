package tenant

import (
	"context"
	"fmt"
	"github.com/saas-mingyang/mingyang-admin-core/rpc/types/core"

	"github.com/saas-mingyang/mingyang-admin-core/api/internal/svc"
	"github.com/saas-mingyang/mingyang-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTenantLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTenantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTenantLogic {
	return &CreateTenantLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTenantLogic) CreateTenant(req *types.TenantInfo) (resp *types.BaseMsgResp, err error) {

	data, err := l.svcCtx.CoreRpc.CreateTenant(l.ctx,
		&core.TenantInfo{
			Status:        req.Status,
			Code:          req.Code,
			Name:          req.Name,
			ParentId:      req.ParentId,
			AdminId:       req.AdminId,
			CompanyName:   req.CompanyName,
			Intro:         req.Intro,
			Level:         req.Level,
			LicenseNumber: req.LicenseNumber,
			Domain:        req.Domain,
			Address:       req.Address,
			PlanId:        req.PlanId,
			ContactPhone:  req.ContactPhone,
			ContactEmail:  req.ContactEmail,
		})
	if err != nil {
		fmt.Printf("create tenant plan error: %v", err)
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
