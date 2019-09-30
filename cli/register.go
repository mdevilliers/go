package cli

import (
	"github.com/spf13/cobra"
)

type cmdCreator func() (*cobra.Command, error)

// RegisterCommands allows you to attach your commands for execution or return the first error found
func RegisterCommands(root *cobra.Command, cmds ...cmdCreator) error {

	for i := range cmds {

		cmd, err := cmds[i]()
		if err != nil {
			return err
		}

		root.AddCommand(cmd)
	}
	return nil
}
