package file

import (
	"io"
	"log"
	"os"
	"path/filepath"

	error2 "goinpractice.com/GoInPractice/main/technique57/error2"
)

type LocalFile struct {
	Base string
}

func (l LocalFile) Load(path string) (io.ReadCloser, error) {
	p := filepath.Join(l.Base, path)
	o, err := os.Open(p)
	if err != nil && os.IsNotExist(err) {
		log.Printf("Unable to find %s", path)
		return nil, error2.ErrFileNotFound
	} else if err != nil {
		log.Printf("Error loading file %s, err: %s", path, err)
		return nil, error2.ErrCannotLoadFile
	}
	return o, nil
}

/*
Logging the original error is important. If a problem occurs when connecting to the
remote system, that problem needs to be logged. A monitoring system can catch
errors communicating with external systems and raise alerts so you have an opportu-
nity to remediate the problems.
*/

func (l LocalFile) Save(path string, body io.ReadSeeker) error {
	p := filepath.Join(l.Base, path)
	d := filepath.Dir(p)
	err := os.Mkdir(d, os.ModeDir|os.ModePerm)
	if err != nil {
		return err
	}
	f, err := os.Create(p)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, body)
	return err
}
