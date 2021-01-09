package main

import (
	"fmt"
	"github.com/StevenPG/database-state-presented-experiment/pkg"
	"time"
)

func main() {
	fmt.Println("Starting Experiment")
	fmt.Println("Description: We're going to write 10 items to the service and continually update the value.\n" +
		"For each update, we will require the entire state of the database, " +
		"which we will retrieve for each request as quickly as possible.")

	// Each routine will get a unique id
	go pkg.LoopRequest(1 * time.Second, pkg.MakeRequestWithRandomUpdate, "testUser1")
	go pkg.LoopRequest(1 * time.Second, pkg.MakeRequestWithRandomUpdate, "testUser2")
	go pkg.LoopRequest(1 * time.Second, pkg.MakeRequestWithRandomUpdate, "testUser3")
	go pkg.LoopRequest(1 * time.Second, pkg.MakeRequestWithRandomUpdate, "testUser4")
	go pkg.LoopRequest(1 * time.Second, pkg.MakeRequestWithRandomUpdate, "testUser5")

	// Run for 60 seconds
	time.Sleep(60 * time.Second)
}
