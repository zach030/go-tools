package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func main(){
	db,err := bolt.Open("my.db",0600,nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	tx,err := db.Begin(true)
	if err != nil {
		log.Fatal(err)
	}
	bucket := tx.Bucket([]byte("bucket"))
	//bucket,err := tx.CreateBucket([]byte("bucket"))
	//if err != nil {
	//	log.Fatal(err)
	//}
	err = bucket.Put([]byte("foo"),[]byte("bar"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bucket.Get([]byte("foo"))))
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

}

func sliceRise(s []int)  {
	s = append(s,0)
	for i, _:= range s {
		s[i]++
	}
}
