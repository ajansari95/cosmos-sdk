package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/mintIbc/testConstants"
	"github.com/cosmos/cosmos-sdk/x/mintIbc/types"
	"github.com/spf13/cobra"
)

func NewTxCmd() *cobra.Command {

	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "IBC mint",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(NewMintIbcCmd())
	return txCmd

}

func NewMintIbcCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint [address] [amount]",
		Short: "mint ibc tokens",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Set(flags.FlagFrom, args[0])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			fmt.Println(1)
			addr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			fmt.Println(2)
			amount, ok := sdk.NewIntFromString(args[1])
			if !ok {
				return err
			}
			msg := types.NewMintIBCMsg(addr, sdk.NewCoin(testConstants.IbcDenom, amount))
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd

}
