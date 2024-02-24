package main

import (
	"flag"
	"fmt"
	"reflect"
	"strconv"
)

func main(){
	//to run this program give a cli arg  like -name=somename -ding=somenumber
	name := flag.String("name","world","the name to greet")
	age := flag.String("ding","0","the age")
	flag.Parse()
	nage,_ := strconv.Atoi(*age)
	
	fmt.Println(*age,nage,"type of",reflect.TypeOf(nage).Kind())
	fmt.Printf("Hello %s!  \n",*name)
}