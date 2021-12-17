package mintIbc

import (
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/mintIbc/keeper"
	"github.com/cosmos/cosmos-sdk/x/mintIbc/testConstants"
	"github.com/cosmos/cosmos-sdk/x/mintIbc/types"
	"time"
)

func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)
	mintedCoins := sdk.NewCoins(sdk.NewCoin(testConstants.IbcDenom, sdk.NewInt(1)))

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.ModuleName,
			sdk.NewAttribute(types.ModuleName, mintedCoins.String()),
		))
}
