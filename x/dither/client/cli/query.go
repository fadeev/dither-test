package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/virgo-project/dither/x/dither/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	ditherQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	ditherQueryCmd.AddCommand(
		flags.GetCommands(
			GetCmdListPosts(queryRoute, cdc),
			GetCmdListChannels(queryRoute, cdc),
		)...,
	)

	return ditherQueryCmd

}

// GetCmdListPosts ...
func GetCmdListPosts(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "listPosts [channel_id]",
		Short: "list posts",
		Args:  cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var channelID string
			if len(args) == 1 {
				channelID = args[0]
			}
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QueryListPosts, channelID), nil)
			if err != nil {
				fmt.Printf("could not get posts\n%s\n", err.Error())
				return nil
			}
			var out types.Posts
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdListChannels ...
func GetCmdListChannels(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "listChannels",
		Short: "list all channels",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/"+types.QueryListChannels, queryRoute), nil)
			if err != nil {
				fmt.Printf("could not list channels\n%s\n", err.Error())
				return nil
			}
			var out types.Channels
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
