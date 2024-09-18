package helm

import (
	"github.com/rwxrob/kt/pkg/helm"
	"github.com/rwxrob/kt/pkg/util"
	"github.com/spf13/cobra"
)

var appVersionForCmd = &cobra.Command{
	Use:   `app-version-for CHART VERSION`,
	Short: `Lookup app version for given chart version`,
	Args:  cobra.ExactArgs(2),
	RunE: func(x *cobra.Command, args []string) error {
		version, err := helm.AppVersionFor(args[0], args[1])
		util.SmartPrint(version)
		return err
	},
}
