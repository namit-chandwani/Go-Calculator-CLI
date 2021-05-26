package cmdloader

import (
	"fmt"
	"strconv"

	"github.com/namit-chandwani/Go-Calculator-CLI/pkg/operations"
	"github.com/spf13/cobra"
)

var x = [...]*operations.CommandConfig{operations.AddConfig, operations.SubtractConfig, operations.MultiplyConfig, operations.DivideConfig, operations.RemainderConfig}

// AddCommands adds the arithmetic operation sub-commands to the CLI
func AddCommands(cmd *cobra.Command) {
	for _, cmdCfg := range x {
		cmd.AddCommand(addCommand(cmdCfg))
	}
}

func addCommand(cmdCfg *operations.CommandConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:           cmdCfg.Name,
		Short:         cmdCfg.ShortDescription,
		SilenceErrors: true,
		Example:       cmdCfg.Examples,
		Args:          cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			x, errX := strconv.Atoi(args[0])
			y, errY := strconv.Atoi(args[1])
			if errX != nil {
				return errX
			} else if errY != nil {
				return errY
			}
			fmt.Printf("The result of %s between numbers (%d and %d) is: %d\n", cmdCfg.ShortDescription, x, y, cmdCfg.OperationFunc(x, y))

			return nil
		},
	}
	return cmd
}
