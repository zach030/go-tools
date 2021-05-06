package ds

import (
	"log"
	"testing"
)

func TestBTree_Insert(t *testing.T) {
	bt := NewBTree()
	err := bt.Insert(1, []byte("world"))
	if err != nil {
		log.Fatal(err)
	}
	_ = bt.Insert(2, []byte("www"))
	_ = bt.Insert(3, []byte("one"))
	_ = bt.Insert(4, []byte("kaka"))
	_ = bt.Insert(5, []byte("no"))
	_ = bt.Insert(6, []byte("you"))
	_ = bt.Insert(7, []byte("you"))
	_ = bt.Insert(8, []byte("w"))
	_ = bt.Insert(9, []byte("r"))
	_ = bt.Insert(10, []byte("w"))
	_ = bt.Insert(11, []byte("you"))
}
