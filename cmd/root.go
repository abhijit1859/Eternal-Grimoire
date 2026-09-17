package cmd

import (
 	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
    Use:   "grimoire",
    Short: "A database backup and recovery CLI",
}


func Execute() error{
	return rootCmd.Execute()
}

func init(){
	rootCmd.AddCommand()
	 
}