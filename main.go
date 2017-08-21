package main

import (
	"flag"
	"fmt"
	"mychat/lib"
	"os"
)

func main() {
	var isHost bool
	flag.BoolVar(&isHost, "listen", false, "Listens on the specified IP address")
	flag.Parse()

	if isHost {
		// go run main.go -listen <ip>
		connIP := os.Args[2]
		lib.RunHost(connIP)
		fmt.Println("Is Host")
	} else {
		// go run main.go <ip>
		connIP := os.Args[1]
		lib.RunGuest(connIP)
		fmt.Println("Is Guest")
	}

}
