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

//var db fmsdata.DataBase

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
		/*db := fmsdata.NewDataBase(settings.Host, settings.Filename, settings.User, settings.Pass)
		if err := db.Login(); err != nil {
			job.err = errors.New("login err")
			results <- job
		} else {
			resp, err := db.GetAllFrom("sudde")
			if err != nil {
				job.err = errors.New("work error")
			} else {
				job.Result = resp
			}
			db.Logout()
			results <- job

		}*/
		resp, err := makeCallToFms("sudde")
		job.Result, job.err = resp, err
		results <- job
	}
}

func makeCallToFms(layout string) ([]byte, error) {
	db := fmsdata.NewDataBase(settings.Host, settings.Filename, settings.User, settings.Pass)
	/*if err := db.Login(); err != nil {
		return nil, err
	}*/
	resp, err := db.GetAllFrom(layout)
	db.Logout()
	return resp, err

}

func analyze(results <-chan *Job) {
	defer wg2.Done()
	for job := range results {
		//fmt.Printf("result:%d %s\n", job.Id, string(job.Result))
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
	//db = fmsdata.NewDataBase(settings.Host, settings.Filename, settings.User, settings.Pass)
	//defer db.Logout()
	jobs := make(chan *Job, 3000)    // Buffered channel
	results := make(chan *Job, 3000) // Buffered channel

	// Start consumers:
	for i := 0; i < settings.NoOfConcurret; i++ { // 5 consumers
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
