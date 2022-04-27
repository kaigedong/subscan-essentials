package service

import (
	"github.com/kaigedong/subscan-plugin/storage"
	"github.com/kaigedong/subscan/plugins/balance/dao"
	"github.com/kaigedong/subscan/plugins/balance/model"
)

type Service struct {
	d storage.Dao
}

func (s *Service) GetAccountListJson(page, row int) ([]model.Account, int) {
	return dao.GetAccountList(s.d, page, row)
}

func New(d storage.Dao) *Service {
	return &Service{
		d: d,
	}
}
