package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra" //go get github.com/spf13/cobra@latest
)

func main() {
	var cmdPrint = &cobra.Command{
		Use:   "print [string to print]",
		Short: "Print anything to the screen",
		Long:  `print is for printing anything back to the screen.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Print: " + strings.Join(args, " "))
		},
	}

	// comando principal
	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdPrint)
	rootCmd.Execute()
}
