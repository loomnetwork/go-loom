package commands

import (
	"github.com/spf13/cobra"
)

func Add(cmd *cobra.Command) {
	AddDPOS(cmd)
}
