package kt

import (
	"github.com/rwxrob/kt/pkg/cmd/helm"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     `kt`,
	Short:   `Kubernetes tools`,
	Version: `v0.0.1`, // TODO change this to something triggered by tag
}

func init() {
	Cmd.AddCommand(helm.Cmd)
}
