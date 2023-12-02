package util

import (
	"bufio"
	"os"
)

type FileScanner struct {
	file    *os.File
	scanner *bufio.Scanner
}

func NewScanner(path string) *FileScanner {
	file, err := os.Open(path)
	CheckErr(err)
	
	return &FileScanner{
		file:    file,
		scanner: bufio.NewScanner(file),
	}
}

func (fs FileScanner) Scan() bool {
	return fs.scanner.Scan()
}

func (fs FileScanner) Text() string {
	return fs.scanner.Text()
}

func (fs FileScanner) Bytes() []byte {
	return fs.scanner.Bytes()
}

func (fs FileScanner) Error() error {
	return fs.scanner.Err()
}

func (fs FileScanner) Close() {
	CheckErr(fs.file.Close())
}