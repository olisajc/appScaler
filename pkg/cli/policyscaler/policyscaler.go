package policyscaler

import (
	"github.com/olisajc/appScaler/pkg/cli/policyscaler/get"
	"github.com/spf13/cobra"
)

func PolicyScalerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "policyscaler",
		Short: "crud operations for PolicyScaler resources",
	}

	cmd.AddCommand(get.GetCmd())

	return cmd

}
