package main

import (
	"github.com/samgozman/go-finra-short-sales-analyzer/internal/mongodb"
)

func main() {
	// Get Client, Context, CalcelFunc and
	// err from connect method.
	client, ctx, cancel, err := mongodb.Connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	// Release resource when the main
	// function is returned.
	defer mongodb.Close(client, ctx, cancel)

	// Ping mongoDB with Ping method
	mongodb.Ping(client, ctx)
}
