package load

import (
	"os"
	"text/template"

	"github.com/scottweitzner/crane/common"
	"github.com/scottweitzner/crane/types"
	"github.com/spf13/cobra"
)

type loadOptions struct {
	manifest string
}

// NewLoadCommand returns the configuration for the load command
func NewLoadCommand() *cobra.Command {

	options := &loadOptions{
		manifest: "~/.crane/manifest.yaml",
	}

	cmd := &cobra.Command{
		Use:   "load",
		Short: "load up! create your dockerfile",
		Long:  "load up! templates your dockerfiles according to the manifest",
		RunE: func(cmd *cobra.Command, args []string) error {
			os.RemoveAll(common.GitClonePath) // wipe the temp clone directory
			return common.RedWrapError(options.run())
		},
	}
	cmd.Flags().StringVarP(&options.manifest, "mainfest", "m", options.manifest, "manifest location")
	return cmd
}

func (options *loadOptions) run() error {
	manifest, err := types.ParseManifest(options.manifest)
	if err != nil {
		return err
	}

	dockerFilePath, err := manifest.FormSourcePath()
	if err != nil {
		return err
	}

	template, err := template.ParseFiles(dockerFilePath)
	if err != nil {
		return err
	}

	outFile, err := os.Create(manifest.FormOutputPath())
	if err != nil {
		return err
	}

	if err = template.Execute(outFile, manifest.Values); err != nil {
		return err
	}

	return nil
}
