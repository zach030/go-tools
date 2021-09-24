package db

import (
	"io"
	"os"
	"sync"
)

type BurgerDB struct {
	path   string
	lock   sync.Mutex
	idx    map[string]int64
	dbFile *DBFile
}

func OpenDB(dir string) (*BurgerDB, error) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return nil, err
		}
	}

	dbFile, err := OpenDBFile(dir)
	if err != nil {
		return nil, err
	}
	db := &BurgerDB{
		path:   dir,
		idx:    make(map[string]int64),
		dbFile: dbFile,
	}
	db.initIndex()
	return db, nil
}

func (b *BurgerDB) initIndex() {
	var offset int64
	for {
		entry, err := b.dbFile.Read(offset)
		if err != nil {
			if err == io.EOF {
				break
			}
			return
		}
		b.idx[entry.Key()] = offset
		if entry.tag == DEL {
			delete(b.idx, entry.Key())
		}
		offset += entry.Size()
	}
	return
}

func (b *BurgerDB) Get(key []byte) (val []byte, err error) {
	if len(key) == 0 {
		return
	}
	off, ok := b.idx[string(key)]
	if !ok {
		return
	}
	entry, err := b.dbFile.Read(off)
	if err != nil {
		return
	}
	if entry != nil {
		val = entry.value
	}
	return
}

func (b *BurgerDB) Set(key, value []byte) error {
	entry := NewEntry(key, value, SET)
	err := b.dbFile.Write(entry)
	if err != nil {
		return err
	}
	b.lock.Lock()
	defer b.lock.Unlock()
	b.idx[string(key)] = b.dbFile.offset
	return nil
}

func (b *BurgerDB) Delete(key []byte) error {
	if _, ok := b.idx[string(key)]; !ok {
		return nil
	}
	entry := NewEntry(key, nil, DEL)
	err := b.dbFile.Write(entry)
	if err != nil {
		return err
	}
	delete(b.idx, string(key))
	return nil
}
