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

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo() (resp *types.UserBaseIDInfoResp, err error) {
	userId := utils.GetUserIdFromCtx(l.ctx)
	if userId == 0 {
		return nil, errorx.NewCodeInvalidArgumentError("invalid user id")
	}

	user, err := l.svcCtx.CoreRpc.GetUserById(l.ctx,
		&core.IDReq{Id: userId})
	if err != nil {
		return nil, err
	}

	return &types.UserBaseIDInfoResp{
		BaseDataInfo: types.BaseDataInfo{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.Success)},
		Data: types.UserBaseIDInfo{
			UserId:         user.Id,
			Username:       user.Username,
			Nickname:       user.Nickname,
			Avatar:         user.Avatar,
			HomePath:       user.HomePath,
			Description:    user.Description,
			DepartmentName: l.svcCtx.Trans.Trans(l.ctx, *user.DepartmentName),
			RoleName:       TransRoleName(l.svcCtx, l.ctx, user.RoleName),
		},
	}, nil
}

// TransRoleName returns the i18n translation of role name slice.
func TransRoleName(svc *svc.ServiceContext, ctx context.Context, data []string) []string {
	var result []string
	for _, v := range data {
		result = append(result, svc.Trans.Trans(ctx, v))
	}
	return result
}
