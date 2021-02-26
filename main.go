package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"sync"
	"github.com/gofmanaa/port_scanertest/data"
	"time"
)

var address string

const ttl = 2000

func init() {
	flag.StringVar(&address, "addr", "localhost", "Host name")
}

func main() {
	flag.Parse()
	scanPorts(address)
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

func scanPorts(address string) {
	var wg sync.WaitGroup
	res := data.NewSet()
	for port := 1; port < 1<<16; port++ {
		wg.Add(1)
		startScan(address, port, &wg, res)
	}

	wg.Wait()
	fmt.Println(*res)
}

func startScan(address string, port int, wg *sync.WaitGroup, out *data.Set) {
	if isOpenPort(address, port) {
		out.Add(port)
		fmt.Printf("Port %d open\n", port)
	}

	wg.Done()
}
