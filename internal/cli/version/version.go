package version

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

func CmdlineMain() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print version/build info",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Version:\t\t n/a\n")
			fmt.Printf("GoVersion:\t\t %s\n", runtime.Version())
		},
	}
}
