package user

import (
	"context"

	"github.com/saas-mingyang/mingyang-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/saas-mingyang/mingyang-admin-core/api/internal/svc"
	"github.com/saas-mingyang/mingyang-admin-core/api/internal/types"
	"github.com/saas-mingyang/mingyang-admin-core/api/internal/utils"
	"github.com/saas-mingyang/mingyang-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserProfileLogic {
	return &GetUserProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserProfileLogic) GetUserProfile() (resp *types.ProfileResp, err error) {
	userId := utils.GetUserIdFromCtx(l.ctx)
	if userId == 0 {
		return nil, errorx.NewCodeInvalidArgumentError("invalid user id")
	}

	data, err := l.svcCtx.CoreRpc.GetUserById(l.ctx, &core.IDReq{Id: userId})
	if err != nil {
		return nil, err
	}

	return &types.ProfileResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.ProfileInfo{
			Nickname: data.Nickname,
			Avatar:   data.Avatar,
			Mobile:   data.Mobile,
			Email:    data.Email,
		},
	}, nil
}
