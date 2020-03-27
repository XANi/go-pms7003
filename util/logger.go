package main

import (
	"github.com/XANi/go-pmsX003"
	"github.com/jacobsa/go-serial/serial"
	"log"
	"os"
)

func main() {
	serialDev := os.Getenv("SERIAL_DEV")
	if len(serialDev) == 0 {
		serialDev = "/dev/ttyUSB0"
	}
	options := serial.OpenOptions{
      PortName: serialDev,
      BaudRate: 9600,
      DataBits: 8,
      StopBits: 1,
      InterCharacterTimeout: 2000,
    }
    port, err := serial.Open(options)
    if err != nil {
      log.Fatalf("serial.Open: %v", err)
    }
    defer port.Close()
	for {
		data, err := pmsX003.DecodeFrame(port)
		if err != nil {
			log.Printf("err:%s", err)
		} else {
			log.Printf("frame:%+v", data)
		}
	}
}
