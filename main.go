package main

import (
	"fmt"
	"log"
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
	for i := 0; i < 2000; i++ {
		resp, err := db.GetAllFrom("sudde")
		if err != nil {
			log.Fatal(err)
		}
		_ = resp
		//fmt.Println(string(resp))
	}
	fmt.Println(time.Now())
}
