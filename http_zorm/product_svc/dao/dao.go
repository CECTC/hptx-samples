package dao

import (
	"context"

	"gitee.com/chunanyong/zorm"
)

const (
	allocateInventorySql = `update product.inventory set available_qty = available_qty - ?, 
		allocated_qty = allocated_qty + ? where product_sysno = ? and available_qty >= ?`
)

type Dao struct {
	*zorm.DBDao
}

type AllocateInventoryReq struct {
	ProductSysNo int64
	Qty          int32
}

func (dao *Dao) AllocateInventory(ctx context.Context, reqs []*AllocateInventoryReq) error {

	_, err := zorm.Transaction(ctx, func(ctx context.Context) (interface{}, error) {

		for _, req := range reqs {
			finder := zorm.NewFinder()
			finder.Append(allocateInventorySql, req.Qty, req.Qty, req.ProductSysNo, req.Qty)
			_, err := zorm.UpdateFinder(ctx, finder)
			if err != nil {
				return nil, err
			}
		}
		return nil, nil
	})

	return err
}
