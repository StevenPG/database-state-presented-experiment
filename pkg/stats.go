package pkg

import (
	"encoding/json"
	"os"
	"time"
)

// ReportStats reports stats on the database every 10s
func ReportStats() {

	// Grab the initial stats.
	prev := DB.Stats()

	for {
		// Wait for 10s.
		time.Sleep(10 * time.Second)

		stats := DB.Stats()
		diff := stats.Sub(&prev)

		// Encode stats to JSON and print to STDERR.
		json.NewEncoder(os.Stderr).Encode(diff)

		// Save stats for the next loop.
		prev = stats
	}
}
