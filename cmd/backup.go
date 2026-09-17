package cmd

import (
	"fmt"
	"time"

	"github.com/abhijit1859/eternal_grimoire/internal/backup"
	"github.com/abhijit1859/eternal_grimoire/internal/compression"
	"github.com/abhijit1859/eternal_grimoire/internal/config"
	"github.com/abhijit1859/eternal_grimoire/internal/storage"
)

var backupCmd=&cobra.Command{
	Use:"Backup Database",
	Short:"Stream database dump, compress with gzip, and store",
	RunE:func(cmd *cobra.Command,args []string)error{
		ctx:=cmd.Context()
		if ctx==""{
			ctx=context.Background()
		}

		startTime:=time.Now()

		engine:=backup.NewPostgresBackup()
		compression:=compression.NewGzipCompressor()

		backupDir:="backups"
		if config.StorageConfig.Path==""{
			backupDir=config.StorageConfig.Path
		}

		storage:=storage.NewLocalStorage(backupDir)


	}
}


 