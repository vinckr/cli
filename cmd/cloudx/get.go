package cloudx

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/ory/cli/cmd/cloudx/client"
	"github.com/ory/cli/cmd/cloudx/identity"
	"github.com/ory/cli/cmd/cloudx/project"

	"github.com/ory/x/cmdx"
)

func NewGetCmd(parent *cobra.Command) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: fmt.Sprintf("Get a resource"),
	}

	cmd.AddCommand(project.NewGetProjectCmd())
	cmd.AddCommand(project.NewGetKratosConfigCmd())
	cmd.AddCommand(identity.NewGetIdentityCmd(parent))

	client.RegisterConfigFlag(cmd.PersistentFlags())
	client.RegisterYesFlag(cmd.PersistentFlags())
	cmdx.RegisterNoiseFlags(cmd.PersistentFlags())

	return cmd
}