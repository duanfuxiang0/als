package als

import (
	"os"
)

type AFile struct {
	fp *os.File
}

func NewAFile(filePath string) (*AFile, error) {
	fp, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return &AFile{fp: fp}, nil
}

func (db *AFile) Read(offset uint32, length uint32) ([]byte, error) {
	/*
	**name   ks  :   Key    : vs :  value
	**size:  1   :  (<=2^8) : 1  :  (<=2^8)
	 */

	val := make([]byte, length)
	if _, err := db.fp.Seek(int64(offset), 0); err != nil {
		return nil, err
	}
	if ret, err := db.fp.Read(val); ret < 0 {
		return nil, err
	} else if ret == 0 {
		return nil, nil
	}
	return val, nil
}

func (db *AFile) Write(key []byte, value []byte) {
}
