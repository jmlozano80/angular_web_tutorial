//special package that can create executables
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {

	//command lines flags
	port := flag.Int("port", 8080, "port to serve on")
	dir := flag.String("directory", "web/", "directory of web files")
	flag.Parse()

	//handle all requests by serving a file of the same name
	fs := http.Dir(*dir)
	handler := http.FileServer(fs)
	http.Handle("/", handler)

	log.Printf("Runnung on port %d\n", *port)

	addr := fmt.Sprintf("127.0.0.1:%d", *port)
	// this call blocks -- the program to run here forever
	err := http.ListenAndServe(addr, nil)
	fmt.Println(err.Error())

}
