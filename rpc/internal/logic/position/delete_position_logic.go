package position

import (
	"context"

	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/saas-mingyang/mingyang-admin-core/rpc/ent/position"
	"github.com/saas-mingyang/mingyang-admin-core/rpc/ent/user"

	"github.com/saas-mingyang/mingyang-admin-core/rpc/internal/svc"
	"github.com/saas-mingyang/mingyang-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/saas-mingyang/mingyang-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/saas-mingyang/mingyang-admin-common/i18n"
)

type DeletePositionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePositionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePositionLogic {
	return &DeletePositionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeletePositionLogic) DeletePosition(in *core.IDsReq) (*core.BaseResp, error) {
	count, err := l.svcCtx.DB.User.Query().Where(user.HasPositionsWith(position.IDIn(in.Ids...))).Count(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	if count != 0 {
		return nil, errorx.NewInvalidArgumentError("position.userExistError")
	}

	_, err = l.svcCtx.DB.Position.Delete().Where(position.IDIn(in.Ids...)).Exec(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
