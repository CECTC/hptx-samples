package dao

import (
	"time"

	"gitee.com/chunanyong/zorm"
)

type SoMasterModel struct {
	zorm.EntityStruct
	SysNo                int64     `column:"sysno"`
	SoID                 string    `column:"so_id"`
	BuyerUserSysNo       int64     `column:"buyer_user_sysNo"`
	SellerCompanyCode    string    `column:"seller_company_code"`
	ReceiveDivisionSysNo int64     `column:"receive_division_sysNo"`
	ReceiveAddress       string    `column:"receive_address"`
	ReceiveZip           string    `column:"receive_zip"`
	ReceiveContact       string    `column:"receive_contact"`
	ReceiveContactPhone  string    `column:"receive_contact_phone"`
	StockSysNo           int64     `column:"stock_sysno"`
	PaymentType          int32     `column:"payment_type"`
	SoAmt                float64   `column:"so_amt"` //10，创建成功，待支付；30；支付成功，待发货；50；发货成功，待收货；70，确认收货，已完成；90，下单失败；100已作废
	Status               int32     `column:"status"`
	OrderDate            time.Time `column:"order_date"`
	PaymentDate          time.Time `column:"payment_date"`
	DeliveryDate         time.Time `column:"delivery_date"`
	ReceiveDate          time.Time `column:"receive_date"`
	AppID                string    `column:"appid"`
	Memo                 string    `column:"memo"`
	CreateUser           string    `column:"create_user"`
	GmtCreate            time.Time `column:"gmt_create"`
	ModifyUser           string    `column:"modify_user"`
	GmtModified          time.Time `column:"gmt_modified"`
}

func (entity *SoMasterModel) GetTableName() string {
	return "order.so_master"
}

func (entity *SoMasterModel) GetPKColumnName() string {
	return "sysno"
}
