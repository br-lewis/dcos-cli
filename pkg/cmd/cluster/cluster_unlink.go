package cluster

import (
	"github.com/dcos/dcos-cli/api"
	"github.com/dcos/dcos-cli/pkg/clusterlinker"
	"github.com/dcos/dcos-cli/pkg/config"
	"github.com/spf13/cobra"
)

// newCmdClusterLink links the attached cluster to another one.
func newCmdClusterUnlink(ctx api.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "unlink",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			attachedCluster, err := ctx.Cluster()
			if err != nil {
				return err
			}

			manager := ctx.ConfigManager()
			linkedClusterConfig, err := manager.Find(args[0], false)
			if err != nil {
				return err
			}
			linkedCluster := config.NewCluster(linkedClusterConfig)

			attachedClient := clusterlinker.NewClient(ctx.HTTPClient(attachedCluster), ctx.Logger())
			return attachedClient.Unlink(linkedCluster.ID())
		},
	}
	return cmd
}
