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

// TODO: 从两张表中查询！查询后结果结合起来返回
func (d *Dao) GetTransferList(c context.Context, hexAddress string, page int, row int) ([]model.TransferJson, int) {
	var outs []model.TransferJson

	// 从Event表中查询
	var events []model.ChainEvent

	query := d.db.Model(model.ChainEvent{}).
		Where(model.ChainEvent{ModuleId: "balances", EventId: "Transfer"}).
		Where("params LIKE ?", "%"+hexAddress+"%").
		Scan(&events)

	if query == nil || query.RecordNotFound() {
		return nil, 0
	}

	// 对于每一个Event，从Extrinsic表中查询对应的extrinsic
	// TODO: 根据event-index,从extrinsic_index进行查询
	for _, aEvent := range events {

		var transferEvent TransferEvent
		if err := json.Unmarshal([]byte(util.ToString(aEvent.Params)), &transferEvent); err != nil {
			fmt.Println("##### ERR: ", err)
			continue
		}
		if len(transferEvent) != 3 {
			fmt.Println("Transfer event decoded len must be 3")
			continue
		}

		var aExtrinsic model.ChainExtrinsic
		newQuery := d.db.Model(model.ChainExtrinsic{}).Where(model.ChainExtrinsic{
			ExtrinsicIndex: aEvent.EventIndex,
		}).Find(&aExtrinsic)

		if newQuery == nil || newQuery.RecordNotFound() {
			return nil, 0
		}

		amount, err := decimal.NewFromString(transferEvent[2].Value)
		if err != nil {
			continue
		}

		outs = append(outs, model.TransferJson{
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

	return outs, len(outs)
}
