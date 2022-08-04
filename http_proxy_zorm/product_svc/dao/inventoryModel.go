package dao

import "gitee.com/chunanyong/zorm"

type InventoryModel struct {
	zorm.EntityStruct

	Sysno           int64 `column:"sysno"`
	ProductSysno    int64 `column:"product_sysno"`
	AccountQty      int64 `column:"account_qty"`
	AvailableQty    int64 `column:"available_qty"`
	AllocatedQty    int64 `column:"allocated_qty"`
	AdjustLockedQty int64 `column:"adjust_locked_qty"`
}

func (entity *InventoryModel) GetTableName() string {
	return "order.inventory"
}

func (entity *InventoryModel) GetPKColumnName() string {
	return "sysno"
}
