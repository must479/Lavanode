package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/lavanet/lava/x/subscription/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdListProjects() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-projects [subscription]",
		Short: "Query to show all the subscription&#39;s projects",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqSubscription := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryListProjectsRequest{
				Subscription: reqSubscription,
			}

			res, err := queryClient.ListProjects(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
