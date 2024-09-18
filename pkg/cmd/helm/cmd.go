package helm

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   `helm COMMAND`,
	Short: `Collection of Helm tools`,
}

func init() {
	Cmd.AddCommand(appVersionForCmd)
}
