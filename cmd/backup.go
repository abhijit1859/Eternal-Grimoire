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

var customBackupName string

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Stream database dump, compress with gzip, and store",

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

 		filename := customBackupName

		if filename == "" {
			filename = fmt.Sprintf(
				"%s_%s.dump.gz",
				cfg.Database.Database,
				time.Now().Format("2006-01-02_15-04-05"),
			)
		}

 		storageR, storageW := io.Pipe()
		dumpR, dumpW := io.Pipe()

 		go func() {
			defer dumpW.Close()

			if err := engine.Backup(ctx, dumpW); err != nil {
				dumpW.CloseWithError(
					fmt.Errorf("database dump failed: %w", err),
				)
			}
		}()

 		go func() {
			defer storageW.Close()

			if err := compressor.Compres(dumpR, storageW); err != nil {
				storageW.CloseWithError(
					fmt.Errorf("compression failed: %w", err),
				)
			}
		}()
 
		if err := store.Store(ctx, filename, storageR); err != nil {
			return fmt.Errorf("storage failed: %w", err)
		}

		fmt.Printf(
			"Backup successfully created in %s (Duration: %s)\n",
			backupDir,
			time.Since(startTime).Round(time.Millisecond),
		)

		return nil
	},
}

func init() {
	backupCmd.Flags().StringVarP(&customBackupName, "name", "n", "", "Custom filename for the backup archive")
	rootCmd.AddCommand(backupCmd)
}
