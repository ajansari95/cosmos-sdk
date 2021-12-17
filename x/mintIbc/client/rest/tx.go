package rest

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/mintIbc/types"
	"github.com/gorilla/mux"
	"net/http"
)

type MintToken struct {
	TargetAdress sdk.Address
	ibcToken     sdk.Coin
}

type SendTxRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	MintToken
}

func mintIbcHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bech32Addr := vars["address"]
		toAddr, err := sdk.AccAddressFromBech32(bech32Addr)
		if rest.CheckBadRequestError(w, err) {
			return
		}
		var req SendTxRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}
		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}
		amount := vars["amount"]
		var IbcToken sdk.Coin
		IbcToken.Denom = "ibc/transfer/channel-0/stake"
		IbcToken.Amount, _ = sdk.NewIntFromString(amount)

		msg := types.NewMintIBCMsg(toAddr, IbcToken)
		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
