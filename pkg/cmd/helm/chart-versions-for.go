package helm

import (
	"fmt"

	"github.com/rwxrob/kt/pkg/helm"
	"github.com/spf13/cobra"
)

var chartVersionsForCmd = &cobra.Command{
	Use:   `chart-versions-for CHART APPVERSION`,
	Short: `Lookup chart versions for specific app version`,
	Args:  cobra.ExactArgs(2),
	RunE: func(x *cobra.Command, args []string) error {
		versions, err := helm.ChartVersionsFor(args[0], args[1])
		fmt.Print(versions)
		return err
	},
}
