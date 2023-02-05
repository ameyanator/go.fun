package file

import "io"

type File interface {
	Load(string) (io.ReadCloser, error)
	Save(string, io.ReadSeeker) error
}
