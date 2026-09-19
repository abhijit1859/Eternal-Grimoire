package cmd

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/abhijit1859/eternal_grimoire/internal/backup"
	"github.com/abhijit1859/eternal_grimoire/internal/compression"
	 
	"github.com/abhijit1859/eternal_grimoire/internal/storage"
	"github.com/spf13/cobra"
)

var restorefile string

var restoreCmd = &cobra.Command{
	Use:   "Restore Database",
	Short: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		if ctx == nil {
			ctx = context.Background()
		}

		startTime := time.Now()

		engine := backup.NewPostgresBackup(cfg.Database)
		compressor := compression.NewGzipCompressor()
		backupDir := "backups"

		if cfg.Storage.Path != "" {
			backupDir = cfg.Storage.Path
		}

		store := storage.NewLocalStorage(backupDir)

		fmt.Printf("Starting to restore database: %s from archive: %s: ...\n", cfg.Database.Database, restorefile)

		archiveStream, err := store.Retrieve(ctx, restorefile)

		if err != nil {
			fmt.Printf("Failed to restore database %w", err)
		}
		defer archiveStream.Close()

		restoreR, restoreW := io.Pipe()

		go func() {
			defer restoreW.Close()
			if err := compressor.Decompress(archiveStream, restoreW); err != nil {
				restoreW.CloseWithError(fmt.Errorf("decompression failed: %w", err))
			}
		}()

		if err := engine.Restore(ctx, restoreR); err != nil {
			return fmt.Errorf("Database restore failed: %w", err)
		}
		fmt.Printf("Database successfully restored Duration %s \n", time.Since(startTime).Round(time.Millisecond))

		return nil

	},
}

func init(){
	restoreCmd.Flags().StringVarP(&restorefile,"file","f","","Backup filename to restore (required)")

	restoreCmd.MarkFlagRequired("file")
	rootCmd.AddCommand(restoreCmd)
}
