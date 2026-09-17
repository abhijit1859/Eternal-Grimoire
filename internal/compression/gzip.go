package compression

import (
	"compress/gzip"
	"fmt"
	"io"
)

 
type GzipCompressor struct{}


func NewGzipCompressor() *GzipCompressor{
	return &GzipCompressor{}
}

func (g *GzipCompressor) Compres(src io.Reader,dst io.Writer) error {
	gzipWriter:=gzip.NewWriter(dst)

	_,err:=io.Copy(gzipWriter,src)

	if err!=nil{
		gzipWriter.Close()
		return fmt.Errorf("compression failed %w",err)
	}

	if err := gzipWriter.Close(); err != nil {
		return fmt.Errorf("failed to finalize gzip: %w", err)
	}

	return nil
}	

func (g *GzipCompressor) Decompress (src io.Reader,dst io.Writer) error{
	reader,err:=gzip.NewReader(src)
	if err!=nil{
	 
		return fmt.Errorf("decompression failed %w",err)
	}

	defer reader.Close()

	if _,err:=io.Copy(dst,reader);err!=nil{
		return fmt.Errorf("decompression failed: %w", err)
	}
	return nil
}