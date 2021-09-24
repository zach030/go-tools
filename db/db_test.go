package db

import (
	"fmt"
	"testing"
)

func TestBurgerDB_Set(t *testing.T) {
	db, err := OpenDB("D:\\AllProjects\\Go\\zburger\\db")
	if err != nil {
		return
	}
	err = db.Set([]byte("name"), []byte("zach"))
	if err != nil {
		fmt.Println(err)
		return
	}
	err = db.Set([]byte("age"), []byte("21"))
	if err != nil {
		fmt.Println(err)
		return
	}

	val, err := db.Get([]byte("name"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("key is: name, val:", string(val))

	val, err = db.Get([]byte("age"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("key is: age, val:", string(val))
}
