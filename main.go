package main

import (
	"fmt"
	"sync"

	"github.com/johansundell/fms-data-api/fmsdata"
)

var wg, wg2 sync.WaitGroup

type Job struct {
	Id     int
	Work   string
	Result []byte
	err    error
}

func produce(jobs chan<- *Job) {
	// Generate jobs:
	for i := 0; i < settings.NoOfRequest; i++ {
		jobs <- &Job{Id: i, Work: ""}
	}
	close(jobs)
}

func consume(id int, jobs <-chan *Job, results chan<- *Job) {
	defer wg.Done()
	// Do work here
	for job := range jobs {
		resp, err := makeCallToFms("sudde")
		job.Result, job.err = resp, err
		results <- job
	}
}

func makeCallToFms(layout string) ([]byte, error) {
	// Login to filemaker
	db := fmsdata.NewDataBase(settings.Host, settings.Filename, settings.User, settings.Pass)
	if err := db.Login(); err != nil {
		return nil, err
	}
	// Get the data from a layout
	resp, err := db.GetAllFrom(layout)
	// Logout from filemaker
	db.Logout()
	return resp, err

}

func analyze(results <-chan *Job) {
	defer wg2.Done()
	for job := range results {
		if job.err != nil {
			fmt.Println("error", job.Id, job.err)
		} else {
			if settings.ShowDone {
				fmt.Println("job", job.Id, "done")
			}
		}
	}
}

func main() {
	jobs := make(chan *Job, 3000)    // Buffered channel
	results := make(chan *Job, 3000) // Buffered channel

	// Start consumers:
	for i := 0; i < settings.NoOfConcurret; i++ { // Prepare our workers
		wg.Add(1)
		go consume(i, jobs, results)
	}
	// Start producing
	go produce(jobs)

	// Start analyzing:
	wg2.Add(1)
	go analyze(results)

	wg.Wait() // Wait all consumers to finish processing jobs

	// All jobs are processed, no more values will be sent on results:
	close(results)

	wg2.Wait() // Wait analyzer to analyze all results
}
