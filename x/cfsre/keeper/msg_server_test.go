package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/vedhavyas/cf-sre/testutil/keeper"
	"github.com/vedhavyas/cf-sre/x/cfsre/keeper"
	"github.com/vedhavyas/cf-sre/x/cfsre/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CfsreKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
