package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("the function of penny")

	var myCh = make(chan int,2)
	wg := &sync.WaitGroup{}

	//myCh <- 5

	wg.Add(2)

	go func(ch <-chan int,wg *sync.WaitGroup){
		for i:=0;i<2;i++{
		fmt.Println(<-ch)
		}
		fmt.Println("kaido dha",<-ch)

		defer wg.Done()
	}(myCh,wg)

	go func(ch chan<- int,wg *sync.WaitGroup){

		ch<-99
		ch<-54
		ch<-34
		fmt.Println(<-ch)
		defer wg.Done()
	}(myCh,wg)

	wg.Wait()

	//fmt.Println(<-myCh)

}