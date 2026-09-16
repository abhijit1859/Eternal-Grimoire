package compression

import "io"


type Compress interface{
	Compress(src io.Reader,dst io.Writer)
	Decompress(src io.Reader,dst io.Writer)
}