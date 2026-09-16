package storage

import (
	"context"
	"io"
)


type StorageProvider interface{
	Store(ctx context.Context,name string,data io.Reader)
	Retrieve(ctx context.Context,name string) (io.ReadCloser,error)
	List(ctx context.Context) ([]string,error)
}