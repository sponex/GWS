// Tool     -   GWS
// Author   -   Simon Whitehouse
// Version  -   0.2
// Description  - GWS also known as GoWebServer, My own version of pythons simpleHTTP server so I can serve files.

package main

import (
	"log" //Logging package
	"net/http" //Networking package
    "flag" //Comand line flag package
	//"fmt"
)

func main() {

// Define Commands Line options
// port for ports
// dir for directorys
port := flag.String("p", "8000", "HTTP Port to use (Defaults to 8000)")

//dir := flag.string("d", ".", "Set Directory to Serve (Defaults to current)")
flag.Parse()

mux := http.NewServeMux()
fileServer := http.FileServer(http.Dir("."))
mux.Handle("/", http.StripPrefix("/", fileServer))


// Logging/ Outpt
log.Printf("Server started on Port %s", *port)
err := http.ListenAndServe(":" + *port, mux)
log.Fatal(err)
}