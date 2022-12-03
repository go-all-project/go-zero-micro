package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go_zero/apps/app/web/internal/svc"
	"go_zero/apps/app/web/internal/types"
	"go_zero/apps/order/rpc/order"
)

type HomeBannerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeBannerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeBannerLogic {
	return &HomeBannerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeBannerLogic) HomeBanner() (resp *types.HomeBannerResponse, err error) {
	logx.Infow("日志打印", logx.Field("uid", "test"))

	if _, err = l.svcCtx.OrderRPCClient.Orders(l.ctx, &order.OrdersRequest{UserId: 1}); err != nil {
		return nil, err
	}

	return &types.HomeBannerResponse{}, nil
}
