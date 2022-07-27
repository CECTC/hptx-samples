package dao

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/cectc/hptx-samples/order_svc/api"
)

const (
	insertSoMaster = `INSERT INTO order.so_master (sysno, so_id, buyer_user_sysno, seller_company_code, 
		receive_division_sysno, receive_address, receive_zip, receive_contact, receive_contact_phone, stock_sysno, 
        payment_type, so_amt, status, order_date, appid, memo) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,now(),?,?)`
	insertSoItem = `INSERT INTO order.so_item(sysno, so_sysno, product_sysno, product_name, cost_price, 
		original_price, deal_price, quantity) VALUES (?,?,?,?,?,?,?,?)`
)

type Dao struct {
	*sql.DB
}

func (dao *Dao) CreateSO(ctx context.Context, soMasters []*api.SoMaster) ([]uint64, error) {
	result := make([]uint64, 0, len(soMasters))
	tx, err := dao.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  false,
	})
	if err != nil {
		return nil, err
	}
	for _, soMaster := range soMasters {
		soid := NextID()
		_, err = tx.Exec(insertSoMaster, soid, soid, soMaster.BuyerUserSysNo, soMaster.SellerCompanyCode, soMaster.ReceiveDivisionSysNo,
			soMaster.ReceiveAddress, soMaster.ReceiveZip, soMaster.ReceiveContact, soMaster.ReceiveContactPhone, soMaster.StockSysNo,
			soMaster.PaymentType, soMaster.SoAmt, soMaster.Status, soMaster.AppID, soMaster.Memo)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		soItems := soMaster.SoItems
		for _, soItem := range soItems {
			soItemID := NextID()
			_, err = tx.Exec(insertSoItem, soItemID, soid, soItem.ProductSysNo, soItem.ProductName, soItem.CostPrice, soItem.OriginalPrice,
				soItem.DealPrice, soItem.Quantity)
			if err != nil {
				tx.Rollback()
				return nil, err
			}
		}
		result = append(result, soid)
	}
	err = tx.Commit()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return result, nil
}

func NextID() uint64 {
	id, _ := uuid.NewUUID()
	return uint64(id.ID())
}
