package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptrace"
	"os"
	"time"
)

var site string

func main() {
	if len(os.Args) < 2 {
		fmt.Println("please supply destination. eg. http://google.com")
		os.Exit(1)
	}
	site = os.Args[1]

	req, _ := http.NewRequest("GET", site, nil)
	trace := &httptrace.ClientTrace{
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("%v Got Conn: %+v\n", time.Now(), connInfo)
		},
		DNSStart: func(dnsInfo httptrace.DNSStartInfo) {
			fmt.Printf("%v DNS Start: %+v\n", time.Now(), dnsInfo)
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			fmt.Printf("%v DNS Done: %+v\n", time.Now(), dnsInfo)
		},
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

	client := http.DefaultClient
	// _, err := http.DefaultTransport.RoundTrip(req)
	_, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

}
