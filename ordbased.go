package main

import (
	"fmt"
	// "log"
	// "os"
	"net/http"


	"github.com/pixelartexchange/artbase.server/artbase"
	"github.com/pixelartexchange/artbase.server/serve"
)





func main() {

  url := "https://github.com/pixelartexchange/ordbase.server/raw/master/collections.csv"

	collections := artbase.DownloadCollections( url )

	serve := serve.NewRouter( collections )

	// default addr to localhost:8080 for now
	//    for windows include localhost to avoid firewall warning/popup
	//       if binding to :8080 only  - why? why not?

  addr := "localhost:8080"

 	http.ListenAndServe( addr, serve )

	fmt.Println( "Bye!")
}

