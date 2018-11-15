package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
)

const usage = `Usage: test-data <options>

  Options:
    -p <port> : webserver listening port
`

var (
	webClient *WebClient
	wg        sync.WaitGroup
)

type entry struct {
	url  string
	safe bool
}

func mainInternal() error {
	var port int
	flag.IntVar(&port, "p", 8080, "listening port")
	flag.Usage = func() {
		fmt.Println(usage)
	}
	flag.Parse()

	// init client
	webClient = &WebClient{fmt.Sprintf("http://localhost:%d", port)}

	// start the test
	wg.Add(1)
	go sendRequests(&wg)

	// wait for completion
	wg.Wait()
	webClient.Close()
	fmt.Println("Test completed")
	return nil
}

func sendRequests(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		url := fmt.Sprintf("/info/v1/data%02d", i+1)
		if code, err := webClient.Write(url); err != nil {
			fmt.Println("Error:", err.Error())
		} else {
			fmt.Println("Status:", code)
		}
	}
}

func main() {
	if err := mainInternal(); err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	os.Exit(0)
}
