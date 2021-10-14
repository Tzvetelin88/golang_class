package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Refer to www.golang.org/pkg/

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	// Example 1 for writing it
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	lw := logWriter{}
	// Example 2 for writing it
	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Bytes in current LogWriter", len(bs))

	return len(bs), nil
}
