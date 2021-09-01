package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/raphaelreyna/cm/internal/maps"
	"github.com/spf13/cobra"
)

var dumpCmd = cobra.Command{
	Use:   "dump",
	Short: "Dump all character maps.",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var w = tabwriter.NewWriter(os.Stdout, 12, 15, 0, '\t', 0)

		for name, cm := range maps.All() {
			w.Write([]byte(fmt.Sprintf("%s\t\t\t\n", name)))
			for charName, char := range cm.Chars {
				w.Write([]byte(fmt.Sprintf("\t%s\t%c\t%c\n", charName, char.U, char.L)))
			}
		}

		w.Flush()
	},
}
