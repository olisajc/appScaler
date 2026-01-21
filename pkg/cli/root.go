package cli

import (
	"fmt"
	"path/filepath"

	"github.com/olisajc/appScaler/pkg/cli/policyscaler"
	kc "github.com/olisajc/appScaler/pkg/kubeclient"
	"github.com/spf13/cobra"
	"k8s.io/client-go/util/homedir"
)

var ConfigPath = filepath.Join(homedir.HomeDir(), ".kube", "config")

func Root() *cobra.Command {

	options := &RootOptions{}

	cmd := &cobra.Command{
		Use:   "",
		Short: "",
		Long:  `.`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := createClient(options); err != nil {
				return err
			}

			return nil
		},
	}
	cmd.AddCommand(policyscaler.PolicyScalerCmd())
	cmd.PersistentFlags().BoolVar(&options.Development, "development", false, "Use development kubeconfig")
	cmd.PersistentFlags().BoolVar(&options.Cluster, "cluster", false, "Use in-cluster kubeconfig")
	cmd.Flags().String("kubeconfig", ConfigPath, "absolute path to the kubeconfig file")

	return cmd
}

func createClient(options *RootOptions) error {

	if !options.Development && !options.Cluster {
		return nil
	}

	if options.Development && options.Cluster {
		return kc.ErrMultipleConfigSources
	}

	fmt.Printf("development %v\n", options.Development)

	_, err := kc.InitKubeClient(options.Development)
	return err
}
