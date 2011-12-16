package main

import (
	"net/http"
	"flag"
	"fmt"
)

var (
	addrFlag = flag.String("a", ":8080", "Address to bind to")
	helpFlag = flag.Bool("h", false, "Show this help")
	fs http.Handler
)

func main() {
	flag.Parse()
	if *helpFlag || flag.NArg() > 1 {
		fmt.Printf("Usage: goserve [options] [dir]\n")
		flag.PrintDefaults()
		return
	}

	if flag.NArg() == 1 {
		fs = http.FileServer(http.Dir(flag.Arg(0)))
	} else {
		fs = http.FileServer(http.Dir("."))
	}
	http.HandleFunc("/", handler)
	e := http.ListenAndServe(*addrFlag, nil)
	if e != nil {
		panic(e)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	fs.ServeHTTP(w, r)
}
