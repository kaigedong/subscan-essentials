package service

import (
	"fmt"
	"strings"

	"github.com/kaigedong/subscan/internal/dao"
	"github.com/kaigedong/subscan/model"
	"github.com/kaigedong/subscan/util"
	"github.com/kaigedong/substrate-api-rpc"
	"github.com/kaigedong/substrate-api-rpc/storage"
)

func (s *Service) EmitLog(txn *dao.GormDB, blockNum int, l []storage.DecoderLog, finalized bool, validatorList []string) (validator string, err error) {
	for index, logData := range l {
		dataStr := util.ToString(logData.Value)

		ce := model.ChainLog{
			LogIndex:  fmt.Sprintf("%d-%d", blockNum, index),
			BlockNum:  blockNum,
			LogType:   logData.Type,
			Data:      dataStr,
			Finalized: finalized,
		}
		if err = s.dao.CreateLog(txn, &ce); err != nil {
			return "", err
		}

		// check validator
		if strings.EqualFold(ce.LogType, "PreRuntime") {
			validator = substrate.ExtractAuthor([]byte(dataStr), validatorList)
		}

	}
	return validator, err
}
