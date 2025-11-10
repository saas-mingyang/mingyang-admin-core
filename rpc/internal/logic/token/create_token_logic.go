package token

import (
	"context"
	"fmt"

	"github.com/saas-mingyang/mingyang-admin-common/utils/pointy"
	"github.com/saas-mingyang/mingyang-admin-core/rpc/internal/svc"
	"github.com/saas-mingyang/mingyang-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/saas-mingyang/mingyang-admin-core/rpc/types/core"

	"github.com/saas-mingyang/mingyang-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTokenLogic {
	return &CreateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateTokenLogic) CreateToken(in *core.TokenInfo) (*core.BaseIDResp, error) {
	result, err := l.svcCtx.DB.Token.Create().
		SetNotNilStatus(pointy.GetStatusPointer(in.Status)).
		SetNotNilUserID(in.UserId).
		SetNotNilToken(in.Token).
		SetNotNilSource(in.Source).
		SetNotNilUsername(in.Username).
		SetNotNilExpiredAt(pointy.GetTimeMilliPointer(in.ExpiredAt)).
		Save(l.ctx)
	if err != nil {
		fmt.Printf("create token error: %v", err)
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil
}
