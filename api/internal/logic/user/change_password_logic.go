package user

import (
	"context"

	"github.com/saas-mingyang/mingyang-admin-common/utils/encrypt"
	"github.com/saas-mingyang/mingyang-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/saas-mingyang/mingyang-admin-core/api/internal/svc"
	"github.com/saas-mingyang/mingyang-admin-core/api/internal/types"
	"github.com/saas-mingyang/mingyang-admin-core/api/internal/utils"
	"github.com/saas-mingyang/mingyang-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordLogic {
	return &ChangePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangePasswordLogic) ChangePassword(req *types.ChangePasswordReq) (resp *types.BaseMsgResp, err error) {
	userId := utils.GetUserIdFromCtx(l.ctx)
	if userId == 0 {
		return nil, errorx.NewCodeInvalidArgumentError("invalid user id")
	}

	userData, err := l.svcCtx.CoreRpc.GetUserById(l.ctx, &core.IDReq{Id: userId})
	if err != nil {
		return nil, err
	}

	if encrypt.BcryptCheck(req.OldPassword, *userData.Password) {
		result, err := l.svcCtx.CoreRpc.UpdateUser(l.ctx, &core.UserInfo{
			Id:       pointy.GetPointer(userId),
			Password: pointy.GetPointer(req.NewPassword),
		})
		if err != nil {
			return nil, err
		}

		return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, result.Msg)}, nil
	}

	return nil, errorx.NewCodeInvalidArgumentError("login.wrongPassword")
}
