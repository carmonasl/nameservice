package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/carmonasl/nameservice/testutil/keeper"
	"github.com/carmonasl/nameservice/x/nameservice/keeper"
	"github.com/carmonasl/nameservice/x/nameservice/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.NameserviceKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
