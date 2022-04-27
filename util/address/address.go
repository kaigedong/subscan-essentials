package address

import (
	"github.com/kaigedong/subscan/util"
	"github.com/kaigedong/subscan/util/ss58"
)

func SS58Address(address string) string {
	return ss58.Encode(address, util.StringToInt(util.AddressType))
}
