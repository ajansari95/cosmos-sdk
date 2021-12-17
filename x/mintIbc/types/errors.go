package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrInvalidAddress = sdkerrors.Register(ModuleName, 2, "invalid address")
	ErrInvalidToken   = sdkerrors.Register(ModuleName, 3, "invalid token")
)
