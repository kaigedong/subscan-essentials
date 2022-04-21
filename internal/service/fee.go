package service

import (
	"github.com/kaigedong/substrate-api-rpc/rpc"
	"github.com/kaigedong/substrate-api-rpc/websocket"
	"github.com/shopspring/decimal"
)

// GetExtrinsicFee
func GetExtrinsicFee(p websocket.WsConn, encodeExtrinsic string) (fee decimal.Decimal, err error) {
	paymentInfo, err := rpc.GetPaymentQueryInfo(p, encodeExtrinsic)
	if paymentInfo != nil {
		return paymentInfo.PartialFee, nil
	}
	return decimal.Zero, err
}
