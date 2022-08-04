package dao

import (
	"context"
	"fmt"
	"time"

	"gitee.com/chunanyong/zorm"
	"github.com/google/uuid"
)

const (
	insertSoMaster = `INSERT INTO order.so_master (sysno, so_id, buyer_user_sysno, seller_company_code, 
		receive_division_sysno, receive_address, receive_zip, receive_contact, receive_contact_phone, stock_sysno, 
        payment_type, so_amt, status, order_date, appid, memo) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,now(),?,?)`
	insertSoItem = `INSERT INTO order.so_item(sysno, so_sysno, product_sysno, product_name, cost_price, 
		original_price, deal_price, quantity) VALUES (?,?,?,?,?,?,?,?)`
)

type Dao struct {
	*zorm.DBDao
}

//现实中涉及金额可能使用长整形，这里使用 float64 仅作测试，不具有参考意义

type SoMaster struct {
	SysNo                int64   `json:"sysNo"`
	SoID                 string  `json:"soID"`
	BuyerUserSysNo       int64   `json:"buyerUserSysNo"`
	SellerCompanyCode    string  `json:"sellerCompanyCode"`
	ReceiveDivisionSysNo int64   `json:"receiveDivisionSysNo"`
	ReceiveAddress       string  `json:"receiveAddress"`
	ReceiveZip           string  `json:"receiveZip"`
	ReceiveContact       string  `json:"receiveContact"`
	ReceiveContactPhone  string  `json:"receiveContactPhone"`
	StockSysNo           int64   `json:"stockSysNo"`
	PaymentType          int32   `json:"paymentType"`
	SoAmt                float64 `json:"soAmt"`
	//10，创建成功，待支付；30；支付成功，待发货；50；发货成功，待收货；70，确认收货，已完成；90，下单失败；100已作废
	Status       int32     `json:"status"`
	OrderDate    time.Time `json:"orderDate"`
	PaymentDate  time.Time `json:"paymentDate"`
	DeliveryDate time.Time `json:"deliveryDate"`
	ReceiveDate  time.Time `json:"receiveDate"`
	AppID        string    `json:"appID"`
	Memo         string    `json:"memo"`
	CreateUser   string    `json:"createUser"`
	GmtCreate    time.Time `json:"gmtCreate"`
	ModifyUser   string    `json:"modifyUser"`
	GmtModified  time.Time `json:"gmtModified"`

	SoItems []*SoItem
}

type SoItem struct {
	SysNo         int64   `json:"sysNo"`
	SoSysNo       int64   `json:"soSysNo"`
	ProductSysNo  int64   `json:"productSysNo"`
	ProductName   string  `json:"productName"`
	CostPrice     float64 `json:"costPrice"`
	OriginalPrice float64 `json:"originalPrice"`
	DealPrice     float64 `json:"dealPrice"`
	Quantity      int32   `json:"quantity"`
}

func (dao *Dao) CreateSO(ctx context.Context, soMasters []*SoMaster) ([]uint64, error) {

	result := make([]uint64, 0, len(soMasters))

	_, err := zorm.Transaction(ctx, func(ctx context.Context) (interface{}, error) {

		for _, soMaster := range soMasters {
			soid := NextID()

			var data SoMasterModel
			data.SysNo = int64(soid)
			data.SoID = fmt.Sprintf("%d", soid)
			data.BuyerUserSysNo = soMaster.BuyerUserSysNo
			data.SellerCompanyCode = soMaster.SoID
			data.ReceiveDivisionSysNo = soMaster.ReceiveDivisionSysNo
			data.ReceiveAddress = soMaster.ReceiveAddress
			data.ReceiveZip = soMaster.ReceiveZip
			data.ReceiveContact = soMaster.ReceiveContact
			data.ReceiveContactPhone = soMaster.ReceiveContactPhone
			data.StockSysNo = soMaster.StockSysNo
			data.PaymentType = soMaster.PaymentType
			data.SoAmt = soMaster.SoAmt
			data.Status = soMaster.Status
			data.OrderDate = time.Now()
			data.PaymentDate = time.Now()
			data.DeliveryDate = time.Now()
			data.ReceiveDate = time.Now()
			data.AppID = soMaster.AppID
			data.Memo = soMaster.Memo
			data.CreateUser = soMaster.CreateUser
			data.GmtCreate = time.Now()
			data.ModifyUser = soMaster.ModifyUser
			data.GmtModified = time.Now()
			_, err := zorm.Insert(ctx, &data)
			if err != nil {
				return nil, err
			}

			soItems := soMaster.SoItems
			for _, soItem := range soItems {
				soItemID := NextID()

				var data SoItemModel
				data.SysNo = int64(soItemID)
				data.SoSysNo = int64(soItemID)
				data.ProductSysNo = soItem.ProductSysNo
				data.ProductName = soItem.ProductName
				data.CostPrice = soItem.CostPrice
				data.OriginalPrice = soItem.OriginalPrice
				data.DealPrice = soItem.DealPrice
				data.Quantity = soItem.Quantity
				_, err := zorm.Insert(ctx, &data)
				if err != nil {
					return nil, err
				}
				if err != nil {
					return nil, err
				}
			}
			result = append(result, soid)
		}
		return nil, nil
	})
	return nil, err
}

func NextID() uint64 {
	id, _ := uuid.NewUUID()
	return uint64(id.ID())
}
