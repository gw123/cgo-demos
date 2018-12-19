package main

import (
	"fmt"
	"github.com/jacobsa/go-serial/serial"
)

func main() {
	// Set up options.
	options := serial.OpenOptions{
		PortName: "COM7",
		BaudRate: 19200,
		DataBits: 8,
		StopBits: 1,
		MinimumReadSize: 4,
	}

	// Open the port.
	port, err := serial.Open(options)
	if err != nil {
		fmt.Println("serial.Open: %v", err)
		return
	}

	// Make sure to close it later.
	defer port.Close()

	// Write 4 bytes to the port.
	b := []byte("hello world \n\n\n")
	n, err := port.Write(b)
	if err != nil {
		fmt.Println("port.Write: %v", err)
	}
	fmt.Println("Wrote", n, "bytes.")
}
