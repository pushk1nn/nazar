package main

import (
	"os"

	"github.com/pushk1nn/nazar/logging"
)

func main() {
	// Get interface name from user input
	args := os.Args[1:]
	iface := args[0]

	// Start listening for new packets
	logging.Listen(iface)
}
