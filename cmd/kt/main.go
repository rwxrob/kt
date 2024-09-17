package main

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     `kt`,
	Short:   `Kubernetes tools`,
	Version: `v0.0.1`, // TODO change this to something triggered by tag
}

func main() {
	Cmd.Execute()
}
