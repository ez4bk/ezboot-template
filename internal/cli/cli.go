package cli

import (
	"github.com/google/wire"
	"github.com/spf13/cobra"

	"example.ezboot/myna/internal/cli/version"
)

// Commands 命令行集合
type Commands []*cobra.Command

// NewCommands 注册命令行
func NewCommands() Commands {
	return Commands{
		version.CmdlineMain(),
	}
}

var ProviderSet = wire.NewSet(
	NewCommands,
)
