package mintIbc

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/mintIbc/keeper"
	"github.com/cosmos/cosmos-sdk/x/mintIbc/types"
)

func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, ak types.AccountKeeper, data types.GenesisState) {
	ak.GetModuleAccount(ctx, types.ModuleName)
}

func ExportGenesis(ctx sdk.Context, keeper keeper.Keeper) types.GenesisState {
	return *types.NewGenesisState()
}
