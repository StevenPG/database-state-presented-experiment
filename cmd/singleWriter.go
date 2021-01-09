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
	pkg.LoopRequest(1 * time.Second, pkg.MakeRequestWithRandomUpdate, "testUser1")
}
