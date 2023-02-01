package SkyLine_Network

import (
	"fmt"
	"net/http"
)

func Server(thread bool) {
	if thread {
		go func() {
			ServerThread()
		}()
	} else {
		ServerThread()
	}
}

func ServerThread() {
	port := VerifyConnection()
	server := &http.Server{
		Addr: ":" + port,
	}
	if server.Addr != "" {
		fmt.Printf("[Information] Server has been started on port %s (http://localhost:%s)\n", port, port)
		http.HandleFunc("/", SkyLine_Network_Handler)
		server.ListenAndServe()
	} else {
		fmt.Println("Error: Server could not find a verified port to work off of")
	}
}
