package svc

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"gitee.com/chunanyong/zorm"

	"github.com/cectc/hptx-samples/order_svc/dao"
	dao2 "github.com/cectc/hptx-samples/product_svc/dao"
)

type Svc struct {
}

func (svc *Svc) CreateSo(ctx context.Context, rollback bool) error {

	// 手动开启zorm全局事务
	globalTransactionCtx, _ := zorm.BindContextEnableGlobalTransaction(ctx)

	_, err := zorm.Transaction(globalTransactionCtx, func(ctx context.Context) (interface{}, error) {

		xid := ctx.Value("XID").(string)
		log.Printf("xid: %s", xid)

		soMasters := []*dao.SoMaster{
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
				SoItems: []*dao.SoItem{
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

		reqs := []*dao2.AllocateInventoryReq{{
			ProductSysNo: 1,
			Qty:          2,
		}}

		type rq1 struct {
			Req []*dao.SoMaster
		}

		type rq2 struct {
			Req []*dao2.AllocateInventoryReq
		}

		q1 := &rq1{Req: soMasters}
		soReq, err := json.Marshal(q1)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		fmt.Println(string(soReq))

		req1, err := http.NewRequest("POST", "http://order-svc:8002/createSo", bytes.NewBuffer(soReq))
		if err != nil {
			log.Println(err)
			return nil, err
		}
		req1.Header.Set("Content-Type", "application/json")
		req1.Header.Set("xid", xid)

		client := &http.Client{}
		result1, err1 := client.Do(req1)
		if err1 != nil {
			return nil, err1
		}

		if result1.StatusCode == 400 {
			return nil, errors.New("err")
		}

		q2 := &rq2{
			Req: reqs,
		}
		ivtReq, _ := json.Marshal(q2)
		fmt.Println(string(ivtReq))
		req2, err := http.NewRequest("POST", "http://product-svc:8001/allocateInventory", bytes.NewBuffer(ivtReq))
		if err != nil {
			panic(err)
		}
		req2.Header.Set("Content-Type", "application/json")
		req2.Header.Set("xid", xid)

		result2, err2 := client.Do(req2)
		if err2 != nil {
			return nil, err2
		}

		if result2.StatusCode == 400 {
			return nil, errors.New("err")
		}

		if rollback {
			return nil, errors.New("there is a error")
		}
		return nil, nil
	})
	return err
}
