package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, mxErr := net.LookupMX(domain)
	if mxErr != nil {
		log.Printf("Error in lookupMX: %v\n", mxErr.Error())
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords, txtErr := net.LookupTXT(domain)
	if txtErr != nil {
		log.Printf("Error in lookupTXT: %v\n", txtErr.Error())
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, dmarcRecordErr := net.LookupTXT("_dmarc." + domain)
	if dmarcRecordErr != nil {
		log.Printf("Error in lookupTXT _dmarc: %v\n", dmarcRecordErr.Error())
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v,", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord\n")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Printf("couldnot read from scanner; Error: %v\n", err.Error())
	}
}
