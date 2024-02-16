package main

import (
	"fmt"
	"net/http"
	"log"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w,"method is not supported",http.StatusNotFound)
		return
	}
	fmt.Fprintf(w,"hello!")
}

func dingHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err!= nil {
		fmt.Fprintf(w, "ParseForm() err: %v",err)
		return
	}
	fmt.Fprintf(w,"Post request successfull")
	name:=r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w,"name = %s \n",name)
	fmt.Fprintf(w,"address = %s \n",address)
}


func main(){
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileserver)
	http.HandleFunc("/hello",helloHandler)
	http.HandleFunc("/form",dingHandler)
	fmt.Printf("starting the server at port 8080")
	if err := http.ListenAndServe(":8080",nil); err != nil{
		log.Fatal(err)
	}
}
