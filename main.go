package main

import (
	"fmt"
	"os"

	sq "github.com/rumblefrog/go-a2s"
)

func main() {
	address := "127.0.0.1:27015"

	if len(os.Args) >= 2 {
		address = os.Args[1]
	}

	client, err := sq.NewClient(address)
	if err != nil {
		fmt.Printf("configure: %v\n", err)
		os.Exit(1)
	}
	defer client.Close()

	_, err = client.QueryInfo()
	if err != nil {
		fmt.Printf("query: %v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
