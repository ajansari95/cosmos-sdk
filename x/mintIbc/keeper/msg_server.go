package keeper

import (
	"context"
	"fmt"
	"github.com/armon/go-metrics"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/mintIbc/testConstants"
	"github.com/cosmos/cosmos-sdk/x/mintIbc/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returnd an implemenatation of the mintIBc interface
var _ types.MsgServer = msgServer{}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

func (m msgServer) MintIBC(goCtx context.Context, msg *types.MintIBCMsg) (*types.MintIBCResp, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	toAddr, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return nil, err
	}
	tokens, err := sdk.ParseCoinNormalized(msg.Amount)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidToken, "invalid TOKENnn: %s", msg.Amount)
	}

	err = m.MintCoins(ctx, sdk.NewCoins(tokens))
	fmt.Println("2345")

	if err != nil {
		return nil, err
	}
	if err := m.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, toAddr, sdk.NewCoins(tokens)); err != nil {
		return nil, err
	}
	defer func() {
		telemetry.IncrCounter(1, types.ModuleName, "delegate")
		telemetry.SetGaugeWithLabels(
			[]string{"mintIbc", "msg", msg.Type()},
			float32(tokens.Amount.Int64()),
			[]metrics.Label{telemetry.NewLabel("denom", testConstants.IbcDenom)},
		)
	}()

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return &types.MintIBCResp{}, nil
}
