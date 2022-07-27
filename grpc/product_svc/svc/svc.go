package svc

import (
	"context"

	hptxGrpc "github.com/cectc/hptx/pkg/contrib/grpc"
	"github.com/cectc/mysql"
	"google.golang.org/grpc/metadata"

	"github.com/cectc/hptx-samples/product_svc/api"
	"github.com/cectc/hptx-samples/product_svc/dao"
)

type Service struct {
	Dao *dao.Dao
}

func (svc *Service) AllocateInventory(ctx context.Context, req *api.AllocateInventoryReq) (*api.AllocateInventoryResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	xid := md.Get(hptxGrpc.XID)[0]
	ctx = context.WithValue(ctx, mysql.XID, xid)

	err := svc.Dao.AllocateInventory(ctx, req.AllocateInventories)
	if err == nil {
		return &api.AllocateInventoryResponse{
			Success: true,
			Message: "success",
		}, nil
	}
	return nil, err
}
