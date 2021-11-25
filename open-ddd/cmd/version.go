package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

const (
	versionNum = "v0.0.1"
)

var (
	VersionCmd = &cobra.Command{
		Use:   "version",
		Short: "Output current version number",
		Long:  `Output current version number`,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			fmt.Println("open-ddd " + versionNum)
			return
		},
	}
)

func init() {
	AddCommand(VersionCmd)
}
