package error

import "errors"

var (
	ErrFileNotFound   = errors.New("File Not Found")
	ErrCannotLoadFile = errors.New("Unable to load file")
	ErrCannotSaveFile = errors.New("Unable to save file")
)
