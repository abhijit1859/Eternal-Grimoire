package cmd

import (
	"context"
	"fmt"
 
	"github.com/abhijit1859/eternal_grimoire/internal/storage"
	"github.com/spf13/cobra"
)



var listCmd=&cobra.Command{
	Use: "List backups",
	Short: "List all available backups in the storage backend",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx:=cmd.Context()
		if ctx==nil{
			ctx=context.Background()
		}

		backupDir:="backups"

		if cfg.Storage.Path!=""{
			backupDir=cfg.Storage.Path
		}

		store:=storage.NewLocalStorage(backupDir)

		backups,err:=store.List(ctx)
		if err!=nil{
			return fmt.Errorf("failed to list backups because of error: %w",err)

		}

		if len(backups)==0{
			fmt.Printf("No backups found.")
			return nil
		}

		fmt.Printf("Found %d backup(s) in %s : \n",len(backups),backupDir)

		for i,name:=range backups{
			fmt.Printf("[%d] %s\n",i+1,name)
		}
		return nil
	},
}

func init(){
	rootCmd.AddCommand(listCmd)
}