package svc

import (
	"context"
	"errors"

	"github.com/cectc/hptx/pkg/constant"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/metadata"

	"github.com/cectc/hptx-samples/aggregation_svc/api"
	soApi "github.com/cectc/hptx-samples/order_svc/api"
	inventoryApi "github.com/cectc/hptx-samples/product_svc/api"
)

type Service struct {
	SoClient        soApi.SoServiceClient
	InventoryClient inventoryApi.InventoryServiceClient
}

func (svc *Service) CreateSoCommit(ctx context.Context, _ *empty.Empty) (*api.Response, error) {
	xid := ctx.Value(constant.XID)
	md := metadata.Pairs(string(constant.XID), xid.(string))
	ctx = metadata.NewOutgoingContext(ctx, md)

	soMasters := []*soApi.SoMaster{
		{
			BuyerUserSysNo:       10001,
			SellerCompanyCode:    "SC001",
			ReceiveDivisionSysNo: 110105,
			ReceiveAddress:       "beijing",
			ReceiveZip:           "000001",
			ReceiveContact:       "scott",
			ReceiveContactPhone:  "18728828296",
			StockSysNo:           1,
			PaymentType:          1,
			SoAmt:                6999 * 2,
			Status:               10,
			AppID:                "dk-order",
			SoItems: []*soApi.SoItem{
				{
					ProductSysNo:  1,
					ProductName:   "apple iphone 13",
					CostPrice:     6799,
					OriginalPrice: 6799,
					DealPrice:     6999,
					Quantity:      2,
				},
			},
		},
	}

	_, err := svc.SoClient.CreateSo(ctx, &soApi.CreateSoReq{
		SoMasters: soMasters,
	})
	if err != nil {
		return nil, err
	}

	reqs := &inventoryApi.AllocateInventoryReq{
		AllocateInventories: []*inventoryApi.AllocateInventory{
			{
				ProductSysNo: 1,
				Qty:          2,
			},
		}}
	_, err = svc.InventoryClient.AllocateInventory(ctx, reqs)
	if err != nil {
		return nil, err
	}

	return &api.Response{
		Success: true,
		Message: "success",
	}, nil
}

func (svc *Service) CreateSoRollback(ctx context.Context, _ *empty.Empty) (*api.Response, error) {
	xid := ctx.Value(constant.XID)
	md := metadata.Pairs(string(constant.XID), xid.(string))
	ctx = metadata.NewOutgoingContext(ctx, md)

	soMasters := []*soApi.SoMaster{
		{
			BuyerUserSysNo:       10001,
			SellerCompanyCode:    "SC001",
			ReceiveDivisionSysNo: 110105,
			ReceiveAddress:       "beijing",
			ReceiveZip:           "000001",
			ReceiveContact:       "scott",
			ReceiveContactPhone:  "18728828296",
			StockSysNo:           1,
			PaymentType:          1,
			SoAmt:                6999 * 2,
			Status:               10,
			AppID:                "dk-order",
			SoItems: []*soApi.SoItem{
				{
					ProductSysNo:  1,
					ProductName:   "apple iphone 13",
					CostPrice:     6799,
					OriginalPrice: 6799,
					DealPrice:     6999,
					Quantity:      2,
				},
			},
		},
	}

	_, err := svc.SoClient.CreateSo(ctx, &soApi.CreateSoReq{
		SoMasters: soMasters,
	})
	if err != nil {
		return nil, err
	}

	reqs := &inventoryApi.AllocateInventoryReq{
		AllocateInventories: []*inventoryApi.AllocateInventory{
			{
				ProductSysNo: 1,
				Qty:          2,
			},
		}}
	_, err = svc.InventoryClient.AllocateInventory(ctx, reqs)
	if err != nil {
		return nil, err
	}

	return nil, errors.New("there is a error")
}
