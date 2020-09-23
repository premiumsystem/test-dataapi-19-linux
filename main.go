package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/johansundell/fms-data-api/fmsdata"
)

func main() {
	// Set the login info to our database
	db := fmsdata.NewDataBase(settings.Host, settings.Filename, settings.User, settings.Pass)
	// Start a new session to the database, abort on fail
	if err := db.Login(); err != nil {
		log.Fatal(err)
	}
	// Wait until this functions ends before logout
	defer db.Logout()
	// Print the current time
	fmt.Println(time.Now())
	// Get a waitgroup so we can handle new routines start up and end to keep track of them
	wg := sync.WaitGroup{}
	// Make a channel array of a fixed size so we can keep track of how many we can run at the same time
	// When all channels in this aray are filled, the loop below will wait until one is set free ;)
	ws := make(chan struct{}, settings.NoOfConcurret)
	// Loop over and start requests
	for i := 0; i < settings.NoOfRequest; i++ {
		// Add one to our waitgroup
		wg.Add(1)
		// Init our dummy struct
		ws <- struct{}{}
		// Start a new routine running this function
		go func(n int, w chan struct{}, wg *sync.WaitGroup) {
			// Wait until this function are done before removing one from our waitgroup
			defer wg.Done()
			// Get the data from FileMaker and print the error on fail
			resp, err := db.GetAllFrom("sudde")
			if err != nil {
				fmt.Println(err)
			}
			// We don't really care about the data we got
			_ = resp
			// Free the channel so it can be used again
			<-w
		}(i, ws, &wg)

	}
	// Wait until all routines in our waitgroup are done
	wg.Wait()
	// Print the current time
	fmt.Println(time.Now())
}
