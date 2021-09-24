package db

import (
	"os"
	"strings"
)

const (
	fileName = "zburger.db"
)

type DBFile struct {
	file   *os.File
	offset int64
}

func OpenDBFile(dir string) (*DBFile, error) {
	builder := strings.Builder{}
	builder.WriteString(dir)
	builder.WriteRune(os.PathSeparator)
	builder.WriteString(fileName)
	return openInternalFile(builder.String())
}

func openInternalFile(path string) (*DBFile, error) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	info, err := os.Stat(fileName)
	if err != nil {
		return nil, err
	}
	return &DBFile{file: file, offset: info.Size()}, nil
}

func (db *DBFile) Read(offset int64) (*Entry, error) {
	headBuf := make([]byte, entryHeader)
	_, err := db.file.ReadAt(headBuf, offset)
	if err != nil {
		return nil, err
	}
	entry := Decode(headBuf)
	offset += entryHeader
	if entry.keySize > 0 {
		keyBuf := make([]byte, entry.keySize)
		_, err = db.file.ReadAt(keyBuf, offset)
		if err != nil {
			return nil, err
		}
		entry.key = keyBuf
	}
	offset += int64(entry.keySize)
	if entry.valueSize > 0 {
		valBuf := make([]byte, entry.valueSize)
		_, err = db.file.ReadAt(valBuf, offset)
		if err != nil {
			return nil, err
		}
		entry.value = valBuf
	}
	return entry, nil
}

func (db *DBFile) Write(entry *Entry) error {
	buf := entry.Encode()
	_, err := db.file.WriteAt(buf, db.offset)
	if err != nil {
		return err
	}
	db.offset += entry.Size()
	return nil
}
