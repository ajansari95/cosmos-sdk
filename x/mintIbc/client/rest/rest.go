package rest

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/rest"
	"github.com/gorilla/mux"
)

//RegisterHandlers registers the handlers for the REST interface

func RegisterHandlers(clientCtx client.Context, rtr *mux.Router) {
	r := rest.WithHTTPDeprecationHeaders(rtr)
	r.HandleFunc("/mintIbc/{address}{amount}", mintIbcHandlerFn(clientCtx)).Methods("POST")
}
