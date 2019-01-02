// +build !evm

package commands

import (
	"github.com/spf13/cobra"
)

func AddValidatorCommands(root *cobra.Command) {
	// validator commands have a dependency on go-ethereum so can't build them in non-evm builds
}
