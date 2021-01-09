package pkg

import (
	"fmt"
	"go.etcd.io/bbolt"
	"log"
	"time"
)

// Create the database file
func StartDatabase() {
	db := openDatabase()
	log.Print("Successfully created/opened experiment.database")

	db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("testBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	log.Print("Successfully created/opened testBucket Bucket in experiment.database")

	defer db.Close()
}

// TODO - get everything from the database, all k/v pairs
// GetBoltContent get everything in the database
func GetBoltContent() []string {
	db := openDatabase()
	defer db.Close()

	emptySlice := make([]string, 0)
	db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("testBucket"))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			emptySlice = append(emptySlice, string(v))
		}
		return nil
	})
	return emptySlice
}

func GetBoltContentAsState() []State {
	states := GetBoltContent()

	statesObjectSlice := make([]State, 0)
	for _, state := range states {
		statesObjectSlice = append(statesObjectSlice, State{Value: state})
	}

	return statesObjectSlice
}

// GetItem get a single item from the database using key
func GetItem() {
	// TODO - get item from database using key
}

// WriteItem write a single item to the database using key
func WriteItem(item *Item) {
	db := openDatabase()
	defer db.Close()
	_ = db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("testBucket"))
		//err := b.Put([]byte(item.Key), []byte(item.Value))
		err := b.Put([]byte(item.Key), []byte(item.Value))

		return err
	})
}

func openDatabase() *bbolt.DB {
	db, err := bbolt.Open("experiment.database", 0600, &bbolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
