package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

func hello(say func() error) error {
	return say()
}

func main() {
	db, err := bolt.Open("db.db", 0600, nil)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("mybucket"))
		if err != nil {
			return fmt.Errorf("Create bucket %s", err)
		}
		err = b.Put([]byte("answer"), []byte("Hello World"))
		err = b.Put([]byte("answer1"), []byte("Shakir Mengrani"))
		return err
	})

	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("%s key %s has been updated with value %s", "mybucket", "answer", "Hello World")
	}

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("mybucket"))
		c := b.Cursor()
		if b == nil {
			return errors.New("Bucket not found")
		}
		log.Printf("Value of %s is %s", "answer", b.Get([]byte("answer")))
		for k, v := c.First(); k != nil; k, v = c.Next() {
			log.Printf("Value of %s is %s", k, v)
		}
		return nil
	})

	err = hello(func() error {
		return errors.New("Callback func error")
	})

	if err != nil {
		log.Fatal(err)
	}

}
