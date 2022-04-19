package main

import (
	"assigment2/config"
	"assigment2/route"

	_ "github.com/lib/pq"
)

func main() {
	config.StartDB()

	route.StartRoute().Run(":8080")
}
