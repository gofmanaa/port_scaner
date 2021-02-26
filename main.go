package main

import (
	"flag"
	"fmt"
	"github.com/gofmanaa/port_scanner/data"
	"net"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var address string

const ttl = 500

func init() {
	flag.StringVar(&address, "addr", "localhost", "Host name")
}

func main() {
	flag.Parse()
	out := make(chan int, 100)

	scanPorts(address, out)
	getChanel(out)
}

func isOpenPort(address string, port int) bool {
	address = address + ":" + strconv.Itoa(port)
	con, err := net.DialTimeout("tcp", address, time.Microsecond*ttl)
	if err != nil {
		return false
	}

	con.Close()

	return true
}

func scanPorts(address string, out chan int) {
	var wg sync.WaitGroup

	wg.Add(1)
	go startScan(address, &wg, out)

	gorutines := runtime.NumGoroutine()
	fmt.Println("gorut: ", gorutines)
	wg.Wait()
}

func startScan(address string, wg *sync.WaitGroup, out chan int) {
	for port := 1; port < 1<<16; port++ {
		if isOpenPort(address, port) {
			out <- port
		}
	}
	close(out)
	wg.Done()
}

func getChanel(in <-chan int) *data.Set {
	res := data.NewSet()

	for openPort := range in {
		res.Add(openPort)
		fmt.Printf("Port %d open\n", openPort)
	}

	fmt.Println(*res)
	return res
}
