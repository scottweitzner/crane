package cmd

import (
	"github.com/scottweitzner/crane/cmd/load"
	"github.com/spf13/cobra"
)

// NewRootCommand returns the command configuration for crane
func NewRootCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "crane",
		Short: "template dockerfiles!",
		Long:  "crane can better template your local dockerfiles or from a source!",
	}
	cmd.AddCommand(load.NewLoadCommand())

	return cmd
}
