package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgMintIbc = "mintibc"
)

var _ sdk.Msg = &MintIBCMsg{}

func NewMintIBCMsg(
	address sdk.AccAddress,
	Token sdk.Coin,
) *MintIBCMsg {
	return &MintIBCMsg{
		FromAddress: address.String(),
		Amount:      Token.String(),
	}
}

func (msg MintIBCMsg) Route() string { return RouterKey }

func (msg MintIBCMsg) Type() string { return TypeMsgMintIbc }

func (msg MintIBCMsg) ValidateBasic() error {
	if msg.FromAddress == "" {
		return sdkerrors.Wrapf(ErrInvalidAddress, "invalid address: %s", msg.FromAddress)
	}
	if msg.Amount == "0" {
		return sdkerrors.Wrapf(ErrInvalidToken, "invalid TOKEN: %s", msg.Amount)
	}
	return nil
}

func (msg MintIBCMsg) GetSigners() []sdk.AccAddress {

	address, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{address}
}
