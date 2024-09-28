package helm

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:     `helm COMMAND`,
	Short:   `Collection of Helm tools`,
	Version: `v0.0.1`,
}

func init() {
	Cmd.AddCommand(appVersionForCmd)
	Cmd.AddCommand(chartVersionsForCmd)
}
