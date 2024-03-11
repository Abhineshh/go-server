package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("domain,hasMX,mxRecord,hasSPF,spfRecord,hasDMARC,dmarcRecord \n")
	for reader.Scan() {
		checkDomain(reader.Text())
	}
	if err := reader.Err(); err!=nil{
		log.Fatal("Error: could not read from input %v \n",err)
	}
}

func checkDomain(domain string) {

	var hasMX,hasSPF,hasDMARC bool
	var spfRecord,dmarcRecord string

	mxRecord,err:=net.LookupMX(domain)
	if err != nil {
		log.Printf("Error  %v \n",err)
	}
	if len(mxRecord)>0{
		hasMX = true
	}

	txtRecord,err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error %v \n",err)
	}

	for _,record := range txtRecord{
		if strings.HasPrefix(record,"v=spf1"){
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords,err := net.LookupTXT("_dmarc."+domain)
	if err != nil {
		log.Printf("Error %v \n",err)
	}
	for _,record := range dmarcRecords{
		if strings.HasPrefix(record,"v=DMARC1"){
			hasDMARC = true
			dmarcRecord = record
			break 
		}
	}

	fmt.Printf("%v \n %v \n %v \n %v \n\n %v \n\n %v",domain,hasMX,hasSPF,hasDMARC,spfRecord,dmarcRecord);
}
