package main

import (
	"encoding/json"
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	"log"
	"strconv"
	"strings"
	"time"
)

type Message struct {
	Temp   float64
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
	temp_raw := make([]byte, 5)
	for counter := 0; counter < 10; counter++ {
		n, err1 := port.Read(temp_raw)
		if err1 != nil {
			log.Fatalf("port.Read: %v", err1)
		}

		//Converting temperature-reading from array to string to float
		temp_string := strings.TrimSpace(string(temp_raw[:n]))
		temp, _ := strconv.ParseFloat(temp_string, 64)

		//Adding timestamp
		now := time.Now()

		//Preparing the message and encoding it to JSON
		m_in := Message{temp, int64(now.Minute()), int64(now.Second())}

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

		//Test output
		fmt.Printf("%v:%v.%v   ", now.Hour(), now.Minute(), now.Second())
		fmt.Printf("%v ºC\n", temp)
		fmt.Printf("%v\n%v\n", m_in, m_out)
	}
}
