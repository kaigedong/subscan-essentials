package service

import (
	"testing"

	"github.com/kaigedong/subscan/model"
	"github.com/shopspring/decimal"
)

func Test_emitEvent(t *testing.T) {
	testSrv.emitEvent(&testBlock, &testEvent, decimal.Zero)
}

func Test_emitExtrinsic(t *testing.T) {
	testSrv.emitExtrinsic(&testBlock, &testSignedExtrinsic, []model.ChainEvent{testEvent})
}
