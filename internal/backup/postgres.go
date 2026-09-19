package backup

import (
	"bytes"
	"context"
	"io"
	"os/exec"

	"fmt"

	"github.com/abhijit1859/eternal_grimoire/internal/config"
)

type PostgresBackup struct {
	config config.DatabaseConfig
}

func NewPostgresBackup(config config.DatabaseConfig) *PostgresBackup {
	return &PostgresBackup{config: config}
}

func (p *PostgresBackup) Backup(ctx context.Context, dst io.Writer) error {
	fmt.Println("config name",p.config.Username)
	//command to backup db
	cmd := exec.CommandContext(ctx, "pg_dump",
		"-h", p.config.Host,
		"-p", fmt.Sprintf("%d", p.config.Port),
		"-U", p.config.Username,
		"-d", p.config.Database,
		"-F", "c",  
	)

	cmd.Stdout = dst

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("backup failed: %w: %s",
			err,
			stderr.String())
	} 

	return nil

}

func (p *PostgresBackup) Restore(ctx context.Context, src io.Reader) error {
	cmd := exec.CommandContext(ctx, "pg_restore",
		"-h", p.config.Host,
		"-p", fmt.Sprintf("%d", p.config.Port),
		"-U", p.config.Username,
		"-d", p.config.Database,
		"-F", "c",
		"-",
	)

	output, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf(
			"restore failed: %w: %s",
			err,
			output,
		)
	}

	return nil
}
