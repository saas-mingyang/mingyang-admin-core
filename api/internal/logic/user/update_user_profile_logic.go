package user

import (
	"context"

	"github.com/saas-mingyang/mingyang-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/saas-mingyang/mingyang-admin-core/api/internal/svc"
	"github.com/saas-mingyang/mingyang-admin-core/api/internal/types"
	"github.com/saas-mingyang/mingyang-admin-core/api/internal/utils"
	"github.com/saas-mingyang/mingyang-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserProfileLogic {
	return &UpdateUserProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserProfileLogic) UpdateUserProfile(req *types.ProfileInfo) (resp *types.BaseMsgResp, err error) {
	userId := utils.GetUserIdFromCtx(l.ctx)
	if userId == 0 {
		return nil, errorx.NewCodeInvalidArgumentError("invalid user id")
	}

	result, err := l.svcCtx.CoreRpc.UpdateUser(l.ctx, &core.UserInfo{
		Id:       pointy.GetPointer(userId),
		Nickname: req.Nickname,
		Email:    req.Email,
		Mobile:   req.Mobile,
		Avatar:   req.Avatar,
	})
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, result.Msg)}, nil
}
