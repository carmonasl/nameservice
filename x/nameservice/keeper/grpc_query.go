package keeper

import (
	"github.com/carmonasl/nameservice/x/nameservice/types"
)

var _ types.QueryServer = Keeper{}
