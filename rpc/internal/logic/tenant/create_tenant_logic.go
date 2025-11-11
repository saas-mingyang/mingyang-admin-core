package tenant

import (
	"context"
	"fmt"
	"github.com/saas-mingyang/mingyang-admin-common/i18n"
	"github.com/saas-mingyang/mingyang-admin-common/utils/pointy"
	"github.com/saas-mingyang/mingyang-admin-core/rpc/internal/utils/dberrorhandler"

	"github.com/saas-mingyang/mingyang-admin-core/rpc/internal/svc"
	"github.com/saas-mingyang/mingyang-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTenantLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateTenantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTenantLogic {
	return &CreateTenantLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateTenant Tenant management
func (l *CreateTenantLogic) CreateTenant(in *core.TenantInfo) (*core.BaseIDResp, error) {
	result, err := l.svcCtx.DB.Tenant.Create().
		SetNotNilStatus(pointy.GetStatusPointer(in.Status)).
		SetNotNilCode(in.Code).
		SetNotNilName(in.Name).
		SetNotNilIntro(in.Intro).
		SetNotNilLevel(in.Level).
		SetNotNilLicenseNumber(in.LicenseNumber).
		SetNotNilContactPhone(in.ContactPhone).
		SetNotNilContactEmail(in.ContactEmail).
		SetNotNilCompanyName(in.CompanyName).
		SetDomain(*in.Domain).
		SetAddress(*in.Address).
		SetAdminID(*in.AdminId).
		Save(l.ctx)

	if err != nil {
		fmt.Printf("create tenant error: %v", err)
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}
	return &core.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil
}
