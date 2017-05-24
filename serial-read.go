package main

import (
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	"gopkg.in/alexcesaro/statsd.v2"
	"log"
	"strconv"
	"strings"
	//"time"
)

type Opt_statsd struct {
	Address     statsd.Option
	Network     statsd.Option
	FlushPeriod statsd.Option
}

func main() {

	options_serial := serial.OpenOptions{
		PortName:              "/dev/cu.usbmodem1411",
		BaudRate:              9600,
		DataBits:              6,
		StopBits:              1,
		InterCharacterTimeout: 200,
		MinimumReadSize:       7,
	}

	options_statsd := Opt_statsd{
		Address:     statsd.Address("localhost:8125"),
		Network:     statsd.Network("udp"),
		FlushPeriod: statsd.FlushPeriod(100),
	}

	// Open port serial (to receive data from the sensor)
	port, err := serial.Open(options_serial)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	// Connect to statsd (to send data)
	client, err := statsd.New(options_statsd.Address,
		options_statsd.Network,
		options_statsd.FlushPeriod)
	if err != nil {
		log.Fatalf("statsd.New: %v", err)
	}

	// Make sure to close the ports later
	defer port.Close()
	defer client.Close()

	// Read from sensor
	temp_raw := make([]byte, 7)
	for {
		n, err1 := port.Read(temp_raw)
		if err1 != nil {
			log.Fatalf("port.Read: %v", err1)
		}

		//Converting temperature-reading from array to string to float
		temp_string := strings.TrimSpace(string(temp_raw[:n]))
		temp_float, _ := strconv.ParseFloat(temp_string, 64)

		//Adding timestamp
		//now := time.Now()

		// Til senere?: gjøre omregning fra sensordata til celsius her rett før
		// det sendes videre til statsd.

		//Sending to statd
		client.Gauge("temperature", temp_float)

		//Console output
		fmt.Printf("temperature: %v\n", temp_float)
	}

}
