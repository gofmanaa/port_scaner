package main

import (
	"flag"
	"fmt"
	"net"
	"runtime"
	"sync"
	"time"

	"github.com/gofmanaa/port_scanner/data"
)

var address string

const ttl = 2

func init() {
	flag.StringVar(&address, "addr", "localhost", "Host name")
}

func main() {
	fmt.Println("Start sacanning...")
	flag.Parse()
	ch := make(chan int, 1<<16)

	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)

	fmt.Println("CpuNum: ", cpus)
	var wg sync.WaitGroup

	for port := 1; port < 1<<16; port++ {
		wg.Add(1)
		go startScan(address, port, &wg, ch)
	}

	wg.Wait()
	close(ch)

	openPortsMap := getChanel(ch)
	openPorts := openPortsMap.GetInt()

	if len(openPorts) > 0 {
		for _, openPort := range openPorts {
			fmt.Printf("Port %d [open]\n", openPort)
		}
	} else {
		fmt.Printf("Host %s have no open ports.\n", address)
	}
	fmt.Println("Done.")
}

func isOpenPort(address string, port int) bool {
	address = fmt.Sprintf("%s:%d", address, port)
	con, err := net.DialTimeout("tcp", address, time.Second*ttl)
	if err == nil {
		_ = con.Close()
		return true
	}

	return false
}

func startScan(address string, port int, wg *sync.WaitGroup, out chan int) {
	if isOpenPort(address, port) {
		out <- port
	}

	wg.Done()
}

func getChanel(in <-chan int) *data.Set {
	res := data.NewSet()

	for openPort := range in {
		res.Add(openPort)
	}

	return res
}
