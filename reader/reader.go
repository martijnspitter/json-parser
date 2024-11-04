package reader

import (
	"errors"
	"os"
)

type reader struct {
	filepath string
	content  string
}

func NewReader(filepath string) *reader {
	return &reader{filepath: filepath}
}

func (r *reader) Read() error {
	return r.getFile()
}

func (r *reader) getFile() error {
	// Open file
	file, err := os.Open(r.filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Read file
	stat, err := file.Stat()
	if err != nil {
		return err
	}
	if stat.IsDir() {
		return errors.New("file is a directory")
	}

	content, err := os.ReadFile(r.filepath)
	if err != nil {
		return err
	}

	r.content = string(content)

	return nil
}

func (r *reader) GetContent() string {
	return r.content
}
