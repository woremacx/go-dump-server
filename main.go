package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

var portNumber = flag.String("port", "80", "port number.")

func main() {
	flag.Parse()

	fmt.Printf("go-dump-server port:%s\n", *portNumber)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "%s", dump)
		fmt.Printf("%s", dump)
	})

	if err := http.ListenAndServe(":"+*portNumber, handler); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
