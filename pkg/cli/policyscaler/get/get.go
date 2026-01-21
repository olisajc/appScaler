package get

import (
	"context"
	"fmt"

	"github.com/olisajc/appScaler/pkg/kubeclient"
	"github.com/olisajc/appScaler/pkg/policyscaler"
	"github.com/spf13/cobra"
)

func GetCmd() *cobra.Command {
	options := &GetOptions{}
	cmd := &cobra.Command{
		Use:     "get",
		Short:   "get a PolicyScaler resource",
		Example: ` appScaler policyscaler get --name my-policyscaler --namespace default`,
		RunE: func(cmd *cobra.Command, args []string) error {
			k8sClient, err := kubeclient.InitKubeClient(true)
			if err != nil {
				return fmt.Errorf("failed to initialize kubeclient: %v", err)
			}
			Service := policyscaler.NewPolicyScalerService()

			return GetPolicyScaler(context.Background(), options, k8sClient, Service)
		},
	}
	cmd.PersistentFlags().StringVarP(&options.Name, "name", "n", "", "Name of the PolicyScaler resource")
	cmd.PersistentFlags().StringVarP(&options.Namespace, "namespace", "s", "default", "Namespace of the PolicyScaler resource")

	return cmd
}

func GetPolicyScaler(ctx context.Context, options *GetOptions, k8client *kubeclient.KubeClient, service *policyscaler.PolicyScalerService) error {
	if options.Name == "" {
		return fmt.Errorf("policy scaler name is required")
	}

	if service == nil || k8client == nil {
		return fmt.Errorf("internal error while retrieving policy scaler")
	}

	pScaler, err := service.Get(ctx, k8client, options.Name, options.Namespace)

	if err != nil {
		return fmt.Errorf("error Getting policy scaler : %s", options.Name)
	}

	if pScaler == nil {
		return fmt.Errorf("nil policy scaler returned for Get Operation on policy scaler : %s", options.Name)
	}

	fmt.Println(pScaler)

	return nil
}
