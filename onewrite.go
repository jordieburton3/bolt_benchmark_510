package main

import (
	"fmt"
	"log"
	"flag"
	"os"
	"../bolt"
	//"github.com/jordieburton3/bolt"
	"time"
)

var (
	myBucket = []byte("perftest")
	dbLocation string
	totalCnt = 0
)


func init() {
	flag.StringVar(&dbLocation, "db", "/tmp/test.db", "location of your boltdb file")
}

func handleErr(err error) {
	if err != nil {
		log.Fatalf("Unable to proceed [%s]", err)
	}
}

func getKey(id int) string {
	user := "6a204bd89f3c8348afd5c77c717a097a"
	typeOf := "details"
	value := "2413fb3709b05939f04cf2e92f7d0897fc2596f9ad0b8a9ea855c7bfebaae892"
	return fmt.Sprintf("%s:%s:%s:%d", user, typeOf, value, id) // makes a key of hefty length
}

func main() {
	flag.Parse()

	log.Printf("Starting with dbpath [%s]", dbLocation)
	startInsert := time.Now()

	// insert 1 record
	insert(0)
	
	elapsedInsert := time.Since(startInsert)


	startRead := time.Now()
	read()
	elapsedRead := time.Since(startRead)

	log.Printf("TOTAL INSERT took %s for %d items", elapsedInsert, totalCnt)
	log.Printf("TOTAL READ took %s", elapsedRead)
}

func read() {

	db, err := bolt.Open(dbLocation, 0644, nil)
	handleErr(err)
	defer db.Close()

	db.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(myBucket))
		v := b.Get([]byte("0"))
		fmt.Printf("The value is %d\n", v)
		return nil
	})
}

func insert(offset int) {
	start := time.Now()
	db, err := bolt.Open(dbLocation, 0644, nil)
	handleErr(err)
	defer db.Close()

	// store some data
	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(myBucket)
		handleErr(err)

		err = bucket.Put([]byte("0"), []byte("8"))
		handleErr(err)
		return nil
	})
	handleErr(err)

	now := time.Now()
	elapsed := time.Since(start)

	f, err := os.OpenFile("testwrite.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
    	log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Printf("Inserted [%d] now: [%s] items took %s\n", totalCnt, now, elapsed)


}

