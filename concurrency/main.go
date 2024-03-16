package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup // pointer is preferabel
var mut sync.Mutex // pointer is preferable

var signals = []string{}

func main() {
	wg.Add(2)
	go greeter("abhinesh")
	greeter("Scientist")
	wg.Wait()
	//time.Sleep(3 * time.Millisecond) 

	websitelist := []string{
		"http://lco.dev",
		"http://google.com",
		"http://github.com",
	}

	

	for _,web := range websitelist{
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}

func greeter(s string) {
	for i:= 0 ;i<6 ;i++{
		time.Sleep(3 * time.Millisecond) 
		fmt.Println(s)
	}
	defer wg.Done()
}

func getStatusCode(endpoint string){
	defer wg.Done()
	resp,err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	mut.Lock()
	signals  = append(signals,endpoint)
	mut.Unlock()
	fmt.Printf(" %d Status Code for %s \n",resp.StatusCode,endpoint)
}