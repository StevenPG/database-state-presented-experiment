package pkg

import (
	"fmt"
	"go.etcd.io/bbolt"
	"log"
	"time"
)

// Create an exported global variable to hold the database connection pool.
var DB *bbolt.DB

// Create the database file
func StartDatabase() {
	db, err := bbolt.Open("experiment.database", 0600, &bbolt.Options{Timeout: 1 * time.Second})
	DB = db
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Successfully created/opened experiment.database")

	DB.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("testBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	log.Print("Successfully created/opened testBucket Bucket in experiment.database")
}

// TODO - get everything from the database, all k/v pairs
// GetBoltContent get everything in the database
func GetBoltContent() []string {
	emptySlice := make([]string, 0)
	DB.View(func(tx *bbolt.Tx) error {
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

// WriteItem write a single item to the database using key
func WriteItem(item *Item) {
	_ = DB.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("testBucket"))
		err := b.Put([]byte(item.Key), []byte(item.Value))
		return err
	})
}
