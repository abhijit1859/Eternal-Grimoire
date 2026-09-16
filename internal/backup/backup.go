package backup

import (
	"context"
	"io"
)

type Engine interface{
	Backup(ctx context.Context,dst io.Writer) error
	Restore(ctx context.Context,src io.Reader) error
}