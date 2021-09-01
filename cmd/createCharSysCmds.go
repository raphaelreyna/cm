package cmd

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/raphaelreyna/cm/internal/maps"
	"github.com/spf13/cobra"
)

func createCharSystemCommands() {
	for name, cm := range maps.All() {
		var (
			charCount = len(cm.Chars)

			newline bool

			cmd = cobra.Command{
				Use:   name,
				Short: fmt.Sprintf("The basic %d %s characters. Capitalized named -> capitalized character.", charCount, name),
				Aliases: []string{
					strings.ToTitle(name),
				},
				ValidArgs: make([]string, charCount),
			}
		)

		cmd.Flags().BoolVarP(&newline, "newline", "n", false, "Print each character on its own line.")

		var idx int
		for charName := range cm.Chars {
			cmd.ValidArgs[idx] = charName
			idx++
		}

		cmd.Run = func(cmd *cobra.Command, args []string) {
			var charSysName = cmd.Use

			switch len(args) {
			case 0:
				cm, err := maps.GetMap(charSysName)
				if err != nil {
					return
				}

				sortedNames := cm.SortedNames()
				writer := tabwriter.NewWriter(os.Stdout, 12, 15, 0, '\t', 0)
				for _, name := range sortedNames {
					c := cm.Chars[name]
					writer.Write([]byte(fmt.Sprintf("%s\t%c\t%c\n", name, c.U, c.L)))
				}
				writer.Flush()
			default:
				var (
					output string
					tmplt  = "%c"
				)

				if newline {
					tmplt += "\n"
				}

				for _, charName := range args {
					var lcName = strings.ToLower(charName)
					char, err := maps.MapChar(charSysName, lcName)
					if err != nil {
						continue
					}

					switch lcName[0] == charName[0] {
					case true:
						output += fmt.Sprintf(tmplt, char.L)
					case false:
						output += fmt.Sprintf(tmplt, char.U)
					}
				}

				if newline {
					output = output[:len(output)-1]
				}

				fmt.Println(output)
			}
		}

		rootCmd.AddCommand(&cmd)
	}
}
