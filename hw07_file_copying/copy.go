package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	// Place your code here.

	// количество копируемых байт limit, по умолчанию - 0 -- мы заранее знаем сколько хотим прочитать.
	// отступ в источнике (offset), по умолчанию - 0.

	info, err := os.Stat(fromPath)
	if err != nil {
		return err
	}
	if info.IsDir() || !info.Mode().IsRegular() {
		return ErrUnsupportedFile
	}
	fileSize := info.Size()
	if offset >= fileSize { // offset больше, чем размер файла - невалидная ситуация.
		return ErrOffsetExceedsFileSize
	}

	file, err := os.Open(fromPath) // открываем файл для чтения.
	if err != nil {
		return err
	}
	defer file.Close()

	dst, err := os.Create(toPath) // открываем файл для записи.
	if err != nil {
		return err
	}
	defer dst.Close()

	bytesLeft := fileSize - offset
	barSize := limit
	// limit больше, чем размер файла - валидная ситуация, копируется исходный файл до его EOF.
	if bytesLeft < limit || limit == 0 {
		barSize = bytesLeft
	}

	bar := pb.Full.Start64(barSize)
	defer bar.Finish()
	reader := bar.NewProxyReader(file)

	_, err = file.Seek(offset, io.SeekStart) // делаем отступ на offset байт вперёд относительно начала файла.
	if err != nil {
		return err
	}

	if limit == 0 {
		_, err = io.Copy(dst, reader)
	} else {
		_, err = io.CopyN(dst, reader, limit)
	}
	if errors.Is(err, io.EOF) && err != nil {
		return err
	}

	return nil
}
