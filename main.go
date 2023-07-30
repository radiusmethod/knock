package main

import (
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"sync"
)

var ports []int
var knocksReceived []int
var listenAddress = "127.0.0.1"

const numberOfKnocks int = 3

func main() {
	banner()
	resultChan := make(chan int)
	var wg sync.WaitGroup

	wg.Add(numberOfKnocks)

	for i := 0; i < numberOfKnocks; i++ {
		ports = append(ports, randomPort())
		go createUdpServer(ports[i], &wg, resultChan)
	}

	fmt.Printf("Knocks sequence: %v\n", ports)

	for {
		result := <-resultChan
		checkKnocks(result)
	}
}

func checkKnocks(portReceived int) {
	knocksReceived = append(knocksReceived, portReceived)
	fmt.Printf("Knocks received so far: %v", knocksReceived)
	if findSubArray(knocksReceived, ports) > -1 {
		webServer()
	}
}

func findSubArray(mainArray, subArray []int) int {
	for i := 0; i <= len(mainArray)-len(subArray); i++ {
		found := true
		for j := 0; j < len(subArray); j++ {
			if mainArray[i+j] != subArray[j] {
				found = false
				break
			}
		}
		if found {
			return i
		}
	}
	return -1 // Subarray not found
}

func webServer() {
	// Define the port number for the server
	fmt.Println("Webserver activated!")
	port := "9999"

	// Set up the route and handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>Found!</h1>")
	})

	// Start the server and handle incoming requests
	fmt.Printf("Server started on port %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

func createUdpServer(port int, wg *sync.WaitGroup, resultChan chan<- int) {
	defer wg.Done()

	listenAddressAndPort := fmt.Sprintf("%s:%d", listenAddress, port)

	// Resolve the UDP address
	udpAddr, err := net.ResolveUDPAddr("udp", listenAddressAndPort)
	if err != nil {
		fmt.Printf("Error resolving UDP address: %v\n", err)
		return
	}

	// Create a UDP connection
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Printf("Error creating UDP connection: %v\n", err)
		return
	}
	defer conn.Close()

	fmt.Printf("UDP server started on %s\n", listenAddressAndPort)

	// Buffer to hold incoming data
	buffer := make([]byte, 1024)

	for {
		// Read data from the UDP connection
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("Error reading from UDP: %v\n", err)
			continue
		}

		data := buffer[:n]
		fmt.Printf("Received from %s on port %d: %s\n", addr.String(), port, string(data))
		resultChan <- port

	}
}

func randomPort() int {
	min := 10000
	max := 49151
	return rand.Intn(max-min+1) + min
}

func banner() {
	fmt.Println("____    ____  _____   ______           _____          _____    ____    ____")
	fmt.Println("|    |  |    ||\\    \\ |\\     \\     ____|\\    \\     ___|\\    \\  |    |  |    |")
	fmt.Println("|    |  |    | \\\\    \\| \\     \\   /     /\\    \\   /    /\\    \\ |    |  |    |")
	fmt.Println("|    | /    //  \\|    \\  \\     | /     /  \\    \\ |    |  |    ||    | /    //")
	fmt.Println("|    |/ _ _//    |     \\  |    ||     |    |    ||    |  |____||    |/ _ _//")
	fmt.Println("|    |\\    \\'    |      \\ |    ||     |    |    ||    |   ____ |    |\\    \\'")
	fmt.Println("|    | \\    \\    |    |\\ \\|    ||\\     \\  /    /||    |  |    ||    | \\    \\")
	fmt.Println("|____|  \\____\\   |____||\\_____/|| \\_____\\/____/ ||\\ ___\\/    /||____|  \\____\\")
	fmt.Println("|    |   |    |  |    |/ \\|   || \\ |    ||    | /| |   /____/ ||    |   |    |")
	fmt.Println("|____|   |____|  |____|   |___|/  \\|____||____|/  \\|___|    | /|____|   |____|")
	fmt.Println("\\(       )/      \\(       )/       \\(    )/       \\( |____|/   \\(       )/")
	fmt.Println("'       '        '       '         '    '         '   )/       '       '")
	fmt.Println("                                                      '")
}
