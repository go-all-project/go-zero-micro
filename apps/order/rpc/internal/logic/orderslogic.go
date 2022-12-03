package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go_zero/apps/order/model"
	"go_zero/apps/order/rpc/internal/svc"
	"go_zero/apps/order/rpc/pb"
	"google.golang.org/grpc/metadata"
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
	if md, ok := metadata.FromIncomingContext(l.ctx); ok {
		tmp := md.Get("username")
		fmt.Printf("rpc metadata 传参 %v \n", tmp)
	}

	user := models.User{
		Mobile: in.UserId,
	}

	_ = l.svcCtx.DB.Create(&user)

	return &pb.OrdersResponse{}, nil
}
