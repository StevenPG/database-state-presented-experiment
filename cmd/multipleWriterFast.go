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

	// Each routine will get a unique id, staggered starts
	executionPause := 50 * time.Millisecond
	go pkg.LoopRequest(executionPause, pkg.MakeRequestWithRandomUpdate, "testUser1")
	time.Sleep(25 * time.Millisecond)
	go pkg.LoopRequest(executionPause, pkg.MakeRequestWithRandomUpdate, "testUser2")
	time.Sleep(36 * time.Millisecond)
	go pkg.LoopRequest(executionPause, pkg.MakeRequestWithRandomUpdate, "testUser3")
	time.Sleep(41 * time.Millisecond)
	go pkg.LoopRequest(executionPause, pkg.MakeRequestWithRandomUpdate, "testUser4")
	time.Sleep(47 * time.Millisecond)
	go pkg.LoopRequest(executionPause, pkg.MakeRequestWithRandomUpdate, "testUser5")
	time.Sleep(61 * time.Millisecond)
	go pkg.LoopRequest(executionPause, pkg.MakeRequestWithRandomUpdate, "testUser6")
	time.Sleep(100 * time.Millisecond)
	go pkg.LoopRequest(executionPause, pkg.MakeRequestWithRandomUpdate, "testUser7")
	time.Sleep(101 * time.Millisecond)
	go pkg.LoopRequest(executionPause, pkg.MakeRequestWithRandomUpdate, "testUser8")
	time.Sleep(102 * time.Millisecond)
	go pkg.LoopRequest(executionPause, pkg.MakeRequestWithRandomUpdate, "testUser9")
	time.Sleep(103 * time.Millisecond)
	go pkg.LoopRequest(executionPause, pkg.MakeRequestWithRandomUpdate, "testUser10")
	// Each routine will get a unique id, staggered starts
	go pkg.LoopRequest(executionPause, pkg.MakeRequestWithRandomUpdate, "testUser11")
	go pkg.LoopRequest(executionPause, pkg.MakeRequestWithRandomUpdate, "testUser12")
	go pkg.LoopRequest(executionPause, pkg.MakeRequestWithRandomUpdate, "testUser13")
	go pkg.LoopRequest(executionPause, pkg.MakeRequestWithRandomUpdate, "testUser14")
	go pkg.LoopRequest(executionPause, pkg.MakeRequestWithRandomUpdate, "testUser15")
	go pkg.LoopRequest(executionPause, pkg.MakeRequestWithRandomUpdate, "testUser16")
	go pkg.LoopRequest(executionPause, pkg.MakeRequestWithRandomUpdate, "testUser17")
	go pkg.LoopRequest(executionPause, pkg.MakeRequestWithRandomUpdate, "testUser18")
	go pkg.LoopRequest(executionPause, pkg.MakeRequestWithRandomUpdate, "testUser19")
	go pkg.LoopRequest(executionPause, pkg.MakeRequestWithRandomUpdate, "testUser20")
	go pkg.LoopRequest(executionPause, pkg.MakeRequestWithRandomUpdate, "testUser21")
	go pkg.LoopRequest(executionPause, pkg.MakeRequestWithRandomUpdate, "testUser22")
	go pkg.LoopRequest(executionPause, pkg.MakeRequestWithRandomUpdate, "testUser23")
	go pkg.LoopRequest(executionPause, pkg.MakeRequestWithRandomUpdate, "testUser24")
	go pkg.LoopRequest(executionPause, pkg.MakeRequestWithRandomUpdate, "testUser25")
	go pkg.LoopRequest(executionPause, pkg.MakeRequestWithRandomUpdate, "testUser26")
	go pkg.LoopRequest(executionPause, pkg.MakeRequestWithRandomUpdate, "testUser27")
	go pkg.LoopRequest(executionPause, pkg.MakeRequestWithRandomUpdate, "testUser28")
	go pkg.LoopRequest(executionPause, pkg.MakeRequestWithRandomUpdate, "testUser29")
	go pkg.LoopRequest(executionPause, pkg.MakeRequestWithRandomUpdate, "testUser30")
	go pkg.LoopRequest(executionPause, pkg.MakeRequestWithRandomUpdate, "testUser31")
	go pkg.LoopRequest(executionPause, pkg.MakeRequestWithRandomUpdate, "testUser32")
	go pkg.LoopRequest(executionPause, pkg.MakeRequestWithRandomUpdate, "testUser33")
	go pkg.LoopRequest(executionPause, pkg.MakeRequestWithRandomUpdate, "testUser34")
	go pkg.LoopRequest(executionPause, pkg.MakeRequestWithRandomUpdate, "testUser35")
	go pkg.LoopRequest(executionPause, pkg.MakeRequestWithRandomUpdate, "testUser36")
	go pkg.LoopRequest(executionPause, pkg.MakeRequestWithRandomUpdate, "testUser37")
	go pkg.LoopRequest(executionPause, pkg.MakeRequestWithRandomUpdate, "testUser38")
	go pkg.LoopRequest(executionPause, pkg.MakeRequestWithRandomUpdate, "testUser39")
	go pkg.LoopRequest(executionPause, pkg.MakeRequestWithRandomUpdate, "testUser40")

	// Run for 60 seconds
	time.Sleep(60 * time.Second)
}
