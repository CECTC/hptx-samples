package svc

import (
	"context"

	hptxGrpc "github.com/cectc/hptx/pkg/contrib/grpc"
	"github.com/cectc/mysql"
	"google.golang.org/grpc/metadata"

	"github.com/cectc/hptx-samples/order_svc/api"
	"github.com/cectc/hptx-samples/order_svc/dao"
)

type Service struct {
	Dao *dao.Dao
}

func (svc *Service) CreateSo(ctx context.Context, req *api.CreateSoReq) (*api.CreateSoResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	xid := md.Get(hptxGrpc.XID)[0]
	ctx = context.WithValue(ctx, mysql.XID, xid)

	_, err := svc.Dao.CreateSO(ctx, req.SoMasters)
	if err == nil {
		return &api.CreateSoResponse{
			Success: true,
			Message: "success",
		}, nil
	}
	return nil, err
}
