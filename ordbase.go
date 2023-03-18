package main

import (
	"fmt"
	// "log"
	"os"
	"net/http"


	"github.com/pixelartexchange/artbase.server/artbase"
	"github.com/pixelartexchange/artbase.server/serve"
)




var Collections = []artbase.Collection{
  {Name: "ordinalpunks",     Width: 24, Height: 24,
	 Path: "./ordinalpunks.png",
	 Url: "https://github.com/pixelartexchange/ordinals.sandbox/raw/master/i/ordinalpunks.png",
	 Count: 100 },

	{Name: "ordinalphunks",     Width: 24, Height: 24,
	 Path: "./ordinalphunks.png",
	 Url: "https://github.com/pixelartexchange/ordinals.sandbox/raw/master/i/ordinalphunks.png",
	 Count: 101 },

	{Name: "lilordinalphunks",     Width: 24, Height: 24,
	 Path: "./lilordinalphunks.png",
	 Url: "https://github.com/pixelartexchange/ordinals.sandbox/raw/master/i/lilordinalphunks.png",
	 Count: 101 },

	{Name: "ordoggies",     Width: 32, Height: 32,
	 Path: "./ordoggies.png",
	 Url: "https://github.com/pixelartexchange/ordinals.sandbox/raw/master/i/ordoggies.png",
	 Count: 101 },


	 {Name: "bitcoinpunks",     Width: 24, Height: 24,
	 Path: "./bitcoinpunks.png",
	 Url: "https://github.com/pixelartexchange/ordinals.sandbox/raw/master/i/bitcoinpunks.png",
	 Count: 100,
	 Background: true },

	 {Name: "gpunks",     Width: 24, Height: 24,
	 Path: "./gpunks.png",
	 Url: "https://github.com/pixelartexchange/ordinals.sandbox/raw/master/i/gpunks.png",
	 Count: 88,
	 Background: true },

	 {Name: "ordinalheropunks",     Width: 24, Height: 24,
	 Path: "./ordinalheropunks.png",
	 Url: "https://github.com/pixelartexchange/ordinals.sandbox/raw/master/i/ordinalheropunks.png",
	 Count: 50,
	 Background: true },

	 {Name: "litecoinpunks",     Width: 24, Height: 24,
	 Path: "./litecoinpunks.png",
	 Url: "https://github.com/pixelartexchange/ordinals.sandbox/raw/master/i/litecoinpunks.png",
	 Count: 100,
	 Background: true },

	 {Name: "ordinaryapes",     Width: 24, Height: 24,
	 Path: "./ordinaryapes.png",
	 Url: "https://github.com/pixelartexchange/ordinals.sandbox/raw/master/i/ordinaryapes.png",
	 Count: 100,
	 Background: true },

	 {Name: "pixelpepes",     Width: 24, Height: 24,
	 Path: "./pixelpepes.png",
	 Url: "https://github.com/pixelartexchange/ordinals.sandbox/raw/master/i/pixelpepes.png",
	 Count: 111,
	 Background: true },

	 {Name: "dogepunks",     Width: 28, Height: 28,
	 Path: "./dogepunks.png",
	 Url: "https://github.com/pixelartexchange/ordinals.sandbox/raw/master/i/dogepunks.png",
	 Count: 69,
	 Background: true },

	 {Name: "extraordinalwomen",     Width: 32, Height: 32,
	 Path: "./extraordinalwomen.png",
	 Url: "https://github.com/pixelartexchange/ordinals.sandbox/raw/master/i/extraordinalwomen.png",
	 Count: 111 },

	 {Name: "ordinalkitties",     Width: 45, Height: 45,
	 Path: "./ordinalkitties.png",
	 Url: "https://github.com/pixelartexchange/ordinals.sandbox/raw/master/i/ordinalkitties.png",
	 Count: 100,
	 Background: true },

}





func main() {

	serve := serve.NewRouter( Collections )

	// default addr to localhost:8080 for now
	//    for windows include localhost to avoid firewall warning/popup
	//       if binding to :8080 only  - why? why not?

  addr := "localhost:8080"

	// check for port in env settings - required by heroku
	port := os.Getenv( "PORT" )
	if port != "" {
		addr = ":" + port
	}


 	http.ListenAndServe( addr, serve )

	fmt.Println( "Bye!")
}

