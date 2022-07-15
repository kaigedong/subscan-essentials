package dao

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kaigedong/subscan/model"
	"github.com/kaigedong/subscan/util"
	"github.com/kaigedong/subscan/util/address"
	"github.com/shopspring/decimal"
)

type TransferEvent []OneField

type OneField struct {
	Type     string `json:"type"`
	TypeName string `json:"type_name"`
	Value    string `json:"value"`
}

// NOTE: 从两张表中查询！查询后结果结合起来返回
// NOTE: 允许page, row参数
func (d *Dao) GetTransferList(c context.Context, hexAddress string, page int, row int) ([]model.TransferJson, int) {
	// 从Event表中查询到的transfer Event
	var events []model.ChainEvent

	// 根据transfer Event获取到的转账记录
	var transfer []model.TransferJson
	blockNum, _ := d.GetFillBestBlockNum(context.TODO())

	for index := blockNum / model.SplitTableBlockNum; index >= 0; index-- {
		query := d.db.Model(model.ChainEvent{BlockNum: index * model.SplitTableBlockNum}).
			Where("module_id='balances' and event_id='Transfer' and params LIKE ?", "%"+hexAddress+"%").
			Scan(&events)

		if query == nil || query.RecordNotFound() {
			continue
		}

		// 对于每一个Event，从Extrinsic表中查询对应的extrinsic
		// 根据event-index,从extrinsic_index进行查询
		for _, aEvent := range events {
			var transferEvent TransferEvent
			if err := json.Unmarshal([]byte(util.ToString(aEvent.Params)), &transferEvent); err != nil {
				continue
			}
			if len(transferEvent) != 3 {
				fmt.Println("Transfer event decoded len must be 3")
				continue
			}
			var aExtrinsic model.ChainExtrinsic
			// NOTE: .Find找不到, .Limit同样会导致找不到
			newQuery := d.db.Model(model.ChainExtrinsic{BlockNum: aEvent.BlockNum}).
				Where("extrinsic_index=?", aEvent.EventIndex).Scan(&aExtrinsic)

			if newQuery == nil || newQuery.RecordNotFound() {
				fmt.Println("### Event Not Found in Extrinsics.....")
				continue
			}

			amount, err := decimal.NewFromString(transferEvent[2].Value)
			if err != nil {
				continue
			}

			transfer = append(transfer, model.TransferJson{
				From:           address.SS58Address(transferEvent[0].Value), // 将From转换成SS58
				To:             address.SS58Address(transferEvent[1].Value),
				Module:         aExtrinsic.CallModule,
				Amount:         amount,
				Hash:           aExtrinsic.ExtrinsicHash,
				BlockTimestamp: aExtrinsic.BlockTimestamp,
				BlockNum:       aExtrinsic.BlockNum,
				ExtrinsicIndex: aExtrinsic.ExtrinsicIndex,
				Success:        true,
				Fee:            aExtrinsic.Fee,
			})
		}

	}

	return transfer, len(transfer)
}
