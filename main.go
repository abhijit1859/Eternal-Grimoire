package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {

	var cli = &cobra.Command{
		Use:   "Grimoire",
		Short: "A backup database",
	}

	cli.AddCommand(&cobra.Command{
		Use:   "backup",
		Short: "Create a database backup",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Creating database backup...")
		},
	})

	if err := cli.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
