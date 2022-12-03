package logic

import (
	"context"

	"go_zero/apps/order/rpc/internal/svc"
	"go_zero/apps/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrdersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrdersLogic {
	return &OrdersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrdersLogic) Orders(in *pb.OrdersRequest) (*pb.OrdersResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.OrdersResponse{}, nil
}
