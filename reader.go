package main

import (
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	"log"
)

func main() {

	options := serial.OpenOptions{
		PortName:              "/dev/cu.usbmodem1411",
		BaudRate:              9600,
		DataBits:              8,
		StopBits:              1,
		InterCharacterTimeout: 200,
		MinimumReadSize:       6,
	}

	// Open the port
	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	// Make sure to close the port later
	defer port.Close()

	// Read from port
	b := make([]byte, 5)
	for counter := 0; counter < 10; counter++ {
		n, err := port.Read(b)
		if err != nil {
			log.Fatalf("port.Read: %v", err)
		}

		fmt.Println("Read", n, "bytes.")
		//fmt.Println(b)
		//fmt.Println("Printf with % x:")
		//fmt.Printf("% x\n", b)
		fmt.Println("Printf with %q:")
		fmt.Printf("%q\n", b)
	}
}
