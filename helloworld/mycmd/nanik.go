package mycmd

import (
	"flamingo.me/dingo"
	"fmt"
	"github.com/spf13/cobra"
)

type (
	// InitModule initial module for basic setup
	Module struct{}
)

func (*Module) Configure(injector *dingo.Injector) {
	injector.BindMulti(new(cobra.Command)).ToProvider(wohoo)
}

func wohoo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "wohoo",
		Short: "wohoo description",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Wohooo I'm alive")
		},
	}
	return cmd
}
