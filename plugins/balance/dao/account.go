package dao

import (
	"fmt"

	"github.com/kaigedong/subscan-plugin/storage"
	"github.com/kaigedong/subscan/plugins/balance/model"
	"github.com/kaigedong/subscan/util"
	"github.com/kaigedong/substrate-api-rpc/rpc"
	"github.com/shopspring/decimal"
)

func GetAccountList(db storage.DB, page, row int) ([]model.Account, int) {
	var accounts []model.Account
	opt := storage.Option{PluginPrefix: "balance", Page: page, PageSize: row}
	db.FindBy(&accounts, nil, &opt)
	return accounts, len(accounts)
}

func NewAccount(db storage.DB, accountId string) error {
	accountId = util.TrimHex(accountId)
	_ = db.Create(&model.Account{Address: accountId})
	return AfterAccountCreate(db, accountId)
}

func AfterAccountCreate(db storage.DB, accountId string) error {
	accountDataRaw, err := rpc.ReadStorage(nil, "system", "account", "", accountId)
	if err != nil {
		return err
	}
	accountData := new(model.AccountData)
	accountDataRaw.ToAny(accountData)
	return db.Update(model.Account{}, fmt.Sprintf("address = '%s'", accountId), map[string]interface{}{
		"nonce":   accountData.Nonce,
		"balance": accountData.Data.Free.Add(accountData.Data.Reserved),
		"lock":    decimal.Max(accountData.Data.MiscFrozen, accountData.Data.FeeFrozen),
	})
}
