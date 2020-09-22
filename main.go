package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/johansundell/mark-gantt/fmsdata"
)

var db fmsdata.DataBase

func main() {
	db = fmsdata.NewDataBase("https://fms1.sudde.eu", "sudde", "admin", "Onsdag5")
	if err := db.Login(); err != nil {
		log.Fatal(err)
	}
	defer db.Logout()
	fmt.Println(time.Now())
	wg := sync.WaitGroup{}
	ws := make(chan struct{}, 100)
	for i := 0; i < 25000; i++ {
		/*resp, err := db.GetAllFrom("sudde")
		if err != nil {
			log.Fatal(err)
		}
		_ = resp*/
		//fmt.Println(string(resp))

		// Get them all at once
		wg.Add(1)
		ws <- struct{}{}
		go func(n int, w chan struct{}, wg *sync.WaitGroup) {
			defer wg.Done()
			resp, err := db.GetAllFrom("sudde")
			if err != nil {
				//log.Fatal(err)
				fmt.Println(err)
			}
			//fmt.Println(n)
			_ = resp
			<-w
		}(i, ws, &wg)

	}
	wg.Wait()
	fmt.Println(time.Now())
}
