package logic

import (
	"context"
	"github.com/593349399/zeroshop/apps/order/rpc/order"
	"github.com/593349399/zeroshop/apps/product/rpc/product"
	"strconv"
	"strings"

	"github.com/593349399/zeroshop/apps/app/api/internal/svc"
	"github.com/593349399/zeroshop/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderListLogic {
	return &OrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderListLogic) OrderList(req *types.OrderListRequest) (resp *types.OrderListResponse, err error) {
	orderRet, err := l.svcCtx.OrderRPC.Orders(l.ctx, &order.OrdersRequest{UserId: req.UID})
	if err != nil {
		return nil, err
	}
	var pids []string
	for _, o := range orderRet.Orders {
		pids = append(pids, strconv.Itoa(int(o.Id)))
	}
	productRet, err := l.svcCtx.ProductRPC.Products(l.ctx, &product.ProductRequest{ProductIds: strings.Join(pids, ",")})
	if err != nil {
		return nil, err
	}
	var orders []*types.Order
	for _, o := range orderRet.Orders {
		if p, ok := productRet.Products[o.Proid]; ok {
			orders = append(orders, &types.Order{
				OrderID:     o.Orderid,
				ProductName: p.Name,
			})
		}
	}
	return &types.OrderListResponse{Orders: orders}, nil
}
