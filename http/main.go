package main

import (
	"fmt"
	"io"
	"net/http"
)

type logWriter struct {}

func main() {
	response, err := http.Get("https://google.com")

	if (err != nil) {
		fmt.Errorf("Error: %s", err)
	}


	lw := logWriter{}

	io.Copy(lw, response.Body)
	// will copy response body to os.Stdout (writer)
	//io.Copy(os.Stdout, response.Body)
	// bs := make([]byte, 99999) // takes a type, and empty spaces
	// response.Body.Read(bs) // read the data in the byte slice
	// fmt.Println(string(bs))
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))
	return len(bs), nil
}