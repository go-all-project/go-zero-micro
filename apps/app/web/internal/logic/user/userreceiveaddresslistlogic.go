package user

import (
	"context"

	"go_zero/apps/app/web/internal/svc"
	"go_zero/apps/app/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserReceiveAddressListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserReceiveAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserReceiveAddressListLogic {
	return &UserReceiveAddressListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserReceiveAddressListLogic) UserReceiveAddressList(req *types.UserReceiveAddressListReq) (resp *types.UserReceiveAddressListRes, err error) {
	// todo: add your logic here and delete this line

	return
}
