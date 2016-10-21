package main

import (
	"encoding/json"
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	"log"
	"time"
)

type Message struct {
	Name   string
	Body   string
	Hour   int64
	Second int64
}

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
	temp := make([]byte, 5)
	for counter := 0; counter < 10; counter++ {
		_, err1 := port.Read(temp)
		if err1 != nil {
			log.Fatalf("port.Read: %v", err1)
		}

		//Adding timestamp
		now := time.Now()
		fmt.Printf("%v:%v.%v   ", now.Hour(), now.Minute(), now.Second())
		//fmt.Println("Read", n, "bytes.") // From port.Read
		fmt.Printf("%s ºC\n", temp)

		//Making a message and encoding to JSON
		m_in := Message{"Temperature", string(temp), int64(now.Minute()), int64(now.Second())}

		m_encoded, err2 := json.Marshal(m_in)
		if err2 != nil {
			log.Fatalf("json.Marshal: %v", err2)
		}

		//Testing by decoding the same message
		var m_out Message

		err3 := json.Unmarshal(m_encoded, &m_out)
		if err3 != nil {
			log.Fatalf("json.Unmarshal: %v", err3)
		}

		//Printing both messages, making sure they are equal
		fmt.Println(m_in)
		fmt.Println(m_out)
	}
}
