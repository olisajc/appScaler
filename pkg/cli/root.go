package cli

import (
	"fmt"
	"path/filepath"

	kc "github.com/olisajc/appScaler/pkg/kubeclient"
	"github.com/spf13/cobra"
	"k8s.io/client-go/util/homedir"
)

type RootOptions struct {
	development bool
	cluster     bool
}

var ConfigPath = filepath.Join(homedir.HomeDir(), ".kube", "config")

func Root() *cobra.Command {

	options := &RootOptions{}

	cmd := &cobra.Command{
		Use:   "",
		Short: "",
		Long:  `.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := createClient(options); err != nil {
				return err
			}

			return nil
		},
	}

	cmd.PersistentFlags().BoolVar(&options.development, "development", false, "Use development kubeconfig")
	cmd.PersistentFlags().BoolVar(&options.cluster, "cluster", false, "Use in-cluster kubeconfig")
	cmd.Flags().String("kubeconfig", ConfigPath, "absolute path to the kubeconfig file")

	return cmd
}

func createClient(options *RootOptions) error {

	if !options.development && !options.cluster {
		return nil
	}

	if options.development && options.cluster {
		return kc.ErrMultipleConfigSources
	}

	fmt.Printf("development %v\n", options.development)

	_, err := kc.InitKubeClient(options.development)
	return err
}
