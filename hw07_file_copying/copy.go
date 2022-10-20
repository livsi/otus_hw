package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrOnCreateDestFile      = errors.New("err on create dest file")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	source, err := os.Open(fromPath)
	if err != nil {
		return fmt.Errorf("open file error %s: %w", fromPath, err)
	}
	defer closeFile(source)

	stat, err := source.Stat()
	if err != nil {
		return fmt.Errorf("stat file error: %w", err)
	}

	if !stat.Mode().IsRegular() {
		return ErrUnsupportedFile
	}

	size := stat.Size()
	if offset > size {
		return ErrOffsetExceedsFileSize
	}
	var mustBytesRead int64

	switch {
	case limit == 0:
		mustBytesRead = size - offset
	case limit < (size - offset):
		mustBytesRead = limit
	default:
		mustBytesRead = size - offset
	}

	dest, err := os.Create(toPath)
	if err != nil {
		return ErrOnCreateDestFile
	}
	defer closeFile(dest)

	if _, err := source.Seek(offset, 0); err != nil {
		return err
	}

	var defaultBuf int64 = 1024
	if limit > 0 && limit < defaultBuf {
		defaultBuf = limit
	}
	buf := make([]byte, defaultBuf)

	var bytesRead int64
	bytesRead = 0
	for {
		chunk, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		bytesRead += int64(chunk)
		if _, err := dest.Write(buf[:chunk]); err != nil {
			return err
		}

		if mustBytesRead == bytesRead {
			break
		}
	}

	return nil
}

func closeFile(f *os.File) {
	if err := f.Close(); err != nil {
		_ = fmt.Errorf("error: %w", err)
		os.Exit(1)
	}
}
