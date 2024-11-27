package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func createSelect() {
	prompt := promptui.Select{
		Label: "Select day",
		Items: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"},
	}

	//int, posicion o id de la lista ; string de la opcion seleccionada; tipo error si es distinta a null fallo el Run.
	//con _ indicamos que no lo vamos a usar.
	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", &result)

}
func main() {

	var cmdPrintSelecet = &cobra.Command{
		Use:   "select",
		Short: "Show a select.",
		Long:  "Display a selector of thge option custom",
		Run: func(cmd *cobra.Command, args []string) {
			createSelect()
		},
	}

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdPrintSelecet)
	rootCmd.Execute()

}
