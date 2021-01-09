package pkg

import (
	"encoding/json"
	"go.etcd.io/bbolt"
	"log"
	"os"
	"time"
)

// ReportStats reports stats on the database every 10s
func ReportStats() {

	// Grab the initial stats.
	db, err := bbolt.Open("experiment.database", 0600, &bbolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	prev := db.Stats()
	db.Close()

	for {
		// Wait for 10s.
		time.Sleep(10 * time.Second)

		// Grab the current stats and diff them.
		db, err := bbolt.Open("experiment.database", 0600, &bbolt.Options{Timeout: 1 * time.Second})
		if err != nil {
			log.Fatal(err)
		}

		stats := db.Stats()
		diff := stats.Sub(&prev)

		// Encode stats to JSON and print to STDERR.
		json.NewEncoder(os.Stderr).Encode(diff)

		// Save stats for the next loop.
		prev = stats
		db.Close()
	}
}
