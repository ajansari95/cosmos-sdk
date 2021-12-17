package types

import sdk "github.com/cosmos/cosmos-sdk/types"
import authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

type BankKeeper interface {
	MintCoins(ctx sdk.Context, name string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
}
type AccountKeeper interface {
	GetModuleAddress(name string) sdk.AccAddress

	SetModuleAccount(ctx sdk.Context, macc authtypes.ModuleAccountI)
	GetModuleAccount(ctx sdk.Context, moduleName string) authtypes.ModuleAccountI
}
