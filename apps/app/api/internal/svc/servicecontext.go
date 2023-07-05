package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zeroshop/app/api/internal/config"
	"zeroshop/order/rpc/order"
	"zeroshop/product/rpc/product"
)

type ServiceContext struct {
	Config     config.Config
	OrderRPC   order.Order
	ProductRPC product.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		OrderRPC:   order.NewOrder(zrpc.MustNewClient(c.OrderRPC)),
		ProductRPC: product.NewProduct(zrpc.MustNewClient(c.ProductRPC)),
	}
}
