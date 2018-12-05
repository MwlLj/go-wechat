package utils

import (
	"bufio"
	"os"
)

func FileIsEixsts(path *string) bool {
	_, err := os.Stat(*path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

type PicturePath2StreamCallback func(buf []byte)

func PicturePath2Stream(path *string, callback PicturePath2StreamCallback) error {
	file, err := os.Open(*path)
	if err != nil {
		return err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		buf := make([]byte, 1024)
		n, err := reader.Read(buf)
		if err != nil && err.Error() != "EOF" {
			return err
		}
		if n == 0 {
			break
		}
		if callback != nil {
			callback(buf)
		}
	}
	return nil
}
