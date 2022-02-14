package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"golang.org/x/net/http2"
	"net"
	"net/http"
	"os"
	"time"
        "io/ioutil"
)

func checkErr(err error, msg string) {
	if err == nil {
		return
	}
	fmt.Printf("ERROR: %s: %s\n", msg, err)
	os.Exit(1)
}

func main() {
	RoundTripExample()
}


func RoundTripExample() {
        url := os.Args[1]
	req, err := http.NewRequest("GET", url, nil)
	checkErr(err, "during new request")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	tr := &http2.Transport{
		AllowHTTP: true,
		DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
			return net.Dial(network, addr)
		},
	}
        for {
	   req.WithContext(ctx)
	   resp, err := tr.RoundTrip(req)
	   checkErr(err, "during roundtrip")
           data,_ := ioutil.ReadAll(resp.Body)
           resp.Body.Close() 
	   fmt.Printf("RoundTrip Proto: %v\n", resp)
           fmt.Printf("%s \n",data)
        }
}
