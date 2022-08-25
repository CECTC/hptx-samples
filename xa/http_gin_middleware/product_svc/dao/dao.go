package dao

import (
	"context"
	"database/sql"

	"github.com/cectc/hptx/pkg/xa"
)

const (
	allocateInventorySql = `update product.inventory set available_qty = available_qty - ?, 
		allocated_qty = allocated_qty + ? where product_sysno = ? and available_qty >= ?`
)

type Dao struct {
	*sql.DB
}

type AllocateInventoryReq struct {
	ProductSysNo int64
	Qty          int32
}

func (dao *Dao) AllocateInventory(ctx context.Context, reqs []*AllocateInventoryReq) error {
	return xa.HandleWithXA(ctx, dao.DB, "productSvc", func(conn *sql.Conn) error {
		for _, req := range reqs {
			_, err := conn.ExecContext(ctx, allocateInventorySql, req.Qty, req.Qty, req.ProductSysNo, req.Qty)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
