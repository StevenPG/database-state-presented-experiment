package main

import (
	"github.com/StevenPG/database-state-presented-experiment/pkg"
)

func main() {
	pkg.StartDatabase()
	//go pkg.ReportStats()
	pkg.RunHttpServer()
}
