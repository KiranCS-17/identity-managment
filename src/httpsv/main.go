package main

import (
	"fmt"
	"golang.org/x/net/http2"
	"net"
	"net/http"
	"os"
        "strconv"
)
var n int64 = 0
var tJsonString  = `{"Originator":"Client","requestId":`

func checkErr(err error, msg string) {
	if err == nil {
		return
	}
	fmt.Printf("ERROR: %s: %s\n", msg, err)
	os.Exit(1)
}

func main() {
	H2CServerPrior()
}

// This server only supports "H2C prior knowledge".
// You can add standard HTTP/2 support by adding a TLS config.
func H2CServerPrior() {
	server := http2.Server{}

	l, err := net.Listen("tcp", "10.53.176.70:8080")
	checkErr(err, "while listening")

	fmt.Printf("Listening [10.53.176.70:8080]...\n")
	for {
                conn, err := l.Accept() 
		checkErr(err, "during accept")
		server.ServeConn(conn, &http2.ServeConnOpts{
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                                n++
                                s := strconv.FormatInt(n, 10)
                                myJsonString := tJsonString + `"` + s + `"}`
                                fmt.Println(myJsonString)
				fmt.Fprintf(w, myJsonString)
			}),
		})
	}
}
