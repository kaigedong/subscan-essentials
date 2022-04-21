package service

import (
	"sync"

	"github.com/itering/subscan/util"
	"github.com/kaigedong/substrate-api-rpc/rpc"
	"github.com/kaigedong/substrate-api-rpc/websocket"
)

var onceToken sync.Once

// Unknown token reg
func (s *Service) unknownToken() {
	websocket.SetEndpoint(util.WSEndPoint)
	onceToken.Do(func() {
		if p, _ := rpc.GetSystemProperties(nil); p != nil {
			util.AddressType = util.IntToString(p.Ss58Format)
			util.BalanceAccuracy = util.IntToString(p.TokenDecimals)
		}
	})
}
