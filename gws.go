// Tool     -   GWS
// Author   -   Simon Whitehouse
// Version  -   0.3
// Description  - GWS also known as GoWebServer, My own version of Python's simpleHTTP server so I can serve files.

package main

import (
	"log" //Logging package
	"net/http" //Networking package
    "flag" //Comand line flag package
	"fmt" // debugging
)

func main() {

// Define Commands Line options
// port for ports, dir for directorys, SSL to force use of SSL. 
port := flag.String("p", "8000", "HTTP Port to use (Defaults to 8000)")
dir := flag.String("d", ".", "Set Directory to Serve (Defaults to current)")
ssl := flag.Bool("s", false, "Forces use of SSL (Defaults to false)")

flag.Parse()

mux := http.NewServeMux()
fileServer := http.FileServer(http.Dir(*dir))
mux.Handle("/", http.StripPrefix("/", fileServer))


if *ssl == false {
	log.Printf("Server started on Port %s", *port)
	err := http.ListenAndServe(":" + *port, mux)
	log.Fatal(err)
} else {
	log.Printf("Server started on SSL Port %s", *port)
	err := http.ListenAndServeTLS(":" + *port, "./tls/cert.pem", "./tls/key.pem", mux)
	log.Fatal(err)
}
}