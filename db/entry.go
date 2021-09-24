package db

import "encoding/binary"

const (
	entryHeader = 8
)

type Entry struct {
	key       []byte
	value     []byte
	keySize   uint32
	valueSize uint32
}

func NewEntry(key, value []byte) *Entry {
	return &Entry{
		key:       key,
		value:     value,
		keySize:   uint32(len(key)),
		valueSize: uint32(len(value)),
	}
}

func (e *Entry) Size() int64 {
	return int64(entryHeader + e.keySize + e.valueSize)
}

func (e *Entry) Key() string {
	return string(e.key)
}

// Encode head+body
// 4(key size) + 4 (body size)  + key + value
func (e *Entry) Encode() []byte {
	buf := make([]byte, e.Size())
	binary.BigEndian.PutUint32(buf[0:4], e.keySize)
	binary.BigEndian.PutUint32(buf[4:8], e.valueSize)
	copy(buf[entryHeader:entryHeader+e.keySize], e.key)
	copy(buf[entryHeader+e.keySize:], e.value)
	return buf
}

func Decode(buf []byte) *Entry {
	kSize := binary.BigEndian.Uint32(buf[0:4])
	vSize := binary.BigEndian.Uint32(buf[4:8])
	return &Entry{
		keySize:   kSize,
		valueSize: vSize,
	}
}
